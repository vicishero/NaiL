// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core/admin"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// GetSysMsgList 获取系统消息列表
func (s *AuthServant) GetSysMsgList(c *gin.Context) {
	var req admin.H5SysMsgListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	total, list, err := s.service.GetSysMsgList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	if list == nil { list = []admin.H5SysMsgItem{} }
	app.NewResponse(c).ToResponse(gin.H{"list": list, "total": total})
}

// CreateSysMsg 创建系统消息
func (s *AuthServant) CreateSysMsg(c *gin.Context) {
	var req admin.H5SysMsgCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.CreateSysMsg(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// DeleteSysMsg 删除系统消息
func (s *AuthServant) DeleteSysMsg(c *gin.Context) {
	var req admin.H5SysMsgDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteSysMsg(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}
