-- 将 p_notice 数据迁回 p_message
INSERT INTO `p_message` SELECT * FROM `p_notice`;
DROP TABLE IF EXISTS `p_notice`;
