// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
	"github.com/vicishero/NaiL/server/internal/core/admin"
)

// GetAuthorityApisByIds 根据API ID列表获取API详情
func (d *adminDao) GetAuthorityApisByIds(ctx context.Context, apiIds []uint) ([]*dbr.SysApi, error) {
	var apis []*dbr.SysApi
	err := d.db.WithContext(ctx).Where("id IN ?", apiIds).Find(&apis).Error
	return apis, err
}

// CreateApi 创建API
func (d *adminDao) CreateApi(ctx context.Context, api *dbr.SysApi) error {
	return d.db.WithContext(ctx).Create(api).Error
}

// UpdateApi 更新API
func (d *adminDao) UpdateApi(ctx context.Context, api *dbr.SysApi) error {
	return d.db.WithContext(ctx).Model(&dbr.SysApi{}).Where("id = ?", api.ID).Updates(api).Error
}

// DeleteApi 删除API
func (d *adminDao) DeleteApi(ctx context.Context, apiId uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysApi{}, apiId).Error
}

// GetApiByID 根据ID获取API
func (d *adminDao) GetApiByID(ctx context.Context, apiId uint) (*dbr.SysApi, error) {
	var api dbr.SysApi
	err := d.db.WithContext(ctx).Where("id = ?", apiId).First(&api).Error
	if err != nil {
		return nil, err
	}
	return &api, nil
}

// GetApiList 获取API列表
func (d *adminDao) GetApiList(ctx context.Context, req *admin.GetApiListReq) (int64, []*dbr.SysApi, error) {
	var total int64
	var list []*dbr.SysApi

	db := d.db.WithContext(ctx).Model(&dbr.SysApi{})
	if req.Path != "" {
		db = db.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.ApiGroup != "" {
		db = db.Where("api_group = ?", req.ApiGroup)
	}
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}

	if err := db.Count(&total).Error; err != nil {
		return 0, nil, err
	}

	offset := (req.Page - 1) * req.PageSize
	if err := db.Offset(offset).Limit(req.PageSize).Find(&list).Error; err != nil {
		return 0, nil, err
	}
	return total, list, nil
}

// GetAllApiGroups 获取所有API分组
func (d *adminDao) GetAllApiGroups(ctx context.Context) ([]string, error) {
	var groups []string
	err := d.db.WithContext(ctx).Model(&dbr.SysApi{}).Distinct().Pluck("api_group", &groups).Error
	return groups, err
}
