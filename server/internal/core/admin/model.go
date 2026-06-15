// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"time"
	"github.com/vicishero/NaiL/server/internal/dao/admin/dbr"
)

// LoginReq 登录请求参数
type LoginReq struct {
	Username  string `json:"username" binding:"required"` // 用户名
	Password  string `json:"password" binding:"required"` // 密码
	Captcha   string `json:"captcha"` // 验证码（暂时不需要）
	CaptchaId string `json:"captchaId"` // 验证码ID（暂时不需要）
}

// LoginResp 登录响应
type LoginResp struct {
	User      *dbr.SysUser `json:"user"` // 用户信息
	Token     string       `json:"token"` // JWT token
	ExpiresAt int64        `json:"expiresAt"` // 过期时间戳
}

// UserInfoResp 用户信息响应
type UserInfoResp struct {
	User        *dbr.SysUser       `json:"user"` // 用户基本信息
	Permissions []string           `json:"permissions"` // 权限列表
	Roles       []*dbr.SysAuthority `json:"roles"` // 角色列表
}

// MenuResp 菜单响应结构
// MenuMeta 菜单元数据
type MenuMeta struct {
	Title       string `json:"title"`       // 菜单标题
	Icon        string `json:"icon"`        // 菜单图标
	Hidden      bool   `json:"hidden"`      // 是否隐藏
	KeepAlive   bool   `json:"keepAlive"`   // 是否缓存
	DefaultMenu bool   `json:"defaultMenu"` // 是否默认菜单
	CloseTab    bool   `json:"closeTab"`    // 是否关闭标签
}

type MenuResp struct {
	ID           uint         `json:"id"`
	ParentID     uint         `json:"parentId"`
	Path         string       `json:"path"`
	Name         string       `json:"name"`
	Hidden       bool         `json:"hidden"`
	Component    string       `json:"component"`
	Sort         int          `json:"sort"`
	Meta         MenuMeta     `json:"meta"`           // 菜单元数据
	Btns         []string     `json:"btns"`           // 按钮权限列表
	Redirect     string       `json:"redirect"`
	AlwaysShow   bool         `json:"alwaysShow"`
	IsKeepAlive  bool         `json:"isKeepAlive"`
	IsAffix      bool         `json:"isAffix"`
	IsIframe     bool         `json:"isIframe"`
	FrameSrc     string       `json:"frameSrc"`
	Children     []*MenuResp  `json:"children,omitempty"` // 子菜单
}

