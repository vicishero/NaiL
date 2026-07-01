// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicishero/NaiL/server/internal/conf"
	adminDao "github.com/vicishero/NaiL/server/internal/dao/admin"
	adminSvc "github.com/vicishero/NaiL/server/internal/service/admin"
	"github.com/vicishero/NaiL/server/pkg/app"
	"github.com/vicishero/NaiL/server/pkg/xerror"
	"github.com/sirupsen/logrus"
)

const kolCacheKey = "explore:kol_categories"
const kolCacheTTL = 300 // 5分钟

func registerExploreRoute(e *gin.Engine) {
	db := conf.MustGormDB()
	aDao := adminDao.NewAdminDao(db)
	aService := adminSvc.NewAdminService(aDao)

	e.POST("/v1/explore/refreshCache", func(c *gin.Context) {
		RefreshKolCategories()
		app.NewResponse(c).ToResponse(gin.H{"msg": "缓存已刷新"})
	})

	e.GET("/v1/explore/kolCategories", func(c *gin.Context) {
		// 尝试从 Redis 缓存读取
		if _wc != nil {
			if cached, err := _wc.Get(kolCacheKey); err == nil && len(cached) > 0 {
				c.Data(200, "application/json; charset=utf-8", cached)
				return
			}
		}

		// 缓存未命中，查数据库
		resp, err := aService.GetExploreKolCategories(c.Request.Context())
		if err != nil {
			app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
			return
		}

		// 构建完整响应并写入缓存（异步）
		jsonResp, _ := json.Marshal(gin.H{"code": 0, "data": resp})
		go func() {
			if err := _wc.Set(kolCacheKey, jsonResp, kolCacheTTL); err != nil {
				logrus.Debugf("cache kol categories failed: %v", err)
			}
		}()

		app.NewResponse(c).ToResponse(resp)
	})
}

// RefreshKolCategories 供管理后台修改KOL后刷新缓存
func RefreshKolCategories() {
	if _wc != nil {
		_ = _wc.Delete(kolCacheKey)
		_ = _wc.DelAny(kolCacheKey + "*")
	}
}

// init 启动时预热缓存
func init() {
	go func() {
		time.Sleep(3 * time.Second) // 等服务完全启动
		RefreshKolCategories()      // 清空旧缓存，首次请求时重建
	}()
}
