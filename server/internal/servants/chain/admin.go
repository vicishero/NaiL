// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"github.com/gin-gonic/gin"
	"github.com/vicishero/NaiL/server/internal/core/ms"
	"github.com/vicishero/NaiL/server/pkg/app"
	"github.com/vicishero/NaiL/server/pkg/xerror"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exist := c.Get("USER")
		if !exist {
			response := app.NewResponse(c)
			response.ToErrorResponse(xerror.Forbidden.WithDetails("需要管理员权限: USER不存在"))
			c.Abort()
			return
		}
		userModel, ok := user.(*ms.User)
		if !ok {
			response := app.NewResponse(c)
			response.ToErrorResponse(xerror.Forbidden.WithDetails("需要管理员权限: USER类型错误"))
			c.Abort()
			return
		}
		if userModel.Status != ms.UserStatusNormal {
			response := app.NewResponse(c)
			response.ToErrorResponse(xerror.Forbidden.WithDetails("需要管理员权限: 账户状态异常"))
			c.Abort()
			return
		}
		if !userModel.IsAdmin {
			response := app.NewResponse(c)
			response.ToErrorResponse(xerror.Forbidden.WithDetails("需要管理员权限: 非管理员账户"))
			c.Abort()
			return
		}
		c.Next()
	}
}
