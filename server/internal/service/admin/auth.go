// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vicishero/NaiL/server/internal/conf"
	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type adminService struct {
	dao admin.AdminDao
	jwtSecret []byte
	jwtIssuer string
	jwtExpire time.Duration
}

func NewAdminService(dao admin.AdminDao) admin.AdminService {
	return &adminService{
		dao: dao,
		jwtSecret: []byte(conf.JWTSetting.Secret),
		jwtIssuer: conf.JWTSetting.Issuer,
		jwtExpire: conf.JWTSetting.Expire,
	}
}

// Login 管理员登录
func (s *adminService) Login(ctx context.Context, username, password string) (*admin.LoginResp, error) {
	// 查询用户信息
	user, err := s.dao.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("用户名或密码错误")
	}

	// 校验密码（GVA使用BCrypt加密）
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("用户名或密码错误")
	}

	// 生成JWT token
	expiresAt := time.Now().Add(s.jwtExpire)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    s.jwtIssuer,
		Subject:   strconv.Itoa(int(user.ID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %v", err)
	}

	return &admin.LoginResp{
		User:      user,
		Token:     tokenString,
		ExpiresAt: expiresAt.Unix(),
	}, nil
}

// Logout 管理员退出登录（暂时空实现，后续可以加token黑名单）
func (s *adminService) Logout(ctx context.Context, userId uint, token string) error {
	// TODO: 实现token加入黑名单逻辑
	return nil
}

// ====== MFA 多因素认证 ======

// IsMfaSystemEnabled 检查系统MFA开关是否开启
func (s *adminService) IsMfaSystemEnabled() bool {
	cfg, _ := s.GetSystemConfig(context.Background())
	if v, ok := cfg.Config["mfa_enabled"]; ok {
		switch val := v.(type) {
		case bool:
			return val
		case float64:
			return val == 1
		case string:
			return val == "true" || val == "1"
		}
	}
	return false
}

// generateMfaToken 生成临时MFA验证token(5分钟有效)
func (s *adminService) generateMfaToken(username string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    s.jwtIssuer,
		Subject:   username,
		ID:        "mfa",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

// verifyMfaToken 验证临时MFA token
func (s *adminService) verifyMfaToken(tokenString string) (string, error) {
	claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("MFA token无效或已过期")
	}
	if claims.ID != "mfa" {
		return "", fmt.Errorf("非MFA token")
	}
	return claims.Subject, nil
}

// LoginMfa MFA认证登录(第二步)
func (s *adminService) LoginMfa(ctx context.Context, username, code, mfaToken string) (*admin.LoginResp, error) {
	// 验证临时MFA token
	tokenUser, err := s.verifyMfaToken(mfaToken)
	if err != nil {
		return nil, err
	}
	if tokenUser != username {
		return nil, fmt.Errorf("MFA token与用户不匹配")
	}

	// 查询用户
	user, err := s.dao.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}

	// 验证MFA码
	if user.MfaSecret == "" || user.MfaBound == 0 {
		return nil, fmt.Errorf("用户未绑定MFA")
	}
	ok, err := VerifyTOTP(user.MfaSecret, code)
	if err != nil || !ok {
		return nil, fmt.Errorf("MFA验证码错误")
	}

	// 生成正式JWT token
	expiresAt := time.Now().Add(s.jwtExpire)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    s.jwtIssuer,
		Subject:   strconv.Itoa(int(user.ID)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %v", err)
	}

	return &admin.LoginResp{
		User:      user,
		Token:     tokenString,
		ExpiresAt: expiresAt.Unix(),
	}, nil
}

// GetMfaStatus 获取MFA状态
func (s *adminService) GetMfaStatus(ctx context.Context, userId uint) (*admin.MfaStatusResp, error) {
	user, err := s.dao.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	sysEnabled := s.IsMfaSystemEnabled()
	bound := user.MfaBound == 1 && user.MfaSecret != ""

	resp := &admin.MfaStatusResp{
		Bound:         bound,
		SystemEnabled: sysEnabled,
	}

	// 如果未绑定，生成新密钥供前端展示二维码
	if !bound {
		secret, err := GenerateMFASecret()
		if err != nil {
			return nil, fmt.Errorf("生成MFA密钥失败: %v", err)
		}
		resp.Secret = secret
		resp.QrURI = GetTOTPURI(secret, user.Username, "NaiL-Admin")
	}

	return resp, nil
}

// BindMfa 绑定MFA
func (s *adminService) BindMfa(ctx context.Context, userId uint, code string) error {
	// 获取当前用户的MFA状态（获取未绑定时生成的secret）
	status, err := s.GetMfaStatus(ctx, userId)
	if err != nil {
		return err
	}
	if status.Bound {
		return fmt.Errorf("MFA已绑定")
	}

	// 验证验证码
	ok, err := VerifyTOTP(status.Secret, code)
	if err != nil || !ok {
		return fmt.Errorf("验证码错误")
	}

	// 保存secret到用户记录
	user, err := s.dao.GetUserByID(ctx, userId)
	if err != nil {
		return err
	}
	user.MfaSecret = status.Secret
	user.MfaBound = 1
	return s.dao.UpdateUser(ctx, user)
}

// UnbindMfa 解绑MFA
func (s *adminService) UnbindMfa(ctx context.Context, userId uint, code string) error {
	user, err := s.dao.GetUserByID(ctx, userId)
	if err != nil {
		return err
	}
	if user.MfaSecret == "" || user.MfaBound == 0 {
		return fmt.Errorf("MFA未绑定")
	}

	// 验证当前验证码
	ok, err := VerifyTOTP(user.MfaSecret, code)
	if err != nil || !ok {
		return fmt.Errorf("验证码错误")
	}

	user.MfaSecret = ""
	user.MfaBound = 0
	return s.dao.UpdateUser(ctx, user)
}

// GetUserInfo 获取用户详细信息，包含权限和角色
func (s *adminService) GetUserInfo(ctx context.Context, userId uint) (*admin.UserInfoResp, error) {
	// 查询用户信息
	user, err := s.dao.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 查询角色信息
	role, err := s.dao.GetAuthorityByID(ctx, user.AuthorityID)
	if err != nil {
		return nil, err
	}

	// 查询权限列表
	permissions, err := s.dao.GetUserPermissions(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &admin.UserInfoResp{
		User:        user,
		Permissions: permissions,
		Roles:       []*dbr.SysAuthority{role},
	}, nil
}

// GetUserMenu 获取用户有权限的菜单树
func (s *adminService) GetUserMenu(ctx context.Context, authorityId uint) ([]*admin.MenuResp, error) {
	// 查询菜单列表
	menus, err := s.dao.GetAuthorityMenus(ctx, authorityId)
	if err != nil {
		return nil, err
	}

	// 构建菜单树
	return s.buildMenuTree(menus, 0), nil
}

// buildMenuTree 构建树形菜单结构
func (s *adminService) buildMenuTree(menus []*dbr.SysBaseMenu, parentID uint) []*admin.MenuResp {
	var tree []*admin.MenuResp
	for _, menu := range menus {
		if menu.ParentID == parentID {
			children := s.buildMenuTree(menus, menu.ID)
			tree = append(tree, &admin.MenuResp{
				ID:        menu.ID,
				ParentID:  menu.ParentID,
				Path:      menu.Path,
				Name:      menu.Name,
				Hidden:    menu.Hidden,
				Component: menu.Component,
				Sort:      menu.Sort,
				Meta: admin.MenuMeta{
					Title:     menu.Title,
					Icon:      menu.Icon,
					Hidden:    menu.Hidden,
					KeepAlive: menu.IsKeepAlive,
				},
				Btns:        []string{},
				Redirect:    menu.Redirect,
				AlwaysShow:  menu.AlwaysShow,
				IsKeepAlive: menu.IsKeepAlive,
				IsAffix:     menu.IsAffix,
				IsIframe:    menu.IsIframe,
				FrameSrc:    menu.FrameSrc,
				Children:    children,
			})
		}
	}
	return tree
}

// CheckPermission 校验用户是否有指定接口的访问权限
func (s *adminService) CheckPermission(ctx context.Context, userId uint, path, method string) (bool, error) {
	// 获取用户信息
	user, err := s.dao.GetUserByID(ctx, userId)
	if err != nil {
		return false, err
	}

	// 校验权限
	return s.dao.CheckApiPermission(ctx, user.AuthorityID, path, method)
}

// ChangePassword 修改用户密码
func (s *adminService) ChangePassword(ctx context.Context, userId uint, oldPwd, newPwd string) error {
	// 获取用户信息
	user, err := s.dao.GetUserByID(ctx, userId)
	if err != nil {
		return err
	}

	// 校验旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPwd))
	if err != nil {
		return fmt.Errorf("旧密码错误")
	}

	// 加密新密码
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(newPwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	user.Password = string(hashPwd)
	return s.dao.UpdateUser(ctx, user)
}

// RecordOperationLog 记录操作日志
func (s *adminService) RecordOperationLog(ctx context.Context, log *dbr.SysOperationRecord) error {
	return s.dao.CreateOperationLog(ctx, log)
}

// CreateUser 创建用户
func (s *adminService) CreateUser(ctx context.Context, req *admin.CreateUserReq) error {
	// 检查用户名是否已存在
	existingUser, err := s.dao.GetUserByUsername(ctx, req.Username)
	if err == nil && existingUser != nil {
		return fmt.Errorf("用户名已存在")
	}

	// 加密密码
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("密码加密失败: %v", err)
	}

	// 创建用户
	user := &dbr.SysUser{
		Username:    req.Username,
		Password:    string(hashPwd),
		NickName:    req.NickName,
		HeaderImg:   req.HeaderImg,
		AuthorityID: req.AuthorityID,
		Phone:       req.Phone,
		Email:       req.Email,
		Enable:      req.Enable,
		UUID:        fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s-%d", req.Username, time.Now().UnixNano())))),
	}

	return s.dao.CreateUser(ctx, user)
}

// UpdateUser 更新用户信息
func (s *adminService) UpdateUser(ctx context.Context, req *admin.UpdateUserReq) error {
	// 查询用户信息
	user, err := s.dao.GetUserByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("用户不存在: %v", err)
	}

	// 更新字段
	user.NickName = req.NickName
	user.HeaderImg = req.HeaderImg
	user.AuthorityID = req.AuthorityID
	user.Phone = req.Phone
	user.Email = req.Email
	user.Enable = req.Enable

	return s.dao.UpdateUser(ctx, user)
}

