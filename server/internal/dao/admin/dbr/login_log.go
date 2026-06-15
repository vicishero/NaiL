// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"
)

// SysLoginLog 登录日志模型
type SysLoginLog struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	Username  string    `gorm:"column:username;comment:用户名" json:"username"`
	Status    *bool     `gorm:"column:status;comment:状态 1成功 2失败" json:"status"`
	Ip        string    `gorm:"column:ip;comment:登录IP" json:"ip"`
	IpAddr    string    `gorm:"column:ip_addr;comment:IP归属地" json:"ipAddr"`
	Browser   string    `gorm:"column:browser;comment:浏览器" json:"browser"`
	Os        string    `gorm:"column:os;comment:操作系统" json:"os"`
	Platform  string    `gorm:"column:platform;comment:平台" json:"platform"`
	LoginTime time.Time `gorm:"column:login_time;comment:登录时间" json:"loginTime"`
	UserID    uint      `gorm:"column:user_id;comment:用户ID" json:"userId"`
}

// TableName 表名
func (SysLoginLog) TableName() string {
	return "sys_login_logs"
}
