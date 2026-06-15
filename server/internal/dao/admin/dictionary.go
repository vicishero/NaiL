// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
	"github.com/vicishero/NaiL/server/internal/core/admin"
)

// CreateSysDictionary 创建字典
func (d *adminDao) CreateSysDictionary(ctx context.Context, dict *dbr.SysDictionary) error {
	return d.db.WithContext(ctx).Create(dict).Error
}

// DeleteSysDictionary 删除字典
func (d *adminDao) DeleteSysDictionary(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysDictionary{}, id).Error
}

// UpdateSysDictionary 更新字典
func (d *adminDao) UpdateSysDictionary(ctx context.Context, dict *dbr.SysDictionary) error {
	return d.db.WithContext(ctx).Save(dict).Error
}

// FindSysDictionary 根据ID查找字典
func (d *adminDao) FindSysDictionary(ctx context.Context, id uint) (*dbr.SysDictionary, error) {
	var dict dbr.SysDictionary
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&dict).Error
	if err != nil {
		return nil, err
	}
	// 加载详情
	var details []dbr.SysDictionaryDetail
	d.db.WithContext(ctx).Where("dictionary_id = ?", id).Find(&details)
	dict.Details = details
	return &dict, nil
}

// GetSysDictionaryList 获取字典列表
func (d *adminDao) GetSysDictionaryList(ctx context.Context, req *admin.GetSysDictionaryListReq) (int64, []*dbr.SysDictionary, error) {
	var total int64
	var list []*dbr.SysDictionary

	db := d.db.WithContext(ctx).Model(&dbr.SysDictionary{})
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Type != "" {
		db = db.Where("type LIKE ?", "%"+req.Type+"%")
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

// GetSysDictionaryByType 根据类型获取字典
func (d *adminDao) GetSysDictionaryByType(ctx context.Context, typeStr string) (*dbr.SysDictionary, error) {
	var dict dbr.SysDictionary
	err := d.db.WithContext(ctx).Where("type = ?", typeStr).First(&dict).Error
	if err != nil {
		return nil, err
	}
	// 加载详情
	var details []dbr.SysDictionaryDetail
	d.db.WithContext(ctx).Where("dictionary_id = ?", dict.ID).Order("sort asc").Find(&details)
	dict.Details = details
	return &dict, nil
}
