// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicishero/NaiL/server/internal/conf"
	"github.com/sirupsen/logrus"
)

type DifyChatReq struct {
	Query     string `json:"query" binding:"required"`
	User      string `json:"user"`
	KolUserID int64  `json:"kol_user_id"`
}

type ConversationReq struct {
	KolUserID int64 `json:"kol_user_id"`
}

type chatConversation struct {
	ID                  int64  `gorm:"column:id"`
	UserID              int64  `gorm:"column:user_id"`
	KolUserID           int64  `gorm:"column:kol_user_id"`
	DifyConversationID  string `gorm:"column:dify_conversation_id"`
	CreatedOn           int64  `gorm:"column:created_on"`
	ModifiedOn          int64  `gorm:"column:modified_on"`
}

func (chatConversation) TableName() string { return "p_chat_conversation" }

// GetOrCreateConversation 获取或创建会话
func (s *coreSrv) GetOrCreateConversation(c *gin.Context) {
	uid, ok := getUID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
		return
	}

	var req ConversationReq
	c.ShouldBindJSON(&req)

	var conv chatConversation
	db := conf.MustGormDB().WithContext(c.Request.Context())
	err := db.Where("user_id = ? AND kol_user_id = ? AND is_del = 0", uid, req.KolUserID).First(&conv).Error
	if err != nil {
		now := time.Now().Unix()
		conv = chatConversation{
			UserID:     uid,
			KolUserID:  req.KolUserID,
			CreatedOn:  now,
			ModifiedOn: now,
		}
		if createErr := db.Create(&conv).Error; createErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "创建会话失败"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"id":                    conv.ID,
			"dify_conversation_id":  conv.DifyConversationID,
			"kol_user_id":           conv.KolUserID,
		},
	})
}

// GetChatHistory 获取聊天历史
func (s *coreSrv) GetChatHistory(c *gin.Context) {
	uid, ok := getUID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
		return
	}

	kolUserID := int64(0)
	if v := c.Query("kol_user_id"); v != "" {
		fmt.Sscanf(v, "%d", &kolUserID)
	}

	var conv chatConversation
	db := conf.MustGormDB().WithContext(c.Request.Context())
	err := db.Where("user_id = ? AND kol_user_id = ? AND is_del = 0", uid, kolUserID).First(&conv).Error
	if err != nil || conv.DifyConversationID == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"messages": []interface{}{}, "dify_conversation_id": ""}})
		return
	}

	messages := fetchDifyMessages(conv.DifyConversationID, fmt.Sprint(uid), kolUserID)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"messages":              messages,
			"dify_conversation_id":  conv.DifyConversationID,
		},
	})
}

