// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	adminDao "github.com/rocboss/paopao-ce/internal/dao/admin"
	adminSvc "github.com/rocboss/paopao-ce/internal/service/admin"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

func registerExploreRoute(e *gin.Engine) {
	db := conf.MustGormDB()
	aDao := adminDao.NewAdminDao(db)
	aService := adminSvc.NewAdminService(aDao)

	e.GET("/v1/explore/kolCategories", func(c *gin.Context) {
		resp, err := aService.GetExploreKolCategories(c.Request.Context())
		if err != nil {
			app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
			return
		}
		app.NewResponse(c).ToResponse(resp)
	})
}
