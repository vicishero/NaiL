// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

type noticeSrv struct {
	db *gorm.DB
}

func newNoticeService(db *gorm.DB) core.NoticeService {
	return &noticeSrv{db: db}
}

func (s *noticeSrv) CreateNotice(notice *dbr.Notice) (*dbr.Notice, error) {
	return notice.Create(s.db)
}

func (s *noticeSrv) GetNotices(userId int64, offset, limit int) (res []*dbr.NoticeFormated, total int64, err error) {
	var notices []*dbr.Notice
	// Count — 当前用户或全员广播
	where := "(receiver_user_id=? OR receiver_user_id=0) AND is_del=0"
	countDB := s.db.Table(_notice_).Where(where, userId)
	if err = countDB.Count(&total).Error; err != nil || total == 0 {
		return
	}
	// Find — independent chain
	findDB := s.db.Table(_notice_).Where(where, userId)
	if offset >= 0 && limit > 0 {
		findDB = findDB.Offset(offset).Limit(limit)
	}
	if err = findDB.Order("id DESC").Find(&notices).Error; err != nil {
		return
	}
	for _, n := range notices {
		res = append(res, n.Format())
	}
	return
}

func (s *noticeSrv) GetUnreadNoticeCount(userId int64) (int64, error) {
	return (&dbr.Notice{}).CountUnread(s.db, userId)
}

func (s *noticeSrv) ReadNotice(noticeId int64) error {
	return s.db.Table(_notice_).Where("id=?", noticeId).Update("is_read", 1).Error
}

func (s *noticeSrv) ReadAllNotice(userId int64) error {
	return (&dbr.Notice{}).ReadAll(s.db, userId)
}

func (s *noticeSrv) DeleteNotice(id int64) error {
	return s.db.Table(_notice_).Where("id=?", id).Updates(map[string]interface{}{
		"is_del": 1, "deleted_on": gorm.Expr("UNIX_TIMESTAMP()"),
	}).Error
}
