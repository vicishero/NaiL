// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"
)

// SysError 系统错误日志模型
type SysError struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `gorm:"index" json:"-"`

	Type        string    `gorm:"column:type;comment:错误类型" json:"type"`
	Message     string    `gorm:"column:message;type:text;comment:错误信息" json:"message"`
	Stack       string    `gorm:"column:stack;type:text;comment:堆栈信息" json:"stack"`
	Path        string    `gorm:"column:path;comment:请求路径" json:"path"`
	Method      string    `gorm:"column:method;comment:请求方法" json:"method"`
	UserID      uint      `gorm:"column:user_id;comment:用户ID" json:"userId"`
	Username    string    `gorm:"column:username;comment:用户名" json:"username"`
	IP          string    `gorm:"column:ip;comment:IP地址" json:"ip"`
	Status      int       `gorm:"column:status;default:1;comment:处理状态 1未处理 2已处理 3已忽略" json:"status"`
	Solution    string    `gorm:"column:solution;type:text;comment:解决方案" json:"solution"`
	HandlerID   uint      `gorm:"column:handler_id;comment:处理人ID" json:"handlerId"`
	HandlerName string    `gorm:"column:handler_name;comment:处理人名称" json:"handlerName"`
	HandleTime  time.Time `gorm:"column:handle_time;comment:处理时间" json:"handleTime"`
	IsPublic    bool      `gorm:"column:is_public;default:false;comment:是否公开可见" json:"isPublic"`
	Version     string    `gorm:"column:version;comment:系统版本" json:"version"`
}

// TableName 表名
func (SysError) TableName() string {
	return "sys_errors"
}
