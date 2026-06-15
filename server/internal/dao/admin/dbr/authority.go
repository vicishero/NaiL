// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "time"

// SysAuthority 角色表模型
type SysAuthority struct {
	ID            uint      `gorm:"primarykey;column:authority_id" json:"authorityId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `gorm:"index" json:"-"`
	ParentID      uint      `gorm:"column:parent_id" json:"parentId"`
	AuthorityName string    `gorm:"column:authority_name" json:"authorityName"`
	DefaultRouter string    `gorm:"column:default_router" json:"defaultRouter"`
	ShowStatus    int       `gorm:"column:show_status" json:"showStatus"`
	AuthorityType int       `gorm:"column:authority_type" json:"authorityType"`
}

func (SysAuthority) TableName() string { return "sys_authorities" }

// SysAuthorityMenu 角色菜单关联表模型
type SysAuthorityMenu struct {
	MenuID      uint `gorm:"primaryKey;column:sys_base_menu_id" json:"menuId"`
	AuthorityID uint `gorm:"primaryKey;column:sys_authority_authority_id" json:"authorityId"`
}

func (SysAuthorityMenu) TableName() string { return "sys_authority_menus" }
