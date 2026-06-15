// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"
	"fmt"
	"time"

	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// CreateSysError 创建错误日志
func (s *adminService) CreateSysError(ctx context.Context, err *dbr.SysError) error {
	return s.dao.CreateSysError(ctx, err)
}

// RecordError 记录错误（供中间件调用）
func (s *adminService) RecordError(ctx context.Context, errType, message, stack, path, method, ip string, userID uint, username string) {
	err := &dbr.SysError{
		Type:     errType,
		Message:  message,
		Stack:    stack,
		Path:     path,
		Method:   method,
		UserID:   userID,
		Username: username,
		IP:       ip,
		Status:   1, // 未处理
	}
	// 异步记录错误，不影响主流程
	go s.dao.CreateSysError(context.Background(), err)
}

// DeleteSysError 删除错误日志
func (s *adminService) DeleteSysError(ctx context.Context, id uint) error {
	return s.dao.DeleteSysError(ctx, id)
}

// DeleteSysErrorByIds 批量删除错误日志
func (s *adminService) DeleteSysErrorByIds(ctx context.Context, ids []uint) error {
	return s.dao.DeleteSysErrorByIds(ctx, ids)
}

// UpdateSysError 更新错误日志（标记处理状态）
func (s *adminService) UpdateSysError(ctx context.Context, req *admin.UpdateSysErrorReq) error {
	errItem, err := s.dao.FindSysError(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("错误日志不存在: %v", err)
	}

	if req.Status != 0 {
		errItem.Status = req.Status
	}
	if req.Solution != "" {
		errItem.Solution = req.Solution
	}
	if req.HandlerID != 0 {
		errItem.HandlerID = req.HandlerID
	}
	if req.HandlerName != "" {
		errItem.HandlerName = req.HandlerName
	}
	if req.Status == 2 { // 已处理
		now := time.Now()
		errItem.HandleTime = now
	}

	return s.dao.UpdateSysError(ctx, errItem)
}

// FindSysError 根据ID查找错误日志
func (s *adminService) FindSysError(ctx context.Context, id uint) (*dbr.SysError, error) {
	return s.dao.FindSysError(ctx, id)
}

// GetSysErrorList 获取错误日志列表
func (s *adminService) GetSysErrorList(ctx context.Context, req *admin.GetSysErrorListReq) (*admin.GetSysErrorListResp, error) {
	total, list, err := s.dao.GetSysErrorList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("获取错误日志列表失败: %v", err)
	}
	return &admin.GetSysErrorListResp{
		Total: total,
		List:  list,
	}, nil
}

// GetSysErrorPublic 获取公开的常见错误
func (s *adminService) GetSysErrorPublic(ctx context.Context) (*admin.GetSysErrorPublicResp, error) {
	list, err := s.dao.GetSysErrorPublic(ctx)
	if err != nil {
		return nil, fmt.Errorf("获取公开错误列表失败: %v", err)
	}
	return &admin.GetSysErrorPublicResp{
		List: list,
	}, nil
}

// GetSysErrorSolution 获取错误解决方案
func (s *adminService) GetSysErrorSolution(ctx context.Context, id uint) (*admin.GetSysErrorSolutionResp, error) {
	solution, err := s.dao.GetSysErrorSolution(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("获取错误解决方案失败: %v", err)
	}
	return &admin.GetSysErrorSolutionResp{
		Solution: solution,
	}, nil
}
