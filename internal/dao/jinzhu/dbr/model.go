// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"

	"github.com/rocboss/paopao-ce/pkg/snowflake"
)

// Model 公共Model
type Model struct {
	ID         int64                 `gorm:"primary_key" json:"id"`
	CreatedOn  int64                 `json:"created_on"`
	ModifiedOn int64                 `json:"modified_on"`
	DeletedOn  int64                 `json:"deleted_on"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag" json:"is_del"`
}

type (
	ConditionsT map[string]any
	Predicates  map[string][]any
)

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	nowTime := time.Now().Unix()
	// 雪花ID：应用层生成全局唯一int64主键，取代数据库自增
	// 实体嵌入 *Model 指针可能为nil，通过 Dest 直接注入 ID
	if m == nil || m.ID == 0 {
		tx.Statement.SetColumn("id", snowflake.Next())
	}
	tx.Statement.SetColumn("created_on", nowTime)
	tx.Statement.SetColumn("modified_on", nowTime)
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	if !tx.Statement.Changed("modified_on") {
		tx.Statement.SetColumn("modified_on", time.Now().Unix())
	}

	return
}
