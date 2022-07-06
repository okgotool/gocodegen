CREATE TABLE IF NOT EXISTS  `test_cat` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cat_name` varchar(256) DEFAULT NULL COMMENT '名称',
  `created_by` varchar(128) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_by` varchar(128) DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试表';