// ChangePwdReq 修改密码请求
type ChangePwdReq struct {
	OldPassword string `json:"oldPassword" binding:"required"` // 旧密码
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

// CreateUserReq 创建用户请求
type CreateUserReq struct {
	Username    string `json:"username" binding:"required"`  // 用户名
	Password    string `json:"password" binding:"required"`  // 密码
	NickName    string `json:"nickName" binding:"required"`  // 昵称
	HeaderImg   string `json:"headerImg"`                    // 头像
	AuthorityID uint   `json:"authorityId" binding:"min=1"` // 角色ID
	Phone       string `json:"phone"`                        // 手机号
	Email       string `json:"email"`                        // 邮箱
	Enable      int    `json:"enable" binding:"oneof=0 1"`  // 是否启用 1启用 0禁用
}

// UpdateUserReq 更新用户请求
type UpdateUserReq struct {
	ID          uint   `json:"id" binding:"required"`        // 用户ID
	NickName    string `json:"nickName" binding:"required"`  // 昵称
	HeaderImg   string `json:"headerImg"`                    // 头像
	AuthorityID uint   `json:"authorityId" binding:"min=1"` // 角色ID
	Phone       string `json:"phone"`                        // 手机号
	Email       string `json:"email"`                        // 邮箱
	Enable      int    `json:"enable" binding:"oneof=0 1"`  // 是否启用 1启用 0禁用
}

// SetUserAuthorityReq 设置用户角色请求
type SetUserAuthorityReq struct {
	ID          uint   `json:"id" binding:"required"`        // 用户ID
	AuthorityID uint   `json:"authorityId" binding:"min=1"` // 角色ID
}

// ResetPasswordReq 重置密码请求
type ResetPasswordReq struct {
	ID          uint   `json:"id" binding:"required"`        // 用户ID
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

// DeleteUserReq 删除用户请求
type DeleteUserReq struct {
	ID uint `json:"id" binding:"required"` // 用户ID
}

// UserListReq 用户列表请求
type UserListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Username string `form:"username" json:"username"` // 用户名搜索
	NickName string `form:"nickName" json:"nickName"` // 昵称搜索
	Enable   *int   `form:"enable" json:"enable"`     // 启用状态搜索
}

// UserListResp 用户列表响应
type UserListResp struct {
	Total int64               `json:"total"` // 总数
	List  []*dbr.SysUser      `json:"list"`  // 用户列表
}

// CreateAuthorityReq 创建角色请求
type CreateAuthorityReq struct {
	ParentID      uint   `json:"parentId"`                   // 父角色ID
	AuthorityName string `json:"authorityName" binding:"required"` // 角色名称
	DefaultRouter string `json:"defaultRouter" binding:"required"` // 默认路由
	ShowStatus    int    `json:"showStatus" binding:"oneof=0 1"`  // 显示状态 1显示 0隐藏
	AuthorityType int    `json:"authorityType"`               // 角色类型
}

// UpdateAuthorityReq 更新角色请求
type UpdateAuthorityReq struct {
	ID            uint   `json:"id" binding:"required"`        // 角色ID
	AuthorityName string `json:"authorityName" binding:"required"` // 角色名称
	ParentID      uint   `json:"parentId"`                   // 父角色ID
	DefaultRouter string `json:"defaultRouter" binding:"required"` // 默认路由
	ShowStatus    int    `json:"showStatus" binding:"oneof=0 1"`  // 显示状态 1显示 0隐藏
	AuthorityType int    `json:"authorityType"`               // 角色类型
}

// DeleteAuthorityReq 删除角色请求
type DeleteAuthorityReq struct {
	ID uint `json:"id" binding:"required"` // 角色ID
}

// AuthorityListReq 角色列表请求
type AuthorityListReq struct {
	Page           int    `form:"page" json:"page"`           // 页码
	PageSize       int    `form:"pageSize" json:"pageSize"`   // 每页数量
	AuthorityName  string `form:"authorityName" json:"authorityName"` // 角色名称搜索
}

// AuthorityListResp 角色列表响应
type AuthorityListResp struct {
	Total int64               `json:"total"` // 总数
	List  []*dbr.SysAuthority `json:"list"`  // 角色列表
}

// SetAuthorityMenuReq 设置角色菜单权限请求
type SetAuthorityMenuReq struct {
	AuthorityID uint     `json:"authorityId" binding:"min=1"` // 角色ID
	MenuIds     []uint `json:"menuIds" binding:"required"`    // 菜单ID列表
}

// GetAuthorityMenuReq 获取角色菜单权限请求
type GetAuthorityMenuReq struct {
	AuthorityID uint `form:"authorityId" binding:"min=1"` // 角色ID
}

// GetAuthorityMenuResp 获取角色菜单权限响应
type GetAuthorityMenuResp struct {
	Menus []*dbr.SysBaseMenu `json:"menus"` // 菜单列表
}

// SetAuthorityApiReq 设置角色API权限请求
type SetAuthorityApiReq struct {
	AuthorityID uint     `json:"authorityId" binding:"min=1"` // 角色ID
	ApiIds      []uint `json:"apiIds" binding:"required"`     // API ID列表
}

// GetAuthorityApiReq 获取角色API权限请求
type GetAuthorityApiReq struct {
	AuthorityID uint `form:"authorityId" binding:"min=1"` // 角色ID
}

// GetAuthorityApiResp 获取角色API权限响应
type GetAuthorityApiResp struct {
	Apis []*dbr.SysApi `json:"apis"` // API列表
}

// CreateBaseMenuReq 创建菜单请求
type CreateBaseMenuReq struct {
	ParentID      uint   `json:"parentId"`                     // 父菜单ID
	Path          string `json:"path" binding:"required"`      // 路由路径
	Name          string `json:"name" binding:"required"`      // 路由名称
	Hidden        bool   `json:"hidden"`                       // 是否隐藏
	Component     string `json:"component" binding:"required"` // 组件路径
	Sort          int    `json:"sort" binding:"min=0"`         // 排序
	Icon          string `json:"icon"`                         // 菜单图标
	Title         string `json:"title" binding:"required"`     // 菜单标题
	Redirect      string `json:"redirect"`                     // 重定向地址
	AlwaysShow    bool   `json:"alwaysShow"`                   // 是否总是显示
	IsKeepAlive   bool   `json:"isKeepAlive"`                  // 是否缓存
	IsAffix       bool   `json:"isAffix"`                      // 是否固定在标签栏
	IsIframe      bool   `json:"isIframe"`                     // 是否iframe嵌入
	FrameSrc      string `json:"frameSrc"`                     // iframe地址
}

// UpdateBaseMenuReq 更新菜单请求
type UpdateBaseMenuReq struct {
	ID            uint   `json:"id" binding:"required"`        // 菜单ID
	ParentID      uint   `json:"parentId"`                     // 父菜单ID
	Path          string `json:"path" binding:"required"`      // 路由路径
	Name          string `json:"name" binding:"required"`      // 路由名称
	Hidden        bool   `json:"hidden"`                       // 是否隐藏
	Component     string `json:"component" binding:"required"` // 组件路径
	Sort          int    `json:"sort" binding:"min=0"`         // 排序
	Icon          string `json:"icon"`                         // 菜单图标
	Title         string `json:"title" binding:"required"`     // 菜单标题
	Redirect      string `json:"redirect"`                     // 重定向地址
	AlwaysShow    bool   `json:"alwaysShow"`                   // 是否总是显示
	IsKeepAlive   bool   `json:"isKeepAlive"`                  // 是否缓存
	IsAffix       bool   `json:"isAffix"`                      // 是否固定在标签栏
	IsIframe      bool   `json:"isIframe"`                     // 是否iframe嵌入
	FrameSrc      string `json:"frameSrc"`                     // iframe地址
}

// DeleteBaseMenuReq 删除菜单请求
type DeleteBaseMenuReq struct {
	ID uint `json:"id" binding:"required"` // 菜单ID
}

// GetBaseMenuTreeReq 获取菜单树请求
type GetBaseMenuTreeReq struct {
	// 暂无参数，返回完整菜单树
}

// GetBaseMenuTreeResp 获取菜单树响应
type GetBaseMenuTreeResp struct {
	Menus []*MenuResp `json:"menus"` // 菜单树
}

// CreateApiReq 创建API请求
type CreateApiReq struct {
	Path        string `json:"path" binding:"required"`      // 接口路径
	Description string `json:"description" binding:"required"` // 接口描述
	ApiGroup    string `json:"apiGroup" binding:"required"`  // 接口分组
	Method      string `json:"method" binding:"required"`    // 请求方法
}

// UpdateApiReq 更新API请求
type UpdateApiReq struct {
	ID          uint   `json:"id" binding:"required"`        // API ID
	Path        string `json:"path" binding:"required"`      // 接口路径
	Description string `json:"description" binding:"required"` // 接口描述
	ApiGroup    string `json:"apiGroup" binding:"required"`  // 接口分组
	Method      string `json:"method" binding:"required"`    // 请求方法
}

// DeleteApiReq 删除API请求
type DeleteApiReq struct {
	ID uint `json:"id" binding:"required"` // API ID
}

// GetApiListReq API列表请求
type GetApiListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Path     string `form:"path" json:"path"`         // 路径搜索
	ApiGroup string `form:"apiGroup" json:"apiGroup"` // 分组搜索
	Method   string `form:"method" json:"method"`     // 请求方法搜索
}

