// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"
)

// SysDictionary 字典模型
type SysDictionary struct {
	ID        uint                    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time               `json:"createdAt"`
	UpdatedAt time.Time               `json:"updatedAt"`
	DeletedAt time.Time               `gorm:"index" json:"-"`
	Name      string                  `gorm:"column:name;index;comment:字典名称" json:"name"`
	Type      string                  `gorm:"column:type;index;comment:字典类型" json:"type"`
	Status    *bool                   `gorm:"column:status;default:1;comment:状态 1启用 2禁用" json:"status"`
	Desc      string                  `gorm:"column:desc;comment:描述" json:"desc"`
	Details   []SysDictionaryDetail `gorm:"-" json:"details,omitempty"`
}

// TableName 表名
func (SysDictionary) TableName() string {
	return "sys_dictionaries"
}
