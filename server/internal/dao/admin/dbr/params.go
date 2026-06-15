// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"
)

// SysParams 参数模型
type SysParams struct {
	ID        uint      `gorm:"primarykey" json:"ID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	Name      string    `gorm:"column:name;comment:参数名称" json:"name"`
	Key       string    `gorm:"column:key;comment:参数键" json:"key"`
	Value     string    `gorm:"column:value;comment:参数值" json:"value"`
	Type      string    `gorm:"column:type;comment:参数类型" json:"type"`
	Status    *bool     `gorm:"column:status;default:1;comment:状态 1启用 2禁用" json:"status"`
	Remark    string    `gorm:"column:remark;comment:备注" json:"remark"`
}

// TableName 表名
func (SysParams) TableName() string {
	return "sys_params"
}
