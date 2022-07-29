CREATE TABLE `official_user` (
     `id` bigint(20) unsigned NOT NULL,
     `enterprise_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '企业id',
     `appid` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公众号appid',
     `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户主表id',
     `union_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户unionid',
     `open_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户openid',
     `nick_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
     `sex` int(11) NOT NULL DEFAULT '0' COMMENT '用户的性别，值为1时是男性，值为2时是女性，值为0时是未知',
     `language` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户的语言，简体中文为zh_CN',
     `city` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户所在城市',
     `province` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户所在省份',
     `country` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户所在国家',
     `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户头像',
     `subscribe_time` int(11) NOT NULL DEFAULT '0' COMMENT '用户关注时间，为时间戳',
     `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公众号运营者对粉丝的备注',
     `groupid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户所在的分组ID',
     `subscribe_scene` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '返回用户关注的渠道来源',
     `qr_scene` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '二维码扫码场景',
     `qr_scene_str` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '二维码扫码场景描述',
     `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
     `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
     `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
     PRIMARY KEY (`id`),
     KEY `enterprise_id_index` (`enterprise_id`),
     KEY `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='公众号用户表结构';

CREATE TABLE `official_user_tag` (
  `id` bigint(20) unsigned NOT NULL,
  `enterprise_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '企业id',
  `tagid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '标签id',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  `count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '标签下用户数',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `enterprise_id_tagid_index` (`enterprise_id`,`tagid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='公众号用户标签表';

CREATE TABLE `official_user_tag_relation` (
  `id` bigint(20) unsigned NOT NULL,
  `enterprise_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `official_user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `tag_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `enterprise_id_index` (`enterprise_id`),
  KEY `officical_user_id_index` (`official_user_id`),
  KEY `tag_id_index` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='公众号用户标签';

CREATE TABLE `user_order_statistics` (
    `id` bigint(20) NOT NULL,
    `user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
    `enterprise_id` bigint(20) unsigned NOT NULL DEFAULT '0',
    `pay_order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '已支付订单数',
    `total_fee` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '实付金额总数',
    `refund_order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '退款订单数',
    `refund_fee` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '退款金额数',
    `per_customer_transaction` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客单价',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
    `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY `enterprise_id_index` (`enterprise_id`),
    KEY `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户订单统计';

alter table user_order_statistics
    add `latest_market_time` int(10) unsigned NOT NULL DEFAULT '0' comment '最新营销时间',
    add `market_count` int(10) unsigned NOT NULL DEFAULT '0' comment '营销次数',
    add KEY `pay_order_index` (`pay_order`, `user_id`),
    add KEY `total_fee_index` (`total_fee`, `user_id`),
    add KEY `refund_order_index` (`refund_order`, `user_id`),
    add KEY `refund_fee_index` (`refund_fee`, `user_id`),
    add KEY `transaction_index` (`per_customer_transaction`, `user_id`),
    add KEY `latest_market_time_INDEX` (`latest_market_time`, `user_id`),
    add KEY `market_count_INDEX` (`market_count`, `user_id`);

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL DEFAULT '0',
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0',
  `union_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `nickname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '性别：0-未知 1-男性 2-女性',
  `mobile` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `first_source_type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '首次访问来源：1-商城，2-企业微信，3-公众号',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `enterprise_id_index` (`enterprise_id`),
  KEY `union_id_index` (`union_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户总表';

CREATE TABLE `user_event_config` (
    `id` bigint(20) NOT NULL,
    `event` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '事件key',
    `desc` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '事件描述',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
    `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `user_events_v1` (
  `id` bigint(20) NOT NULL,
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户主表id',
  `event` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户行为事件',
  `object` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户行为对象值',
  `event_desc` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户行为描述',
  `extra` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '附加信息',
  `created_at` int(11) NOT NULL DEFAULT '0',
  `updated_at` int(11) NOT NULL DEFAULT '0',
  `deleted_at` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_unique` (`enterprise_id`,`user_id`,`event`,`event_desc`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户行为事件记录表';

CREATE TABLE `work_external_tag` (
  `id` bigint(20) NOT NULL,
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0',
  `group_id` bigint(20) NOT NULL COMMENT 'group表标签组id',
  `tag_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签id',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1-企业设置, 2-用户自定义',
  `member_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '自定义标签对应员工id',
  `order` int(10) NOT NULL DEFAULT '0' COMMENT '次序值',
  `create_time` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `created_at` int(10) NOT NULL DEFAULT '0',
  `updated_at` int(10) NOT NULL DEFAULT '0',
  `deleted_at` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `enterprise_id_index` (`enterprise_id`),
  KEY `group_id_index` (`group_id`),
  KEY `tag_id_index` (`tag_id`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='外部联系人标签表';

CREATE TABLE `work_external_tag_group` (
  `id` bigint(20) NOT NULL,
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0',
  `group_id` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '企业微信标签组id',
  `group_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签组名称',
  `order` int(10) NOT NULL DEFAULT '0' COMMENT '标签组次序值',
  `create_time` int(10) NOT NULL DEFAULT '0' COMMENT '标签组创建时间',
  `created_at` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(10) NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `enterprise_id_idx` (`enterprise_id`),
  KEY `group_id_index` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='企业微信外部练习人标签组表';


CREATE TABLE `work_external_user_tag` (
  `id` bigint(20) NOT NULL,
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0',
  `relation_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'work_external_user_relation表id',
  `external_tag_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '标签表标签id',
  `created_at` int(10) NOT NULL DEFAULT '0',
  `updated_at` int(10) NOT NULL DEFAULT '0',
  `deleted_at` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `enterprise_id_index` (`enterprise_id`),
  KEY `relation_id_index` (`relation_id`),
  KEY `tag_id_index` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='外部联系人用户标签关系表';

CREATE TABLE `work_group_chat` (
  `id` bigint(20) NOT NULL,
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0',
  `chat_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '企业微信群聊id',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '群类型 1：外部群 2：内部群',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '群聊名称',
  `owner_member_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '群创建人id',
  `notice` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `create_time` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '群状态 1=群存在 2=群已解散',
  `created_at` int(10) NOT NULL DEFAULT '0',
  `updated_at` int(10) NOT NULL DEFAULT '0',
  `deleted_at` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `enterprise_id_chat_id_index` (`enterprise_id`,`chat_id`(191)),
  KEY `owner_member_id_index` (`owner_member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='企业微信群聊表';

CREATE TABLE `work_group_chat_member` (
  `id` bigint(20) NOT NULL,
  `group_chat_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '群聊表id',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '成员类型:1 - 企业成员,2 - 外部联系人',
  `userid` varchar(64) NOT NULL DEFAULT '0' COMMENT '群成员id',
  `member_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '内部成员ID',
  `unionid` varchar(32) NOT NULL DEFAULT '' COMMENT '微信unionid',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '群状态，1=入群 2=退群',
  `join_time` int(10) NOT NULL DEFAULT '0' COMMENT '入群时间',
  `leave_time` int(10) NOT NULL DEFAULT '0' COMMENT '退群时间',
  `join_scene` int(10) NOT NULL DEFAULT '1' COMMENT '入群方式。\n1 - 由成员邀请入群（直接邀请入群）\n2 - 由成员邀请入群（通过邀请链接入群）\n3 - 通过扫描群二维码入群',
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '企业ID',
  `created_at` int(10) NOT NULL DEFAULT '0',
  `updated_at` int(10) NOT NULL DEFAULT '0',
  `deleted_at` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `group_chat_id_index` (`group_chat_id`),
  KEY `type_userid_index` (`type`,`userid`),
  KEY `leave_time_index` (`leave_time`),
  KEY `join_time_index` (`join_time`),
  KEY `status_index` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='企业微信群聊成员表';

CREATE TABLE `work_external_relation_transfer_records` (
  `id` bigint(20) NOT NULL,
  `handover_member_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '原跟进成员ID',
  `handover_member_userid` varchar(64) NOT NULL DEFAULT '' COMMENT '原跟进成员企业微信userid',
  `takeover_member_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '接替成员ID',
  `takeover_member_userid` varchar(64) NOT NULL DEFAULT '' COMMENT '接替成员企业微信userid',
  `transfer_external_work_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '外部联系人主键ID',
  `transfer_external_userid` varchar(255) NOT NULL DEFAULT '' COMMENT '外部联系人企业微信userid',
  `transfer_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '转移状态 0：初始状态 1：=转移成功 2：转移失败',
  `transfer_success_msg` varchar(255) NOT NULL '' COMMENT '转移成功之后的文案',
  `transfer_fail_msg` varchar(255) NOT NULL DEFAULT '' COMMENT '接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限',
  `created_at` int(10) NOT NULL DEFAULT '0',
  `updated_at` int(10) NOT NULL DEFAULT '0',
  `deleted_at` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `handover_member_id_index` (`handover_member_id`),
  KEY `takeover_member_id_index` (`takeover_member_id`),
  KEY `handover_member_userid_index` (`handover_member_userid`),
  KEY `takeover_member_userid_index` (takeover_member_userid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='离职继承转移记录表';

CREATE TABLE `marketing_user_statistics` (
    `id` bigint(20) unsigned NOT NULL,
    `enterprise_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '企业id',
    `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户营销触达总次数',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
    `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `enterprise_id_index` (`enterprise_id`),
    KEY `user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户营销策略统计';

CREATE TABLE `user_group` (
    `id` bigint(20) NOT NULL,
    `enterprise_id` bigint(20) unsigned NOT NULL DEFAULT '0',
    `user_group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'cdp客户分组id',
    `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '客户分组名称',
    `type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '客户分组类型：1-忠诚度模型AIPL,2-RFM模型，3-自定义',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
    `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='客户分组-同步cdp客户分组数据';

CREATE TABLE `tag` (
    `id` bigint(20) NOT NULL,
    `enterprise_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '企业id',
    `tag_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'cdp标签id',
    `tag_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签名称',
    `tag_group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'cdp标签组id',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
    `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签-同步cdp标签数据';

CREATE TABLE `tag_group` (
    `id` bigint(20) NOT NULL,
    `enterprise_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '企业id',
    `group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'cdp标签组id',
    `group_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签组名称',
    `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '标签组类型：1-自动标签，2-手动标签',
    `created_at` int(10) unsigned NOT NULL DEFAULT '0',
    `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
    `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签组-同步cdp标签组数据';

CREATE TABLE `user_category` (
  `id` bigint(20) NOT NULL,
  `tag_or_group_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '标签id or 用户分组id',
  `type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '数据类型:tag，group',
  `enterprise_id` bigint(20) NOT NULL,
  `union_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL DEFAULT '0',
  `updated_at` int(11) NOT NULL DEFAULT '0',
  `deleted_at` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx0` (`tag_or_group_id`,`type`),
  KEY `idx-union-id` (`union_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='客户分类（标签或分组）表';

CREATE TABLE `moment_topics` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '主题名称',
  `publish_time` int(10) NOT NULL DEFAULT '0' COMMENT '发布时间',
  `publish_type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '发布类型 1立即发送 2定时发送',
  `content` varchar(5000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '发布内容',
  `created_member_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '朋友圈创建员工ID',
  `member_count` int(10) NOT NULL DEFAULT '0' COMMENT '创建员工数',
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '企业ID',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `name` (`name`) USING BTREE,
  KEY `enterprise_id` (`enterprise_id`) USING BTREE,
  KEY `created_member_id` (`created_member_id`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE,
  KEY `updated_at` (`updated_at`) USING BTREE,
  KEY `deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='朋友圈模块朋友圈表';

CREATE TABLE `moment_topic_members` (
  `id` bigint(20) unsigned NOT NULL COMMENT '自增主键',
  `moment_topic_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'circle_topics表ID',
  `moment_member_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'circle_members表ID',
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '企业ID',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  KEY `moment_topic_id` (`moment_topic_id`) USING BTREE,
  KEY `moment_member_id` (`moment_member_id`) USING BTREE,
  KEY `enterprise_id` (`enterprise_id`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE,
  KEY `updated_at` (`updated_at`) USING BTREE,
  KEY `deleted_at` (`deleted_at`) USING BTREE,
  KEY `id` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='朋友圈模块朋友圈和员工关系表';

CREATE TABLE `moment_members` (
  `id` bigint(20) unsigned NOT NULL COMMENT '自增主键',
  `member_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '员工ID',
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sign` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '签名',
  `mark` char(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '链接加密mark',
  `avatar` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `background` varchar(2500) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '背景图',
  `enterprise_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '企业ID',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `enterprise_id` (`enterprise_id`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE,
  KEY `updated_at` (`updated_at`) USING BTREE,
  KEY `deleted_at` (`deleted_at`) USING BTREE,
  KEY `mark` (`mark`) USING BTREE,
  KEY `member_id` (`member_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='朋友圈模块员工表';


ALTER TABLE
  `crs_prod_scrm`.`work_external_user_relation`
ADD
  INDEX `enterprise_delete_external_member_index`(
    `enterprise_id`,
    `deleted_at`,
    `work_external_user_id`,
    `member_id`
  );