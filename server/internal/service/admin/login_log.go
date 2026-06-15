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

// CreateLoginLog 创建登录日志
func (s *adminService) CreateLoginLog(ctx context.Context, log *dbr.SysLoginLog) error {
	return s.dao.CreateLoginLog(ctx, log)
}

// DeleteLoginLog 删除登录日志
func (s *adminService) DeleteLoginLog(ctx context.Context, id uint) error {
	return s.dao.DeleteLoginLog(ctx, id)
}

// DeleteLoginLogByIds 批量删除登录日志
func (s *adminService) DeleteLoginLogByIds(ctx context.Context, ids []uint) error {
	return s.dao.DeleteLoginLogByIds(ctx, ids)
}

// FindLoginLog 根据ID查找登录日志
func (s *adminService) FindLoginLog(ctx context.Context, id uint) (*dbr.SysLoginLog, error) {
	return s.dao.FindLoginLog(ctx, id)
}

// GetLoginLogList 获取登录日志列表
func (s *adminService) GetLoginLogList(ctx context.Context, req *admin.GetLoginLogListReq) (*admin.GetLoginLogListResp, error) {
	total, list, err := s.dao.GetLoginLogList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("获取登录日志列表失败: %v", err)
	}
	return &admin.GetLoginLogListResp{
		Total: total,
		List:  list,
	}, nil
}
