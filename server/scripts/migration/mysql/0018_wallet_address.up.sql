-- 添加钱包地址字段
ALTER TABLE `p_user` ADD COLUMN `address` VARCHAR(64) DEFAULT '' COMMENT '钱包地址' AFTER `is_admin`;

-- 添加钱包地址索引
ALTER TABLE `p_user` ADD INDEX `idx_address` (`address`);
