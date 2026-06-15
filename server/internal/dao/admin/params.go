// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
	"github.com/vicishero/NaiL/server/internal/core/admin"
)

// CreateSysParams 创建参数
func (d *adminDao) CreateSysParams(ctx context.Context, params *dbr.SysParams) error {
	return d.db.WithContext(ctx).Create(params).Error
}

// DeleteSysParams 删除参数
func (d *adminDao) DeleteSysParams(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysParams{}, id).Error
}

// DeleteSysParamsByIds 批量删除参数
func (d *adminDao) DeleteSysParamsByIds(ctx context.Context, ids []uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysParams{}, ids).Error
}

// UpdateSysParams 更新参数
func (d *adminDao) UpdateSysParams(ctx context.Context, params *dbr.SysParams) error {
	return d.db.WithContext(ctx).Save(params).Error
}

// FindSysParams 根据ID查找参数
func (d *adminDao) FindSysParams(ctx context.Context, id uint) (*dbr.SysParams, error) {
	var params dbr.SysParams
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&params).Error
	if err != nil {
		return nil, err
	}
	return &params, nil
}

// GetSysParamsList 获取参数列表
func (d *adminDao) GetSysParamsList(ctx context.Context, req *admin.GetSysParamsListReq) (int64, []*dbr.SysParams, error) {
	var total int64
	var list []*dbr.SysParams

	db := d.db.WithContext(ctx).Model(&dbr.SysParams{})
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Key != "" {
		db = db.Where("key LIKE ?", "%"+req.Key+"%")
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

// GetSysParamByKey 根据键获取参数值
func (d *adminDao) GetSysParamByKey(ctx context.Context, key string) (string, error) {
	var params dbr.SysParams
	err := d.db.WithContext(ctx).Where("key = ?", key).First(&params).Error
	if err != nil {
		return "", err
	}
	return params.Value, nil
}
