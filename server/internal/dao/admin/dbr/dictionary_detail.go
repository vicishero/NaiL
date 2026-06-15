// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"
)

// SysDictionaryDetail 字典详情模型
type SysDictionaryDetail struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `gorm:"index" json:"-"`
	Label        string    `gorm:"column:label;comment:展示值" json:"label"`
	Value        string    `gorm:"column:value;comment:字典值" json:"value"`
	Status       *bool     `gorm:"column:status;default:1;comment:状态 1启用 2禁用" json:"status"`
	Sort         int       `gorm:"column:sort;comment:排序标记" json:"sort"`
	DictionaryID uint      `gorm:"column:dictionary_id;comment:关联标记" json:"dictionaryId"`
}

// TableName 表名
func (SysDictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}
