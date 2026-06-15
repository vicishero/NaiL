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

// CreateSysParams 创建参数
func (s *adminService) CreateSysParams(ctx context.Context, params *dbr.SysParams) error {
	return s.dao.CreateSysParams(ctx, params)
}

// DeleteSysParams 删除参数
func (s *adminService) DeleteSysParams(ctx context.Context, id uint) error {
	return s.dao.DeleteSysParams(ctx, id)
}

// DeleteSysParamsByIds 批量删除参数
func (s *adminService) DeleteSysParamsByIds(ctx context.Context, ids []uint) error {
	return s.dao.DeleteSysParamsByIds(ctx, ids)
}

// UpdateSysParams 更新参数
func (s *adminService) UpdateSysParams(ctx context.Context, params *dbr.SysParams) error {
	return s.dao.UpdateSysParams(ctx, params)
}

// FindSysParams 根据ID查找参数
func (s *adminService) FindSysParams(ctx context.Context, id uint) (*dbr.SysParams, error) {
	return s.dao.FindSysParams(ctx, id)
}

// GetSysParamsList 获取参数列表
func (s *adminService) GetSysParamsList(ctx context.Context, req *admin.GetSysParamsListReq) (*admin.GetSysParamsListResp, error) {
	total, list, err := s.dao.GetSysParamsList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("获取参数列表失败: %v", err)
	}
	return &admin.GetSysParamsListResp{
		Total: total,
		List:  list,
	}, nil
}

// GetSysParamByKey 根据键获取参数值
func (s *adminService) GetSysParamByKey(ctx context.Context, key string) (string, error) {
	return s.dao.GetSysParamByKey(ctx, key)
}
