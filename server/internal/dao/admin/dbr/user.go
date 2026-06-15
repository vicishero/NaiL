// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "time"

// SysUser 管理员用户表模型
type SysUser struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `gorm:"index" json:"-"`
	UUID        string    `gorm:"column:uuid" json:"uuid"`
	Username    string    `gorm:"column:username" json:"userName"`
	Password    string    `gorm:"column:password" json:"-"`
	NickName    string    `gorm:"column:nick_name" json:"nickName"`
	HeaderImg   string    `gorm:"column:header_img" json:"headerImg"`
	AuthorityID uint      `gorm:"column:authority_id" json:"authorityId"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	Email       string    `gorm:"column:email" json:"email"`
	Enable      int       `gorm:"column:enable" json:"enable"`
	MfaSecret   string    `gorm:"column:mfa_secret" json:"mfaSecret"`
	MfaBound    int       `gorm:"column:mfa_bound" json:"mfaBound"`
}

func (SysUser) TableName() string { return "sys_users" }
