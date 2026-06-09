// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"bytes"
	"context"
	"encoding/base64"
	"image/color"
	"image/png"
	"strings"
	"time"

	"github.com/afocus/captcha"
	"github.com/gofrs/uuid/v5"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core/admin"
	"github.com/rocboss/paopao-ce/internal/dao/admin/dbr"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/servants/base"
	"github.com/rocboss/paopao-ce/internal/servants/web/assets"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/utils"
	"github.com/rocboss/paopao-ce/pkg/xerror"
	"github.com/sirupsen/logrus"
)

type AuthServant struct {
	service     admin.AdminService
	daoServant  *base.DaoServant
}

func NewAuthServant(service admin.AdminService, daoServant *base.DaoServant) *AuthServant {
	return &AuthServant{
		service:    service,
		daoServant: daoServant,
	}
}

// Login 管理员登录
// @Summary 管理员登录
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body admin.LoginReq true "登录参数"
// @Success 200 {object} app.DataResponse{data=admin.LoginResp} "登录成功"
// @Router /user/login [post]
func (s *AuthServant) Login(c *gin.Context) {
	var req admin.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	resp, err := s.service.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthFailed.WithDetails(err.Error()))
		return
	}

	// 检查MFA: 系统开关开启 且 用户已绑定MFA
	if s.service.IsMfaSystemEnabled() && resp.User.MfaBound == 1 {
		// 生成临时MFA验证token(5分钟有效)
		mfaClaims := jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    conf.JWTSetting.Issuer,
			Subject:   req.Username,
			ID:        "mfa",
		}
		mfaToken := jwt.NewWithClaims(jwt.SigningMethodHS256, mfaClaims)
		mfaTokenStr, err := mfaToken.SignedString([]byte(conf.JWTSetting.Secret))
		if err != nil {
			app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails("生成MFA token失败"))
			return
		}
		app.NewResponse(c).ToResponse(gin.H{
			"mfaRequired": true,
			"mfaToken":    mfaTokenStr,
			"username":    req.Username,
		})
		return
	}

	// 获取用户完整信息（含角色），构建GVA兼容响应
	userInfo, _ := s.service.GetUserInfo(c.Request.Context(), resp.User.ID)
	user := gin.H{
		"ID":        resp.User.ID,
		"uuid":      resp.User.UUID,
		"userName":  resp.User.Username,
		"nickName":  resp.User.NickName,
		"headerImg": resp.User.HeaderImg,
		"phone":     resp.User.Phone,
		"email":     resp.User.Email,
		"enable":    resp.User.Enable,
	}
	if userInfo != nil && len(userInfo.Roles) > 0 {
		user["authority"] = gin.H{
			"authorityId":   userInfo.Roles[0].ID,
			"authorityName": userInfo.Roles[0].AuthorityName,
			"defaultRouter": userInfo.Roles[0].DefaultRouter,
		}
	} else {
		user["authority"] = gin.H{"defaultRouter": "dashboard"}
	}

	app.NewResponse(c).ToResponse(gin.H{
		"user":      user,
		"token":     resp.Token,
		"expiresAt": resp.ExpiresAt,
	})
}

