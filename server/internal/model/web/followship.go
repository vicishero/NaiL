// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/vicishero/NaiL/server/internal/model/joint"
	"github.com/vicishero/NaiL/server/internal/servants/base"
)

type FollowUserReq struct {
	BaseInfo `json:"-" binding:"-"`
	UserId   int64 `json:"user_id,string" binding:"required"`
}

type UnfollowUserReq struct {
	BaseInfo `json:"-" binding:"-"`
	UserId   int64 `json:"user_id,string" binding:"required"`
}

type ListFollowsReq struct {
	BaseInfo `json:"-" binding:"-"`
	joint.BasePageInfo
	Username string `form:"username" binding:"required"`
}

type ListFollowsResp base.PageResp

type ListFollowingsReq struct {
	BaseInfo `form:"-" binding:"-"`
	joint.BasePageInfo
	Username string `form:"username" binding:"required"`
}

type ListFollowingsResp base.PageResp
