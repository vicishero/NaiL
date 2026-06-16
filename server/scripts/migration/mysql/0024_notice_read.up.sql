-- 系统消息已读记录表
-- 用户点开消息时才产生记录，通过 LEFT JOIN 判断是否已读
CREATE TABLE `p_notice_read` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `msg_id` BIGINT NOT NULL COMMENT '关联p_notice.id',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `read_time` DATETIME NOT NULL DEFAULT NOW() COMMENT '阅读时间',
    UNIQUE KEY `uk_msg_user` (`msg_id`, `user_id`),
    INDEX `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统消息阅读记录';