// ChatWithDify Dify AI 聊天代理（SSE 流式 + 会话持久化 + KOL 个性化）
func (s *coreSrv) ChatWithDify(c *gin.Context) {
	var req DifyChatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误"})
		return
	}

	uid, ok := getUID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请先登录"})
		return
	}
	userID := fmt.Sprint(uid)

	difyURL, difyKey := getDifyConfig(req.KolUserID)
	if difyKey == "" {
		c.JSON(http.StatusBadGateway, gin.H{"code": 502, "msg": "Dify API Key 未配置"})
		return
	}

	// 读取 KOL 的 system_prompt
	systemPrompt := getKOLSystemPrompt(req.KolUserID)
	inputs := map[string]interface{}{}
	if systemPrompt != "" {
		inputs["system_prompt"] = systemPrompt
	}

	// 查找已有会话
	difyConversationID := ""
	db := conf.MustGormDB().WithContext(c.Request.Context())
	var conv chatConversation
	if err := db.Where("user_id = ? AND kol_user_id = ? AND is_del = 0", uid, req.KolUserID).First(&conv).Error; err == nil {
		difyConversationID = conv.DifyConversationID
	}

	difyReq := map[string]interface{}{
		"inputs":          inputs,
		"query":           req.Query,
		"user":            userID,
		"response_mode":   "streaming",
		"conversation_id": difyConversationID,
	}
	body, _ := json.Marshal(difyReq)

	proxyReq, _ := http.NewRequestWithContext(c.Request.Context(), "POST",
		strings.TrimRight(difyURL, "/")+"/v1/chat-messages", strings.NewReader(string(body)))
	proxyReq.Header.Set("Content-Type", "application/json")
	proxyReq.Header.Set("Authorization", "Bearer "+difyKey)

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"code": 502, "msg": "Dify 服务不可用"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		logrus.Errorf("Dify error status=%d body=%s", resp.StatusCode, string(bodyBytes))
		c.JSON(http.StatusBadGateway, gin.H{"code": 502, "msg": "Dify 返回错误"})
		return
	}

	// SSE 流式转发 + 解析 conversation_id
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Flush()

	newConvID := ""
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		// 解析 conversation_id
		if newConvID == "" && strings.HasPrefix(string(line), "data: ") {
			var event map[string]interface{}
			if json.Unmarshal([]byte(strings.TrimPrefix(string(line), "data: ")), &event) == nil {
				if cid, ok := event["conversation_id"].(string); ok && cid != "" {
					newConvID = cid
				}
			}
		}
		c.Writer.Write(line)
		c.Writer.Flush()
	}

	// 持久化 conversation_id
	if newConvID != "" {
		now := time.Now().Unix()
		if conv.ID > 0 {
			db.Model(&conv).Updates(map[string]interface{}{
				"dify_conversation_id": newConvID, "modified_on": now,
			})
		} else {
			db.Create(&chatConversation{
				UserID: uid, KolUserID: req.KolUserID,
				DifyConversationID: newConvID, CreatedOn: now, ModifiedOn: now,
			})
		}
	}
}

// getUID 从 context 获取用户 ID
func getUID(c *gin.Context) (int64, bool) {
	if uid, exists := c.Get("UID"); exists {
		if v, ok := uid.(int64); ok {
			return v, true
		}
	}
	// fallback: 从 JWT 中间件设置的 USER 中读取
	if u, exists := c.Get("USER"); exists {
		if user, ok := u.(interface{ GetID() int64 }); ok {
			return user.GetID(), true
		}
	}
	return 0, false
}

// getDifyConfig 获取 Dify 配置（优先使用 KOL 独立 key）
func getDifyConfig(kolUserID int64) (url, apiKey string) {
	url = "https://api.dify.ai"
	if conf.DifySetting != nil {
		if conf.DifySetting.URL != "" {
			url = conf.DifySetting.URL
		}
		apiKey = conf.DifySetting.APIKey
	}
	// KOL 独立 API Key
	if kolUserID > 0 {
		var profile struct{ ApiKey string `gorm:"column:api_key"` }
		db := conf.MustGormDB()
		if err := db.Table("p_kol_profile").Where("user_id = ? AND is_del = 0", kolUserID).First(&profile).Error; err == nil && profile.ApiKey != "" {
			apiKey = profile.ApiKey
		}
	}
	return
}

// getKOLSystemPrompt 获取 KOL 系统提示词
func getKOLSystemPrompt(kolUserID int64) string {
	if kolUserID <= 0 {
		return ""
	}
	var profile struct{ SystemPrompt string `gorm:"column:system_prompt"` }
	db := conf.MustGormDB()
	if err := db.Table("p_kol_profile").Where("user_id = ? AND is_del = 0", kolUserID).First(&profile).Error; err == nil {
		return profile.SystemPrompt
	}
	return ""
}

// fetchDifyMessages 从 Dify 获取历史消息
func fetchDifyMessages(conversationID, userID string, kolUserID int64) []interface{} {
	difyURL, difyKey := getDifyConfig(kolUserID)
	if difyKey == "" {
		return []interface{}{}
	}

	url := fmt.Sprintf("%s/v1/messages?conversation_id=%s&user=%s&limit=100",
		strings.TrimRight(difyURL, "/"), conversationID, userID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []interface{}{}
	}
	req.Header.Set("Authorization", "Bearer "+difyKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorf("fetchDifyMessages failed: %v", err)
		return []interface{}{}
	}
	defer resp.Body.Close()

	var result struct {
		Data []interface{} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []interface{}{}
	}
	return result.Data
}
