// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// CreateApiToken 创建API Token
func (d *adminDao) CreateApiToken(ctx context.Context, token *dbr.SysApiToken) error {
	return d.db.WithContext(ctx).Create(token).Error
}

// DeleteApiToken 删除API Token
func (d *adminDao) DeleteApiToken(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysApiToken{}, id).Error
}

// DeleteApiTokenByIds 批量删除API Token
func (d *adminDao) DeleteApiTokenByIds(ctx context.Context, ids []uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysApiToken{}, ids).Error
}

// UpdateApiToken 更新API Token
func (d *adminDao) UpdateApiToken(ctx context.Context, token *dbr.SysApiToken) error {
	return d.db.WithContext(ctx).Save(token).Error
}

// GetApiTokenByID 根据ID获取API Token
func (d *adminDao) GetApiTokenByID(ctx context.Context, id uint) (*dbr.SysApiToken, error) {
	var token dbr.SysApiToken
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// GetApiTokenList 获取API Token列表
func (d *adminDao) GetApiTokenList(ctx context.Context, req *admin.GetApiTokenListReq) (int64, []*dbr.SysApiToken, error) {
	var total int64
	var list []*dbr.SysApiToken

	db := d.db.WithContext(ctx).Model(&dbr.SysApiToken{})
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return 0, nil, err
	}

	offset := (req.Page - 1) * req.PageSize
	if err := db.Order("created_at desc").Offset(offset).Limit(req.PageSize).Find(&list).Error; err != nil {
		return 0, nil, err
	}
	return total, list, nil
}

// VerifyApiToken 验证Token有效性
func (d *adminDao) VerifyApiToken(ctx context.Context, tokenStr string) (*dbr.SysApiToken, error) {
	var token dbr.SysApiToken
	now := "NOW()" // 使用数据库时间
	err := d.db.WithContext(ctx).Where("token = ? AND status = 1 AND expire_time > ?", tokenStr, now).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}
