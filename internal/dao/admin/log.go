// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/rocboss/paopao-ce/internal/dao/admin/dbr"
	"github.com/rocboss/paopao-ce/internal/core/admin"
)

// CreateOperationLog 创建操作日志
func (d *adminDao) CreateOperationLog(ctx context.Context, log *dbr.SysOperationRecord) error {
	return d.db.WithContext(ctx).Create(log).Error
}

// GetOperationLogList 获取操作日志列表
func (d *adminDao) GetOperationLogList(ctx context.Context, req *admin.GetOperationLogListReq) (int64, []*dbr.SysOperationRecord, error) {
	var total int64
	var list []*dbr.SysOperationRecord

	db := d.db.WithContext(ctx).Model(&dbr.SysOperationRecord{})
	if req.Username != "" {
		db = db.Joins("LEFT JOIN sys_users ON sys_users.id = sys_operation_records.user_id").
			Where("sys_users.username LIKE ?", "%"+req.Username+"%")
	}
	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
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

// GetOperationLogByID 根据ID获取操作日志
func (d *adminDao) GetOperationLogByID(ctx context.Context, logId uint) (*dbr.SysOperationRecord, error) {
	var log dbr.SysOperationRecord
	err := d.db.WithContext(ctx).Where("id = ?", logId).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// DeleteOperationLog 删除操作日志
func (d *adminDao) DeleteOperationLog(ctx context.Context, logId uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysOperationRecord{}, logId).Error
}
