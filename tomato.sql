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

CREATE TABLE IF NOT EXISTS `busi_task` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '任务 ID',
  `uid` BIGINT NOT NULL COMMENT '任务创建人 ID',

  `task_title` varchar(256) NOT NULL COMMENT '任务标题',
  `task_description` varchar(512) DEFAULT NULL COMMENT '任务描述',

  `priority` SMALLINT DEFAULT 3 COMMENT '重要程度 0-5',
  `urgency` SMALLINT DEFAULT 3 COMMENT '紧急程度 0-5',
  `status` SMALLINT DEFAULT 0 COMMENT '任务状态 0 或者 1',

  `task_score` SMALLINT DEFAULT 3 COMMENT '总结评分 0-5',
  `task_summary` varchar(512) DEFAULT NULL COMMENT '任务总结',

  `begin_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `due_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `finished_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '数据更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  CONSTRAINT `sys_user_foreign_id` FOREIGN KEY (`uid`) REFERENCES `sys_user`(`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
