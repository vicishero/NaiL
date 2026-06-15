// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// CreateSysDictionaryDetail 创建字典详情
func (d *adminDao) CreateSysDictionaryDetail(ctx context.Context, detail *dbr.SysDictionaryDetail) error {
	return d.db.WithContext(ctx).Create(detail).Error
}

// DeleteSysDictionaryDetail 删除字典详情
func (d *adminDao) DeleteSysDictionaryDetail(ctx context.Context, id uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysDictionaryDetail{}, id).Error
}

// UpdateSysDictionaryDetail 更新字典详情
func (d *adminDao) UpdateSysDictionaryDetail(ctx context.Context, detail *dbr.SysDictionaryDetail) error {
	return d.db.WithContext(ctx).Save(detail).Error
}

// FindSysDictionaryDetail 根据ID查找字典详情
func (d *adminDao) FindSysDictionaryDetail(ctx context.Context, id uint) (*dbr.SysDictionaryDetail, error) {
	var detail dbr.SysDictionaryDetail
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&detail).Error
	if err != nil {
		return nil, err
	}
	return &detail, nil
}

// GetSysDictionaryDetailList 获取字典详情列表
func (d *adminDao) GetSysDictionaryDetailList(ctx context.Context, dictionaryID uint) ([]*dbr.SysDictionaryDetail, error) {
	var list []*dbr.SysDictionaryDetail
	err := d.db.WithContext(ctx).Where("dictionary_id = ?", dictionaryID).Order("sort asc").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
