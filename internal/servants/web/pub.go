// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/color"
	"image/png"
	"regexp"
	"unicode/utf8"

	"github.com/afocus/captcha"
	"github.com/gofrs/uuid/v5"
	api "github.com/rocboss/paopao-ce/auto/api/v1"
	"github.com/rocboss/paopao-ce/internal/core/ms"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/web/assets"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/rocboss/paopao-ce/pkg/version"
	"github.com/rocboss/paopao-ce/pkg/web3"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

const (
	_MaxLoginErrTimes = 10
	_MaxPhoneCaptcha  = 10
)

type pubSrv struct {
	api.UnimplementedPubServant
	*base.DaoServant
}

func (s *pubSrv) SendCaptcha(req *web.SendCaptchaReq) error {
	ctx := context.Background()

	// 验证图片验证码
	if captcha, err := s.Redis.GetImgCaptcha(ctx, req.ImgCaptchaID); err != nil || string(captcha) != req.ImgCaptcha {
		logrus.Debugf("get captcha err:%s expect:%s got:%s", err, captcha, req.ImgCaptcha)
		return web.ErrErrorCaptchaPassword
	}
	s.Redis.DelImgCaptcha(ctx, req.ImgCaptchaID)

	// 今日频次限制
	if count, _ := s.Redis.GetCountSmsCaptcha(ctx, req.Phone); count >= _MaxPhoneCaptcha {
		return web.ErrTooManyPhoneCaptchaSend
	}

	if err := s.Ds.SendPhoneCaptcha(req.Phone); err != nil {
		return xerror.ServerError
	}
	// 写入计数缓存
	s.Redis.IncrCountSmsCaptcha(ctx, req.Phone)

	return nil
}