// SetUserAuthority 设置用户角色
func (s *adminService) SetUserAuthority(ctx context.Context, req *admin.SetUserAuthorityReq) error {
	// 检查角色是否存在
	_, err := s.dao.GetAuthorityByID(ctx, req.AuthorityID)
	if err != nil {
		return fmt.Errorf("角色不存在: %v", err)
	}

	// 查询用户信息
	user, err := s.dao.GetUserByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("用户不存在: %v", err)
	}

	// 更新角色
	user.AuthorityID = req.AuthorityID
	return s.dao.UpdateUser(ctx, user)
}

// ResetPassword 重置用户密码
func (s *adminService) ResetPassword(ctx context.Context, req *admin.ResetPasswordReq) error {
	// 查询用户信息
	user, err := s.dao.GetUserByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("用户不存在: %v", err)
	}

	// 加密新密码
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("密码加密失败: %v", err)
	}

	// 更新密码
	user.Password = string(hashPwd)
	return s.dao.UpdateUser(ctx, user)
}

// DeleteUser 删除用户
func (s *adminService) DeleteUser(ctx context.Context, userId uint) error {
	// 检查用户是否存在
	_, err := s.dao.GetUserByID(ctx, userId)
	if err != nil {
		return fmt.Errorf("用户不存在: %v", err)
	}

	return s.dao.DeleteUser(ctx, userId)
}

// GetUserList 获取用户列表
func (s *adminService) GetUserList(ctx context.Context, req *admin.UserListReq) (*admin.UserListResp, error) {
	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	total, list, err := s.dao.GetUserList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("查询用户列表失败: %v", err)
	}

	// 清空密码字段，不返回给前端
	for _, user := range list {
		user.Password = ""
	}

	return &admin.UserListResp{
		Total: total,
		List:  list,
	}, nil
}

// CreateAuthority 创建角色
func (s *adminService) CreateAuthority(ctx context.Context, req *admin.CreateAuthorityReq) error {
	// 检查角色名称是否已存在
	exists, err := s.dao.CheckAuthorityNameExists(ctx, req.AuthorityName, 0)
	if err != nil {
		return fmt.Errorf("检查角色名称失败: %v", err)
	}
	if exists {
		return fmt.Errorf("角色名称已存在")
	}

	// 如果有父角色，检查父角色是否存在
	if req.ParentID > 0 {
		_, err := s.dao.GetAuthorityByID(ctx, req.ParentID)
		if err != nil {
			return fmt.Errorf("父角色不存在: %v", err)
		}
	}

	// 创建角色
	authority := &dbr.SysAuthority{
		ParentID:      req.ParentID,
		AuthorityName: req.AuthorityName,
		DefaultRouter: req.DefaultRouter,
		ShowStatus:    req.ShowStatus,
		AuthorityType: req.AuthorityType,
	}

	return s.dao.CreateAuthority(ctx, authority)
}

// UpdateAuthority 更新角色
func (s *adminService) UpdateAuthority(ctx context.Context, req *admin.UpdateAuthorityReq) error {
	// 检查角色是否存在
	existingAuthority, err := s.dao.GetAuthorityByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("角色不存在: %v", err)
	}

	// 检查角色名称是否已被其他角色使用
	exists, err := s.dao.CheckAuthorityNameExists(ctx, req.AuthorityName, req.ID)
	if err != nil {
		return fmt.Errorf("检查角色名称失败: %v", err)
	}
	if exists {
		return fmt.Errorf("角色名称已存在")
	}

	// 如果有父角色，检查父角色是否存在且不能是自己或子角色
	if req.ParentID > 0 {
		if req.ParentID == req.ID {
			return fmt.Errorf("父角色不能是自己")
		}
		_, err := s.dao.GetAuthorityByID(ctx, req.ParentID)
		if err != nil {
			return fmt.Errorf("父角色不存在: %v", err)
		}
	}

	// 更新角色
	existingAuthority.ParentID = req.ParentID
	existingAuthority.AuthorityName = req.AuthorityName
	existingAuthority.DefaultRouter = req.DefaultRouter
	existingAuthority.ShowStatus = req.ShowStatus
	existingAuthority.AuthorityType = req.AuthorityType

	return s.dao.UpdateAuthority(ctx, existingAuthority)
}

// DeleteAuthority 删除角色
func (s *adminService) DeleteAuthority(ctx context.Context, authorityId uint) error {
	// 检查角色是否存在
	_, err := s.dao.GetAuthorityByID(ctx, authorityId)
	if err != nil {
		return fmt.Errorf("角色不存在: %v", err)
	}

	// 检查是否有用户使用该角色
	userCount, err := s.dao.CountUsersByAuthorityId(ctx, authorityId)
	if err != nil {
		return fmt.Errorf("检查角色使用情况失败: %v", err)
	}
	if userCount > 0 {
		return fmt.Errorf("该角色下还有用户，无法删除")
	}

	// 检查是否有子角色
	childCount, err := s.dao.CountChildAuthoritiesByParentId(ctx, authorityId)
	if err != nil {
		return fmt.Errorf("检查子角色失败: %v", err)
	}
	if childCount > 0 {
		return fmt.Errorf("该角色下还有子角色，无法删除")
	}

	return s.dao.DeleteAuthority(ctx, authorityId)
}

// GetAuthorityList 获取角色列表
func (s *adminService) GetAuthorityList(ctx context.Context, req *admin.AuthorityListReq) (*admin.AuthorityListResp, error) {
	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	total, list, err := s.dao.GetAuthorityList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("查询角色列表失败: %v", err)
	}

	return &admin.AuthorityListResp{
		Total: total,
		List:  list,
	}, nil
}

// GetAllAuthorities 获取所有角色
func (s *adminService) GetAllAuthorities(ctx context.Context) ([]*dbr.SysAuthority, error) {
	list, err := s.dao.GetAllAuthorities(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询所有角色失败: %v", err)
	}
	return list, nil
}

// SetAuthorityMenu 设置角色菜单权限
func (s *adminService) SetAuthorityMenu(ctx context.Context, req *admin.SetAuthorityMenuReq) error {
	// 检查角色是否存在
	_, err := s.dao.GetAuthorityByID(ctx, req.AuthorityID)
	if err != nil {
		return fmt.Errorf("角色不存在: %v", err)
	}

	// 检查所有菜单ID是否合法
	if len(req.MenuIds) > 0 {
		valid, err := s.dao.CheckMenuIdsValid(ctx, req.MenuIds)
		if err != nil {
			return fmt.Errorf("检查菜单ID失败: %v", err)
		}
		if !valid {
			return fmt.Errorf("存在无效的菜单ID")
		}
	}

	return s.dao.SetAuthorityMenu(ctx, req.AuthorityID, req.MenuIds)
}

