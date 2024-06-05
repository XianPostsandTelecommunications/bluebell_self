CREATE TABLE `bad_words` (
                             `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
                             `content` text COMMENT '敏感词内容',
                             `create_user_id`bigint(20) unsigned COMMENT '创建用户id',
                             `extra` text COMMENT '扩展信息',
                             `create_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
                             `modify_time` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
                             `status` int DEFAULT 0 COMMENT '0存在，1删除',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8 COLLATE=utf8_general_ci;
