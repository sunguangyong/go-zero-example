CREATE TABLE `admin_list` (
                              `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `uid` BIGINT NOT NULL DEFAULT '0' COMMENT '币虎用户id',
                              `chat_id` BIGINT NOT NULL DEFAULT '0' COMMENT 'tg_id',
                              `group_id` BIGINT NOT NULL DEFAULT '0' COMMENT '群id',
                              `group_name` VARCHAR ( 255 ) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '群名称',
                              `permission` VARCHAR ( 50 ) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '权限',
                              `user_name` VARCHAR ( 255 ) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'tg user_name',
                              `first_name` VARCHAR ( 255 ) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'tg first_name',
                              `last_name` VARCHAR ( 255 ) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'tg last_name',
                              `is_delete` INT NOT NULL DEFAULT '0' COMMENT '删除标识 0 未删除 1 已删除',
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                              `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
                              PRIMARY KEY ( `id` ),
                              KEY `idx_chat_id` ( `chat_id` )
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT '管理员表';