// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/pkg/app"
	"github.com/vicishero/NaiL/server/pkg/xerror"
)

// MfaStatus 获取MFA绑定状态
func (s *AuthServant) MfaStatus(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthNotExist)
		return
	}

	resp, err := s.service.GetMfaStatus(c.Request.Context(), userId.(uint))
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// MfaBind 绑定MFA
func (s *AuthServant) MfaBind(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthNotExist)
		return
	}

	var req admin.MfaBindReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.BindMfa(c.Request.Context(), userId.(uint), req.Code); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// MfaUnbind 解绑MFA
func (s *AuthServant) MfaUnbind(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthNotExist)
		return
	}

	var req admin.MfaUnbindReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.UnbindMfa(c.Request.Context(), userId.(uint), req.Code); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// MfaVerify 登录第二步：MFA验证
func (s *AuthServant) MfaVerify(c *gin.Context) {
	var req admin.MfaVerifyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	resp, err := s.service.LoginMfa(c.Request.Context(), req.Username, req.Code, req.MfaToken)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthFailed.WithDetails(err.Error()))
		return
	}

	// 构建GVA兼容响应
	userInfo, _ := s.service.GetUserInfo(c.Request.Context(), resp.User.ID)
	user := gin.H{
		"ID":        resp.User.ID,
		"uuid":      resp.User.UUID,
		"userName":  resp.User.Username,
		"nickName":  resp.User.NickName,
		"headerImg": resp.User.HeaderImg,
		"phone":     resp.User.Phone,
		"email":     resp.User.Email,
		"enable":    resp.User.Enable,
	}
	if userInfo != nil && len(userInfo.Roles) > 0 {
		user["authority"] = gin.H{
			"authorityId":   userInfo.Roles[0].ID,
			"authorityName": userInfo.Roles[0].AuthorityName,
			"defaultRouter": userInfo.Roles[0].DefaultRouter,
		}
	} else {
		user["authority"] = gin.H{"defaultRouter": "dashboard"}
	}

	app.NewResponse(c).ToResponse(gin.H{
		"user":      user,
		"token":     resp.Token,
		"expiresAt": resp.ExpiresAt,
	})
}
