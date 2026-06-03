# Gin-Vue-Admin 管理后台API接口列表

本文档整理了admin前端项目中所有用到的API接口，后端开发可以按照此列表实现对应的服务端接口。

---

## 🔐 基础认证模块

### 1. 用户相关 (user.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/user/login` | 用户登录 |
| POST | `/user/register` | 用户注册 |
| POST | `/user/change_password` | 修改密码 |
| POST | `/user/set_user_authority` | 设置用户角色 |
| POST | `/user/set_user_info` | 修改用户信息 |
| POST | `/user/set_self_info` | 修改个人信息 |
| GET | `/user/user_list` | 获取用户列表 |
| DELETE | `/user/delete_user` | 删除用户 |
| GET | `/user/get_user_info` | 获取登录用户信息 |

### 2. JWT相关 (jwt.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/jwt/jsonInBlacklist` | 退出登录（JWT加入黑名单） |

---

## 🎛️ 权限管理模块

### 1. 角色管理 (authority.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/authority/createAuthority` | 创建角色 |
| POST | `/authority/copyAuthority` | 复制角色 |
| POST | `/authority/deleteAuthority` | 删除角色 |
| POST | `/authority/updateAuthority` | 更新角色信息 |
| GET | `/authority/getAuthorityList` | 获取角色列表 |
| POST | `/authority/setDataAuthority` | 设置角色数据权限 |
| POST | `/authority/getPolicyPathByAuthorityId` | 获取角色关联的API路径 |

### 2. 按钮权限管理 (authorityBtn.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/authorityBtn/createAuthorityBtn` | 创建按钮权限 |
| POST | `/authorityBtn/deleteAuthorityBtn` | 删除按钮权限 |
| POST | `/authorityBtn/updateAuthorityBtn` | 更新按钮权限 |
| GET | `/authorityBtn/getAuthorityBtn` | 获取按钮权限列表 |
| GET | `/authorityBtn/canRemoveAuthorityBtn` | 检查是否可删除按钮权限 |

### 3. 菜单管理 (menu.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/menu/addBaseMenu` | 添加基础菜单 |
| POST | `/menu/deleteBaseMenu` | 删除基础菜单 |
| POST | `/menu/updateBaseMenu` | 更新基础菜单 |
| GET | `/menu/getMenuList` | 获取动态菜单 |
| GET | `/menu/getBaseMenuTree` | 获取基础菜单树 |
| POST | `/menu/addMenuAuthority` | 添加菜单权限关联 |
| GET | `/menu/getMenuAuthority` | 获取菜单权限关联 |
| POST | `/menu/getUserMenu` | 获取用户有权限的菜单 |

### 4. 接口管理 (api.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/api/createApi` | 创建接口 |
| POST | `/api/deleteApi` | 删除接口 |
| POST | `/api/updateApi` | 更新接口 |
| GET | `/api/getApiList` | 获取接口列表 |
| GET | `/api/getAllApiGroups` | 获取所有API分组 |
| GET | `/api/getApiById` | 根据ID获取接口详情 |
| GET | `/api/getApiPaths` | 获取所有API路径 |

### 5. Casbin权限管理 (casbin.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/casbin/updateCasbin` | 更新权限策略 |
| POST | `/casbin/addPolicy` | 添加权限策略 |
| POST | `/casbin/removePolicy` | 移除权限策略 |

### 6. API Token管理 (sysApiToken.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/sysApiToken/updateApiTokenStatus` | 更新API Token状态 |

---

## 🗄️ 系统管理模块

### 1. 系统配置 (system.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| GET | `/system/getServerInfo` | 获取服务器信息 |
| GET | `/system/getSystemConfig` | 获取系统配置 |
| POST | `/system/setSystemConfig` | 设置系统配置 |
| POST | `/system/reloadSystem` | 重载系统配置 |
| GET | `/system/getVersion` | 获取系统版本 |

### 2. 系统参数 (sysParams.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/sysParams/createSysParams` | 创建系统参数 |
| POST | `/sysParams/deleteSysParams` | 删除系统参数 |
| POST | `/sysParams/updateSysParams` | 更新系统参数 |
| GET | `/sysParams/getSysParamsList` | 获取系统参数列表 |
| GET | `/sysParams/getSysParams` | 根据键名获取系统参数 |

### 3. 数据字典 (sysDictionary.js + sysDictionaryDetail.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/sysDictionary/createSysDictionary` | 创建字典 |
| POST | `/sysDictionary/deleteSysDictionary` | 删除字典 |
| POST | `/sysDictionary/updateSysDictionary` | 更新字典 |
| GET | `/sysDictionary/getSysDictionaryList` | 获取字典列表 |
| GET | `/sysDictionary/getSysDictionary` | 获取字典详情 |
| POST | `/sysDictionaryDetail/createSysDictionaryDetail` | 创建字典项 |
| POST | `/sysDictionaryDetail/deleteSysDictionaryDetail` | 删除字典项 |
| POST | `/sysDictionaryDetail/updateSysDictionaryDetail` | 更新字典项 |
| GET | `/sysDictionaryDetail/getSysDictionaryDetailList` | 获取字典项列表 |

### 4. 操作日志 (sysOperationRecord.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/sysOperationRecord/deleteSysOperationRecord` | 删除操作日志 |
| POST | `/sysOperationRecord/deleteSysOperationRecordByIds` | 批量删除操作日志 |
| GET | `/sysOperationRecord/getSysOperationRecordList` | 获取操作日志列表 |

### 5. 登录日志 (sysLoginLog.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/sysLoginLog/deleteSysLoginLog` | 删除登录日志 |
| POST | `/sysLoginLog/deleteSysLoginLogByIds` | 批量删除登录日志 |
| GET | `/sysLoginLog/getSysLoginLogList` | 获取登录日志列表 |

