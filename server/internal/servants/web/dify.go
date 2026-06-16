// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicishero/NaiL/server/internal/conf"
	"github.com/sirupsen/logrus"
)

// DifyChatReq Dify 聊天请求
type DifyChatReq struct {
	Query          string   `json:"query" binding:"required"`
	User           string   `json:"user"`
	ConversationID string   `json:"conversation_id"`
}

// ChatWithDify Dify AI 聊天代理（SSE 流式转发）
func (s *coreSrv) ChatWithDify(c *gin.Context) {
	var req DifyChatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	if req.User == "" {
		req.User = "anonymous"
	}

	difyURL := "https://api.dify.ai"
	difyKey := ""
	if conf.DifySetting != nil {
		if conf.DifySetting.URL != "" {
			difyURL = conf.DifySetting.URL
		}
		difyKey = conf.DifySetting.APIKey
	}
	if difyKey == "" {
		c.JSON(http.StatusBadGateway, gin.H{"code": 502, "msg": "Dify API Key 未配置"})
		return
	}

	difyReq := map[string]interface{}{
		"inputs":          map[string]interface{}{},
		"query":           req.Query,
		"user":            req.User,
		"response_mode":   "streaming",
		"conversation_id": req.ConversationID,
	}
	body, _ := json.Marshal(difyReq)

	proxyReq, err := http.NewRequestWithContext(c.Request.Context(), "POST",
		strings.TrimRight(difyURL, "/")+"/v1/chat-messages",
		strings.NewReader(string(body)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "创建请求失败"})
		return
	}
	proxyReq.Header.Set("Content-Type", "application/json")
	proxyReq.Header.Set("Authorization", "Bearer "+difyKey)

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		logrus.Errorf("Dify request failed: %v", err)
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

	// SSE 流式转发
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Flush()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		c.Writer.Write(line)
		c.Writer.Flush()
	}
}
