// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"
	"fmt"

	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// CreateSysDictionary 创建字典
func (s *adminService) CreateSysDictionary(ctx context.Context, dict *dbr.SysDictionary) error {
	return s.dao.CreateSysDictionary(ctx, dict)
}

// DeleteSysDictionary 删除字典
func (s *adminService) DeleteSysDictionary(ctx context.Context, id uint) error {
	return s.dao.DeleteSysDictionary(ctx, id)
}

// UpdateSysDictionary 更新字典
func (s *adminService) UpdateSysDictionary(ctx context.Context, dict *dbr.SysDictionary) error {
	return s.dao.UpdateSysDictionary(ctx, dict)
}

// FindSysDictionary 根据ID查找字典
func (s *adminService) FindSysDictionary(ctx context.Context, id uint) (*dbr.SysDictionary, error) {
	return s.dao.FindSysDictionary(ctx, id)
}

// GetSysDictionaryList 获取字典列表
func (s *adminService) GetSysDictionaryList(ctx context.Context, req *admin.GetSysDictionaryListReq) (*admin.GetSysDictionaryListResp, error) {
	total, list, err := s.dao.GetSysDictionaryList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("获取字典列表失败: %v", err)
	}
	return &admin.GetSysDictionaryListResp{
		Total: total,
		List:  list,
	}, nil
}

// GetSysDictionaryByType 根据类型获取字典
func (s *adminService) GetSysDictionaryByType(ctx context.Context, typeStr string) (*dbr.SysDictionary, error) {
	return s.dao.GetSysDictionaryByType(ctx, typeStr)
}

// CreateSysDictionaryDetail 创建字典详情
func (s *adminService) CreateSysDictionaryDetail(ctx context.Context, detail *dbr.SysDictionaryDetail) error {
	return s.dao.CreateSysDictionaryDetail(ctx, detail)
}

// DeleteSysDictionaryDetail 删除字典详情
func (s *adminService) DeleteSysDictionaryDetail(ctx context.Context, id uint) error {
	return s.dao.DeleteSysDictionaryDetail(ctx, id)
}

// UpdateSysDictionaryDetail 更新字典详情
func (s *adminService) UpdateSysDictionaryDetail(ctx context.Context, detail *dbr.SysDictionaryDetail) error {
	return s.dao.UpdateSysDictionaryDetail(ctx, detail)
}

// FindSysDictionaryDetail 根据ID查找字典详情
func (s *adminService) FindSysDictionaryDetail(ctx context.Context, id uint) (*dbr.SysDictionaryDetail, error) {
	return s.dao.FindSysDictionaryDetail(ctx, id)
}

// GetSysDictionaryDetailList 获取字典详情列表
func (s *adminService) GetSysDictionaryDetailList(ctx context.Context, dictionaryID uint) ([]*dbr.SysDictionaryDetail, error) {
	return s.dao.GetSysDictionaryDetailList(ctx, dictionaryID)
}
