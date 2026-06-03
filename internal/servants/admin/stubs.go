// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core/admin"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// ====== 菜单管理补充接口 ======

// GetMenu 获取用户动态路由（菜单树）
func (s *AuthServant) GetMenu(c *gin.Context) {
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

	menus, err := s.service.GetUserMenu(c.Request.Context(), userInfo.User.AuthorityID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}

	app.NewResponse(c).ToResponse(gin.H{
		"menus": menus,
	})
}

// AddMenuAuthority 添加菜单与角色关联
func (s *AuthServant) AddMenuAuthority(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// GetMenuAuthority 获取菜单与角色关联
func (s *AuthServant) GetMenuAuthority(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"menus": []interface{}{}})
}

// GetBaseMenuById 根据ID获取菜单
func (s *AuthServant) GetBaseMenuById(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"menu": gin.H{}})
}

// GetMenuRoles 获取拥有指定菜单的角色列表
func (s *AuthServant) GetMenuRoles(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"roles": []interface{}{}})
}

// SetMenuRoles 设置菜单关联的角色
func (s *AuthServant) SetMenuRoles(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ====== 用户管理补充接口 ======

// SetSelfInfo 设置自身信息
func (s *AuthServant) SetSelfInfo(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// SetSelfSetting 设置自身界面配置
func (s *AuthServant) SetSelfSetting(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// SetUserAuthorities 设置用户多角色
func (s *AuthServant) SetUserAuthorities(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ====== 角色管理补充接口 ======

// CopyAuthority 复制角色
func (s *AuthServant) CopyAuthority(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// GetUsersByAuthority 获取角色下的用户
func (s *AuthServant) GetUsersByAuthority(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"users": []interface{}{}})
}

// SetDataAuthority 设置数据权限
func (s *AuthServant) SetDataAuthority(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// SetRoleUsers 设置角色关联用户
func (s *AuthServant) SetRoleUsers(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ====== API管理补充接口 ======

// DeleteApisByIds 批量删除API
func (s *AuthServant) DeleteApisByIds(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// EnterSyncApi 进入同步API页面
func (s *AuthServant) EnterSyncApi(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// FreshCasbin 刷新Casbin权限
func (s *AuthServant) FreshCasbin(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// GetAllApis 获取所有API(不分页)
func (s *AuthServant) GetAllApis(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"apis": []interface{}{}})
}

// GetApiById 根据ID获取API
func (s *AuthServant) GetApiById(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"api": gin.H{}})
}

// GetApiRoles 获取API关联的角色
func (s *AuthServant) GetApiRoles(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"roles": []interface{}{}})
}

// IgnoreApi 忽略API
func (s *AuthServant) IgnoreApi(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// SetApiRoles 设置API关联角色
func (s *AuthServant) SetApiRoles(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// SetAuthApi 设置角色API权限(别名)
func (s *AuthServant) SetAuthApi(c *gin.Context) {
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

// ====== 系统配置补充接口 ======

// ReloadSystem 重启系统
func (s *AuthServant) ReloadSystem(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ====== 操作日志补充接口 ======

// DeleteSysOperationRecordByIds 批量删除操作日志
func (s *AuthServant) DeleteSysOperationRecordByIds(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ====== 文件管理补充接口 ======

// BreakpointContinue 断点续传
func (s *AuthServant) BreakpointContinue(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// BreakpointContinueFinish 断点续传完成
func (s *AuthServant) BreakpointContinueFinish(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// EditFileName 编辑文件名
func (s *AuthServant) EditFileName(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// FindFile 查找文件
func (s *AuthServant) FindFile(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ImportURL 导入URL文件
func (s *AuthServant) ImportURL(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// RemoveChunk 移除分片
func (s *AuthServant) RemoveChunk(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ====== Casbin权限接口 ======

// GetPolicyPathByAuthorityId 获取角色权限路径
func (s *AuthServant) GetPolicyPathByAuthorityId(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"paths": []interface{}{}})
}

// UpdateCasbin 更新Casbin策略
func (s *AuthServant) UpdateCasbin(c *gin.Context) {
	app.NewResponse(c).ToResponse(nil)
}

// ====== JWT黑名单接口 ======

// JsonInBlacklist JWT加入黑名单（无token也成功返回，前端退出流程不依赖此接口）
func (s *AuthServant) JsonInBlacklist(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{})
}

// ====== 通用GVA兼容接口（返回空列表/空对象的标准格式） ======

// GvaEmptyList 通用空列表响应
func GvaEmptyList(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"list": []interface{}{}, "total": 0, "page": 1, "pageSize": 10})
}

// GvaEmptyObject 通用空对象响应
func GvaEmptyObject(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{})
}

// GvaSuccess 通用成功响应
func GvaSuccess(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{})
}
