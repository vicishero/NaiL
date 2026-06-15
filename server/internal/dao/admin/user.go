// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
	"github.com/vicishero/NaiL/server/internal/core/admin"
)

// GetUserByUsername 根据用户名获取用户信息
func (d *adminDao) GetUserByUsername(ctx context.Context, username string) (*dbr.SysUser, error) {
	var user dbr.SysUser
	err := d.db.WithContext(ctx).Where("username = ? and enable = 1", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID 根据用户ID获取用户信息
func (d *adminDao) GetUserByID(ctx context.Context, userId uint) (*dbr.SysUser, error) {
	var user dbr.SysUser
	err := d.db.WithContext(ctx).Where("id = ? and enable = 1", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func (d *adminDao) UpdateUser(ctx context.Context, user *dbr.SysUser) error {
	return d.db.WithContext(ctx).Save(user).Error
}

// CreateUser 创建用户
func (d *adminDao) CreateUser(ctx context.Context, user *dbr.SysUser) error {
	return d.db.WithContext(ctx).Create(user).Error
}

// DeleteUser 删除用户
func (d *adminDao) DeleteUser(ctx context.Context, userId uint) error {
	return d.db.WithContext(ctx).Delete(&dbr.SysUser{}, userId).Error
}

// GetUserList 获取用户列表
func (d *adminDao) GetUserList(ctx context.Context, req *admin.UserListReq) (int64, []*dbr.SysUser, error) {
	var total int64
	var list []*dbr.SysUser

	db := d.db.WithContext(ctx).Model(&dbr.SysUser{})
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+req.NickName+"%")
	}
	if req.Enable != nil {
		db = db.Where("enable = ?", *req.Enable)
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

// CountUsersByAuthorityId 统计使用该角色的用户数量
func (d *adminDao) CountUsersByAuthorityId(ctx context.Context, authorityId uint) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&dbr.SysUser{}).Where("authority_id = ?", authorityId).Count(&count).Error
	return count, err
}
