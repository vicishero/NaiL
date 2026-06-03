// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"github.com/rocboss/paopao-ce/internal/dao/admin/dbr"
)

// GetAuthorityMenus 获取角色有权限的菜单列表
func (d *adminDao) GetAuthorityMenus(ctx context.Context, authorityId uint) ([]*dbr.SysBaseMenu, error) {
	var menus []*dbr.SysBaseMenu
	if authorityId == 888 {
		err := d.db.WithContext(ctx).Order("sort asc").Find(&menus).Error
		return menus, err
	}
	err := d.db.WithContext(ctx).Table("sys_base_menus").
		Joins("LEFT JOIN sys_authority_menus ON sys_base_menus.id = sys_authority_menus.sys_base_menu_id").
		Where("sys_authority_menus.sys_authority_authority_id = ?", authorityId).
		Order("sys_base_menus.sort asc").
		Find(&menus).Error
	return menus, err
}

// GetAuthorityMenusByIds 根据菜单ID列表获取菜单详情
func (d *adminDao) GetAuthorityMenusByIds(ctx context.Context, menuIds []uint) ([]*dbr.SysBaseMenu, error) {
	var menus []*dbr.SysBaseMenu
	err := d.db.WithContext(ctx).Where("id IN ?", menuIds).Find(&menus).Error
	return menus, err
}

// CheckMenuIdsValid 检查菜单ID列表是否都有效
func (d *adminDao) CheckMenuIdsValid(ctx context.Context, menuIds []uint) (bool, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&dbr.SysBaseMenu{}).Where("id IN ?", menuIds).Count(&count).Error
	if err != nil {
		return false, err
	}
	return int(count) == len(menuIds), nil
}

// GetAllMenus 获取所有菜单
func (d *adminDao) GetAllMenus(ctx context.Context) ([]*dbr.SysBaseMenu, error) {
	var menus []*dbr.SysBaseMenu
	err := d.db.WithContext(ctx).Order("sort asc").Find(&menus).Error
	return menus, err
}

// CreateBaseMenu 创建菜单
func (d *adminDao) CreateBaseMenu(ctx context.Context, menu *dbr.SysBaseMenu) error {
	return d.db.WithContext(ctx).Create(menu).Error
}

// UpdateBaseMenu 更新菜单
func (d *adminDao) UpdateBaseMenu(ctx context.Context, menu *dbr.SysBaseMenu) error {
	return d.db.WithContext(ctx).Model(&dbr.SysBaseMenu{}).Where("id = ?", menu.ID).Updates(menu).Error
}

// DeleteBaseMenu 删除菜单
func (d *adminDao) DeleteBaseMenu(ctx context.Context, menuId uint) error {
	tx := d.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Where("sys_base_menu_id = ?", menuId).Delete(&dbr.SysAuthorityMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&dbr.SysBaseMenu{}, menuId).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// GetBaseMenuByID 根据ID获取菜单
func (d *adminDao) GetBaseMenuByID(ctx context.Context, menuId uint) (*dbr.SysBaseMenu, error) {
	var menu dbr.SysBaseMenu
	err := d.db.WithContext(ctx).Where("id = ?", menuId).First(&menu).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// CountChildMenusByParentId 统计该菜单的子菜单数量
func (d *adminDao) CountChildMenusByParentId(ctx context.Context, parentId uint) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&dbr.SysBaseMenu{}).Where("parent_id = ?", parentId).Count(&count).Error
	return count, err
}

// CountAuthorityMenusByMenuId 统计使用该菜单的角色数量
func (d *adminDao) CountAuthorityMenusByMenuId(ctx context.Context, menuId uint) (int64, error) {
	var count int64
	err := d.db.WithContext(ctx).Model(&dbr.SysAuthorityMenu{}).Where("sys_base_menu_id = ?", menuId).Count(&count).Error
	return count, err
}