// Captcha 生成验证码
// @Summary 生成验证码
// @Tags 管理员
// @Accept json
// @Produce json
// @Success 200 {object} app.DataResponse{data=admin.CaptchaResp} "生成成功"
// @Router /base/captcha [post]
func (s *AuthServant) Captcha(c *gin.Context) {
	cap := captcha.New()
	if err := cap.AddFontFromBytes(assets.ComicBytes); err != nil {
		logrus.Errorf("cap.AddFontFromBytes err:%s", err)
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	cap.SetSize(160, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	cap.SetBkgColor(color.RGBA{218, 240, 228, 255})
	img, password := cap.Create(6, captcha.NUM)
	emptyBuff := bytes.NewBuffer(nil)
	if err := png.Encode(emptyBuff, img); err != nil {
		logrus.Errorf("png.Encode err:%s", err)
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	key := utils.EncodeMD5(uuid.Must(uuid.NewV4()).String())
	// 五分钟有效期
	redisCache := cache.NewRedisCache()
	_ = redisCache.SetImgCaptcha(context.Background(), key, password)
	app.NewResponse(c).ToResponse(&admin.CaptchaResp{
		CaptchaId: key,
		PicPath:   "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuff.Bytes()),
	})
}

// Logout 管理员退出登录
// @Summary 管理员退出登录
// @Tags 管理员
// @Accept json
// @Produce json
// @Success 200 {object} app.Response "退出成功"
// @Router /user/logout [post]
func (s *AuthServant) Logout(c *gin.Context) {
	userId, _ := c.Get("user_id")
	token := c.GetHeader("Authorization")
	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}

	err := s.service.Logout(c.Request.Context(), userId.(uint), token)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetUserInfo 获取当前登录用户信息
func (s *AuthServant) GetUserInfo(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthNotExist)
		return
	}

	resp, err := s.service.GetUserInfo(c.Request.Context(), userId.(uint))
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	// 返回GVA前端兼容格式: {userInfo: {...}}
	userInfo := gin.H{
		"ID":        resp.User.ID,
		"uuid":      resp.User.UUID,
		"userName":  resp.User.Username,
		"nickName":  resp.User.NickName,
		"headerImg": resp.User.HeaderImg,
		"phone":     resp.User.Phone,
		"email":     resp.User.Email,
		"enable":    resp.User.Enable,
		"authority": gin.H{
			"authorityId":   resp.Roles[0].ID,
			"authorityName": resp.Roles[0].AuthorityName,
			"defaultRouter": resp.Roles[0].DefaultRouter,
		},
	}
	app.NewResponse(c).ToResponse(gin.H{"userInfo": userInfo})
}

// GetMenuList 获取用户有权限的菜单列表
// @Summary 获取用户有权限的菜单列表
// @Tags 管理员
// @Accept json
// @Produce json
// @Success 200 {object} app.DataResponse{data=[]admin.MenuResp} "获取成功"
// @Router /menu/getMenuList [get]
func (s *AuthServant) GetMenuList(c *gin.Context) {
	// 从JWT中获取用户ID，然后获取角色ID
	userId, exists := c.Get("user_id")
	if !exists {
		app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthNotExist)
		return
	}

	userInfo, err := s.service.GetUserInfo(c.Request.Context(), userId.(uint))
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	// 根据角色ID获取菜单
	menus, err := s.service.GetUserMenu(c.Request.Context(), userInfo.User.AuthorityID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	if menus == nil {
		menus = []*admin.MenuResp{}
	}

	app.NewResponse(c).ToResponse(menus)
}

// ChangePassword 修改用户密码
// @Summary 修改用户密码
// @Tags 管理员
// @Accept json
// @Produce json
// @Param data body admin.ChangePwdReq true "修改密码参数"
// @Success 200 {object} app.Response "修改成功"
// @Router /user/changePassword [post]
func (s *AuthServant) ChangePassword(c *gin.Context) {
	var req admin.ChangePwdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		app.NewResponse(c).ToErrorResponse(xerror.UnauthorizedAuthNotExist)
		return
	}

	err := s.service.ChangePassword(c.Request.Context(), userId.(uint), req.OldPassword, req.NewPassword)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// operationLogMiddleware 操作日志中间件，自动记录管理员操作
func (s *AuthServant) OperationLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start).Milliseconds()

		// 获取用户ID
		userId, exists := c.Get("user_id")
		if !exists {
			return
		}

		// 记录操作日志
		log := &dbr.SysOperationRecord{
			Ip:         c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			Status:     c.Writer.Status(),
			Latency:    latency,
			UserAgent:  c.Request.UserAgent(),
			UserID:     userId.(uint),
			ErrorMessage: c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}

		// 异步记录日志
		go func() {
			_ = s.service.RecordOperationLog(c.Request.Context(), log)
		}()
	}
}

