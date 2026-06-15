// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
	"github.com/vicishero/NaiL/server/internal/core/admin"
)

// CreateLoginLog 创建登录日志
func (d *adminDao) CreateLoginLog(ctx context.Context, log *dbr.SysLoginLog) error {
	return d.db.WithContext(ctx).Create(log).Error
}

// DeleteLoginLog 删除登录日志
func (d *adminDao) DeleteLoginLog(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysLoginLog{}, id).Error
}

// DeleteLoginLogByIds 批量删除登录日志
func (d *adminDao) DeleteLoginLogByIds(ctx context.Context, ids []uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysLoginLog{}, ids).Error
}

// FindLoginLog 根据ID查找登录日志
func (d *adminDao) FindLoginLog(ctx context.Context, id uint) (*dbr.SysLoginLog, error) {
	var log dbr.SysLoginLog
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// GetLoginLogList 获取登录日志列表
func (d *adminDao) GetLoginLogList(ctx context.Context, req *admin.GetLoginLogListReq) (int64, []*dbr.SysLoginLog, error) {
	var total int64
	var list []*dbr.SysLoginLog

	db := d.db.WithContext(ctx).Model(&dbr.SysLoginLog{})
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Ip != "" {
		db = db.Where("ip LIKE ?", "%"+req.Ip+"%")
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
