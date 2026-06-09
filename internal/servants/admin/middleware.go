// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// JWTMiddleware JWT认证中间件，解析token获取用户ID
// 支持两种token传递方式:
//   1. Authorization: Bearer <token>
//   2. x-token: <token>
func (s *AuthServant) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// 优先从 Authorization Bearer 获取
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			}
		}

		// 兼容GVA前端 x-token 头
		if tokenString == "" {
			tokenString = c.GetHeader("x-token")
		}

		if tokenString == "" {
			app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthNotExist)
			c.Abort()
			return
		}

		// 解析token
		claims := jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.JWTSetting.Secret), nil
		})

		if err != nil || !token.Valid {
			app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedTokenError)
			c.Abort()
			return
		}

		// 解析用户ID
		userId, err := strconv.Atoi(claims.Subject)
		if err != nil {
			app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedTokenError)
			c.Abort()
			return
		}

		// 将用户ID存入上下文
		c.Set("user_id", uint(userId))
		c.Next()
	}
}

// PermissionMiddleware RBAC权限校验中间件
func (s *AuthServant) PermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID
		userId, exists := c.Get("user_id")
		if !exists {
			app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthNotExist)
			c.Abort()
			return
		}

		// 获取请求路径和方法
		path := c.Request.URL.Path
		method := c.Request.Method

		// 校验权限
		hasPermission, err := s.service.CheckPermission(c.Request.Context(), userId.(uint), path, method)
		if err != nil {
			app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
			c.Abort()
			return
		}

		if !hasPermission {
			app.NewResponse(c).ToErrorResponse(xerror.Forbidden)
			c.Abort()
			return
		}

		c.Next()
	}
}