// GetAuthorityMenu 获取角色菜单权限
func (s *adminService) GetAuthorityMenu(ctx context.Context, req *admin.GetAuthorityMenuReq) (*admin.GetAuthorityMenuResp, error) {
	// 检查角色是否存在
	_, err := s.dao.GetAuthorityByID(ctx, req.AuthorityID)
	if err != nil {
		return nil, fmt.Errorf("角色不存在: %v", err)
	}

	// 超级管理员拥有所有菜单
	if req.AuthorityID == 888 {
		allMenus, err := s.dao.GetAllMenus(ctx)
		if err != nil {
			return nil, fmt.Errorf("查询所有菜单失败: %v", err)
		}
		return &admin.GetAuthorityMenuResp{
			Menus: allMenus,
		}, nil
	}

	// 获取角色已授权的菜单ID
	menuIds, err := s.dao.GetAuthorityMenuIds(ctx, req.AuthorityID)
	if err != nil {
		return nil, fmt.Errorf("获取角色菜单ID失败: %v", err)
	}

	// 如果没有授权菜单，返回空列表
	if len(menuIds) == 0 {
		return &admin.GetAuthorityMenuResp{
			Menus: []*dbr.SysBaseMenu{},
		}, nil
	}

	// 获取菜单详情
	menus, err := s.dao.GetAuthorityMenusByIds(ctx, menuIds)
	if err != nil {
		return nil, fmt.Errorf("获取菜单详情失败: %v", err)
	}

	return &admin.GetAuthorityMenuResp{
		Menus: menus,
	}, nil
}

// SetAuthorityApi 设置角色API权限
func (s *adminService) SetAuthorityApi(ctx context.Context, req *admin.SetAuthorityApiReq) error {
	// 检查角色是否存在
	_, err := s.dao.GetAuthorityByID(ctx, req.AuthorityID)
	if err != nil {
		return fmt.Errorf("角色不存在: %v", err)
	}

	return s.dao.SetAuthorityApi(ctx, req.AuthorityID, req.ApiIds)
}

// GetAuthorityApi 获取角色API权限
func (s *adminService) GetAuthorityApi(ctx context.Context, req *admin.GetAuthorityApiReq) (*admin.GetAuthorityApiResp, error) {
	// 检查角色是否存在
	_, err := s.dao.GetAuthorityByID(ctx, req.AuthorityID)
	if err != nil {
		return nil, fmt.Errorf("角色不存在: %v", err)
	}

	// 获取角色已授权的API ID
	apiIds, err := s.dao.GetAuthorityApiIds(ctx, req.AuthorityID)
	if err != nil {
		return nil, fmt.Errorf("获取角色API ID失败: %v", err)
	}

	// 如果没有授权API，返回空列表
	if len(apiIds) == 0 {
		return &admin.GetAuthorityApiResp{
			Apis: []*dbr.SysApi{},
		}, nil
	}

	// 获取API详情
	apis, err := s.dao.GetAuthorityApisByIds(ctx, apiIds)
	if err != nil {
		return nil, fmt.Errorf("获取API详情失败: %v", err)
	}

	return &admin.GetAuthorityApiResp{
		Apis: apis,
	}, nil
}

// CreateBaseMenu 创建菜单
func (s *adminService) CreateBaseMenu(ctx context.Context, req *admin.CreateBaseMenuReq) error {
	// 如果有父菜单，检查父菜单是否存在
	if req.ParentID > 0 {
		_, err := s.dao.GetBaseMenuByID(ctx, req.ParentID)
		if err != nil {
			return fmt.Errorf("父菜单不存在: %v", err)
		}
	}

	// 创建菜单
	menu := &dbr.SysBaseMenu{
		ParentID:    req.ParentID,
		Path:        req.Path,
		Name:        req.Name,
		Hidden:      req.Hidden,
		Component:   req.Component,
		Sort:        req.Sort,
		Icon:        req.Icon,
		Title:       req.Title,
		Redirect:    req.Redirect,
		AlwaysShow:  req.AlwaysShow,
		IsKeepAlive: req.IsKeepAlive,
		IsAffix:     req.IsAffix,
		IsIframe:    req.IsIframe,
		FrameSrc:    req.FrameSrc,
	}

	return s.dao.CreateBaseMenu(ctx, menu)
}

// UpdateBaseMenu 更新菜单
func (s *adminService) UpdateBaseMenu(ctx context.Context, req *admin.UpdateBaseMenuReq) error {
	// 检查菜单是否存在
	existingMenu, err := s.dao.GetBaseMenuByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("菜单不存在: %v", err)
	}

	// 如果有父菜单，检查父菜单是否存在且不能是自己或子菜单
	if req.ParentID > 0 {
		if req.ParentID == req.ID {
			return fmt.Errorf("父菜单不能是自己")
		}
		_, err := s.dao.GetBaseMenuByID(ctx, req.ParentID)
		if err != nil {
			return fmt.Errorf("父菜单不存在: %v", err)
		}
	}

	// 更新菜单
	existingMenu.ParentID = req.ParentID
	existingMenu.Path = req.Path
	existingMenu.Name = req.Name
	existingMenu.Hidden = req.Hidden
	existingMenu.Component = req.Component
	existingMenu.Sort = req.Sort
	existingMenu.Icon = req.Icon
	existingMenu.Title = req.Title
	existingMenu.Redirect = req.Redirect
	existingMenu.AlwaysShow = req.AlwaysShow
	existingMenu.IsKeepAlive = req.IsKeepAlive
	existingMenu.IsAffix = req.IsAffix
	existingMenu.IsIframe = req.IsIframe
	existingMenu.FrameSrc = req.FrameSrc

	return s.dao.UpdateBaseMenu(ctx, existingMenu)
}

// DeleteBaseMenu 删除菜单
func (s *adminService) DeleteBaseMenu(ctx context.Context, menuId uint) error {
	// 检查菜单是否存在
	_, err := s.dao.GetBaseMenuByID(ctx, menuId)
	if err != nil {
		return fmt.Errorf("菜单不存在: %v", err)
	}

	// 检查是否有子菜单
	childCount, err := s.dao.CountChildMenusByParentId(ctx, menuId)
	if err != nil {
		return fmt.Errorf("检查子菜单失败: %v", err)
	}
	if childCount > 0 {
		return fmt.Errorf("该菜单下还有子菜单，无法删除")
	}

	// 检查是否有角色使用该菜单
	authorityCount, err := s.dao.CountAuthorityMenusByMenuId(ctx, menuId)
	if err != nil {
		return fmt.Errorf("检查菜单使用情况失败: %v", err)
	}
	if authorityCount > 0 {
		return fmt.Errorf("该菜单正在被角色使用，无法删除")
	}

	return s.dao.DeleteBaseMenu(ctx, menuId)
}

// GetBaseMenuTree 获取菜单树
func (s *adminService) GetBaseMenuTree(ctx context.Context, req *admin.GetBaseMenuTreeReq) (*admin.GetBaseMenuTreeResp, error) {
	// 获取所有菜单
	allMenus, err := s.dao.GetAllMenus(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询所有菜单失败: %v", err)
	}

	// 构建菜单树
	tree := s.buildMenuTree(allMenus, 0)

	return &admin.GetBaseMenuTreeResp{
		Menus: tree,
	}, nil
}

// CreateApi 创建API
func (s *adminService) CreateApi(ctx context.Context, req *admin.CreateApiReq) error {
	// 创建API
	api := &dbr.SysApi{
		Path:        req.Path,
		Description: req.Description,
		ApiGroup:    req.ApiGroup,
		Method:      req.Method,
	}

	return s.dao.CreateApi(ctx, api)
}

// UpdateApi 更新API
func (s *adminService) UpdateApi(ctx context.Context, req *admin.UpdateApiReq) error {
	// 检查API是否存在
	existingApi, err := s.dao.GetApiByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("API不存在: %v", err)
	}

	// 更新API
	existingApi.Path = req.Path
	existingApi.Description = req.Description
	existingApi.ApiGroup = req.ApiGroup
	existingApi.Method = req.Method

	return s.dao.UpdateApi(ctx, existingApi)
}

