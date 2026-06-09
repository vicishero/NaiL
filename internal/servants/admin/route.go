// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	adminDao "github.com/rocboss/paopao-ce/internal/dao/admin"
	adminService "github.com/rocboss/paopao-ce/internal/service/admin"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

// RegisterAdminRoutes 注册管理后台路由
func RegisterAdminRoutes(e *gin.Engine, basePath string) {
	// 初始化服务
	db := conf.MustGormDB()
	aDao := adminDao.NewAdminDao(db)
	aService := adminService.NewAdminService(aDao)
	s := NewAuthServant(aService, base.NewDaoServant())

	adminGroup := e.Group(basePath)

	// ====== 公开接口，不需要认证 ======
	baseGroup := adminGroup.Group("/base")
	{
		baseGroup.POST("/captcha", s.Captcha)
		baseGroup.POST("/login", s.Login)
	}
	// JWT黑名单接口（退出登录用，无需认证）
	adminGroup.POST("/jwt/jsonInBlacklist", s.JsonInBlacklist)

	// ====== 需要JWT认证的接口 ======
	authGroup := adminGroup.Group("")
	authGroup.Use(s.JWTMiddleware())
	{
		// ---- 用户基础操作 ----
		authGroup.POST("/user/logout", s.Logout)
		authGroup.GET("/user/getUserInfo", s.GetUserInfo)
		authGroup.POST("/user/changePassword", s.ChangePassword)

		// ---- 菜单（auth级别，登录即可获取） ----
		registerGetPost(authGroup, "/menu/getMenuList", s.GetMenuList)
		authGroup.POST("/menu/getMenu", s.GetMenu)

		// ---- 通用GVA兼容接口（认证即可，无权限校验） ----
		registerGvaAuthRoutes(authGroup, s)

		// ---- H5运维管理后台（JWT认证即可，运维管理功能） ----
		h5Group := authGroup.Group("/h5Admin")
		{
			// 用户管理
			h5Group.GET("/userList", s.GetH5UserList)
			h5Group.GET("/user", s.GetH5User)
			h5Group.PUT("/user", s.UpdateH5User)
			h5Group.DELETE("/user", s.DeleteH5User)
			// 贴文管理
			h5Group.GET("/postList", s.GetH5PostList)
			h5Group.GET("/post", s.GetH5Post)
			h5Group.PUT("/post", s.UpdateH5Post)
			h5Group.DELETE("/post", s.DeleteH5Post)
			h5Group.POST("/syncIndex", s.SyncSearchIndex)
			// 话题管理
			h5Group.GET("/tagList", s.GetH5TagList)
			h5Group.PUT("/tag", s.UpdateH5Tag)
			h5Group.DELETE("/tag", s.DeleteH5Tag)
			h5Group.GET("/commentList", s.GetH5CommentList)
			h5Group.DELETE("/comment", s.DeleteH5Comment)
			h5Group.GET("/collectionList", GvaEmptyList)
			h5Group.GET("/collection", GvaEmptyObject)
			h5Group.DELETE("/collection", GvaSuccess)
			h5Group.GET("/followingList", GvaEmptyList)
			h5Group.GET("/following", GvaEmptyObject)
			h5Group.DELETE("/following", GvaSuccess)
		}
	}

	// ====== 需要权限校验的接口 ======
	permGroup := authGroup.Group("")
	permGroup.Use(s.PermissionMiddleware(), s.OperationLogMiddleware())
	{
		registerUserRoutes(permGroup, s)
		registerAuthorityRoutes(permGroup, s)
		registerMenuRoutes(permGroup, s)
		registerApiRoutes(permGroup, s)
		registerOperationLogRoutes(permGroup, s)
		registerSystemRoutes(permGroup, s)
		registerFileRoutes(permGroup, s)
		registerGvaPermissionRoutes(permGroup, s)
	}
}

func registerUserRoutes(g *gin.RouterGroup, s *AuthServant) {
	userGroup := g.Group("/user")
	{
		registerGetPost(userGroup, "/getUserList", s.GetUserList)
		userGroup.POST("/admin_register", s.CreateUser)
		userGroup.POST("/setUserInfo", s.UpdateUser)
		userGroup.PUT("/setUserInfo", s.UpdateUser)
		userGroup.POST("/setUserAuthority", s.SetUserAuthority)
		userGroup.POST("/setUserAuthorities", s.SetUserAuthorities)
		userGroup.DELETE("/deleteUser", s.DeleteUser)
		userGroup.POST("/resetPassword", s.ResetPassword)
		userGroup.PUT("/setSelfInfo", s.SetSelfInfo)
		userGroup.PUT("/setSelfSetting", s.SetSelfSetting)
	}
}

