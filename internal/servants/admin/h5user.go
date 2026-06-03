// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core/admin"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// GetH5UserList 获取运维用户列表
func (s *AuthServant) GetH5UserList(c *gin.Context) {
	var req admin.H5UserListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	total, list, err := s.service.GetH5UserList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	if list == nil {
		list = []admin.H5UserItem{}
	}

	app.NewResponse(c).ToResponse(gin.H{"list": list, "total": total})
}

// GetH5User 获取单个运维用户
func (s *AuthServant) GetH5User(c *gin.Context) {
	idStr := c.Query("ID")
	if idStr == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("ID不能为空"))
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("ID格式错误"))
		return
	}

	user, err := s.service.GetH5User(c.Request.Context(), id)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(user)
}

// UpdateH5User 更新运维用户
func (s *AuthServant) UpdateH5User(c *gin.Context) {
	var req admin.H5UserUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.UpdateH5User(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(gin.H{})
}

// DeleteH5User 删除运维用户
func (s *AuthServant) DeleteH5User(c *gin.Context) {
	var req admin.H5UserDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.DeleteH5User(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(gin.H{})
}