// DeleteApi 删除API
func (s *adminService) DeleteApi(ctx context.Context, apiId uint) error {
	// 检查API是否存在
	_, err := s.dao.GetApiByID(ctx, apiId)
	if err != nil {
		return fmt.Errorf("API不存在: %v", err)
	}

	return s.dao.DeleteApi(ctx, apiId)
}

// GetApiList 获取API列表
func (s *adminService) GetApiList(ctx context.Context, req *admin.GetApiListReq) (*admin.GetApiListResp, error) {
	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	total, list, err := s.dao.GetApiList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("查询API列表失败: %v", err)
	}

	return &admin.GetApiListResp{
		Total: total,
		List:  list,
	}, nil
}

// GetAllApiGroups 获取所有API分组
func (s *adminService) GetAllApiGroups(ctx context.Context) (*admin.GetAllApiGroupsResp, error) {
	groups, err := s.dao.GetAllApiGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("查询API分组失败: %v", err)
	}

	return &admin.GetAllApiGroupsResp{
		Groups: groups,
	}, nil
}

// SyncApi 同步接口
func (s *adminService) SyncApi(ctx context.Context, req *admin.SyncApiReq) error {
	// TODO: 实现同步接口逻辑，扫描所有路由并同步到数据库
	// 暂时返回成功
	return nil
}

// GetOperationLogList 获取操作日志列表
func (s *adminService) GetOperationLogList(ctx context.Context, req *admin.GetOperationLogListReq) (*admin.GetOperationLogListResp, error) {
	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	total, list, err := s.dao.GetOperationLogList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("查询操作日志列表失败: %v", err)
	}

	return &admin.GetOperationLogListResp{
		Total: total,
		List:  list,
	}, nil
}

// DeleteOperationLog 删除操作日志
func (s *adminService) DeleteOperationLog(ctx context.Context, logId uint) error {
	// 检查操作日志是否存在
	_, err := s.dao.GetOperationLogByID(ctx, logId)
	if err != nil {
		return fmt.Errorf("操作日志不存在: %v", err)
	}

	return s.dao.DeleteOperationLog(ctx, logId)
}

// GetSystemConfig 获取系统配置
func (s *adminService) GetSystemConfig(ctx context.Context) (*admin.GetSystemConfigResp, error) {
	// TODO: 实现从数据库或配置文件获取系统配置
	// 暂时返回示例配置
	return &admin.GetSystemConfigResp{
		Config: map[string]interface{}{
			"site_name":     "PaoPao社区",
			"site_desc":     "一个有趣的社区",
			"site_url":      "https://paopao.info",
			"site_logo":     "/static/logo.png",
			"site_icp":      "",
			"site_copyright": "© 2024 PaoPao",
			"register_enable": true,
			"email_verify":  false,
			"default_avatar": "/static/avatar.png",
			"max_upload_size": 10,
			"allowed_upload_types": []string{"jpg", "jpeg", "png", "gif", "mp4", "avi", "mov"},
			"mfa_enabled":   true,
			"system": map[string]interface{}{
				"use-redis": true,
				"cache-type": "redis",
			},
			"redis": map[string]interface{}{
				"db":       0,
				"addr":     "127.0.0.1:6379",
				"password": "",
			},
			"email": map[string]interface{}{
				"to":      "",
				"port":    587,
				"from":    "",
				"host":    "",
				"is-ssl":  true,
				"user":    "",
				"pass":    "",
			},
		},
	}, nil
}

// gvaSiteSetting site_settings 表行结构
type gvaSiteSetting struct {
	Key         string `gorm:"column:key;primaryKey"`
	Value       string `gorm:"column:value;type:text"`
	IsEncrypted bool   `gorm:"column:is_encrypted"`
	CreatedOn   int64  `gorm:"column:created_on"`
	ModifiedOn  int64  `gorm:"column:modified_on"`
	DeletedOn   int64  `gorm:"column:deleted_on"`
	IsDel       int8   `gorm:"column:is_del"`
}

func (gvaSiteSetting) TableName() string {
	return "site_settings"
}

// SetSystemConfig 设置系统配置（保存到数据库，实现"立即更新"功能）
func (s *adminService) SetSystemConfig(ctx context.Context, req *admin.SetSystemConfigReq) error {
	configJSON, err := json.Marshal(req.Config)
	if err != nil {
		return fmt.Errorf("序列化配置失败: %v", err)
	}

	db := s.dao.DB().WithContext(ctx)
	now := time.Now().Unix()
	var existing gvaSiteSetting

	if err := db.Where("`key` = ?", "gva_system_config").First(&existing).Error; err == nil {
		return db.Model(&existing).Updates(map[string]interface{}{
			"value":       string(configJSON),
			"modified_on": now,
		}).Error
	}

	return db.Create(&gvaSiteSetting{
		Key:         "gva_system_config",
		Value:       string(configJSON),
		IsEncrypted: false,
		CreatedOn:   now,
		ModifiedOn:  now,
	}).Error
}

// GetServerInfo 获取服务器信息（实时获取，实现"系统状态"监控）
func (s *adminService) GetServerInfo(ctx context.Context) (*admin.GetServerInfoResp, error) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	cpuUsage := float64(runtime.NumGoroutine()) / float64(runtime.NumCPU())
	memUsedMB := float64(memStats.Alloc) / 1024 / 1024
	hostname, _ := os.Hostname()

	return &admin.GetServerInfoResp{
		CPU:      cpuUsage,
		Memory:   memUsedMB,
		Disk:     0,
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		Go:       runtime.Version(),
		Hostname: hostname,
		NumCPU:   runtime.NumCPU(),
		Goroutines: runtime.NumGoroutine(),
	}, nil
}

// ReloadSystem 重载系统配置（从数据库重新读取配置，实现"重载服务"功能）
func (s *adminService) ReloadSystem(ctx context.Context) error {
	db := s.dao.DB().WithContext(ctx)
	var record gvaSiteSetting
	if err := db.Where("`key` = ?", "gva_system_config").First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return fmt.Errorf("读取系统配置失败: %v", err)
	}

	var configMap map[string]interface{}
	if err := json.Unmarshal([]byte(record.Value), &configMap); err != nil {
		return fmt.Errorf("解析系统配置失败: %v", err)
	}

	logrus.Infof("系统配置重载成功，共 %d 个配置项", len(configMap))
	return nil
}

// UploadFile 上传文件
func (s *adminService) UploadFile(ctx context.Context, file interface{}, filename string) (*admin.UploadFileResp, error) {
	// TODO: 实现真实的文件上传逻辑
	// 暂时返回示例数据
	ext := filepath.Ext(filename)
	fileType := "unknown"
	if ext != "" {
		ext = strings.ToLower(ext[1:])
		switch ext {
		case "jpg", "jpeg", "png", "gif", "bmp", "webp":
			fileType = "image"
		case "mp4", "avi", "mov", "wmv", "flv", "mkv":
			fileType = "video"
		case "mp3", "wav", "flac", "aac":
			fileType = "audio"
		case "pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "txt", "md":
			fileType = "document"
		case "zip", "rar", "7z", "tar", "gz":
			fileType = "archive"
		}
	}

	// 生成保存路径
	savePath := fmt.Sprintf("upload/%s/%d_%s", time.Now().Format("2006/01/02"), time.Now().UnixNano(), filename)
	fileURL := fmt.Sprintf("/%s", savePath)

	return &admin.UploadFileResp{
		URL:      fileURL,
		Name:     filename,
		Size:     0, // 暂时返回0
		Type:     fileType,
		FilePath: savePath,
	}, nil
}

// GetFileList 获取文件列表
func (s *adminService) GetFileList(ctx context.Context, req *admin.GetFileListReq) (*admin.GetFileListResp, error) {
	// TODO: 实现从数据库获取文件列表
	// 暂时返回示例数据
	var list []*admin.FileInfo
	for i := 0; i < 10; i++ {
		list = append(list, &admin.FileInfo{
			ID:         uint(i + 1),
			Name:       fmt.Sprintf("example_%d.jpg", i+1),
			Size:       1024 * 1024 * 2, // 2MB
			Type:       "image",
			URL:        fmt.Sprintf("/upload/2024/06/03/example_%d.jpg", i+1),
			FilePath:   fmt.Sprintf("upload/2024/06/03/example_%d.jpg", i+1),
			Uploader:   "admin",
			UploadTime: time.Now().Add(-time.Duration(i) * time.Hour),
		})
	}

	return &admin.GetFileListResp{
		Total: int64(len(list)),
		List:  list,
	}, nil
}

