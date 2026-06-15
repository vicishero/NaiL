// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"
)

// SysVersion 系统版本模型
type SysVersion struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `gorm:"index" json:"-"`

	Version     string    `gorm:"column:version;comment:版本号" json:"version"`
	Name        string    `gorm:"column:name;comment:版本名称" json:"name"`
	Content     string    `gorm:"column:content;type:text;comment:版本内容说明" json:"content"`
	Type        string    `gorm:"column:type;comment:版本类型 1正式版 2测试版" json:"type"`
	Status      *int      `gorm:"column:status;default:1;comment:状态 1启用 2禁用" json:"status"`
	UserID      uint      `gorm:"column:user_id;comment:发布人ID" json:"userId"`
	Username    string    `gorm:"column:username;comment:发布人" json:"username"`
	ReleaseTime time.Time `gorm:"column:release_time;comment:发布时间" json:"releaseTime"`
	FileUrl     string    `gorm:"column:file_url;comment:下载地址" json:"fileUrl"`
	FileSize    int64     `gorm:"column:file_size;comment:文件大小" json:"fileSize"`
	Md5         string    `gorm:"column:md5;comment:MD5校验值" json:"md5"`
	Remarks     string    `gorm:"column:remarks;comment:备注" json:"remarks"`
}

// TableName 表名
func (SysVersion) TableName() string {
	return "sys_versions"
}
