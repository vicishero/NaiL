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

// GetH5CommentList 获取评论列表
func (s *AuthServant) GetH5CommentList(c *gin.Context) {
	var req admin.H5CommentListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	total, list, err := s.service.GetH5CommentList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	if list == nil {
		list = []admin.H5CommentItem{}
	}

	app.NewResponse(c).ToResponse(gin.H{"list": list, "total": total})
}

// DeleteH5Comment 删除评论
func (s *AuthServant) DeleteH5Comment(c *gin.Context) {
	var req admin.H5CommentDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if req.ID <= 0 {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("ID无效"))
		return
	}

	if err := s.service.DeleteH5Comment(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(gin.H{})
}
