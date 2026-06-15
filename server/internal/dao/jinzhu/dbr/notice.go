// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dbr

import "gorm.io/gorm"

type Notice struct {
	*Model
	SenderUserID   int64  `json:"sender_user_id"`
	ReceiverUserID int64  `json:"receiver_user_id"`
	Type           int8   `json:"type"`
	Brief          string `json:"brief"`
	Content        string `json:"content"`
	PostID         int64  `json:"post_id"`
	CommentID      int64  `json:"comment_id"`
	ReplyID        int64  `json:"reply_id"`
	IsRead         int8   `json:"is_read"`
}

type NoticeFormated struct {
	ID             int64  `json:"id,string"`
	SenderUserID   int64  `json:"sender_user_id,string"`
	ReceiverUserID int64  `json:"receiver_user_id,string"`
	Brief          string `json:"brief"`
	Content        string `json:"content"`
	IsRead         int8   `json:"is_read"`
	CreatedOn      int64  `json:"created_on"`
}

func (n *Notice) Format() *NoticeFormated {
	if n.Model == nil || n.Model.ID == 0 {
		return nil
	}
	return &NoticeFormated{
		ID:             n.ID,
		SenderUserID:   n.SenderUserID,
		ReceiverUserID: n.ReceiverUserID,
		Brief:          n.Brief,
		Content:        n.Content,
		IsRead:         n.IsRead,
		CreatedOn:      n.CreatedOn,
	}
}

func (n *Notice) Create(db *gorm.DB) (*Notice, error) {
	err := db.Create(&n).Error
	return n, err
}

func (n *Notice) List(db *gorm.DB, userId int64, offset, limit int) (res []*Notice, err error) {
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	err = db.Where("(receiver_user_id=? OR receiver_user_id=0) AND is_del=0", userId).Order("id DESC").Find(&res).Error
	return
}

func (n *Notice) CountUnread(db *gorm.DB, userId int64) (res int64, err error) {
	err = db.Model(n).Where("(receiver_user_id=? OR receiver_user_id=0) AND is_read=0 AND is_del=0", userId).Count(&res).Error
	return
}

func (n *Notice) ReadAll(db *gorm.DB, userId int64) error {
	return db.Table("p_notice").Where("(receiver_user_id=? OR receiver_user_id=0) AND is_del=0", userId).Update("is_read", 1).Error
}
