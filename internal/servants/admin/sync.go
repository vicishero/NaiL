// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/pkg/app"
)

// SyncSearchIndex 同步搜索索引（管理后台触发）
func (s *AuthServant) SyncSearchIndex(c *gin.Context) {
	s.daoServant.PushAllPostToSearch()
	app.NewResponse(c).ToResponse("search index sync started")
}
