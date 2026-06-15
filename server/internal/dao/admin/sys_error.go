// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// CreateSysError 创建错误日志
func (d *adminDao) CreateSysError(ctx context.Context, err *dbr.SysError) error {
	return d.db.WithContext(ctx).Create(err).Error
}

// DeleteSysError 删除错误日志
func (d *adminDao) DeleteSysError(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysError{}, id).Error
}

// DeleteSysErrorByIds 批量删除错误日志
func (d *adminDao) DeleteSysErrorByIds(ctx context.Context, ids []uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysError{}, ids).Error
}

// UpdateSysError 更新错误日志
func (d *adminDao) UpdateSysError(ctx context.Context, err *dbr.SysError) error {
	return d.db.WithContext(ctx).Save(err).Error
}

// FindSysError 根据ID查找错误日志
func (d *adminDao) FindSysError(ctx context.Context, id uint) (*dbr.SysError, error) {
	var err dbr.SysError
	dbErr := d.db.WithContext(ctx).Where("id = ?", id).First(&err).Error
	if dbErr != nil {
		return nil, dbErr
	}
	return &err, nil
}

// GetSysErrorList 获取错误日志列表
func (d *adminDao) GetSysErrorList(ctx context.Context, req *admin.GetSysErrorListReq) (int64, []*dbr.SysError, error) {
	var total int64
	var list []*dbr.SysError

	db := d.db.WithContext(ctx).Model(&dbr.SysError{})
	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
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

// GetSysErrorPublic 获取公开的常见错误
func (d *adminDao) GetSysErrorPublic(ctx context.Context) ([]*dbr.SysError, error) {
	var list []*dbr.SysError
	err := d.db.WithContext(ctx).Where("is_public = ?", true).Order("created_at desc").Limit(50).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// GetSysErrorSolution 获取错误解决方案
func (d *adminDao) GetSysErrorSolution(ctx context.Context, id uint) (string, error) {
	var err dbr.SysError
	dbErr := d.db.WithContext(ctx).Where("id = ?", id).First(&err).Error
	if dbErr != nil {
		return "", dbErr
	}
	return err.Solution, nil
}