func (s *pubSrv) GetCaptcha() (*web.GetCaptchaResp, error) {
	cap := captcha.New()
	if err := cap.AddFontFromBytes(assets.ComicBytes); err != nil {
		logrus.Errorf("cap.AddFontFromBytes err:%s", err)
		return nil, xerror.ServerError
	}
	cap.SetSize(160, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	cap.SetBkgColor(color.RGBA{218, 240, 228, 255})
	img, password := cap.Create(6, captcha.NUM)
	emptyBuff := bytes.NewBuffer(nil)
	if err := png.Encode(emptyBuff, img); err != nil {
		logrus.Errorf("png.Encode err:%s", err)
		return nil, xerror.ServerError
	}
	key := utils.EncodeMD5(uuid.Must(uuid.NewV4()).String())
	// 五分钟有效期
	s.Redis.SetImgCaptcha(context.Background(), key, password)
	return &web.GetCaptchaResp{
		Id:      key,
		Content: "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuff.Bytes()),
	}, nil
}

func (s *pubSrv) Register(req *web.RegisterReq) (*web.RegisterResp, error) {
	if _disallowUserRegister {
		return nil, web.ErrDisallowUserRegister
	}
	// 用户名检查
	if err := s.validUsername(req.Username); err != nil {
		return nil, err
	}
	// 密码检查
	if err := checkPassword(req.Password); err != nil {
		logrus.Errorf("scheckPassword err: %v", err)
		return nil, web.ErrUserRegisterFailed
	}
	password, salt := encryptPasswordAndSalt(req.Password)
	user := &ms.User{
		Nickname: req.Username,
		Username: req.Username,
		Password: password,
		Avatar:   getRandomAvatar(),
		Salt:     salt,
		Status:   ms.UserStatusNormal,
	}
	user, err := s.Ds.CreateUser(user)
	if err != nil {
		logrus.Errorf("Ds.CreateUser err: %s", err)
		return nil, web.ErrUserRegisterFailed
	}
	return &web.RegisterResp{
		UserId:   user.ID,
		Username: user.Username,
	}, nil
}

func (s *pubSrv) Login(req *web.LoginReq) (*web.LoginResp, error) {
	ctx := context.Background()
	user, err := s.Ds.GetUserByUsername(req.Username)
	if err != nil {
		logrus.Errorf("Ds.GetUserByUsername err:%s", err)
		return nil, xerror.UnauthorizedAuthNotExist
	}

	if user.Model != nil && user.ID > 0 {
		if count, err := s.Redis.GetCountLoginErr(ctx, user.ID); err == nil && count >= _MaxLoginErrTimes {
			return nil, web.ErrTooManyLoginError
		}
		// 对比密码是否正确
		if validPassword(user.Password, req.Password, user.Salt) {
			if user.Status == ms.UserStatusClosed {
				return nil, web.ErrUserHasBeenBanned
			}
			// 清空登录计数
			s.Redis.DelCountLoginErr(ctx, user.ID)
		} else {
			// 登录错误计数
			s.Redis.IncrCountLoginErr(ctx, user.ID)
			return nil, xerror.UnauthorizedAuthFailed
		}
	} else {
		return nil, xerror.UnauthorizedAuthNotExist
	}

	token, err := app.GenerateToken(user)
	if err != nil {
		logrus.Errorf("app.GenerateToken err: %v", err)
		return nil, xerror.UnauthorizedTokenGenerate
	}
	return &web.LoginResp{
		Token: token,
	}, nil
}

func (s *pubSrv) Version() (*web.VersionResp, error) {
	return &web.VersionResp{
		BuildInfo: version.ReadBuildInfo(),
	}, nil
}

func (s *pubSrv) WalletNonce(req *web.WalletNonceReq) (*web.WalletNonceResp, error) {
	ctx := context.Background()

	// 生成随机 nonce
	key := utils.EncodeMD5(uuid.Must(uuid.NewV4()).String())
	nonce := utils.EncodeMD5(key + req.Address)
	// 使用简单的消息格式，避免换行符问题
	message := fmt.Sprintf("PaoPao Login: %s", nonce)

	logrus.Debugf("WalletNonce - address=%s, nonce=%s, message=%s", req.Address, nonce, message)

	// 将 nonce 存入 Redis，5 分钟过期
	s.Redis.SetWalletNonce(ctx, req.Address, nonce)

	return &web.WalletNonceResp{
		Nonce:   nonce,
		Message: message,
	}, nil
}

func (s *pubSrv) WalletLogin(req *web.WalletLoginReq) (*web.WalletLoginResp, error) {
	ctx := context.Background()

	// 从 Redis 获取 nonce 并验证
	storedNonce, err := s.Redis.GetWalletNonce(ctx, req.Address)
	if err != nil || storedNonce != req.Nonce {
		logrus.Debugf("wallet nonce mismatch: address=%s expected=%s got=%s err=%v", req.Address, storedNonce, req.Nonce, err)
		return nil, web.ErrInvalidNonce
	}

	// 构造验证消息（必须与 WalletNonce 返回的消息完全一致）
	message := fmt.Sprintf("PaoPao Login: %s", req.Nonce)

	logrus.Debugf("WalletLogin - address=%s, nonce=%s, message=%s, sigLen=%d", req.Address, req.Nonce, message, len(req.Signature))

	// 验证签名
	valid, err := web3.VerifySignature(message, req.Signature, req.Address)
	if err != nil {
		logrus.Errorf("verify signature failed: %v", err)
		return nil, web.ErrInvalidSignature
	}
	if !valid {
		logrus.Debugf("invalid signature: address=%s", req.Address)
		return nil, web.ErrInvalidSignature
	}

	// 验证通过，删除 nonce
	s.Redis.DelWalletNonce(ctx, req.Address)

	// 查找用户是否存在
	user, err := s.Ds.GetUserByAddress(req.Address)
	isNewUser := false

	if err != nil || user == nil || user.Model == nil || user.ID <= 0 {
		// 用户不存在，自动创建新用户
		isNewUser = true
		// 使用钱包地址的后 10 位作为用户名
		username := "user_" + req.Address[len(req.Address)-10:]
		user = &ms.User{
			Username: username,
			Nickname: username,
			Password: "", // 钱包登录无需密码
			Salt:     "", // 钱包登录无需盐
			Avatar:   getRandomAvatar(),
			Status:   ms.UserStatusNormal,
			Address:  req.Address,
		}
		user, err = s.Ds.CreateUser(user)
		if err != nil {
			logrus.Errorf("create user from wallet failed: %v", err)
			return nil, web.ErrUserRegisterFailed
		}
	}

	// 检查用户状态
	if user.Status == ms.UserStatusClosed {
		return nil, web.ErrUserHasBeenBanned
	}

	// 生成 token
	token, err := app.GenerateToken(user)
	if err != nil {
		logrus.Errorf("generate token failed: %v", err)
		return nil, xerror.UnauthorizedTokenGenerate
	}

	return &web.WalletLoginResp{
		Token:     token,
		IsNewUser: isNewUser,
	}, nil
}

// validUsername 验证用户
func (s *pubSrv) validUsername(username string) error {
	// 检测用户是否合规
	if utf8.RuneCountInString(username) < 3 || utf8.RuneCountInString(username) > 12 {
		return web.ErrUsernameLengthLimit
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(username) {
		return web.ErrUsernameCharLimit
	}

	// 重复检查
	user, _ := s.Ds.GetUserByUsername(username)
	if user.Model != nil && user.ID > 0 {
		return web.ErrUsernameHasExisted
	}
	return nil
}

func newPubSrv(s *base.DaoServant) api.Pub {
	return &pubSrv{
		DaoServant: s,
	}
}
