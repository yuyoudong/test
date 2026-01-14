CREATE TABLE
    IF NOT EXISTS `demo` (
                                  `id` VARCHAR (36) NOT NULL COMMENT '主键，uuid',
                                  `created_at` DATETIME (3) NULL COMMENT '创建时间',
                                  `created_by_uid` VARCHAR (36) NOT NULL DEFAULT '' COMMENT '创建用户ID',
                                  `updated_at` DATETIME (3) NULL COMMENT '更新时间',
                                  `updated_by_uid` VARCHAR (36) NOT NULL DEFAULT '' COMMENT '更新用户ID',
                                  `deleted_at` DATETIME (3) COMMENT '删除时间(逻辑删除)',
                                  INDEX `idx_deleted_at` (`deleted_at`),
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE = INNODB CHARACTER
    SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'Demo';
