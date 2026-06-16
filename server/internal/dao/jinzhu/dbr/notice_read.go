// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "time"

// NoticeRead 系统消息阅读记录
type NoticeRead struct {
	ID       int64     `gorm:"primaryKey;autoIncrement" json:"id,string"`
	MsgID    int64     `gorm:"column:msg_id;not null;uniqueIndex:uk_msg_user" json:"msg_id,string"`
	UserID   int64     `gorm:"column:user_id;not null;uniqueIndex:uk_msg_user;index:idx_user" json:"user_id,string"`
	ReadTime time.Time `gorm:"column:read_time;not null" json:"read_time"`
}

// TableName p_notice_read
func (NoticeRead) TableName() string {
	return "p_notice_read"
}
