-- KOL分类表
CREATE TABLE `p_kol_category` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '分类名称',
  `sort` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `created_on` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='KOL分类';

-- 插入默认10个KOL分类
INSERT INTO `p_kol_category` (`name`, `sort`, `created_on`) VALUES
('初恋学生', 1, UNIX_TIMESTAMP()),
('都市御姐', 2, UNIX_TIMESTAMP()),
('温柔少妇', 3, UNIX_TIMESTAMP()),
('运动少女', 4, UNIX_TIMESTAMP()),
('暗黑病娇', 5, UNIX_TIMESTAMP()),
('慵懒纯欲', 6, UNIX_TIMESTAMP()),
('精灵奇幻', 7, UNIX_TIMESTAMP()),
('街头辣妹', 8, UNIX_TIMESTAMP()),
('文艺知性', 9, UNIX_TIMESTAMP()),
('邻家治愈', 10, UNIX_TIMESTAMP());

-- KOL属性表增加分类ID
ALTER TABLE `p_kol_profile` ADD COLUMN `category_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'KOL分类ID' AFTER `user_id`;
