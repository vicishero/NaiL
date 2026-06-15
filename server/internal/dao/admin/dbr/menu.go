// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "time"

// SysBaseMenu 菜单表模型
type SysBaseMenu struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `gorm:"index" json:"-"`
	ParentID    uint      `gorm:"column:parent_id" json:"parentId"`
	Path        string    `gorm:"column:path" json:"path"`
	Name        string    `gorm:"column:name" json:"name"`
	Hidden      bool      `gorm:"column:hidden" json:"hidden"`
	Component   string    `gorm:"column:component" json:"component"`
	Sort        int       `gorm:"column:sort" json:"sort"`
	Icon        string    `gorm:"column:icon" json:"icon"`
	Title       string    `gorm:"column:title" json:"title"`
	Redirect    string    `gorm:"column:redirect" json:"redirect"`
	AlwaysShow  bool      `gorm:"column:always_show" json:"alwaysShow"`
	IsKeepAlive bool      `gorm:"column:keep_alive" json:"isKeepAlive"`
	IsAffix     bool      `gorm:"column:is_affix" json:"isAffix"`
	IsIframe    bool      `gorm:"column:is_iframe" json:"isIframe"`
	FrameSrc    string    `gorm:"column:frame_src" json:"frameSrc"`
}

func (SysBaseMenu) TableName() string { return "sys_base_menus" }
