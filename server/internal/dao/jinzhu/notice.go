// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/vicishero/NaiL/server/internal/core"
	"github.com/vicishero/NaiL/server/internal/dao/jinzhu/dbr"
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
	where := "(n.receiver_user_id=? OR n.receiver_user_id=0) AND n.is_del=0"

	// Count
	if err = s.db.Table(_notice_+" n").Where(where, userId).Count(&total).Error; err != nil || total == 0 {
		return
	}

	// Find with LEFT JOIN p_notice_read to determine is_read
	type noticeWithRead struct {
		dbr.Notice
		ReadID *int64 `gorm:"column:read_id"`
	}
	var rows []noticeWithRead
	db := s.db.Table(_notice_+" n").
		Select("n.*, nr.id AS read_id").
		Joins("LEFT JOIN p_notice_read nr ON n.id = nr.msg_id AND nr.user_id = ?", userId).
		Where(where, userId)
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	if err = db.Order("n.id DESC").Find(&rows).Error; err != nil {
		return
	}

	for _, row := range rows {
		f := row.Notice.Format()
		if f != nil {
			// 有 read_id = 已读，无 = 未读
			if row.ReadID != nil {
				f.IsRead = 1
			} else {
				f.IsRead = 0
			}
			res = append(res, f)
		}
	}
	return
}

func (s *noticeSrv) GetUnreadNoticeCount(userId int64) (int64, error) {
	var count int64
	err := s.db.Table(_notice_+" n").
		Joins("LEFT JOIN p_notice_read nr ON n.id = nr.msg_id AND nr.user_id = ?", userId).
		Where("(n.receiver_user_id=? OR n.receiver_user_id=0) AND n.is_del=0 AND nr.id IS NULL", userId).
		Count(&count).Error
	return count, err
}

func (s *noticeSrv) ReadNotice(noticeId int64, userId int64) error {
	return s.db.Exec(
		"INSERT IGNORE INTO p_notice_read (msg_id, user_id, read_time) VALUES (?, ?, NOW())",
		noticeId, userId,
	).Error
}

func (s *noticeSrv) ReadAllNotice(userId int64) error {
	return s.db.Exec(
		`INSERT IGNORE INTO p_notice_read (msg_id, user_id, read_time)
		 SELECT n.id, ?, NOW()
		 FROM p_notice n
		 LEFT JOIN p_notice_read nr ON n.id = nr.msg_id AND nr.user_id = ?
		 WHERE (n.receiver_user_id=? OR n.receiver_user_id=0) AND n.is_del=0 AND nr.id IS NULL`,
		userId, userId, userId,
	).Error
}

func (s *noticeSrv) DeleteNotice(id int64) error {
	return s.db.Table(_notice_).Where("id=?", id).Updates(map[string]interface{}{
		"is_del": 1, "deleted_on": gorm.Expr("UNIX_TIMESTAMP()"),
	}).Error
}