// GetApiListResp API列表响应
type GetApiListResp struct {
	Total int64        `json:"total"` // 总数
	List  []*dbr.SysApi `json:"list"`  // API列表
}

// GetAllApiGroupsResp 获取所有API分组响应
type GetAllApiGroupsResp struct {
	Groups []string `json:"groups"` // 分组列表
}

// SyncApiReq 同步接口请求
type SyncApiReq struct {
	// 暂时不需要参数，同步所有路由
}

// GetOperationLogListReq 操作日志列表请求
type GetOperationLogListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Username string `form:"username" json:"username"` // 用户名搜索
	Path     string `form:"path" json:"path"`         // 路径搜索
	Method   string `form:"method" json:"method"`     // 请求方法搜索
	Status   *int   `form:"status" json:"status"`     // 状态搜索
}

// GetOperationLogListResp 操作日志列表响应
type GetOperationLogListResp struct {
	Total int64                     `json:"total"` // 总数
	List  []*dbr.SysOperationRecord `json:"list"`  // 操作日志列表
}

// DeleteOperationLogReq 删除操作日志请求
type DeleteOperationLogReq struct {
	ID uint `json:"id" binding:"required"` // 操作日志ID
}

// GetSystemConfigResp 获取系统配置响应
type GetSystemConfigResp struct {
	Config map[string]interface{} `json:"config"` // 系统配置
}

// SetSystemConfigReq 设置系统配置请求
type SetSystemConfigReq struct {
	Config map[string]interface{} `json:"config" binding:"required"` // 系统配置
}

