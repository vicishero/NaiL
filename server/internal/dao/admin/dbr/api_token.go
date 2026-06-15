// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"
)

// SysApiToken API Token 模型
type SysApiToken struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `gorm:"index" json:"-"`

	Token       string    `gorm:"column:token;comment:Token值" json:"token"`
	Name        string    `gorm:"column:name;comment:Token名称" json:"name"`
	UserID      uint      `gorm:"column:user_id;comment:创建用户ID" json:"userId"`
	Username    string    `gorm:"column:username;comment:创建用户名" json:"username"`
	Status      *int      `gorm:"column:status;default:1;comment:状态 1启用 2禁用" json:"status"`
	ExpireTime  time.Time `gorm:"column:expire_time;comment:过期时间" json:"expireTime"`
	Permissions string    `gorm:"column:permissions;type:text;comment:权限范围JSON" json:"permissions"`
	Remarks     string    `gorm:"column:remarks;comment:备注" json:"remarks"`
}

// TableName 表名
func (SysApiToken) TableName() string {
	return "sys_api_tokens"
}
