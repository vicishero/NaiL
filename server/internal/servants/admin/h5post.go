// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/pkg/app"
	"github.com/vicishero/NaiL/server/pkg/xerror"
)

// GetH5PostList 获取贴文列表
func (s *AuthServant) GetH5PostList(c *gin.Context) {
	var req admin.H5PostListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	total, list, err := s.service.GetH5PostList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	if list == nil {
		list = []admin.H5PostItem{}
	}

	app.NewResponse(c).ToResponse(gin.H{"list": list, "total": total})
}

// GetH5Post 获取单个贴文
func (s *AuthServant) GetH5Post(c *gin.Context) {
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

	post, err := s.service.GetH5Post(c.Request.Context(), id)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(post)
}

// UpdateH5Post 更新贴文状态
func (s *AuthServant) UpdateH5Post(c *gin.Context) {
	var req admin.H5PostUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if req.ID <= 0 {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("ID无效"))
		return
	}

	if err := s.service.UpdateH5Post(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(gin.H{})
}

// DeleteH5Post 删除贴文
func (s *AuthServant) DeleteH5Post(c *gin.Context) {
	var req admin.H5PostDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if req.ID <= 0 {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("ID无效"))
		return
	}

	if err := s.service.DeleteH5Post(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(gin.H{})
}