// GetServerInfoResp 获取服务器信息响应
type GetServerInfoResp struct {
	CPU    float64 `json:"cpu"`    // CPU使用率
	Memory float64 `json:"memory"` // 内存使用率
	Disk   float64 `json:"disk"`   // 磁盘使用率
	OS     string  `json:"os"`     // 操作系统
	Arch   string  `json:"arch"`   // 系统架构
	Go     string  `json:"go"`     // Go版本
}

// UploadFileResp 上传文件响应
type UploadFileResp struct {
	URL      string `json:"url"`      // 文件访问URL
	Name     string `json:"name"`     // 文件名
	Size     int64  `json:"size"`     // 文件大小
	Type     string `json:"type"`     // 文件类型
	FilePath string `json:"filePath"` // 文件存储路径
}

// GetFileListReq 获取文件列表请求
type GetFileListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Name     string `form:"name" json:"name"`         // 文件名搜索
	Type     string `form:"type" json:"type"`         // 文件类型搜索
}

// FileInfo 文件信息
type FileInfo struct {
	ID         uint      `json:"id"`         // 文件ID
	Name       string    `json:"name"`       // 文件名
	Size       int64     `json:"size"`       // 文件大小
	Type       string    `json:"type"`       // 文件类型
	URL        string    `json:"url"`        // 文件访问URL
	FilePath   string    `json:"filePath"`   // 文件存储路径
	Uploader   string    `json:"uploader"`   // 上传者
	UploadTime time.Time `json:"uploadTime"` // 上传时间
}

// GetFileListResp 获取文件列表响应
type GetFileListResp struct {
	Total int64       `json:"total"` // 总数
	List  []*FileInfo `json:"list"`  // 文件列表
}

// DeleteFileReq 删除文件请求
type DeleteFileReq struct {
	ID uint `json:"id" binding:"required"` // 文件ID
}

// CaptchaResp 验证码响应
type CaptchaResp struct {
	CaptchaId string `json:"captchaId"` // 验证码ID
	PicPath   string `json:"picPath"`   // base64图片内容
}

// MfaStatusResp MFA状态响应
type MfaStatusResp struct {
	Bound   bool   `json:"bound"`   // 是否已绑定MFA
	Secret  string `json:"secret"`  // MFA密钥(未绑定时返回新生成的)
	QrURI   string `json:"qrUri"`   // 二维码URI
	SystemEnabled bool `json:"systemEnabled"` // 系统MFA开关
}

// MfaBindReq 绑定MFA请求
type MfaBindReq struct {
	Code string `json:"code" binding:"required"` // 验证码
}

// MfaVerifyReq 验证MFA请求
type MfaVerifyReq struct {
	Username string `json:"username" binding:"required"` // 用户名
	Code     string `json:"code" binding:"required"`     // MFA验证码
	MfaToken string `json:"mfaToken" binding:"required"` // 临时MFA token
}

// MfaUnbindReq 解绑MFA请求
type MfaUnbindReq struct {
	Code string `json:"code" binding:"required"` // 当前验证码
}

// MfaRequiredResp MFA认证要求响应(登录时MFA开关开启且用户已绑定MFA)
type MfaRequiredResp struct {
	MfaRequired bool   `json:"mfaRequired"` // 是否需要MFA验证
	MfaToken    string `json:"mfaToken"`    // 临时token用于MFA验证
}

// ========== 字典管理 ==========

// GetSysDictionaryListReq 获取字典列表请求
type GetSysDictionaryListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Name     string `form:"name" json:"name"`         // 字典名称搜索
	Type     string `form:"type" json:"type"`         // 字典类型搜索
	Status   *bool  `form:"status" json:"status"`     // 状态搜索
}

// GetSysDictionaryListResp 获取字典列表响应
type GetSysDictionaryListResp struct {
	Total int64                 `json:"total"` // 总数
	List  []*dbr.SysDictionary `json:"list"`  // 字典列表
}

// ========== 参数管理 ==========

// GetSysParamsListReq 获取参数列表请求
type GetSysParamsListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Name     string `form:"name" json:"name"`         // 参数名称搜索
	Key      string `form:"key" json:"key"`           // 参数键搜索
	Status   *bool  `form:"status" json:"status"`     // 状态搜索
}

// GetSysParamsListResp 获取参数列表响应
type GetSysParamsListResp struct {
	Total int64            `json:"total"` // 总数
	List  []*dbr.SysParams `json:"list"`  // 参数列表
}

// ========== 登录日志管理 ==========

// GetLoginLogListReq 获取登录日志列表请求
type GetLoginLogListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Username string `form:"username" json:"username"` // 用户名搜索
	Status   *bool  `form:"status" json:"status"`     // 状态搜索
	Ip       string `form:"ip" json:"ip"`             // IP搜索
}

