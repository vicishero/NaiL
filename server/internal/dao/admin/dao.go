// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"gorm.io/gorm"
	"github.com/vicishero/NaiL/server/internal/core/admin"
)

type adminDao struct {
	db *gorm.DB
}

func NewAdminDao(db *gorm.DB) admin.AdminDao {
	return &adminDao{db: db}
}

// DB 返回底层GORM DB(用于直接查询非sys表)
func (d *adminDao) DB() *gorm.DB {
	return d.db
}