// GetUserList 获取用户列表
// @Summary 获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param username query string false "用户名搜索"
// @Param nickName query string false "昵称搜索"
// @Param enable query int false "启用状态搜索"
// @Success 200 {object} app.DataResponse{data=admin.UserListResp} "获取成功"
// @Router /user/user_list [get]
func (s *AuthServant) GetUserList(c *gin.Context) {
	var req admin.UserListReq
	// 优先绑定JSON body(GVA前端POST请求)，失败则回退到Query参数(GET请求)
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBindQuery(&req); err2 != nil {
			app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err2.Error()))
			return
		}
	}

	resp, err := s.service.GetUserList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// CreateUser 创建用户
// @Summary 创建用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body admin.CreateUserReq true "用户信息"
// @Success 200 {object} app.Response "创建成功"
// @Router /user/register [post]
func (s *AuthServant) CreateUser(c *gin.Context) {
	var req admin.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.CreateUser(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// UpdateUser 更新用户信息
// @Summary 更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body admin.UpdateUserReq true "用户信息"
// @Success 200 {object} app.Response "更新成功"
// @Router /user/set_user_info [post]
func (s *AuthServant) UpdateUser(c *gin.Context) {
	var req admin.UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.UpdateUser(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// SetUserAuthority 设置用户角色
// @Summary 设置用户角色
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body admin.SetUserAuthorityReq true "用户角色信息"
// @Success 200 {object} app.Response "设置成功"
// @Router /user/set_user_authority [post]
func (s *AuthServant) SetUserAuthority(c *gin.Context) {
	var req admin.SetUserAuthorityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.SetUserAuthority(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// ResetPassword 重置用户密码
// @Summary 重置用户密码
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body admin.ResetPasswordReq true "重置密码信息"
// @Success 200 {object} app.Response "重置成功"
// @Router /user/reset_password [post]
func (s *AuthServant) ResetPassword(c *gin.Context) {
	var req admin.ResetPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.ResetPassword(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param data body admin.DeleteUserReq true "删除用户信息"
// @Success 200 {object} app.Response "删除成功"
// @Router /user/delete_user [delete]
func (s *AuthServant) DeleteUser(c *gin.Context) {
	var req admin.DeleteUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.DeleteUser(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// CreateAuthority 创建角色
// @Summary 创建角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param data body admin.CreateAuthorityReq true "角色信息"
// @Success 200 {object} app.Response "创建成功"
// @Router /authority/createAuthority [post]
func (s *AuthServant) CreateAuthority(c *gin.Context) {
	var req admin.CreateAuthorityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.CreateAuthority(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// UpdateAuthority 更新角色
// @Summary 更新角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param data body admin.UpdateAuthorityReq true "角色信息"
// @Success 200 {object} app.Response "更新成功"
// @Router /authority/updateAuthority [post]
func (s *AuthServant) UpdateAuthority(c *gin.Context) {
	var req admin.UpdateAuthorityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.UpdateAuthority(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// DeleteAuthority 删除角色
// @Summary 删除角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param data body admin.DeleteAuthorityReq true "删除角色信息"
// @Success 200 {object} app.Response "删除成功"
// @Router /authority/deleteAuthority [post]
func (s *AuthServant) DeleteAuthority(c *gin.Context) {
	var req admin.DeleteAuthorityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.DeleteAuthority(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetAuthorityList 获取角色列表
// @Summary 获取角色列表
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param authorityName query string false "角色名称搜索"
// @Success 200 {object} app.DataResponse{data=admin.AuthorityListResp} "获取成功"
// @Router /authority/getAuthorityList [get]
func (s *AuthServant) GetAuthorityList(c *gin.Context) {
	var req admin.AuthorityListReq
	// 优先绑定JSON body(GVA前端POST请求)，失败则回退到Query参数(GET请求)
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBindQuery(&req); err2 != nil {
			app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err2.Error()))
			return
		}
	}

	resp, err := s.service.GetAuthorityList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	if resp.List == nil {
		resp.List = []*dbr.SysAuthority{}
	}
	app.NewResponse(c).ToResponse(resp.List)
}

// GetAllAuthorities 获取所有角色
// @Summary 获取所有角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Success 200 {object} app.DataResponse{data=[]dbr.SysAuthority} "获取成功"
// @Router /authority/getAllAuthorities [get]
func (s *AuthServant) GetAllAuthorities(c *gin.Context) {
	list, err := s.service.GetAllAuthorities(c.Request.Context())
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	if list == nil {
		list = []*dbr.SysAuthority{}
	}

	app.NewResponse(c).ToResponse(list)
}

// SetAuthorityMenu 设置角色菜单权限
// @Summary 设置角色菜单权限
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param data body admin.SetAuthorityMenuReq true "角色菜单权限信息"
// @Success 200 {object} app.Response "设置成功"
// @Router /authority/setAuthorityMenu [post]
func (s *AuthServant) SetAuthorityMenu(c *gin.Context) {
	var req admin.SetAuthorityMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.SetAuthorityMenu(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetAuthorityMenu 获取角色菜单权限
// @Summary 获取角色菜单权限
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param authorityId query uint true "角色ID"
// @Success 200 {object} app.DataResponse{data=admin.GetAuthorityMenuResp} "获取成功"
// @Router /authority/getAuthorityMenu [get]
func (s *AuthServant) GetAuthorityMenu(c *gin.Context) {
	var req admin.GetAuthorityMenuReq
	// 优先绑定JSON body(GVA前端POST请求)，失败则回退到Query参数(GET请求)
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBindQuery(&req); err2 != nil {
			app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err2.Error()))
			return
		}
	}

	resp, err := s.service.GetAuthorityMenu(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// SetAuthorityApi 设置角色API权限
// @Summary 设置角色API权限
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param data body admin.SetAuthorityApiReq true "角色API权限信息"
// @Success 200 {object} app.Response "设置成功"
// @Router /authority/setAuthorityApi [post]
func (s *AuthServant) SetAuthorityApi(c *gin.Context) {
	var req admin.SetAuthorityApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.SetAuthorityApi(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetAuthorityApi 获取角色API权限
// @Summary 获取角色API权限
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param authorityId query uint true "角色ID"
// @Success 200 {object} app.DataResponse{data=admin.GetAuthorityApiResp} "获取成功"
// @Router /authority/getAuthorityApi [get]
func (s *AuthServant) GetAuthorityApi(c *gin.Context) {
	var req admin.GetAuthorityApiReq
	// 优先绑定JSON body(GVA前端POST请求)，失败则回退到Query参数(GET请求)
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBindQuery(&req); err2 != nil {
			app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err2.Error()))
			return
		}
	}

	resp, err := s.service.GetAuthorityApi(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// CreateBaseMenu 创建菜单
// @Summary 创建菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param data body admin.CreateBaseMenuReq true "菜单信息"
// @Success 200 {object} app.Response "创建成功"
// @Router /menu/addBaseMenu [post]
func (s *AuthServant) CreateBaseMenu(c *gin.Context) {
	var req admin.CreateBaseMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.CreateBaseMenu(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// UpdateBaseMenu 更新菜单
// @Summary 更新菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param data body admin.UpdateBaseMenuReq true "菜单信息"
// @Success 200 {object} app.Response "更新成功"
// @Router /menu/updateBaseMenu [post]
func (s *AuthServant) UpdateBaseMenu(c *gin.Context) {
	var req admin.UpdateBaseMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.UpdateBaseMenu(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// DeleteBaseMenu 删除菜单
// @Summary 删除菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param data body admin.DeleteBaseMenuReq true "删除菜单信息"
// @Success 200 {object} app.Response "删除成功"
// @Router /menu/deleteBaseMenu [post]
func (s *AuthServant) DeleteBaseMenu(c *gin.Context) {
	var req admin.DeleteBaseMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.DeleteBaseMenu(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetBaseMenuTree 获取菜单树
// @Summary 获取菜单树
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Success 200 {object} app.DataResponse{data=admin.GetBaseMenuTreeResp} "获取成功"
// @Router /menu/getBaseMenuTree [get]
func (s *AuthServant) GetBaseMenuTree(c *gin.Context) {
	var req admin.GetBaseMenuTreeReq
	// 优先绑定JSON body(GVA前端POST请求)，失败则回退到Query参数(GET请求)
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBindQuery(&req); err2 != nil {
			app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err2.Error()))
			return
		}
	}

	resp, err := s.service.GetBaseMenuTree(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	menus := resp.Menus
	if menus == nil {
		menus = []*admin.MenuResp{}
	}

	app.NewResponse(c).ToResponse(menus)
}

// CreateApi 创建API
// @Summary 创建API
// @Tags API管理
// @Accept json
// @Produce json
// @Param data body admin.CreateApiReq true "API信息"
// @Success 200 {object} app.Response "创建成功"
// @Router /api/createApi [post]
func (s *AuthServant) CreateApi(c *gin.Context) {
	var req admin.CreateApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.CreateApi(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// UpdateApi 更新API
// @Summary 更新API
// @Tags API管理
// @Accept json
// @Produce json
// @Param data body admin.UpdateApiReq true "API信息"
// @Success 200 {object} app.Response "更新成功"
// @Router /api/updateApi [post]
func (s *AuthServant) UpdateApi(c *gin.Context) {
	var req admin.UpdateApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.UpdateApi(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// DeleteApi 删除API
// @Summary 删除API
// @Tags API管理
// @Accept json
// @Produce json
// @Param data body admin.DeleteApiReq true "删除API信息"
// @Success 200 {object} app.Response "删除成功"
// @Router /api/deleteApi [post]
func (s *AuthServant) DeleteApi(c *gin.Context) {
	var req admin.DeleteApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.DeleteApi(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetApiList 获取API列表
// @Summary 获取API列表
// @Tags API管理
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param path query string false "路径搜索"
// @Param apiGroup query string false "分组搜索"
// @Param method query string false "请求方法搜索"
// @Success 200 {object} app.DataResponse{data=admin.GetApiListResp} "获取成功"
// @Router /api/getApiList [get]
func (s *AuthServant) GetApiList(c *gin.Context) {
	var req admin.GetApiListReq
	// 优先绑定JSON body(GVA前端POST请求)，失败则回退到Query参数(GET请求)
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBindQuery(&req); err2 != nil {
			app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err2.Error()))
			return
		}
	}

	resp, err := s.service.GetApiList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// GetAllApiGroups 获取所有API分组
// @Summary 获取所有API分组
// @Tags API管理
// @Accept json
// @Produce json
// @Success 200 {object} app.DataResponse{data=admin.GetAllApiGroupsResp} "获取成功"
// @Router /api/getAllApiGroups [get]
func (s *AuthServant) GetAllApiGroups(c *gin.Context) {
	resp, err := s.service.GetAllApiGroups(c.Request.Context())
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// SyncApi 同步接口
// @Summary 同步接口
// @Tags API管理
// @Accept json
// @Produce json
// @Success 200 {object} app.Response "同步成功"
// @Router /api/syncApi [post]
func (s *AuthServant) SyncApi(c *gin.Context) {
	var req admin.SyncApiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.SyncApi(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetOperationLogList 获取操作日志列表
// @Summary 获取操作日志列表
// @Tags 操作日志
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param username query string false "用户名搜索"
// @Param path query string false "路径搜索"
// @Param method query string false "请求方法搜索"
// @Param status query int false "状态搜索"
// @Success 200 {object} app.DataResponse{data=admin.GetOperationLogListResp} "获取成功"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (s *AuthServant) GetOperationLogList(c *gin.Context) {
	var req admin.GetOperationLogListReq
	// 优先绑定JSON body(GVA前端POST请求)，失败则回退到Query参数(GET请求)
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBindQuery(&req); err2 != nil {
			app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err2.Error()))
			return
		}
	}

	resp, err := s.service.GetOperationLogList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// DeleteOperationLog 删除操作日志
// @Summary 删除操作日志
// @Tags 操作日志
// @Accept json
// @Produce json
// @Param data body admin.DeleteOperationLogReq true "删除操作日志信息"
// @Success 200 {object} app.Response "删除成功"
// @Router /sysOperationRecord/deleteSysOperationRecord [post]
func (s *AuthServant) DeleteOperationLog(c *gin.Context) {
	var req admin.DeleteOperationLogReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.DeleteOperationLog(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetSystemConfig 获取系统配置
// @Summary 获取系统配置
// @Tags 系统配置
// @Accept json
// @Produce json
// @Success 200 {object} app.DataResponse{data=admin.GetSystemConfigResp} "获取成功"
// @Router /system/getSystemConfig [get]
func (s *AuthServant) GetSystemConfig(c *gin.Context) {
	resp, err := s.service.GetSystemConfig(c.Request.Context())
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// SetSystemConfig 设置系统配置
// @Summary 设置系统配置
// @Tags 系统配置
// @Accept json
// @Produce json
// @Param data body admin.SetSystemConfigReq true "系统配置信息"
// @Success 200 {object} app.Response "设置成功"
// @Router /system/setSystemConfig [post]
func (s *AuthServant) SetSystemConfig(c *gin.Context) {
	var req admin.SetSystemConfigReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.SetSystemConfig(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}

// GetServerInfo 获取服务器信息
// @Summary 获取服务器信息
// @Tags 系统配置
// @Accept json
// @Produce json
// @Success 200 {object} app.DataResponse{data=admin.GetServerInfoResp} "获取成功"
// @Router /system/getServerInfo [get]
func (s *AuthServant) GetServerInfo(c *gin.Context) {
	resp, err := s.service.GetServerInfo(c.Request.Context())
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// UploadFile 上传文件
// @Summary 上传文件
// @Tags 文件管理
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "上传的文件"
// @Success 200 {object} app.DataResponse{data=admin.UploadFileResp} "上传成功"
// @Router /fileUploadAndDownload/upload [post]
func (s *AuthServant) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("获取文件失败: " + err.Error()))
		return
	}

	// 打开文件
	src, err := file.Open()
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails("打开文件失败: " + err.Error()))
		return
	}
	defer src.Close()

	resp, err := s.service.UploadFile(c.Request.Context(), src, file.Filename)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// GetFileList 获取文件列表
// @Summary 获取文件列表
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页数量"
// @Param name query string false "文件名搜索"
// @Param type query string false "文件类型搜索"
// @Success 200 {object} app.DataResponse{data=admin.GetFileListResp} "获取成功"
// @Router /fileUploadAndDownload/getFileList [get]
func (s *AuthServant) GetFileList(c *gin.Context) {
	var req admin.GetFileListReq
	// 优先绑定JSON body(GVA前端POST请求)，失败则回退到Query参数(GET请求)
	if err := c.ShouldBindJSON(&req); err != nil {
		if err2 := c.ShouldBindQuery(&req); err2 != nil {
			app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err2.Error()))
			return
		}
	}

	resp, err := s.service.GetFileList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(resp)
}

// DeleteFile 删除文件
// @Summary 删除文件
// @Tags 文件管理
// @Accept json
// @Produce json
// @Param data body admin.DeleteFileReq true "删除文件信息"
// @Success 200 {object} app.Response "删除成功"
// @Router /fileUploadAndDownload/deleteFile [post]
func (s *AuthServant) DeleteFile(c *gin.Context) {
	var req admin.DeleteFileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}

	if err := s.service.DeleteFile(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(nil)
}
