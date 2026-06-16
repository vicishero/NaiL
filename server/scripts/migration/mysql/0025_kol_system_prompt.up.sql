-- KOL 人物属性增加系统提示词字段
ALTER TABLE `p_kol_profile` ADD COLUMN `system_prompt` TEXT COMMENT '系统提示词' AFTER `category_id`;

-- p_user 增加聊天开关
ALTER TABLE `p_user` ADD COLUMN `chat_enabled` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否启用AI聊天' AFTER `is_kol`;
