-- 复制 p_message 表结构为 p_notice
CREATE TABLE `p_notice` LIKE `p_message`;
ALTER TABLE `p_notice` COMMENT='系统通知';

-- 迁移 type=99 的数据到 p_notice
INSERT INTO `p_notice` SELECT * FROM `p_message` WHERE `type` = 99;

-- 删除 p_message 中 type=99 的数据
DELETE FROM `p_message` WHERE `type` = 99;
