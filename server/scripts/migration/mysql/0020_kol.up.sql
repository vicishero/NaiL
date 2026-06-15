-- 用户表增加 is_kol 字段
ALTER TABLE `p_user` ADD COLUMN `is_kol` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否为KOL' AFTER `is_admin`;

-- KOL人物属性表
CREATE TABLE `p_kol_profile` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `height` VARCHAR(16) NOT NULL DEFAULT '160cm' COMMENT '身高',
  `weight` VARCHAR(16) NOT NULL DEFAULT '44kg' COMMENT '体重',
  `measurements` VARCHAR(32) NOT NULL DEFAULT '84/58/84' COMMENT '三围',
  `skin_tone` VARCHAR(32) NOT NULL DEFAULT '冷白病态肌' COMMENT '肤色',
  `eye_color` VARCHAR(16) NOT NULL DEFAULT '酒红' COMMENT '瞳色',
  `orientation` VARCHAR(64) NOT NULL DEFAULT '偏双性恋（情感依赖向）' COMMENT '性向',
  `preferences` VARCHAR(255) NOT NULL DEFAULT '独占欲、暗调氛围、偏执温柔' COMMENT '喜好',
  `favorite_foods` VARCHAR(255) NOT NULL DEFAULT '黑森林、红酒、冷食' COMMENT '喜欢食物',
  `clothing_style` VARCHAR(128) NOT NULL DEFAULT '黑裙、蕾丝、丝带、哥特风' COMMENT '服装风格',
  `makeup_style` VARCHAR(128) NOT NULL DEFAULT '苍白底妆、下垂眼、暗红眼影、冷唇' COMMENT '妆风格',
  `created_on` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `modified_on` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_on` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='KOL人物属性';
