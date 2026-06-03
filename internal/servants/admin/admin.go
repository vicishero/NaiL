// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/m/v1"
)

// RouteWeb register Manager route
func RouteManager(e *gin.Engine) {
	api.RegisterUserServant(e, newUserSrv())
	// 注册GVA管理后台路由，使用/api/v1作为基础路径
	RegisterAdminRoutes(e, "/api/v1")
}
