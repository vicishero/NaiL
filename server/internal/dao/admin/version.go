// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// CreateSysVersion 创建版本
func (d *adminDao) CreateSysVersion(ctx context.Context, version *dbr.SysVersion) error {
	return d.db.WithContext(ctx).Create(version).Error
}

// DeleteSysVersion 删除版本
func (d *adminDao) DeleteSysVersion(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysVersion{}, id).Error
}

// DeleteSysVersionByIds 批量删除版本
func (d *adminDao) DeleteSysVersionByIds(ctx context.Context, ids []uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysVersion{}, ids).Error
}

// UpdateSysVersion 更新版本
func (d *adminDao) UpdateSysVersion(ctx context.Context, version *dbr.SysVersion) error {
	return d.db.WithContext(ctx).Save(version).Error
}

// FindSysVersion 根据ID查找版本
func (d *adminDao) FindSysVersion(ctx context.Context, id uint) (*dbr.SysVersion, error) {
	var version dbr.SysVersion
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&version).Error
	if err != nil {
		return nil, err
	}
	return &version, nil
}

// GetSysVersionList 获取版本列表
func (d *adminDao) GetSysVersionList(ctx context.Context, req *admin.GetSysVersionListReq) (int64, []*dbr.SysVersion, error) {
	var total int64
	var list []*dbr.SysVersion

	db := d.db.WithContext(ctx).Model(&dbr.SysVersion{})
	if req.Version != "" {
		db = db.Where("version LIKE ?", "%"+req.Version+"%")
	}
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
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
