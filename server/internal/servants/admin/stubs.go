// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/vicishero/NaiL/server/internal/core/admin"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
	"github.com/vicishero/NaiL/server/pkg/app"
	"github.com/vicishero/NaiL/server/pkg/xerror"
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

// ReloadSystem 重载系统配置
func (s *AuthServant) ReloadSystem(c *gin.Context) {
	if err := s.service.ReloadSystem(c.Request.Context()); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{"msg": "重载成功"})
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

// ====== 字典管理接口 ======

// CreateSysDictionary 创建字典
func (s *AuthServant) CreateSysDictionary(c *gin.Context) {
	var dict dbr.SysDictionary
	if err := c.ShouldBindJSON(&dict); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.CreateSysDictionary(c.Request.Context(), &dict); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(dict)
}

// DeleteSysDictionary 删除字典
func (s *AuthServant) DeleteSysDictionary(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteSysDictionary(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// UpdateSysDictionary 更新字典
func (s *AuthServant) UpdateSysDictionary(c *gin.Context) {
	var dict dbr.SysDictionary
	if err := c.ShouldBindJSON(&dict); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.UpdateSysDictionary(c.Request.Context(), &dict); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(dict)
}

// FindSysDictionary 根据ID查找字典
func (s *AuthServant) FindSysDictionary(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	dict, err := s.service.FindSysDictionary(c.Request.Context(), req.ID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{"sysDictionary": dict})
}

// GetSysDictionaryList 获取字典列表
func (s *AuthServant) GetSysDictionaryList(c *gin.Context) {
	var req admin.GetSysDictionaryListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10000 // 前端不分页，获取全部
	}
	resp, err := s.service.GetSysDictionaryList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	// 前端期望 res.data 直接是字典数组
	app.NewResponse(c).ToResponse(resp.List)
}

// ====== 参数管理接口 ======

// CreateSysParams 创建参数
func (s *AuthServant) CreateSysParams(c *gin.Context) {
	var params dbr.SysParams
	if err := c.ShouldBindJSON(&params); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.CreateSysParams(c.Request.Context(), &params); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(params)
}

// DeleteSysParams 删除参数
func (s *AuthServant) DeleteSysParams(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteSysParams(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// DeleteSysParamsByIds 批量删除参数
func (s *AuthServant) DeleteSysParamsByIds(c *gin.Context) {
	var req struct {
		Ids []uint `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteSysParamsByIds(c.Request.Context(), req.Ids); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// UpdateSysParams 更新参数
func (s *AuthServant) UpdateSysParams(c *gin.Context) {
	var params dbr.SysParams
	if err := c.ShouldBindJSON(&params); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.UpdateSysParams(c.Request.Context(), &params); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(params)
}

// FindSysParams 根据ID查找参数
func (s *AuthServant) FindSysParams(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	params, err := s.service.FindSysParams(c.Request.Context(), req.ID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{"sysParams": params})
}

// GetSysParamsList 获取参数列表
func (s *AuthServant) GetSysParamsList(c *gin.Context) {
	var req admin.GetSysParamsListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := s.service.GetSysParamsList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{
		"list":     resp.List,
		"total":    resp.Total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

// ====== 登录日志管理接口 ======

// DeleteLoginLog 删除登录日志
func (s *AuthServant) DeleteLoginLog(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteLoginLog(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// DeleteLoginLogByIds 批量删除登录日志
func (s *AuthServant) DeleteLoginLogByIds(c *gin.Context) {
	var req struct {
		Ids []uint `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteLoginLogByIds(c.Request.Context(), req.Ids); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// FindLoginLog 根据ID查找登录日志
func (s *AuthServant) FindLoginLog(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	log, err := s.service.FindLoginLog(c.Request.Context(), req.ID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{"sysLoginLog": log})
}

// GetLoginLogList 获取登录日志列表
func (s *AuthServant) GetLoginLogList(c *gin.Context) {
	var req admin.GetLoginLogListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := s.service.GetLoginLogList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{
		"list":     resp.List,
		"total":    resp.Total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

// ========== API Token 管理接口 ==========

// CreateApiToken 创建Token
func (s *AuthServant) CreateApiToken(c *gin.Context) {
	var req admin.CreateApiTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	token, err := s.service.CreateApiToken(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(token)
}

// DeleteApiToken 删除Token
func (s *AuthServant) DeleteApiToken(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteApiToken(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// DeleteApiTokenByIds 批量删除Token
func (s *AuthServant) DeleteApiTokenByIds(c *gin.Context) {
	var req struct {
		Ids []uint `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteApiTokenByIds(c.Request.Context(), req.Ids); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// GetApiTokenList 获取Token列表
func (s *AuthServant) GetApiTokenList(c *gin.Context) {
	var req admin.GetApiTokenListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := s.service.GetApiTokenList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{
		"list":     resp.List,
		"total":    resp.Total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

// ========== 错误日志管理接口 ==========

// CreateSysError 创建错误日志
func (s *AuthServant) CreateSysError(c *gin.Context) {
	var err dbr.SysError
	if err := c.ShouldBindJSON(&err); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.CreateSysError(c.Request.Context(), &err); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(err)
}

// DeleteSysError 删除错误日志
func (s *AuthServant) DeleteSysError(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteSysError(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// DeleteSysErrorByIds 批量删除错误日志
func (s *AuthServant) DeleteSysErrorByIds(c *gin.Context) {
	var req struct {
		Ids []uint `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteSysErrorByIds(c.Request.Context(), req.Ids); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// UpdateSysError 更新错误日志（处理状态）
func (s *AuthServant) UpdateSysError(c *gin.Context) {
	var req admin.UpdateSysErrorReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.UpdateSysError(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// FindSysError 根据ID查找错误日志
func (s *AuthServant) FindSysError(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	errItem, err := s.service.FindSysError(c.Request.Context(), req.ID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{"sysError": errItem})
}

// GetSysErrorList 获取错误日志列表
func (s *AuthServant) GetSysErrorList(c *gin.Context) {
	var req admin.GetSysErrorListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := s.service.GetSysErrorList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{
		"list":     resp.List,
		"total":    resp.Total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

// GetSysErrorPublic 获取公开的常见错误
func (s *AuthServant) GetSysErrorPublic(c *gin.Context) {
	resp, err := s.service.GetSysErrorPublic(c.Request.Context())
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(resp)
}

// GetSysErrorSolution 获取错误解决方案
func (s *AuthServant) GetSysErrorSolution(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	resp, err := s.service.GetSysErrorSolution(c.Request.Context(), req.ID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(resp)
}

// ========== 版本管理接口 ==========

// CreateSysVersion 创建版本
func (s *AuthServant) CreateSysVersion(c *gin.Context) {
	var req admin.CreateSysVersionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	version, err := s.service.CreateSysVersion(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(version)
}

// UpdateSysVersion 更新版本
func (s *AuthServant) UpdateSysVersion(c *gin.Context) {
	var version dbr.SysVersion
	if err := c.ShouldBindJSON(&version); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.UpdateSysVersion(c.Request.Context(), &version); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(version)
}

// DeleteSysVersion 删除版本
func (s *AuthServant) DeleteSysVersion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteSysVersion(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// DeleteSysVersionByIds 批量删除版本
func (s *AuthServant) DeleteSysVersionByIds(c *gin.Context) {
	var req struct {
		Ids []uint `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteSysVersionByIds(c.Request.Context(), req.Ids); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// FindSysVersion 根据ID查找版本
func (s *AuthServant) FindSysVersion(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	version, err := s.service.FindSysVersion(c.Request.Context(), req.ID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{"sysVersion": version})
}

// GetSysVersionList 获取版本列表
func (s *AuthServant) GetSysVersionList(c *gin.Context) {
	var req admin.GetSysVersionListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	resp, err := s.service.GetSysVersionList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{
		"list":     resp.List,
		"total":    resp.Total,
		"page":     req.Page,
		"pageSize": req.PageSize,
	})
}

// ExportVersion 导出版本
func (s *AuthServant) ExportVersion(c *gin.Context) {
	var req admin.ExportVersionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		// 如果没有传参数，默认导出全部
		req.Ids = []uint{}
	}
	data, err := s.service.ExportVersion(c.Request.Context(), req.Ids)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	// 直接返回JSON数据供下载
	c.Header("Content-Type", "application/json")
	c.Header("Content-Disposition", "attachment; filename=versions.json")
	c.String(200, string(data))
}

// ImportVersion 导入版本
func (s *AuthServant) ImportVersion(c *gin.Context) {
	var req admin.ImportVersionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.ImportVersion(c.Request.Context(), []byte(req.Data)); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// DownloadVersionJson 下载版本JSON
func (s *AuthServant) DownloadVersionJson(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams)
		return
	}
	var req struct {
		ID uint `json:"id" uri:"id"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	data, err := s.service.GenerateVersionJson(c.Request.Context(), req.ID)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	c.Header("Content-Type", "application/json")
	c.Header("Content-Disposition", "attachment; filename=version.json")
	c.String(200, string(data))
}
