// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"
	"crypto/md5"
	"time"
	"fmt"
	"strconv"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core/admin"
	"github.com/rocboss/paopao-ce/internal/dao/admin/dbr"
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
			"site_name": "PaoPao社区",
			"site_desc": "一个有趣的社区",
			"site_url": "https://paopao.info",
			"site_logo": "/static/logo.png",
			"site_icp": "",
			"site_copyright": "© 2024 PaoPao",
			"register_enable": true,
			"email_verify": false,
			"default_avatar": "/static/avatar.png",
			"max_upload_size": 10,
			"allowed_upload_types": []string{"jpg", "jpeg", "png", "gif", "mp4", "avi", "mov"},
		},
	}, nil
}

// SetSystemConfig 设置系统配置
func (s *adminService) SetSystemConfig(ctx context.Context, req *admin.SetSystemConfigReq) error {
	// TODO: 实现保存系统配置到数据库或配置文件
	// 暂时返回成功
	return nil
}

// GetServerInfo 获取服务器信息
func (s *adminService) GetServerInfo(ctx context.Context) (*admin.GetServerInfoResp, error) {
	// TODO: 实现获取真实的服务器信息
	// 暂时返回示例数据
	return &admin.GetServerInfoResp{
		CPU:    12.5,
		Memory: 45.2,
		Disk:   67.8,
		OS:     "Linux",
		Arch:   "amd64",
		Go:     "go1.22.0",
	}, nil
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
	}

	var total int64
	var rows []pUser

	db := s.dao.DB().WithContext(ctx).Table("p_user").Where("is_del = 0")
	if req.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+req.Nickname+"%")
	}
	if req.Username != "" {
		db = db.Where("username LIKE ?", "%"+req.Username+"%")
	}
	if req.WalletAddress != "" {
		db = db.Where("phone LIKE ?", "%"+req.WalletAddress+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	if err := db.Count(&total).Error; err != nil {
		return 0, nil, err
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	offset := (req.Page - 1) * req.PageSize
	if err := db.Order("id DESC").Offset(offset).Limit(req.PageSize).Find(&rows).Error; err != nil {
		return 0, nil, err
	}

	items := make([]admin.H5UserItem, len(rows))
	for i, r := range rows {
		items[i] = admin.H5UserItem{
			ID:            r.ID,
			Nickname:      r.Nickname,
			Username:      r.Username,
			Phone:         r.Phone,
			WalletAddress: "",
			Bio:           "",
			Avatar:        r.Avatar,
			Status:        r.Status,
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
		WalletAddress: "",
		Bio:           "",
		Avatar:        row.Avatar,
		Status:        row.Status,
	}, nil
}

// UpdateH5User 更新运维用户
func (s *adminService) UpdateH5User(ctx context.Context, req *admin.H5UserUpdateReq) error {
	updates := map[string]interface{}{}
	if req.Nickname != "" {
		updates["nickname"] = req.Nickname
	}
	updates["status"] = req.Status
	return s.dao.DB().WithContext(ctx).Table("p_user").Where("id = ? AND is_del = 0", req.ID).Updates(updates).Error
}

// DeleteH5User 删除运维用户(软删除)
func (s *adminService) DeleteH5User(ctx context.Context, userID int64) error {
	return s.dao.DB().WithContext(ctx).Table("p_user").Where("id = ?", userID).Updates(map[string]interface{}{
		"is_del":     1,
		"deleted_on": time.Now().Unix(),
	}).Error
}