// GetLoginLogListResp 获取登录日志列表响应
type GetLoginLogListResp struct {
	Total int64               `json:"total"` // 总数
	List  []*dbr.SysLoginLog `json:"list"`  // 登录日志列表
}

// ========== API Token 管理 ==========

// CreateApiTokenReq 创建Token请求
type CreateApiTokenReq struct {
	Name        string    `json:"name" binding:"required"` // Token名称
	UserID      uint      `json:"userId"`                   // 创建用户ID
	Username    string    `json:"username"`                 // 创建用户名
	ExpireTime  time.Time `json:"expireTime"`               // 过期时间
	Permissions string    `json:"permissions"`              // 权限范围JSON
	Remarks     string    `json:"remarks"`                  // 备注
}

// GetApiTokenListReq 获取Token列表请求
type GetApiTokenListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Name     string `form:"name" json:"name"`         // Token名称搜索
	Username string `form:"username" json:"username"` // 用户名搜索
	Status   *int   `form:"status" json:"status"`     // 状态搜索
}

// GetApiTokenListResp 获取Token列表响应
type GetApiTokenListResp struct {
	Total int64               `json:"total"` // 总数
	List  []*dbr.SysApiToken `json:"list"`  // Token列表
}

// DeleteApiTokenReq 删除Token请求
type DeleteApiTokenReq struct {
	ID uint `json:"id" binding:"required"` // TokenID
}

// ========== 错误日志管理 ==========

// UpdateSysErrorReq 更新错误日志请求
type UpdateSysErrorReq struct {
	ID          uint   `json:"id" binding:"required"` // 错误日志ID
	Status      int    `json:"status"`                // 处理状态 1未处理 2已处理 3已忽略
	Solution    string `json:"solution"`              // 解决方案
	HandlerID   uint   `json:"handlerId"`             // 处理人ID
	HandlerName string `json:"handlerName"`           // 处理人名称
}

// GetSysErrorListReq 获取错误日志列表请求
type GetSysErrorListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Type     string `form:"type" json:"type"`         // 错误类型筛选
	Status   *int   `form:"status" json:"status"`     // 处理状态筛选
	Username string `form:"username" json:"username"` // 用户名搜索
}

// GetSysErrorListResp 获取错误日志列表响应
type GetSysErrorListResp struct {
	Total int64             `json:"total"` // 总数
	List  []*dbr.SysError `json:"list"`  // 错误日志列表
}

// GetSysErrorPublicResp 获取公开错误响应
type GetSysErrorPublicResp struct {
	List []*dbr.SysError `json:"list"` // 公开错误列表
}

// GetSysErrorSolutionResp 获取错误解决方案响应
type GetSysErrorSolutionResp struct {
	Solution string `json:"solution"` // 解决方案
}

// ========== 版本管理 ==========

// CreateSysVersionReq 创建版本请求
type CreateSysVersionReq struct {
	Version  string `json:"version" binding:"required"` // 版本号
	Name     string `json:"name" binding:"required"`    // 版本名称
	Content  string `json:"content"`                    // 版本内容说明
	Type     string `json:"type"`                       // 版本类型 1正式版 2测试版
	UserID   uint   `json:"userId"`                     // 发布人ID
	Username string `json:"username"`                   // 发布人
	FileUrl  string `json:"fileUrl"`                    // 下载地址
	FileSize int64  `json:"fileSize"`                   // 文件大小
	Md5      string `json:"md5"`                        // MD5校验值
	Remarks  string `json:"remarks"`                    // 备注
}

// GetSysVersionListReq 获取版本列表请求
type GetSysVersionListReq struct {
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 每页数量
	Version  string `form:"version" json:"version"`   // 版本号搜索
	Name     string `form:"name" json:"name"`         // 版本名称搜索
	Type     string `form:"type" json:"type"`         // 版本类型筛选
	Status   *int   `form:"status" json:"status"`     // 状态筛选
}

// GetSysVersionListResp 获取版本列表响应
type GetSysVersionListResp struct {
	Total int64               `json:"total"` // 总数
	List  []*dbr.SysVersion `json:"list"`  // 版本列表
}

// ExportVersionReq 导出版本请求
type ExportVersionReq struct {
	Ids []uint `json:"ids"` // 要导出的版本ID列表，空则全部导出
}

// ImportVersionReq 导入版本请求
type ImportVersionReq struct {
	Data string `json:"data" binding:"required"` // 版本JSON数据
}
