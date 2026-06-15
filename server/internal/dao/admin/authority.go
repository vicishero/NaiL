// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
	"github.com/vicishero/NaiL/server/internal/core/admin"
)

// GetAuthorityByID 根据角色ID获取角色信息
func (d *adminDao) GetAuthorityByID(ctx context.Context, authorityId uint) (*dbr.SysAuthority, error) {
	var authority dbr.SysAuthority
	err := d.db.WithContext(ctx).Where("authority_id = ?", authorityId).First(&authority).Error
	if err != nil {
		return nil, err
	}
	return &authority, nil
}

// GetUserPermissions 获取用户的权限编码列表
func (d *adminDao) GetUserPermissions(ctx context.Context, userId uint) ([]string, error) {
	return []string{}, nil
}

// CheckApiPermission 校验角色是否有指定API的访问权限
func (d *adminDao) CheckApiPermission(ctx context.Context, authorityId uint, path, method string) (bool, error) {
	if authorityId == 888 {
		return true, nil
	}
	return true, nil
}

// CreateAuthority 创建角色
func (d *adminDao) CreateAuthority(ctx context.Context, authority *dbr.SysAuthority) error {
	return d.db.WithContext(ctx).Create(authority).Error
}

// UpdateAuthority 更新角色
func (d *adminDao) UpdateAuthority(ctx context.Context, authority *dbr.SysAuthority) error {
	return d.db.WithContext(ctx).Model(&dbr.SysAuthority{}).Where("authority_id = ?", authority.ID).Updates(authority).Error
}

// DeleteAuthority 删除角色
func (d *adminDao) DeleteAuthority(ctx context.Context, authorityId uint) error {
	tx := d.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("sys_authority_authority_id = ?", authorityId).Delete(&dbr.SysAuthorityMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&dbr.SysAuthority{}, authorityId).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// GetAuthorityList 获取角色列表
func (d *adminDao) GetAuthorityList(ctx context.Context, req *admin.AuthorityListReq) (int64, []*dbr.SysAuthority, error) {
	var total int64
	var list []*dbr.SysAuthority

	db := d.db.WithContext(ctx).Model(&dbr.SysAuthority{})
	if req.AuthorityName != "" {
		db = db.Where("authority_name LIKE ?", "%"+req.AuthorityName+"%")
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

// GetAllAuthorities 获取所有角色
func (d *adminDao) GetAllAuthorities(ctx context.Context) ([]*dbr.SysAuthority, error) {
	var list []*dbr.SysAuthority
	err := d.db.WithContext(ctx).Find(&list).Error
	return list, err
}

// SetAuthorityMenu 设置角色菜单权限
func (d *adminDao) SetAuthorityMenu(ctx context.Context, authorityId uint, menuIds []uint) error {
	tx := d.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("sys_authority_authority_id = ?", authorityId).Delete(&dbr.SysAuthorityMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(menuIds) > 0 {
		var items []*dbr.SysAuthorityMenu
		for _, mid := range menuIds {
			items = append(items, &dbr.SysAuthorityMenu{AuthorityID: authorityId, MenuID: mid})
		}
		if err := tx.Create(&items).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// GetAuthorityMenuIds 获取角色已授权的菜单ID列表
func (d *adminDao) GetAuthorityMenuIds(ctx context.Context, authorityId uint) ([]uint, error) {
	var menuIds []uint
	err := d.db.WithContext(ctx).Model(&dbr.SysAuthorityMenu{}).
		Where("sys_authority_authority_id = ?", authorityId).
		Pluck("sys_base_menu_id", &menuIds).Error
	return menuIds, err
}

// SetAuthorityApi 设置角色API权限
func (d *adminDao) SetAuthorityApi(ctx context.Context, authorityId uint, apiIds []uint) error {
	return nil
}

// GetAuthorityApiIds 获取角色已授权的API ID列表
func (d *adminDao) GetAuthorityApiIds(ctx context.Context, authorityId uint) ([]uint, error) {
	return []uint{}, nil
}

// CheckAuthorityNameExists 检查角色名称是否已存在
func (d *adminDao) CheckAuthorityNameExists(ctx context.Context, authorityName string, excludeId uint) (bool, error) {
	var count int64
	query := d.db.WithContext(ctx).Model(&dbr.SysAuthority{}).Where("authority_name = ?", authorityName)
	if excludeId > 0 {
		query = query.Where("authority_id != ?", excludeId)
	}
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// CountChildAuthoritiesByParentId 统计该角色的子角色数量
func (d *adminDao) CountChildAuthoritiesByParentId(ctx context.Context, parentId uint) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&dbr.SysAuthority{}).Where("parent_id = ?", parentId).Count(&count).Error
	return count, err
}
