// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

type GetCaptchaResp struct {
	Id      string `json:"id"`
	Content string `json:"b64s"`
}

type SendCaptchaReq struct {
	Phone        string `json:"phone" form:"phone" binding:"required"`
	ImgCaptcha   string `json:"img_captcha" form:"img_captcha" binding:"required"`
	ImgCaptchaID string `json:"img_captcha_id" form:"img_captcha_id" binding:"required"`
}

type LoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RegisterResp struct {
	UserId   int64 `json:"id,string"`
	Username string `json:"username"`
}

// WalletNonceReq 获取签名nonce请求
type WalletNonceReq struct {
	Address string `json:"address" form:"address" binding:"required"`
}

// WalletNonceResp 获取签名nonce响应
type WalletNonceResp struct {
	Nonce   string `json:"nonce"`
	Message string `json:"message"`
}

// WalletLoginReq 钱包登录请求
type WalletLoginReq struct {
	Address   string `json:"address" form:"address" binding:"required"`
	Signature string `json:"signature" form:"signature" binding:"required"`
	Nonce     string `json:"nonce" form:"nonce" binding:"required"`
}

// WalletLoginResp 钱包登录响应
type WalletLoginResp struct {
	Token    string `json:"token"`
	IsNewUser bool   `json:"is_new_user"` // 是否是新注册用户
}
