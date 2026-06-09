// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model/joint"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/sirupsen/logrus"
)

type trendsSrv struct {
	api.UnimplementedTrendsServant
	*base.DaoServant
	ac                core.AppCache
	indexTrendsExpire int64
	prefixTrends      string
}

func (s *trendsSrv) Chain() gin.HandlersChain {
	// 不需要认证，未登录用户也能访问
	return nil
}

func (s *trendsSrv) GetIndexTrends(req *web.GetIndexTrendsReq) (res *web.GetIndexTrendsResp, _ error) {
	limit, offset := req.PageSize, (req.Page-1)*req.PageSize
	// 获取最近注册的用户（无需认证，公开接口）
	trends, totalRows, err := s.Ds.GetRecentUsers(limit, offset)
	if err != nil {
		logrus.Errorf("Ds.GetRecentUsers err: %s", err)
		return nil, web.ErrGetIndexTrendsFailed
	}
	resp := joint.PageRespFrom(trends, req.Page, req.PageSize, totalRows)
	return &web.GetIndexTrendsResp{
		CachePageResp: joint.CachePageResp{
			Data: resp,
		},
	}, nil
}

func (s *trendsSrv) trendsFromCache(req *web.GetIndexTrendsReq, limit int, offset int) (res *web.GetIndexTrendsResp, key string, ok bool) {
	key = fmt.Sprintf("%s%d:%d:%d", s.prefixTrends, req.Uid, limit, offset)
	if data, err := s.ac.Get(key); err == nil {
		ok, res = true, &web.GetIndexTrendsResp{
			CachePageResp: joint.CachePageResp{
				JsonResp: data,
			},
		}
	}
	return
}

func newTrendsSrv(s *base.DaoServant) api.Trends {
	cs := conf.CacheSetting
	return &trendsSrv{
		DaoServant:        s,
		ac:                _ac,
		indexTrendsExpire: cs.IndexTrendsExpire,
		prefixTrends:      conf.PrefixIdxTrends,
	}
}
