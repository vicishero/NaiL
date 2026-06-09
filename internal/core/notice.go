// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import "github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"

// NoticeService 系统通知服务
type NoticeService interface {
	CreateNotice(notice *dbr.Notice) (*dbr.Notice, error)
	GetNotices(userId int64, offset, limit int) ([]*dbr.NoticeFormated, int64, error)
	GetUnreadNoticeCount(userId int64) (int64, error)
	ReadNotice(noticeId int64) error
	ReadAllNotice(userId int64) error
	DeleteNotice(id int64) error
}
