// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "time"

// SysOperationRecord 操作日志表模型
type SysOperationRecord struct {
	ID           uint      `gorm:"primarykey" json:"ID"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `gorm:"index" json:"-"`
	Ip           string    `gorm:"column:ip" json:"ip"`
	Method       string    `gorm:"column:method" json:"method"`
	Path         string    `gorm:"column:path" json:"path"`
	Status       int       `gorm:"column:status" json:"status"`
	Latency      int64     `gorm:"column:latency" json:"latency"`
	UserAgent    string    `gorm:"column:user_agent" json:"userAgent"`
	UserID       uint      `gorm:"column:user_id" json:"userId"`
	ErrorMessage string    `gorm:"column:error_message" json:"errorMessage"`
	Body         string    `gorm:"column:body" json:"body"`
	Resp         string    `gorm:"column:resp" json:"resp"`
}

func (SysOperationRecord) TableName() string { return "sys_operation_records" }
