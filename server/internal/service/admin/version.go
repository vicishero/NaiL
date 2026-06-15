// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// CreateSysVersion 创建版本
func (s *adminService) CreateSysVersion(ctx context.Context, req *admin.CreateSysVersionReq) (*dbr.SysVersion, error) {
	status := 1 // 默认启用
	version := &dbr.SysVersion{
		Version:     req.Version,
		Name:        req.Name,
		Content:     req.Content,
		Type:        req.Type,
		Status:      &status,
		UserID:      req.UserID,
		Username:    req.Username,
		ReleaseTime: time.Now(),
		FileUrl:     req.FileUrl,
		FileSize:    req.FileSize,
		Md5:         req.Md5,
		Remarks:     req.Remarks,
	}

	if err := s.dao.CreateSysVersion(ctx, version); err != nil {
		return nil, err
	}
	return version, nil
}

// DeleteSysVersion 删除版本
func (s *adminService) DeleteSysVersion(ctx context.Context, id uint) error {
	return s.dao.DeleteSysVersion(ctx, id)
}

// DeleteSysVersionByIds 批量删除版本
func (s *adminService) DeleteSysVersionByIds(ctx context.Context, ids []uint) error {
	return s.dao.DeleteSysVersionByIds(ctx, ids)
}

// UpdateSysVersion 更新版本
func (s *adminService) UpdateSysVersion(ctx context.Context, version *dbr.SysVersion) error {
	return s.dao.UpdateSysVersion(ctx, version)
}

// FindSysVersion 根据ID查找版本
func (s *adminService) FindSysVersion(ctx context.Context, id uint) (*dbr.SysVersion, error) {
	return s.dao.FindSysVersion(ctx, id)
}

// GetSysVersionList 获取版本列表
func (s *adminService) GetSysVersionList(ctx context.Context, req *admin.GetSysVersionListReq) (*admin.GetSysVersionListResp, error) {
	total, list, err := s.dao.GetSysVersionList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("获取版本列表失败: %v", err)
	}
	return &admin.GetSysVersionListResp{
		Total: total,
		List:  list,
	}, nil
}

// ExportVersion 导出版本信息
func (s *adminService) ExportVersion(ctx context.Context, ids []uint) ([]byte, error) {
	// 如果指定了ids，导出指定的；否则导出全部
	var list []*dbr.SysVersion
	db := s.dao.DB().WithContext(ctx).Model(&dbr.SysVersion{})
	if len(ids) > 0 {
		db = db.Where("id IN ?", ids)
	}
	if err := db.Order("created_at desc").Find(&list).Error; err != nil {
		return nil, fmt.Errorf("获取版本列表失败: %v", err)
	}

	// 转成JSON
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("JSON序列化失败: %v", err)
	}
	return data, nil
}

// ImportVersion 导入版本信息
func (s *adminService) ImportVersion(ctx context.Context, data []byte) error {
	var versions []*dbr.SysVersion
	if err := json.Unmarshal(data, &versions); err != nil {
		return fmt.Errorf("JSON反序列化失败: %v", err)
	}

	// 批量导入（注意：跳过已存在的版本号，或者根据需求处理）
	for _, v := range versions {
		// 重置ID，作为新记录插入
		v.ID = 0
		v.CreatedAt = time.Now()
		v.UpdatedAt = time.Now()
		if err := s.dao.CreateSysVersion(ctx, v); err != nil {
			// 继续下一个，不中断整个流程
			continue
		}
	}
	return nil
}

// GenerateVersionJson 生成版本JSON供下载
func (s *adminService) GenerateVersionJson(ctx context.Context, id uint) ([]byte, error) {
	version, err := s.dao.FindSysVersion(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("版本不存在: %v", err)
	}

	data, err := json.MarshalIndent(version, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("JSON序列化失败: %v", err)
	}
	return data, nil
}
