// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core/admin"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

// GetH5TagList 获取话题列表
func (s *AuthServant) GetH5TagList(c *gin.Context) {
	var req admin.H5TagListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	total, list, err := s.service.GetH5TagList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	if list == nil {
		list = []admin.H5TagItem{}
	}

	logrus.Infof("GetH5TagList HANDLER called: total=%d len(list)=%d", total, len(list))
	app.NewResponse(c).ToResponse(gin.H{"list": list, "total": total, "_handler": "GetH5TagList"})
}

// UpdateH5Tag 更新话题
func (s *AuthServant) UpdateH5Tag(c *gin.Context) {
	var req admin.H5TagUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if req.ID <= 0 {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("ID无效"))
		return
	}

	if err := s.service.UpdateH5Tag(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(gin.H{})
}

// DeleteH5Tag 删除话题
func (s *AuthServant) DeleteH5Tag(c *gin.Context) {
	var req admin.H5TagDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if req.ID <= 0 {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("ID无效"))
		return
	}

	if err := s.service.DeleteH5Tag(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(gin.H{})
}
