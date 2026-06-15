// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// CreateApiToken 创建API Token
func (s *adminService) CreateApiToken(ctx context.Context, req *admin.CreateApiTokenReq) (*dbr.SysApiToken, error) {
	// 生成随机Token字符串
	tokenStr, err := generateApiToken()
	if err != nil {
		return nil, fmt.Errorf("生成Token失败: %v", err)
	}

	// 设置默认过期时间（30天）
	expireTime := req.ExpireTime
	if expireTime.IsZero() {
		expireTime = time.Now().AddDate(0, 0, 30)
	}

	status := 1 // 默认启用
	token := &dbr.SysApiToken{
		Token:       tokenStr,
		Name:        req.Name,
		UserID:      req.UserID,
		Username:    req.Username,
		Status:      &status,
		ExpireTime:  expireTime,
		Permissions: req.Permissions,
		Remarks:     req.Remarks,
	}

	if err := s.dao.CreateApiToken(ctx, token); err != nil {
		return nil, err
	}
	return token, nil
}

// DeleteApiToken 删除API Token
func (s *adminService) DeleteApiToken(ctx context.Context, id uint) error {
	return s.dao.DeleteApiToken(ctx, id)
}

// DeleteApiTokenByIds 批量删除API Token
func (s *adminService) DeleteApiTokenByIds(ctx context.Context, ids []uint) error {
	return s.dao.DeleteApiTokenByIds(ctx, ids)
}

// UpdateApiTokenStatus 更新API Token状态
func (s *adminService) UpdateApiTokenStatus(ctx context.Context, id uint, status int) error {
	token, err := s.dao.GetApiTokenByID(ctx, id)
	if err != nil {
		return err
	}
	*token.Status = status
	return s.dao.UpdateApiToken(ctx, token)
}

// GetApiTokenList 获取API Token列表
func (s *adminService) GetApiTokenList(ctx context.Context, req *admin.GetApiTokenListReq) (*admin.GetApiTokenListResp, error) {
	total, list, err := s.dao.GetApiTokenList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("获取Token列表失败: %v", err)
	}
	return &admin.GetApiTokenListResp{
		Total: total,
		List:  list,
	}, nil
}

// VerifyApiToken 验证Token有效性
func (s *adminService) VerifyApiToken(ctx context.Context, tokenStr string) (*dbr.SysApiToken, error) {
	return s.dao.VerifyApiToken(ctx, tokenStr)
}

// generateApiToken 生成随机Token字符串
func generateApiToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
