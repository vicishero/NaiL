// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"context"

	"gorm.io/gorm"
	"github.com/rocboss/paopao-ce/internal/dao/admin/dbr"
)

// AdminService 管理后台服务接口
type AdminService interface {
	// Login 管理员登录
	Login(ctx context.Context, username, password string) (*LoginResp, error)

	// Logout 管理员退出登录
	Logout(ctx context.Context, userId uint, token string) error

	// GetUserInfo 获取用户详细信息，包含权限和角色
	GetUserInfo(ctx context.Context, userId uint) (*UserInfoResp, error)

	// GetUserMenu 获取用户有权限的菜单树
	GetUserMenu(ctx context.Context, authorityId uint) ([]*MenuResp, error)

	// CheckPermission 校验用户是否有指定接口的访问权限
	CheckPermission(ctx context.Context, userId uint, path, method string) (bool, error)

	// ChangePassword 修改用户密码
	ChangePassword(ctx context.Context, userId uint, oldPwd, newPwd string) error

	// RecordOperationLog 记录操作日志
	RecordOperationLog(ctx context.Context, log *dbr.SysOperationRecord) error

	// CreateUser 创建用户
	CreateUser(ctx context.Context, req *CreateUserReq) error

	// UpdateUser 更新用户信息
	UpdateUser(ctx context.Context, req *UpdateUserReq) error

	// SetUserAuthority 设置用户角色
	SetUserAuthority(ctx context.Context, req *SetUserAuthorityReq) error

	// ResetPassword 重置用户密码
	ResetPassword(ctx context.Context, req *ResetPasswordReq) error

	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, userId uint) error

	// GetUserList 获取用户列表
	GetUserList(ctx context.Context, req *UserListReq) (*UserListResp, error)

	// CreateAuthority 创建角色
	CreateAuthority(ctx context.Context, req *CreateAuthorityReq) error

	// UpdateAuthority 更新角色
	UpdateAuthority(ctx context.Context, req *UpdateAuthorityReq) error

	// DeleteAuthority 删除角色
	DeleteAuthority(ctx context.Context, authorityId uint) error

	// GetAuthorityList 获取角色列表
	GetAuthorityList(ctx context.Context, req *AuthorityListReq) (*AuthorityListResp, error)

	// GetAllAuthorities 获取所有角色
	GetAllAuthorities(ctx context.Context) ([]*dbr.SysAuthority, error)

	// SetAuthorityMenu 设置角色菜单权限
	SetAuthorityMenu(ctx context.Context, req *SetAuthorityMenuReq) error

	// GetAuthorityMenu 获取角色菜单权限
	GetAuthorityMenu(ctx context.Context, req *GetAuthorityMenuReq) (*GetAuthorityMenuResp, error)

	// SetAuthorityApi 设置角色API权限
	SetAuthorityApi(ctx context.Context, req *SetAuthorityApiReq) error

	// GetAuthorityApi 获取角色API权限
	GetAuthorityApi(ctx context.Context, req *GetAuthorityApiReq) (*GetAuthorityApiResp, error)

	// CreateBaseMenu 创建菜单
	CreateBaseMenu(ctx context.Context, req *CreateBaseMenuReq) error

	// UpdateBaseMenu 更新菜单
	UpdateBaseMenu(ctx context.Context, req *UpdateBaseMenuReq) error

	// DeleteBaseMenu 删除菜单
	DeleteBaseMenu(ctx context.Context, menuId uint) error

	// GetBaseMenuTree 获取菜单树
	GetBaseMenuTree(ctx context.Context, req *GetBaseMenuTreeReq) (*GetBaseMenuTreeResp, error)

	// CreateApi 创建API
	CreateApi(ctx context.Context, req *CreateApiReq) error

	// UpdateApi 更新API
	UpdateApi(ctx context.Context, req *UpdateApiReq) error

	// DeleteApi 删除API
	DeleteApi(ctx context.Context, apiId uint) error

	// GetApiList 获取API列表
	GetApiList(ctx context.Context, req *GetApiListReq) (*GetApiListResp, error)

	// GetAllApiGroups 获取所有API分组
	GetAllApiGroups(ctx context.Context) (*GetAllApiGroupsResp, error)

	// SyncApi 同步接口
	SyncApi(ctx context.Context, req *SyncApiReq) error

	// GetOperationLogList 获取操作日志列表
	GetOperationLogList(ctx context.Context, req *GetOperationLogListReq) (*GetOperationLogListResp, error)

	// DeleteOperationLog 删除操作日志
	DeleteOperationLog(ctx context.Context, logId uint) error

	// GetSystemConfig 获取系统配置
	GetSystemConfig(ctx context.Context) (*GetSystemConfigResp, error)

	// SetSystemConfig 设置系统配置
	SetSystemConfig(ctx context.Context, req *SetSystemConfigReq) error

	// GetServerInfo 获取服务器信息
	GetServerInfo(ctx context.Context) (*GetServerInfoResp, error)

	// UploadFile 上传文件
	UploadFile(ctx context.Context, file interface{}, filename string) (*UploadFileResp, error)

	// GetFileList 获取文件列表
	GetFileList(ctx context.Context, req *GetFileListReq) (*GetFileListResp, error)

	// DeleteFile 删除文件
	DeleteFile(ctx context.Context, fileId uint) error

	// H5运维用户管理
	GetH5UserList(ctx context.Context, req *H5UserListReq) (int64, []H5UserItem, error)
	GetH5User(ctx context.Context, userID int64) (*H5UserItem, error)
	UpdateH5User(ctx context.Context, req *H5UserUpdateReq) error
	DeleteH5User(ctx context.Context, userID int64) error
}

