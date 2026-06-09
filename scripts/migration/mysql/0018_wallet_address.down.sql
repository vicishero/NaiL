-- 回滚钱包地址字段
ALTER TABLE `p_users` DROP INDEX `idx_address`;
ALTER TABLE `p_users` DROP COLUMN `address`;
