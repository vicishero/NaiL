// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/vicishero/NaiL/server/internal/core"
	"github.com/vicishero/NaiL/server/internal/core/cs"
	"github.com/vicishero/NaiL/server/internal/core/ms"
	"github.com/vicishero/NaiL/server/internal/dao/jinzhu/dbr"
	"gorm.io/gorm"
)

type messageSrv struct {
	db *gorm.DB
}

func newMessageService(db *gorm.DB) core.MessageService {
	return &messageSrv{
		db: db,
	}
}

func (s *messageSrv) CreateMessage(msg *ms.Message) (*ms.Message, error) {
	return msg.Create(s.db)
}

func (s *messageSrv) GetUnreadCount(userID int64) (int64, error) {
	// p_message 未读
	msgCount, err := (&dbr.Message{}).CountUnread(s.db, userID)
	if err != nil {
		return 0, err
	}
	// p_notice 未读（通过 p_notice_read 判断）
	var noticeCount int64
	err = s.db.Table("p_notice n").
		Joins("LEFT JOIN p_notice_read nr ON n.id = nr.msg_id AND nr.user_id = ?", userID).
		Where("(n.receiver_user_id=? OR n.receiver_user_id=0) AND n.is_del=0 AND nr.id IS NULL", userID).
		Count(&noticeCount).Error
	if err != nil {
		return msgCount, nil // 容错：notice 查询失败时只返回 message 未读数
	}
	return msgCount + noticeCount, nil
}

func (s *messageSrv) GetMessageByID(id int64) (*ms.Message, error) {
	return (&dbr.Message{
		Model: &dbr.Model{
			ID: id,
		},
	}).Get(s.db)
}

func (s *messageSrv) ReadMessage(message *ms.Message) error {
	message.IsRead = 1
	return message.Update(s.db)
}

func (s *messageSrv) ReadAllMessage(userId int64) error {
	return s.db.Table(_message_).Where("(receiver_user_id=? OR receiver_user_id=0) AND is_del=0", userId).Update("is_read", 1).Error
}

func (s *messageSrv) GetMessages(userId int64, style cs.MessageStyle, limit int, offset int) (res []*ms.MessageFormated, total int64, err error) {
	var messages []*dbr.Message
	db := s.db.Table(_message_)
	// 1动态，2评论，3回复，4私信，5好友申请，99系统通知'
	switch style {
	case cs.StyleMsgSystem:
		// 系统消息已迁移到 p_notice，由 NoticeService 处理
		return nil, 0, nil
	case cs.StyleMsgWhisper:
		db = db.Where("(receiver_user_id=? OR sender_user_id=?) AND type=4", userId, userId)
	case cs.StyleMsgRequesting:
		db = db.Where("receiver_user_id=? AND type=5", userId)
	case cs.StyleMsgUnread:
		db = db.Where("(receiver_user_id=? OR receiver_user_id=0) AND is_read=0", userId)
	case cs.StyleMsgAll:
		fallthrough
	default:
		db = db.Where("receiver_user_id=? OR receiver_user_id=0 OR (sender_user_id=? AND type=4)", userId, userId)
	}
	if err = db.Count(&total).Error; err != nil || total == 0 {
		return
	}
	if offset >= 0 && limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	if err = db.Order("id DESC").Find(&messages).Error; err != nil {
		return
	}
	for _, message := range messages {
		res = append(res, message.Format())
	}
	return
}