// AdminDao 管理后台数据访问接口
type AdminDao interface {
	// DB 返回底层GORM DB实例(用于直接查询非sys表)
	DB() *gorm.DB

	// GetUserByUsername 根据用户名获取用户信息
	GetUserByUsername(ctx context.Context, username string) (*dbr.SysUser, error)

	// GetUserByID 根据用户ID获取用户信息
	GetUserByID(ctx context.Context, userId uint) (*dbr.SysUser, error)

	// UpdateUser 更新用户信息
	UpdateUser(ctx context.Context, user *dbr.SysUser) error

	// GetAuthorityByID 根据角色ID获取角色信息
	GetAuthorityByID(ctx context.Context, authorityId uint) (*dbr.SysAuthority, error)

	// GetUserPermissions 获取用户的权限编码列表
	GetUserPermissions(ctx context.Context, userId uint) ([]string, error)

	// GetAuthorityMenus 获取角色有权限的菜单列表
	GetAuthorityMenus(ctx context.Context, authorityId uint) ([]*dbr.SysBaseMenu, error)

	// CheckApiPermission 校验角色是否有指定API的访问权限
	CheckApiPermission(ctx context.Context, authorityId uint, path, method string) (bool, error)

	// CreateOperationLog 创建操作日志
	CreateOperationLog(ctx context.Context, log *dbr.SysOperationRecord) error

	// CreateUser 创建用户
	CreateUser(ctx context.Context, user *dbr.SysUser) error

	// DeleteUser 删除用户
	DeleteUser(ctx context.Context, userId uint) error

	// GetUserList 获取用户列表
	GetUserList(ctx context.Context, req *UserListReq) (int64, []*dbr.SysUser, error)

	// CreateAuthority 创建角色
	CreateAuthority(ctx context.Context, authority *dbr.SysAuthority) error

	// UpdateAuthority 更新角色
	UpdateAuthority(ctx context.Context, authority *dbr.SysAuthority) error

	// DeleteAuthority 删除角色
	DeleteAuthority(ctx context.Context, authorityId uint) error

	// GetAuthorityList 获取角色列表
	GetAuthorityList(ctx context.Context, req *AuthorityListReq) (int64, []*dbr.SysAuthority, error)

	// GetAllAuthorities 获取所有角色
	GetAllAuthorities(ctx context.Context) ([]*dbr.SysAuthority, error)

	// SetAuthorityMenu 设置角色菜单权限
	SetAuthorityMenu(ctx context.Context, authorityId uint, menuIds []uint) error

	// GetAuthorityMenuIds 获取角色已授权的菜单ID列表
	GetAuthorityMenuIds(ctx context.Context, authorityId uint) ([]uint, error)

	// GetAuthorityMenusByIds 根据菜单ID列表获取菜单详情
	GetAuthorityMenusByIds(ctx context.Context, menuIds []uint) ([]*dbr.SysBaseMenu, error)

	// SetAuthorityApi 设置角色API权限
	SetAuthorityApi(ctx context.Context, authorityId uint, apiIds []uint) error

	// GetAuthorityApiIds 获取角色已授权的API ID列表
	GetAuthorityApiIds(ctx context.Context, authorityId uint) ([]uint, error)

	// GetAuthorityApisByIds 根据API ID列表获取API详情
	GetAuthorityApisByIds(ctx context.Context, apiIds []uint) ([]*dbr.SysApi, error)

	// CheckAuthorityNameExists 检查角色名称是否已存在
	CheckAuthorityNameExists(ctx context.Context, authorityName string, excludeId uint) (bool, error)

	// CountUsersByAuthorityId 统计使用该角色的用户数量
	CountUsersByAuthorityId(ctx context.Context, authorityId uint) (int64, error)

	// CountChildAuthoritiesByParentId 统计该角色的子角色数量
	CountChildAuthoritiesByParentId(ctx context.Context, parentId uint) (int64, error)

	// CheckMenuIdsValid 检查菜单ID列表是否都有效
	CheckMenuIdsValid(ctx context.Context, menuIds []uint) (bool, error)

	// GetAllMenus 获取所有菜单
	GetAllMenus(ctx context.Context) ([]*dbr.SysBaseMenu, error)

	// CreateBaseMenu 创建菜单
	CreateBaseMenu(ctx context.Context, menu *dbr.SysBaseMenu) error

	// UpdateBaseMenu 更新菜单
	UpdateBaseMenu(ctx context.Context, menu *dbr.SysBaseMenu) error

	// DeleteBaseMenu 删除菜单
	DeleteBaseMenu(ctx context.Context, menuId uint) error

	// GetBaseMenuByID 根据ID获取菜单
	GetBaseMenuByID(ctx context.Context, menuId uint) (*dbr.SysBaseMenu, error)

	// CountChildMenusByParentId 统计该菜单的子菜单数量
	CountChildMenusByParentId(ctx context.Context, parentId uint) (int64, error)

	// CountAuthorityMenusByMenuId 统计使用该菜单的角色数量
	CountAuthorityMenusByMenuId(ctx context.Context, menuId uint) (int64, error)

	// CreateApi 创建API
	CreateApi(ctx context.Context, api *dbr.SysApi) error

	// UpdateApi 更新API
	UpdateApi(ctx context.Context, api *dbr.SysApi) error

	// DeleteApi 删除API
	DeleteApi(ctx context.Context, apiId uint) error

	// GetApiByID 根据ID获取API
	GetApiByID(ctx context.Context, apiId uint) (*dbr.SysApi, error)

	// GetApiList 获取API列表
	GetApiList(ctx context.Context, req *GetApiListReq) (int64, []*dbr.SysApi, error)

	// GetAllApiGroups 获取所有API分组
	GetAllApiGroups(ctx context.Context) ([]string, error)

	// GetOperationLogList 获取操作日志列表
	GetOperationLogList(ctx context.Context, req *GetOperationLogListReq) (int64, []*dbr.SysOperationRecord, error)

	// GetOperationLogByID 根据ID获取操作日志
	GetOperationLogByID(ctx context.Context, logId uint) (*dbr.SysOperationRecord, error)

	// DeleteOperationLog 删除操作日志
	DeleteOperationLog(ctx context.Context, logId uint) error
}