func registerAuthorityRoutes(g *gin.RouterGroup, s *AuthServant) {
	authorityGroup := g.Group("/authority")
	{
		authorityGroup.POST("/createAuthority", s.CreateAuthority)
		authorityGroup.POST("/updateAuthority", s.UpdateAuthority)
		authorityGroup.POST("/deleteAuthority", s.DeleteAuthority)
		registerGetPost(authorityGroup, "/getAuthorityList", s.GetAuthorityList)
		registerGetPost(authorityGroup, "/getAllAuthorities", s.GetAllAuthorities)
		authorityGroup.POST("/setAuthorityMenu", s.SetAuthorityMenu)
		registerGetPost(authorityGroup, "/getAuthorityMenu", s.GetAuthorityMenu)
		authorityGroup.POST("/setAuthorityApi", s.SetAuthorityApi)
		registerGetPost(authorityGroup, "/getAuthorityApi", s.GetAuthorityApi)
		// 补充接口
		authorityGroup.POST("/copyAuthority", s.CopyAuthority)
		registerGetPost(authorityGroup, "/getUsersByAuthority", s.GetUsersByAuthority)
		authorityGroup.POST("/setDataAuthority", s.SetDataAuthority)
		authorityGroup.POST("/setRoleUsers", s.SetRoleUsers)
	}
}

func registerMenuRoutes(g *gin.RouterGroup, s *AuthServant) {
	menuGroup := g.Group("/menu")
	{
		menuGroup.POST("/addBaseMenu", s.CreateBaseMenu)
		menuGroup.POST("/updateBaseMenu", s.UpdateBaseMenu)
		menuGroup.POST("/deleteBaseMenu", s.DeleteBaseMenu)
		registerGetPost(menuGroup, "/getBaseMenuTree", s.GetBaseMenuTree)
		menuGroup.POST("/getBaseMenuById", s.GetBaseMenuById)
		menuGroup.POST("/addMenuAuthority", s.AddMenuAuthority)
		menuGroup.POST("/getMenuAuthority", s.GetMenuAuthority)
		registerGetPost(menuGroup, "/getMenuRoles", s.GetMenuRoles)
		menuGroup.POST("/setMenuRoles", s.SetMenuRoles)
	}
}

func registerApiRoutes(g *gin.RouterGroup, s *AuthServant) {
	apiGroup := g.Group("/api")
	{
		apiGroup.POST("/createApi", s.CreateApi)
		apiGroup.POST("/updateApi", s.UpdateApi)
		apiGroup.POST("/deleteApi", s.DeleteApi)
		apiGroup.POST("/deleteApisByIds", s.DeleteApisByIds)
		registerGetPost(apiGroup, "/getApiList", s.GetApiList)
		registerGetPost(apiGroup, "/getAllApiGroups", s.GetAllApiGroups)
		apiGroup.POST("/syncApi", s.SyncApi)
		apiGroup.POST("/enterSyncApi", s.EnterSyncApi)
		apiGroup.POST("/freshCasbin", s.FreshCasbin)
		registerGetPost(apiGroup, "/getAllApis", s.GetAllApis)
		registerGetPost(apiGroup, "/getApiById", s.GetApiById)
		registerGetPost(apiGroup, "/getApiRoles", s.GetApiRoles)
		apiGroup.POST("/ignoreApi", s.IgnoreApi)
		apiGroup.POST("/setApiRoles", s.SetApiRoles)
		apiGroup.POST("/setAuthApi", s.SetAuthApi)
	}
}

func registerOperationLogRoutes(g *gin.RouterGroup, s *AuthServant) {
	opGroup := g.Group("/sysOperationRecord")
	{
		registerGetPost(opGroup, "/getSysOperationRecordList", s.GetOperationLogList)
		opGroup.POST("/deleteSysOperationRecord", s.DeleteOperationLog)
		opGroup.POST("/deleteSysOperationRecordByIds", s.DeleteSysOperationRecordByIds)
	}
}

func registerSystemRoutes(g *gin.RouterGroup, s *AuthServant) {
	sysGroup := g.Group("/system")
	{
		registerGetPost(sysGroup, "/getSystemConfig", s.GetSystemConfig)
		sysGroup.POST("/setSystemConfig", s.SetSystemConfig)
		registerGetPost(sysGroup, "/getServerInfo", s.GetServerInfo)
		sysGroup.POST("/reloadSystem", s.ReloadSystem)
	}
}