// DeleteFile 删除文件
func (s *adminService) DeleteFile(ctx context.Context, fileId uint) error {
	// TODO: 实现真实的文件删除逻辑，包括删除数据库记录和物理文件
	// 暂时返回成功
	return nil
}

// ====== H5运维用户管理 ======

// GetH5UserList 获取运维用户列表
func (s *adminService) GetH5UserList(ctx context.Context, req *admin.H5UserListReq) (int64, []admin.H5UserItem, error) {
	type pUser struct {
		ID        int64  `gorm:"column:id"`
		Nickname  string `gorm:"column:nickname"`
		Username  string `gorm:"column:username"`
		Phone     string `gorm:"column:phone"`
		Avatar    string `gorm:"column:avatar"`
		Status    int    `gorm:"column:status"`
		IsAdmin   bool   `gorm:"column:is_admin"`
		IsKOL     bool   `gorm:"column:is_kol"`
		ChatEnabled bool   `gorm:"column:chat_enabled"`
		CreatedOn int64  `gorm:"column:created_on"`
		Address   string `gorm:"column:address"`
	}

	var total int64
	var rows []pUser

	countDB := s.dao.DB().WithContext(ctx).Table("p_user").Where("is_del = 0")
	if req.Nickname != "" {
		countDB = countDB.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}
	if req.Username != "" {
		countDB = countDB.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.WalletAddress != "" {
		countDB = countDB.Where("address LIKE ?", "%"+req.WalletAddress+"%")
	}
	if req.Status != nil {
		countDB = countDB.Where("status = ?", *req.Status)
	}
	if err := countDB.Count(&total).Error; err != nil {
		logrus.WithError(err).Error("GetH5UserList count failed")
		return 0, nil, err
	}
	logrus.Debugf("GetH5UserList total=%d", total)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	offset := (req.Page - 1) * req.PageSize

	findDB := s.dao.DB().WithContext(ctx).Table("p_user").Where("is_del = 0")
	if req.Nickname != "" {
		findDB = findDB.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}
	if req.Username != "" {
		findDB = findDB.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.WalletAddress != "" {
		findDB = findDB.Where("address LIKE ?", "%"+req.WalletAddress+"%")
	}
	if req.Status != nil {
		findDB = findDB.Where("status = ?", *req.Status)
	}
	if err := findDB.Order("id DESC").Offset(offset).Limit(req.PageSize).Find(&rows).Error; err != nil {
		logrus.WithError(err).Error("GetH5UserList find failed")
		return 0, nil, err
	}
	logrus.Debugf("GetH5UserList rows=%d", len(rows))

	items := make([]admin.H5UserItem, len(rows))
	for i, r := range rows {
		items[i] = admin.H5UserItem{
			ID:            r.ID,
			Nickname:      r.Nickname,
			Username:      r.Username,
			Phone:         r.Phone,
			WalletAddress: r.Address,
			Bio:           "",
			Avatar:        r.Avatar,
			Status:        r.Status,
			IsAdmin:       r.IsAdmin,
			IsKOL:         r.IsKOL,
				ChatEnabled:    r.ChatEnabled,
			CreatedAt:     time.Unix(r.CreatedOn, 0).Format("2006-01-02 15:04:05"),
		}
	}

	return total, items, nil
}

// GetH5User 获取单个运维用户
func (s *adminService) GetH5User(ctx context.Context, userID int64) (*admin.H5UserItem, error) {
	type pUser struct {
		ID        int64  `gorm:"column:id"`
		Nickname  string `gorm:"column:nickname"`
		Username  string `gorm:"column:username"`
		Phone     string `gorm:"column:phone"`
		Avatar    string `gorm:"column:avatar"`
		Status    int    `gorm:"column:status"`
		IsAdmin   bool   `gorm:"column:is_admin"`
		IsKOL     bool   `gorm:"column:is_kol"`
		ChatEnabled bool   `gorm:"column:chat_enabled"`
		CreatedOn int64  `gorm:"column:created_on"`
		Address   string `gorm:"column:address"`
	}
	var row pUser
	err := s.dao.DB().WithContext(ctx).Table("p_user").Where("id = ? AND is_del = 0", userID).First(&row).Error
	if err != nil {
		return nil, err
	}
	return &admin.H5UserItem{
		ID:            row.ID,
		Nickname:      row.Nickname,
		Username:      row.Username,
		Phone:         row.Phone,
		WalletAddress: row.Address,
		Bio:           "",
		Avatar:        row.Avatar,
		Status:        row.Status,
		CreatedAt:     time.Unix(row.CreatedOn, 0).Format("2006-01-02 15:04:05"),
	}, nil
}

// UpdateH5User 更新运维用户
func (s *adminService) UpdateH5User(ctx context.Context, req *admin.H5UserUpdateReq) error {
	updates := map[string]interface{}{}
	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.IsAdmin != nil {
		updates["is_admin"] = *req.IsAdmin
	}
		if req.IsKOL != nil {
			updates["is_kol"] = *req.IsKOL
		}
		if req.ChatEnabled != nil {
			isKOL := *req.ChatEnabled
			if isKOL {
				var user struct{ IsKOL bool `gorm:"column:is_kol"` }
				if err := s.dao.DB().WithContext(ctx).Table("p_user").Where("id = ? AND is_del = 0", req.ID).First(&user).Error; err != nil {
					return err
				}
				if !user.IsKOL && (req.IsKOL == nil || !*req.IsKOL) {
					return fmt.Errorf("只有 KOL 用户才能启用聊天功能")
				}
			}
			updates["chat_enabled"] = *req.ChatEnabled
		}
		return s.dao.DB().WithContext(ctx).Table("p_user").Where("id = ? AND is_del = 0", req.ID).Updates(updates).Error
}

// DeleteH5User 删除运维用户(软删除)
func (s *adminService) DeleteH5User(ctx context.Context, userID int64) error {
	return s.dao.DB().WithContext(ctx).Table("p_user").Where("id = ?", userID).Updates(map[string]interface{}{
		"is_del":     1,
		"deleted_on": time.Now().Unix(),
	}).Error
}

// ====== H5运维贴文管理 ======

// GetH5PostList 获取贴文列表
func (s *adminService) GetH5PostList(ctx context.Context, req *admin.H5PostListReq) (int64, []admin.H5PostItem, error) {
	type pPost struct {
		ID             int64  `gorm:"column:id"`
		UserID         int64  `gorm:"column:user_id"`
		CommentCount   int64  `gorm:"column:comment_count"`
		CollectionCount int64 `gorm:"column:collection_count"`
		UpvoteCount    int64  `gorm:"column:upvote_count"`
		ShareCount     int64  `gorm:"column:share_count"`
		Visibility     int    `gorm:"column:visibility"`
		IsTop          int8   `gorm:"column:is_top"`
		IsEssence      int8   `gorm:"column:is_essence"`
		IsLock         int8   `gorm:"column:is_lock"`
		CreatedOn      int64  `gorm:"column:created_on"`
	}

	type pUser struct {
		ID       int64  `gorm:"column:id"`
		Nickname string `gorm:"column:nickname"`
		Username string `gorm:"column:username"`
		Avatar   string `gorm:"column:avatar"`
	}

	var total int64
	var rows []pPost

	// Build independent count query
	countDB := s.dao.DB().WithContext(ctx).Table("p_post").Where("is_del = 0")
	if req.UserID > 0 {
		countDB = countDB.Where("user_id = ?", req.UserID)
	}
	if req.Visibility != nil {
		countDB = countDB.Where("visibility = ?", *req.Visibility)
	}
	if err := countDB.Count(&total).Error; err != nil {
		logrus.WithError(err).Error("GetH5PostList count failed")
		return 0, nil, err
	}
	logrus.Debugf("GetH5PostList total=%d", total)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	offset := (req.Page - 1) * req.PageSize

	// Build independent find query (no shared state with count)
	findDB := s.dao.DB().WithContext(ctx).Table("p_post").Where("is_del = 0")
	if req.UserID > 0 {
		findDB = findDB.Where("user_id = ?", req.UserID)
	}
	if req.Visibility != nil {
		findDB = findDB.Where("visibility = ?", *req.Visibility)
	}
	if err := findDB.Order("id DESC").Offset(offset).Limit(req.PageSize).Find(&rows).Error; err != nil {
		logrus.WithError(err).Error("GetH5PostList find failed")
		return 0, nil, err
	}
	logrus.Debugf("GetH5PostList rows=%d", len(rows))

	// 获取所有用户ID
	userIds := make([]int64, len(rows))
	for i, r := range rows {
		userIds[i] = r.UserID
	}

	// 批量查询用户信息
	userMap := make(map[int64]pUser)
	if len(userIds) > 0 {
		var users []pUser
		s.dao.DB().WithContext(ctx).Table("p_user").Where("id IN ?", userIds).Find(&users)
		for _, u := range users {
			userMap[u.ID] = u
		}
	}

	// 获取所有贴文ID
	postIds := make([]int64, len(rows))
	for i, r := range rows {
		postIds[i] = r.ID
	}

	// 批量查询贴文内容
	contentMap := make(map[int64][]admin.H5PostContent)
	if len(postIds) > 0 {
		type pContent struct {
			PostID  int64  `gorm:"column:post_id"`
			Type    int    `gorm:"column:type"`
			Content string `gorm:"column:content"`
			Sort    int    `gorm:"column:sort"`
		}
		var contents []pContent
		s.dao.DB().WithContext(ctx).Table("p_post_content").Where("post_id IN ?", postIds).Order("sort ASC").Find(&contents)
		for _, c := range contents {
			contentMap[c.PostID] = append(contentMap[c.PostID], admin.H5PostContent{
				Type:    c.Type,
				Content: c.Content,
				Sort:    c.Sort,
			})
		}
	}

	items := make([]admin.H5PostItem, len(rows))
	for i, r := range rows {
		user := userMap[r.UserID]
		items[i] = admin.H5PostItem{
			ID:              r.ID,
			UserID:          r.UserID,
			User: &admin.H5UserItem{
				ID:       user.ID,
				Nickname: user.Nickname,
				Username: user.Username,
				Avatar:   user.Avatar,
			},
			Contents:         contentMap[r.ID],
			CommentCount:     r.CommentCount,
			CollectionCount:  r.CollectionCount,
			UpvoteCount:      r.UpvoteCount,
			ShareCount:       r.ShareCount,
			Visibility:       r.Visibility,
			IsTop:            r.IsTop,
			IsEssence:        r.IsEssence,
			IsLock:           r.IsLock,
			CreatedAt:        time.Unix(r.CreatedOn, 0).Format("2006-01-02 15:04:05"),
		}
	}

	return total, items, nil
}

// GetH5Post 获取单个贴文
func (s *adminService) GetH5Post(ctx context.Context, postID int64) (*admin.H5PostItem, error) {
	type pPost struct {
		ID             int64  `gorm:"column:id"`
		UserID         int64  `gorm:"column:user_id"`
		CommentCount   int64  `gorm:"column:comment_count"`
		CollectionCount int64 `gorm:"column:collection_count"`
		UpvoteCount    int64  `gorm:"column:upvote_count"`
		ShareCount     int64  `gorm:"column:share_count"`
		Visibility     int    `gorm:"column:visibility"`
		IsTop          int8   `gorm:"column:is_top"`
		IsEssence      int8   `gorm:"column:is_essence"`
		IsLock         int8   `gorm:"column:is_lock"`
		CreatedOn      int64  `gorm:"column:created_on"`
	}

	var row pPost
	err := s.dao.DB().WithContext(ctx).Table("p_post").Where("id = ? AND is_del = 0", postID).First(&row).Error
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	type pUser struct {
		ID       int64  `gorm:"column:id"`
		Nickname string `gorm:"column:nickname"`
		Username string `gorm:"column:username"`
		Avatar   string `gorm:"column:avatar"`
	}
	var user pUser
	s.dao.DB().WithContext(ctx).Table("p_user").Where("id = ?", row.UserID).First(&user)

	// 获取贴文内容
	type pContent struct {
		Type    int    `gorm:"column:type"`
		Content string `gorm:"column:content"`
		Sort    int    `gorm:"column:sort"`
	}
	var contents []pContent
	s.dao.DB().WithContext(ctx).Table("p_post_content").Where("post_id = ?", postID).Order("sort ASC").Find(&contents)

	contentItems := make([]admin.H5PostContent, len(contents))
	for i, c := range contents {
		contentItems[i] = admin.H5PostContent{
			Type:    c.Type,
			Content: c.Content,
			Sort:    c.Sort,
		}
	}

	return &admin.H5PostItem{
		ID:     row.ID,
		UserID: row.UserID,
		User: &admin.H5UserItem{
			ID:       user.ID,
			Nickname: user.Nickname,
			Username: user.Username,
			Avatar:   user.Avatar,
		},
		Contents:         contentItems,
		CommentCount:     row.CommentCount,
		CollectionCount:  row.CollectionCount,
		UpvoteCount:      row.UpvoteCount,
		ShareCount:       row.ShareCount,
		Visibility:       row.Visibility,
		IsTop:            row.IsTop,
		IsEssence:        row.IsEssence,
		IsLock:           row.IsLock,
		CreatedAt:        time.Unix(row.CreatedOn, 0).Format("2006-01-02 15:04:05"),
	}, nil
}

// UpdateH5Post 更新贴文状态
func (s *adminService) UpdateH5Post(ctx context.Context, req *admin.H5PostUpdateReq) error {
	updates := map[string]interface{}{
		"visibility": req.Visibility,
		"is_top":     req.IsTop,
		"is_essence": req.IsEssence,
		"is_lock":    req.IsLock,
	}
	return s.dao.DB().WithContext(ctx).Table("p_post").Where("id = ? AND is_del = 0", req.ID).Updates(updates).Error
}

// DeleteH5Post 删除贴文(软删除)
func (s *adminService) DeleteH5Post(ctx context.Context, postID int64) error {
	return s.dao.DB().WithContext(ctx).Table("p_post").Where("id = ?", postID).Updates(map[string]interface{}{
		"is_del":     1,
		"deleted_on": time.Now().Unix(),
	}).Error
}

// ====== H5运维话题管理 ======

// GetH5TagList 获取话题列表
func (s *adminService) GetH5TagList(ctx context.Context, req *admin.H5TagListReq) (int64, []admin.H5TagItem, error) {
	type pTag struct {
		ID        int64  `gorm:"column:id"`
		Tag       string `gorm:"column:tag"`
		QuoteNum  int64  `gorm:"column:quote_num"`
		UserID    int64  `gorm:"column:user_id"`
		CreatedOn int64  `gorm:"column:created_on"`
	}

	var total int64
	var rows []pTag

	countDB := s.dao.DB().WithContext(ctx).Table("p_tag").Where("is_del = 0")
	if err := countDB.Count(&total).Error; err != nil {
		logrus.WithError(err).Error("GetH5TagList count failed")
		return 0, nil, err
	}
	logrus.Debugf("GetH5TagList total=%d", total)

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	offset := (req.Page - 1) * req.PageSize

	findDB := s.dao.DB().WithContext(ctx).Table("p_tag").Where("is_del = 0")
	if err := findDB.Order("id DESC").Offset(offset).Limit(req.PageSize).Find(&rows).Error; err != nil {
		logrus.WithError(err).Error("GetH5TagList find failed")
		return 0, nil, err
	}
	logrus.Debugf("GetH5TagList rows=%d", len(rows))

	items := make([]admin.H5TagItem, len(rows))
	for i, r := range rows {
		items[i] = admin.H5TagItem{
			ID:        r.ID,
			Tag:       r.Tag,
			QuoteNum:  r.QuoteNum,
			UserID:    r.UserID,
			CreatedAt: time.Unix(r.CreatedOn, 0).Format("2006-01-02 15:04:05"),
		}
	}

	return total, items, nil
}

// UpdateH5Tag 更新话题
func (s *adminService) UpdateH5Tag(ctx context.Context, req *admin.H5TagUpdateReq) error {
	updates := map[string]interface{}{}
	if req.Tag != "" {
		updates["tag"] = req.Tag
	}
	updates["quote_num"] = req.QuoteNum
	return s.dao.DB().WithContext(ctx).Table("p_tag").Where("id = ? AND is_del = 0", req.ID).Updates(updates).Error
}

// DeleteH5Tag 删除话题(软删除)
func (s *adminService) DeleteH5Tag(ctx context.Context, tagID int64) error {
	return s.dao.DB().WithContext(ctx).Table("p_tag").Where("id = ?", tagID).Updates(map[string]interface{}{
		"is_del":     1,
		"deleted_on": time.Now().Unix(),
	}).Error
}

// ====== H5运维评论管理 ======

// GetH5CommentList 获取评论列表
func (s *adminService) GetH5CommentList(ctx context.Context, req *admin.H5CommentListReq) (int64, []admin.H5CommentItem, error) {
	type pComment struct {
		ID        int64  `gorm:"column:id"`
		PostID    int64  `gorm:"column:post_id"`
		UserID    int64  `gorm:"column:user_id"`
		CreatedOn int64  `gorm:"column:created_on"`
	}

	var total int64
	var rows []pComment

	baseDB := s.dao.DB().WithContext(ctx).Table("p_comment").Where("is_del = 0")
	if req.PostID > 0 {
		baseDB = baseDB.Where("post_id = ?", req.PostID)
	}
	if err := baseDB.Count(&total).Error; err != nil {
		return 0, nil, err
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	offset := (req.Page - 1) * req.PageSize

	findDB := s.dao.DB().WithContext(ctx).Table("p_comment").Where("is_del = 0")
	if req.PostID > 0 {
		findDB = findDB.Where("post_id = ?", req.PostID)
	}
	if err := findDB.Order("id ASC").Offset(offset).Limit(req.PageSize).Find(&rows).Error; err != nil {
		return 0, nil, err
	}

	// 收集用户ID
	userIds := make([]int64, len(rows))
	commentIds := make([]int64, len(rows))
	for i, r := range rows {
		userIds[i] = r.UserID
		commentIds[i] = r.ID
	}

	// 批量查询用户昵称
	type pUser struct {
		ID       int64  `gorm:"column:id"`
		Nickname string `gorm:"column:nickname"`
	}
	userMap := make(map[int64]string)
	if len(userIds) > 0 {
		var users []pUser
		s.dao.DB().WithContext(ctx).Table("p_user").Where("id IN ?", userIds).Find(&users)
		for _, u := range users {
			userMap[u.ID] = u.Nickname
		}
	}

	// 批量查询评论内容（取第一条文本内容）
	type pContent struct {
		CommentID int64  `gorm:"column:comment_id"`
		Content   string `gorm:"column:content"`
	}
	contentMap := make(map[int64]string)
	if len(commentIds) > 0 {
		var contents []pContent
		s.dao.DB().WithContext(ctx).Table("p_comment_content").Where("comment_id IN ?", commentIds).Order("sort ASC").Find(&contents)
		for _, c := range contents {
			if _, ok := contentMap[c.CommentID]; !ok {
				contentMap[c.CommentID] = c.Content
			}
		}
	}

	items := make([]admin.H5CommentItem, len(rows))
	for i, r := range rows {
		items[i] = admin.H5CommentItem{
			ID:        r.ID,
			PostID:    r.PostID,
			UserID:    r.UserID,
			Nickname:  userMap[r.UserID],
			Content:   contentMap[r.ID],
			CreatedAt: time.Unix(r.CreatedOn, 0).Format("2006-01-02 15:04:05"),
		}
	}

	return total, items, nil
}

// DeleteH5Comment 删除评论(软删除)
func (s *adminService) DeleteH5Comment(ctx context.Context, commentID int64) error {
	return s.dao.DB().WithContext(ctx).Table("p_comment").Where("id = ?", commentID).Updates(map[string]interface{}{
		"is_del":     1,
		"deleted_on": time.Now().Unix(),
	}).Error
}

// ====== H5运维KOL属性管理 ======

// GetKolProfile 获取KOL人物属性
func (s *adminService) GetKolProfile(ctx context.Context, userID int64) (*admin.H5KolProfileItem, error) {
	type pKolProfile struct {
		UserID        int64  `gorm:"column:user_id"`
		Height        string `gorm:"column:height"`
		Weight        string `gorm:"column:weight"`
		Measurements  string `gorm:"column:measurements"`
		SkinTone      string `gorm:"column:skin_tone"`
		EyeColor      string `gorm:"column:eye_color"`
		Orientation   string `gorm:"column:orientation"`
		Preferences   string `gorm:"column:preferences"`
		FavoriteFoods string `gorm:"column:favorite_foods"`
		ClothingStyle string `gorm:"column:clothing_style"`
		MakeupStyle   string `gorm:"column:makeup_style"`
		CategoryID    int64  `gorm:"column:category_id"`
		SystemPrompt  string `gorm:"column:system_prompt"`
		ApiKey        string `gorm:"column:api_key"`
	}
	var row pKolProfile
	err := s.dao.DB().WithContext(ctx).Table("p_kol_profile").Where("user_id = ? AND is_del = 0", userID).First(&row).Error
	if err != nil {
		return nil, err
	}
	return &admin.H5KolProfileItem{
		UserID:        row.UserID,
		Height:        row.Height,
		Weight:        row.Weight,
		Measurements:  row.Measurements,
		SkinTone:      row.SkinTone,
		EyeColor:      row.EyeColor,
		Orientation:   row.Orientation,
		Preferences:   row.Preferences,
		FavoriteFoods: row.FavoriteFoods,
		ClothingStyle: row.ClothingStyle,
		MakeupStyle:   row.MakeupStyle,
		CategoryID:    row.CategoryID,
			SystemPrompt:  row.SystemPrompt,
			ApiKey:        row.ApiKey,
	}, nil
}

// SaveKolProfile 保存KOL人物属性(upsert)
func (s *adminService) SaveKolProfile(ctx context.Context, req *admin.H5KolProfileSaveReq) error {
	type pKolProfile struct {
		ID           int64 `gorm:"column:id"`
		UserID       int64 `gorm:"column:user_id"`
	}
	var existing pKolProfile
	err := s.dao.DB().WithContext(ctx).Table("p_kol_profile").Where("user_id = ? AND is_del = 0", req.UserID).First(&existing).Error
	updates := map[string]interface{}{
		"user_id":        req.UserID,
		"height":         req.Height,
		"weight":         req.Weight,
		"measurements":   req.Measurements,
		"skin_tone":      req.SkinTone,
		"eye_color":      req.EyeColor,
		"orientation":    req.Orientation,
		"preferences":    req.Preferences,
		"favorite_foods": req.FavoriteFoods,
		"clothing_style": req.ClothingStyle,
		"makeup_style":   req.MakeupStyle,
		"category_id":    req.CategoryID,
		"system_prompt":  req.SystemPrompt,
		"api_key":       req.ApiKey,
	}
	if err != nil {
		// 不存在则创建
		updates["created_on"] = time.Now().Unix()
		updates["modified_on"] = time.Now().Unix()
		return s.dao.DB().WithContext(ctx).Table("p_kol_profile").Create(updates).Error
	}
	updates["modified_on"] = time.Now().Unix()
	return s.dao.DB().WithContext(ctx).Table("p_kol_profile").Where("id = ?", existing.ID).Updates(updates).Error
}

// ====== H5运维KOL分类管理 ======

// GetKolCategoryList 获取KOL分类列表
func (s *adminService) GetKolCategoryList(ctx context.Context) ([]admin.H5KolCategoryItem, error) {
	type pCat struct {
		ID   int64  `gorm:"column:id"`
		Name string `gorm:"column:name"`
		Sort int    `gorm:"column:sort"`
	}
	var rows []pCat
	if err := s.dao.DB().WithContext(ctx).Table("p_kol_category").Where("is_del = 0").Order("sort ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	items := make([]admin.H5KolCategoryItem, len(rows))
	for i, r := range rows {
		var cnt int64
		s.dao.DB().WithContext(ctx).Table("p_kol_profile").Where("category_id = ? AND is_del = 0", r.ID).Count(&cnt)
		items[i] = admin.H5KolCategoryItem{ID: r.ID, Name: r.Name, Sort: r.Sort, UserCount: cnt}
	}
	return items, nil
}

// SaveKolCategory 保存KOL分类(新增/更新)
func (s *adminService) SaveKolCategory(ctx context.Context, req *admin.H5KolCategorySaveReq) error {
	if req.ID > 0 {
		return s.dao.DB().WithContext(ctx).Table("p_kol_category").Where("id = ?", req.ID).Updates(map[string]interface{}{
			"name": req.Name, "sort": req.Sort, "modified_on": time.Now().Unix(),
		}).Error
	}
	return s.dao.DB().WithContext(ctx).Table("p_kol_category").Create(map[string]interface{}{
		"name": req.Name, "sort": req.Sort, "created_on": time.Now().Unix(),
	}).Error
}

// DeleteKolCategory 删除KOL分类
func (s *adminService) DeleteKolCategory(ctx context.Context, id int64) error {
	return s.dao.DB().WithContext(ctx).Table("p_kol_category").Where("id = ?", id).Updates(map[string]interface{}{
		"is_del": 1, "deleted_on": time.Now().Unix(),
	}).Error
}

// ====== H5运维KOL用户管理 ======

// GetKolManageList 获取KOL管理列表
func (s *adminService) GetKolManageList(ctx context.Context, req *admin.H5KolManageListReq) (int64, []admin.H5KolManageItem, error) {
	type pRow struct {
		ID           int64  `gorm:"column:id"`
		Nickname     string `gorm:"column:nickname"`
		Username     string `gorm:"column:username"`
		Avatar       string `gorm:"column:avatar"`
		CreatedOn    int64  `gorm:"column:created_on"`
		CategoryID   int64  `gorm:"column:category_id"`
		SystemPrompt  string `gorm:"column:system_prompt"`
		ApiKey        string `gorm:"column:api_key"`
		Height       string `gorm:"column:height"`
		Weight       string `gorm:"column:weight"`
		Measurements string `gorm:"column:measurements"`
		Address      string `gorm:"column:address"`
	}

	var total int64
	db := s.dao.DB().WithContext(ctx).Table("p_user u").
		Joins("LEFT JOIN p_kol_profile p ON u.id = p.user_id AND p.is_del = 0").
		Where("u.is_del = 0 AND u.is_kol = 1")
	if req.CategoryID > 0 {
		db = db.Where("p.category_id = ?", req.CategoryID)
	}
	if req.Keyword != "" {
		db = db.Where("u.nickname LIKE ?", "%"+req.Keyword+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		return 0, nil, err
	}

	if req.Page <= 0 { req.Page = 1 }
	if req.PageSize <= 0 { req.PageSize = 10 }
	offset := (req.Page - 1) * req.PageSize

	var rows []pRow
	if err := db.Select("u.id, u.nickname, u.username, u.avatar, u.address, u.created_on, COALESCE(p.category_id,0) AS category_id, COALESCE(p.height,'') AS height, COALESCE(p.weight,'') AS weight, COALESCE(p.measurements,'') AS measurements").
		Order("u.id DESC").Offset(offset).Limit(req.PageSize).Find(&rows).Error; err != nil {
		return 0, nil, err
	}

	// 获取分类名称映射
	catNames := make(map[int64]string)
	var cats []struct {
		ID   int64  `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}
	s.dao.DB().WithContext(ctx).Table("p_kol_category").Where("is_del = 0").Find(&cats)
	for _, c := range cats { catNames[c.ID] = c.Name }

	items := make([]admin.H5KolManageItem, len(rows))
	for i, r := range rows {
		items[i] = admin.H5KolManageItem{
			ID: r.ID, Nickname: r.Nickname, Username: r.Username, Avatar: r.Avatar,
			CategoryID: r.CategoryID, CategoryName: catNames[r.CategoryID],
			Height: r.Height, Weight: r.Weight, Measurements: r.Measurements,
			WalletAddress: r.Address,
			CreatedAt: time.Unix(r.CreatedOn, 0).Format("2006-01-02 15:04:05"),
		}
	}
	return total, items, nil
}

// AssignKolCategory 分配KOL分类
func (s *adminService) AssignKolCategory(ctx context.Context, req *admin.H5KolAssignCategoryReq) error {
	// 检查profile是否存在
	var cnt int64
	s.dao.DB().WithContext(ctx).Table("p_kol_profile").Where("user_id = ? AND is_del = 0", req.UserID).Count(&cnt)
	if cnt > 0 {
		return s.dao.DB().WithContext(ctx).Table("p_kol_profile").Where("user_id = ?", req.UserID).Updates(map[string]interface{}{
			"category_id": req.CategoryID, "modified_on": time.Now().Unix(),
		}).Error
	}
	return s.dao.DB().WithContext(ctx).Table("p_kol_profile").Create(map[string]interface{}{
		"user_id": req.UserID, "category_id": req.CategoryID, "created_on": time.Now().Unix(),
	}).Error
}

// ====== 探索页KOL分类 ======

// GetExploreKolCategories 获取探索页KOL分类及用户
func (s *adminService) GetExploreKolCategories(ctx context.Context) (*admin.ExploreKolCategoryResp, error) {
	type pCat struct {
		ID   int64  `gorm:"column:id"`
		Name string `gorm:"column:name"`
	}
	var cats []pCat
	s.dao.DB().WithContext(ctx).Table("p_kol_category").Where("is_del = 0").Order("sort ASC").Find(&cats)

	type pUser struct {
		ID       int64  `gorm:"column:id"`
		Nickname string `gorm:"column:nickname"`
		Username string `gorm:"column:username"`
		Avatar   string `gorm:"column:avatar"`
	}

	result := make([]admin.ExploreKolCategory, 0, len(cats))
	for _, c := range cats {
		var users []pUser
		s.dao.DB().WithContext(ctx).Table("p_user u").
			Joins("INNER JOIN p_kol_profile p ON u.id = p.user_id AND p.is_del = 0 AND p.category_id = ?", c.ID).
			Where("u.is_del = 0 AND u.is_kol = 1").Limit(6).Find(&users)
		eu := make([]admin.ExploreKolUser, len(users))
		for i, u := range users {
			eu[i] = admin.ExploreKolUser{ID: u.ID, Nickname: u.Nickname, Username: u.Username, Avatar: u.Avatar}
		}
		result = append(result, admin.ExploreKolCategory{ID: c.ID, Name: c.Name, Users: eu})
	}
	return &admin.ExploreKolCategoryResp{Categories: result}, nil
}

// ====== H5运维系统消息管理 ======

// GetSysMsgList 获取系统消息列表
func (s *adminService) GetSysMsgList(ctx context.Context, req *admin.H5SysMsgListReq) (int64, []admin.H5SysMsgItem, error) {
	type pMsg struct {
		ID             int64  `gorm:"column:id"`
		SenderUserID   int64  `gorm:"column:sender_user_id"`
		ReceiverUserID int64  `gorm:"column:receiver_user_id"`
		Brief          string `gorm:"column:brief"`
		Content        string `gorm:"column:content"`
		IsRead         int8   `gorm:"column:is_read"`
		CreatedOn      int64  `gorm:"column:created_on"`
	}
	var total int64
	var rows []pMsg

	db := s.dao.DB().WithContext(ctx).Table("p_notice").Where("is_del = 0")
	if err := db.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if req.Page <= 0 { req.Page = 1 }
	if req.PageSize <= 0 { req.PageSize = 10 }
	offset := (req.Page - 1) * req.PageSize
	if err := db.Order("id DESC").Offset(offset).Limit(req.PageSize).Find(&rows).Error; err != nil {
		return 0, nil, err
	}

	items := make([]admin.H5SysMsgItem, len(rows))
	for i, r := range rows {
		items[i] = admin.H5SysMsgItem{
			ID: r.ID, SenderID: r.SenderUserID, ReceiverID: r.ReceiverUserID,
			Brief: r.Brief, Content: r.Content, IsRead: r.IsRead,
			CreatedAt: time.Unix(r.CreatedOn, 0).Format("2006-01-02 15:04:05"),
		}
	}
	return total, items, nil
}

// CreateSysMsg 创建系统消息
func (s *adminService) CreateSysMsg(ctx context.Context, req *admin.H5SysMsgCreateReq) error {
	// ReceiverID int64，0表示全员
	if req.ReceiverID < 0 {
		req.ReceiverID = 0
	}
	return s.dao.DB().WithContext(ctx).Table("p_notice").Create(map[string]interface{}{
		"sender_user_id":   0,
		"receiver_user_id": req.ReceiverID,
		"brief":            req.Brief,
		"content":          req.Content,
		"is_read":          0,
		"created_on":       time.Now().Unix(),
	}).Error
}

// DeleteSysMsg 删除系统消息
func (s *adminService) DeleteSysMsg(ctx context.Context, id int64) error {
	return s.dao.DB().WithContext(ctx).Table("p_notice").Where("id = ?", id).Updates(map[string]interface{}{
		"is_del": 1, "deleted_on": time.Now().Unix(),
	}).Error
}
