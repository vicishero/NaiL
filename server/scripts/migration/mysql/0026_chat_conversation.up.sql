-- 用户聊天会话映射表
CREATE TABLE `p_chat_conversation` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `kol_user_id` BIGINT NOT NULL DEFAULT 0 COMMENT 'KOL用户ID, 0=默认助手',
    `dify_conversation_id` VARCHAR(64) NOT NULL DEFAULT '' COMMENT 'Dify会话ID',
    `title` VARCHAR(128) DEFAULT '' COMMENT '会话标题',
    `created_on` BIGINT NOT NULL,
    `modified_on` BIGINT NOT NULL,
    `deleted_on` BIGINT NOT NULL DEFAULT 0,
    `is_del` TINYINT NOT NULL DEFAULT 0,
    UNIQUE KEY `uk_user_kol` (`user_id`, `kol_user_id`),
    INDEX `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户聊天会话映射';
