-- DROP DATABASE IF EXISTS tomato_db;
CREATE DATABASE IF NOT EXISTS tomato_db;

USE tomato_db;

CREATE TABLE IF NOT EXISTS `sys_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户 id，int64',
  `username` varchar(32) NOT NULL UNIQUE COMMENT 'username unique',
  `password` varchar(64) NOT NULL COMMENT 'password sha256',
  `mobile` varchar(32) NOT NULL UNIQUE COMMENT 'unique mobile',

  `email` varchar(128) DEFAULT NULL,
  `address` varchar(255) NULL DEFAULT NULL,
  `gender` smallint NULL DEFAULT NULL COMMENT '性别 int8, 0 女性，1 男性',
  `avatar` varchar(255) NULL DEFAULT NULL COMMENT '用户头像地址',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '数据更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `index_user_username`(`username`) USING BTREE COMMENT '对 username 进行索引',
  UNIQUE INDEX `index_user_mobile`(`mobile`) USING BTREE COMMENT '对 mobile 进行索引'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
INSERT INTO `sys_user` (`id`, `username`, `password`, `mobile`, `email`, `address`) VALUES(1, 'admin', 'admin', '13865758743', 'admin@tomato.com', 'Xlab, Tsinghua Street, Beijing');
