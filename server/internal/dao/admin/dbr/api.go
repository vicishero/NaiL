// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "time"

// SysApi API接口表模型
type SysApi struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `gorm:"index" json:"-"`
	Path        string    `gorm:"column:path" json:"path"`
	Description string    `gorm:"column:description" json:"description"`
	ApiGroup    string    `gorm:"column:api_group" json:"apiGroup"`
	Method      string    `gorm:"column:method" json:"method"`
}

func (SysApi) TableName() string { return "sys_apis" }
