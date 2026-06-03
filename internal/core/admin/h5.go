// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

// H5UserListReq 运维用户列表请求
type H5UserListReq struct {
	Page          int    `form:"page" json:"page"`
	PageSize      int    `form:"pageSize" json:"pageSize"`
	Nickname      string `form:"nickname" json:"nickname"`
	Username      string `form:"username" json:"username"`
	WalletAddress string `form:"walletAddress" json:"walletAddress"`
	Status        *int   `form:"status" json:"status"`
}

// H5UserListResp 运维用户列表响应
type H5UserListResp struct {
	List  []H5UserItem `json:"list"`
	Total int64        `json:"total"`
}

// H5UserItem 运维用户列表项
type H5UserItem struct {
	ID             int64  `json:"ID,string"`
	Nickname       string `json:"nickname"`
	Username       string `json:"username"`
	Phone          string `json:"phone"`
	WalletAddress  string `json:"walletAddress"`
	Bio            string `json:"bio"`
	Avatar         string `json:"avatar"`
	Status         int    `json:"status"`
	CreatedAt      string `json:"CreatedAt"`
	FollowingCount int64  `json:"followingCount"`
	FollowerCount  int64  `json:"followerCount"`
}

// H5UserGetReq 获取单个运维用户
type H5UserGetReq struct {
	ID int64 `form:"ID" json:"ID"`
}

// H5UserUpdateReq 更新运维用户
type H5UserUpdateReq struct {
	ID       int64  `json:"ID"`
	Nickname string `json:"nickname"`
	Bio      string `json:"bio"`
	Status   int    `json:"status"`
}

// H5UserDeleteReq 删除运维用户
type H5UserDeleteReq struct {
	ID int64 `json:"ID"`
}