func registerFileRoutes(g *gin.RouterGroup, s *AuthServant) {
	fileGroup := g.Group("/fileUploadAndDownload")
	{
		fileGroup.POST("/upload", s.UploadFile)
		registerGetPost(fileGroup, "/getFileList", s.GetFileList)
		fileGroup.POST("/deleteFile", s.DeleteFile)
		fileGroup.POST("/breakpointContinue", s.BreakpointContinue)
		fileGroup.POST("/breakpointContinueFinish", s.BreakpointContinueFinish)
		fileGroup.POST("/editFileName", s.EditFileName)
		fileGroup.GET("/findFile", s.FindFile)
		fileGroup.POST("/importURL", s.ImportURL)
		fileGroup.POST("/removeChunk", s.RemoveChunk)
	}
}

// registerGvaAuthRoutes 需要认证但不需要权限校验的GVA兼容接口
func registerGvaAuthRoutes(g *gin.RouterGroup, s *AuthServant) {
	// Casbin
	registerGetPost(g, "/casbin/getPolicyPathByAuthorityId", s.GetPolicyPathByAuthorityId)
	g.POST("/casbin/updateCasbin", s.UpdateCasbin)

}

// registerGvaPermissionRoutes 需要权限校验的GVA兼容通用接口
func registerGvaPermissionRoutes(g *gin.RouterGroup, s *AuthServant) {
	// ---- 数据字典 ----
	sysDictGroup := g.Group("/sysDictionary")
	{
		sysDictGroup.POST("/createSysDictionary", GvaSuccess)
		sysDictGroup.POST("/deleteSysDictionary", GvaSuccess)
		sysDictGroup.POST("/updateSysDictionary", GvaSuccess)
		sysDictGroup.GET("/getSysDictionaryList", GvaEmptyList)
		sysDictGroup.GET("/findSysDictionary", GvaEmptyObject)
		sysDictGroup.POST("/importSysDictionary", GvaSuccess)
		sysDictGroup.POST("/exportSysDictionary", GvaSuccess)
	}

	sysDictDetailGroup := g.Group("/sysDictionaryDetail")
	{
		sysDictDetailGroup.POST("/createSysDictionaryDetail", GvaSuccess)
		sysDictDetailGroup.POST("/deleteSysDictionaryDetail", GvaSuccess)
		sysDictDetailGroup.POST("/updateSysDictionaryDetail", GvaSuccess)
		sysDictDetailGroup.GET("/getSysDictionaryDetailList", GvaEmptyList)
		sysDictDetailGroup.GET("/findSysDictionaryDetail", GvaEmptyObject)
		sysDictDetailGroup.GET("/getDictionaryDetailsByParent", GvaEmptyList)
		sysDictDetailGroup.GET("/getDictionaryPath", GvaEmptyList)
		sysDictDetailGroup.GET("/getDictionaryTreeList", GvaEmptyList)
		sysDictDetailGroup.GET("/getDictionaryTreeListByType", GvaEmptyList)
	}

	// ---- 系统参数 ----
	sysParamsGroup := g.Group("/sysParams")
	{
		sysParamsGroup.POST("/createSysParams", GvaSuccess)
		sysParamsGroup.POST("/deleteSysParams", GvaSuccess)
		sysParamsGroup.POST("/deleteSysParamsByIds", GvaSuccess)
		sysParamsGroup.POST("/updateSysParams", GvaSuccess)
		sysParamsGroup.GET("/getSysParamsList", GvaEmptyList)
		sysParamsGroup.GET("/findSysParams", GvaEmptyObject)
		sysParamsGroup.GET("/getSysParam", GvaEmptyObject)
	}

	// ---- API Token ----
	sysApiTokenGroup := g.Group("/sysApiToken")
	{
		sysApiTokenGroup.POST("/createApiToken", GvaSuccess)
		sysApiTokenGroup.POST("/deleteApiToken", GvaSuccess)
		sysApiTokenGroup.GET("/getApiTokenList", GvaEmptyList)
	}

	// ---- 登录日志 ----
	sysLoginLogGroup := g.Group("/sysLoginLog")
	{
		sysLoginLogGroup.POST("/deleteLoginLog", GvaSuccess)
		sysLoginLogGroup.POST("/deleteLoginLogByIds", GvaSuccess)
		sysLoginLogGroup.GET("/getLoginLogList", GvaEmptyList)
		sysLoginLogGroup.GET("/findLoginLog", GvaEmptyObject)
	}

	// ---- 导出模板 ----
	sysExportGroup := g.Group("/sysExportTemplate")
	{
		sysExportGroup.POST("/createSysExportTemplate", GvaSuccess)
		sysExportGroup.POST("/deleteSysExportTemplate", GvaSuccess)
		sysExportGroup.POST("/deleteSysExportTemplateByIds", GvaSuccess)
		sysExportGroup.POST("/updateSysExportTemplate", GvaSuccess)
		sysExportGroup.GET("/getSysExportTemplateList", GvaEmptyList)
		sysExportGroup.GET("/findSysExportTemplate", GvaEmptyObject)
		sysExportGroup.POST("/exportExcel", GvaSuccess)
		sysExportGroup.POST("/exportTemplate", GvaSuccess)
		sysExportGroup.POST("/previewSQL", GvaSuccess)
	}

	// ---- 版本管理 ----
	sysVersionGroup := g.Group("/sysVersion")
	{
		sysVersionGroup.POST("/deleteSysVersion", GvaSuccess)
		sysVersionGroup.POST("/deleteSysVersionByIds", GvaSuccess)
		sysVersionGroup.GET("/getSysVersionList", GvaEmptyList)
		sysVersionGroup.GET("/findSysVersion", GvaEmptyObject)
		sysVersionGroup.POST("/exportVersion", GvaSuccess)
		sysVersionGroup.POST("/importVersion", GvaSuccess)
		sysVersionGroup.GET("/downloadVersionJson", GvaSuccess)
	}

	// ---- 系统错误记录 ----
	sysErrorGroup := g.Group("/sysError")
	{
		sysErrorGroup.POST("/createSysError", GvaSuccess)
		sysErrorGroup.POST("/deleteSysError", GvaSuccess)
		sysErrorGroup.POST("/deleteSysErrorByIds", GvaSuccess)
		sysErrorGroup.POST("/updateSysError", GvaSuccess)
		sysErrorGroup.GET("/getSysErrorList", GvaEmptyList)
		sysErrorGroup.GET("/findSysError", GvaEmptyObject)
		sysErrorGroup.GET("/getSysErrorPublic", GvaEmptyObject)
		sysErrorGroup.GET("/getSysErrorSolution", GvaEmptyObject)
	}

	// ---- 代码生成器 ----
	autoCodeGroup := g.Group("/autoCode")
	{
		autoCodeGroup.POST("/createPackage", GvaSuccess)
		autoCodeGroup.POST("/delPackage", GvaSuccess)
		autoCodeGroup.GET("/getPackage", GvaEmptyList)
		autoCodeGroup.POST("/createTemp", GvaSuccess)
		autoCodeGroup.GET("/getTemplates", GvaEmptyList)
		autoCodeGroup.GET("/getDB", GvaEmptyObject)
		autoCodeGroup.GET("/getTables", GvaEmptyList)
		autoCodeGroup.GET("/getColumn", GvaEmptyObject)
		autoCodeGroup.GET("/getMeta", GvaEmptyObject)
		autoCodeGroup.POST("/preview", GvaSuccess)
		autoCodeGroup.GET("/getSysHistory", GvaEmptyList)
		autoCodeGroup.POST("/delSysHistory", GvaSuccess)
		autoCodeGroup.POST("/rollback", GvaSuccess)
		autoCodeGroup.POST("/addFunc", GvaSuccess)
		autoCodeGroup.POST("/initMenu", GvaSuccess)
		autoCodeGroup.POST("/initAPI", GvaSuccess)
		autoCodeGroup.POST("/initDictionary", GvaSuccess)
		autoCodeGroup.GET("/getPluginList", GvaEmptyList)
		autoCodeGroup.POST("/installPlug", GvaSuccess)
		autoCodeGroup.POST("/pubPlug", GvaSuccess)
		autoCodeGroup.POST("/removePlugin", GvaSuccess)
		autoCodeGroup.POST("/mcp", GvaSuccess)
		autoCodeGroup.GET("/mcpList", GvaEmptyList)
		autoCodeGroup.POST("/mcpStart", GvaSuccess)
		autoCodeGroup.POST("/mcpStop", GvaSuccess)
		autoCodeGroup.GET("/mcpStatus", GvaEmptyObject)
		autoCodeGroup.POST("/mcpTest", GvaSuccess)
		autoCodeGroup.POST("/deleteAIWorkflowSession", GvaSuccess)
		autoCodeGroup.GET("/getAIWorkflowSessionList", GvaEmptyList)
		autoCodeGroup.GET("/getAIWorkflowSessionDetail", GvaEmptyObject)
		autoCodeGroup.POST("/dumpAIWorkflowMarkdown", GvaSuccess)
		autoCodeGroup.POST("/saveAIWorkflowSession", GvaSuccess)
	}

	// ---- 邮件测试 ----
	g.POST("/email/emailTest", GvaSuccess)

	// ---- 客户管理 ----
	g.GET("/customer/customerList", GvaEmptyList)
	g.POST("/customer/customer", GvaSuccess)

	// ---- BSC区块链合约 ----
	bscGroup := g.Group("/bsc")
	{
		bscGroup.POST("/createBscContractConfig", GvaSuccess)
		bscGroup.POST("/deleteBscContractConfig", GvaSuccess)
		bscGroup.POST("/updateBscContractConfig", GvaSuccess)
		bscGroup.GET("/getBscContractConfigList", GvaEmptyList)
		bscGroup.GET("/findBscContractConfig", GvaEmptyObject)
	}
	bscSyncGroup := g.Group("/bscSyncInfo")
	{
		bscSyncGroup.POST("/createBscSyncInfo", GvaSuccess)
		bscSyncGroup.POST("/deleteBscSyncInfo", GvaSuccess)
		bscSyncGroup.POST("/updateBscSyncInfo", GvaSuccess)
		bscSyncGroup.GET("/getBscSyncInfoList", GvaEmptyList)
		bscSyncGroup.GET("/findBscSyncInfo", GvaEmptyObject)
	}
	bscEventGroup := g.Group("/bscSyncedEvent")
	{
		bscEventGroup.GET("/getBscSyncedEventList", GvaEmptyList)
		bscEventGroup.POST("/updateStatus", GvaSuccess)
	}

	// ---- 技能管理 ----
	skillsGroup := g.Group("/skills")
	{
		skillsGroup.POST("/createScript", GvaSuccess)
		skillsGroup.POST("/createTemplate", GvaSuccess)
		skillsGroup.POST("/createReference", GvaSuccess)
		skillsGroup.POST("/createResource", GvaSuccess)
		skillsGroup.POST("/saveSkill", GvaSuccess)
		skillsGroup.POST("/saveScript", GvaSuccess)
		skillsGroup.POST("/saveTemplate", GvaSuccess)
		skillsGroup.POST("/saveReference", GvaSuccess)
		skillsGroup.POST("/saveResource", GvaSuccess)
		skillsGroup.POST("/saveGlobalConstraint", GvaSuccess)
		skillsGroup.POST("/deleteSkill", GvaSuccess)
		skillsGroup.POST("/packageSkill", GvaSuccess)
		skillsGroup.POST("/downloadOnlineSkill", GvaSuccess)
		skillsGroup.GET("/getSkillList", GvaEmptyList)
		skillsGroup.GET("/getSkillDetail", GvaEmptyObject)
		skillsGroup.GET("/getScript", GvaEmptyObject)
		skillsGroup.GET("/getTemplate", GvaEmptyObject)
		skillsGroup.GET("/getReference", GvaEmptyObject)
		skillsGroup.GET("/getResource", GvaEmptyObject)
		skillsGroup.GET("/getGlobalConstraint", GvaEmptyObject)
		skillsGroup.GET("/getTools", GvaEmptyList)
	}

	// ---- 附件分类 ----
	g.POST("/attachmentCategory/addCategory", GvaSuccess)
	g.POST("/attachmentCategory/deleteCategory", GvaSuccess)
	g.GET("/attachmentCategory/getCategoryList", GvaEmptyList)

	// ---- 权限按钮 ----
	g.POST("/authorityBtn/setAuthorityBtn", GvaSuccess)
	g.GET("/authorityBtn/getAuthorityBtn", GvaEmptyObject)
	g.POST("/authorityBtn/canRemoveAuthorityBtn", GvaSuccess)

	// ---- 商城插件 ----
	g.GET("/shopPlugin/getShopPluginList", GvaEmptyList)
}

// registerGetPost 同时注册GET和POST，兼容GVA前端POST请求习惯
func registerGetPost(g *gin.RouterGroup, path string, handler gin.HandlerFunc) {
	g.GET(path, handler)
	g.POST(path, handler)
}