---

## 📁 文件管理模块

### 文件上传下载 (fileUploadAndDownload.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/fileUploadAndDownload/upload` | 上传文件 |
| GET | `/fileUploadAndDownload/download` | 下载文件 |
| GET | `/fileUploadAndDownload/getFileList` | 获取文件列表 |
| POST | `/fileUploadAndDownload/deleteFile` | 删除文件 |
| POST | `/fileUploadAndDownload/editFileName` | 编辑文件名 |
| POST | `/fileUploadAndDownload/breakpointContinue` | 断点续传 |
| POST | `/fileUploadAndDownload/breakpointContinueFinish` | 断点续传完成 |
| POST | `/fileUploadAndDownload/removeChunk` | 删除分片文件 |

### 附件分类 (attachmentCategory.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/attachmentCategory/createAttachmentCategory` | 创建附件分类 |
| POST | `/attachmentCategory/deleteAttachmentCategory` | 删除附件分类 |
| POST | `/attachmentCategory/updateAttachmentCategory` | 更新附件分类 |
| GET | `/attachmentCategory/getAttachmentCategoryList` | 获取附件分类列表 |

---

## 📧 邮件管理模块

| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/email/emailTest` | 发送测试邮件 |
| POST | `/email/sendEmail` | 发送邮件 |

---

## 🛠️ 代码生成器模块 (autoCode.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/autoCode/createTemp` | 生成代码 |
| GET | `/autoCode/getDB` | 获取数据库列表 |
| GET | `/autoCode/getTables` | 获取表列表 |
| GET | `/autoCode/getColumn` | 获取表字段 |
| POST | `/autoCode/createPackage` | 安装插件 |
| POST | `/autoCode/pubPlug` | 发布插件 |
| GET | `/autoCode/pluginList` | 插件列表 |
| POST | `/autoCode/installPlugin` | 安装插件 |
| GET | `/autoCode/mcpServers` | 获取MCP服务列表 |
| POST | `/autoCode/mcpAddServer` | 添加MCP服务 |
| POST | `/autoCode/mcpTestServer` | 测试MCP服务 |
| POST | `/autoCode/mcpGetTools` | 获取MCP工具列表 |
| POST | `/autoCode/mcpCallTool` | 调用MCP工具 |

### AI技能管理 (skills.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/skill/createSkill` | 创建AI技能 |
| POST | `/skill/deleteSkill` | 删除AI技能 |
| POST | `/skill/updateSkill` | 更新AI技能 |
| GET | `/skill/getSkillList` | 获取AI技能列表 |
| GET | `/skill/getSkillById` | 获取技能详情 |
| POST | `/skill/callSkill` | 调用AI技能 |

---

## 👥 客户管理模块 (customer.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| POST | `/customer/createCustomer` | 创建客户 |
| POST | `/customer/deleteCustomer` | 删除客户 |
| POST | `/customer/updateCustomer` | 更新客户信息 |
| GET | `/customer/getCustomerList` | 获取客户列表 |
| GET | `/customer/getCustomer` | 获取客户详情 |

---

## 📱 H5管理模块 (h5Admin.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| GET | `/h5/user/list` | 获取H5用户列表 |
| GET | `/h5/post/list` | 获取帖子列表 |
| POST | `/h5/post/audit` | 审核帖子 |
| GET | `/h5/comment/list` | 获取评论列表 |
| POST | `/h5/comment/audit` | 审核评论 |
| GET | `/h5/tag/list` | 获取标签列表 |
| POST | `/h5/tag/create` | 创建标签 |
| POST | `/h5/tag/update` | 更新标签 |
| POST | `/h5/tag/delete` | 删除标签 |
| GET | `/h5/following/list` | 获取关注列表 |
| GET | `/h5/collection/list` | 获取收藏列表 |
| GET | `/h5/system/state` | 获取系统状态 |

---

## 🔗 三方集成模块
### GitHub集成 (github.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| GET | `/github/githubLogin` | GitHub登录 |

### BSC区块链 (bsc.js)
| 请求方法 | 路径 | 功能说明 |
|---------|------|---------|
| GET | `/bsc/contract/events` | 获取合约事件 |
| POST | `/bsc/contract/sync` | 同步合约事件 |
| GET | `/bsc/sync/info` | 获取同步信息 |
| GET | `/bsc/sync/events` | 获取已同步事件 |

---

## 📌 API通用说明
1. **请求格式**：统一使用JSON格式
2. **认证方式**：请求头中携带`Authorization: Bearer <token>`
3. **响应格式**：
   ```json
   {
     "code": 0,      // 0成功，非0失败
     "msg": "success",// 提示信息
     "data": {},     // 返回数据
     "total": 0      // 列表总数，分页接口返回
   }
   ```
4. **错误码**：
   - 0：成功
   - 400：参数错误
   - 401：未认证/登录失效
   - 403：无权限
   - 404：资源不存在
   - 500：服务器内部错误

---

## 🎯 接口实现优先级建议
### 第一优先级（基础功能必备）
1. 登录、退出接口
2. 用户信息、用户管理接口
3. 角色权限、菜单管理接口
4. 文件上传下载接口

### 第二优先级（系统管理必备）
1. 操作日志、登录日志接口
2. 数据字典、系统参数接口
3. 系统配置接口

### 第三优先级（高级功能）
1. 代码生成器相关接口
2. 邮件通知接口
3. 其他业务相关接口

可以根据实际需求逐步实现，先把基础的权限体系跑通，再实现其他业务功能。
