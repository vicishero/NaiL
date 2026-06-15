-- KOL管理相关API注册到权限管理
INSERT IGNORE INTO `sys_apis` (`path`, `description`, `api_group`, `method`, `created_at`, `updated_at`) VALUES
('/api/v1/h5Admin/kolCategory', '保存KOL分类', 'KOL管理', 'PUT', NOW(), NOW()),
('/api/v1/h5Admin/kolCategory', '删除KOL分类', 'KOL管理', 'DELETE', NOW(), NOW()),
('/api/v1/h5Admin/kolManageList', 'KOL用户管理列表', 'KOL管理', 'GET', NOW(), NOW()),
('/api/v1/h5Admin/kolAssignCategory', '分配KOL分类', 'KOL管理', 'PUT', NOW(), NOW());
