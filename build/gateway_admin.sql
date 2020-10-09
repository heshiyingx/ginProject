CREATE TABLE `gateway_admin` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `passwd` varchar(255) NOT NULL COMMENT '密码',
  `created_at` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
  `status` int NOT NULL DEFAULT '0' COMMENT '账号状态',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '删除标识，1为删除',
  PRIMARY KEY (`id`),
  KEY `passwd` (`passwd`),
  KEY `user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员表';

CREATE TABLE `gateway_app` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `app_id` bigint NOT NULL DEFAULT '0' COMMENT '租户ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '租户名称',
  `secret` varchar(255) NOT NULL DEFAULT '' COMMENT '秘钥',
  `white_ips` varchar(1000) NOT NULL DEFAULT '' COMMENT 'ip白名单，支持前缀匹配',
  `qpd` bigint NOT NULL DEFAULT '0' COMMENT '每日请求限制',
  `qps` bigint NOT NULL DEFAULT '0' COMMENT '每秒请求限制',
  `created_at` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '删除标识',
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='网关租户表';

CREATE TABLE `gateway_service_access_control` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint NOT NULL DEFAULT '0' COMMENT '服务ID',
  `open_auth` tinyint NOT NULL DEFAULT '0' COMMENT '是否开启权限，1为开启',
  `black_list` varchar(1000) NOT NULL DEFAULT '' COMMENT '黑名单IP',
  `white_list` varchar(1000) NOT NULL DEFAULT '' COMMENT '白名单IP',
  `white_host_name` varchar(1000) NOT NULL DEFAULT '' COMMENT '白名单主机',
  `clientip_flow_limit` int NOT NULL DEFAULT '0' COMMENT '客户端IP限流',
  `service_flow_limit` int NOT NULL DEFAULT '0' COMMENT '服务端限流',
  UNIQUE KEY `gateway_service_access_control_pk` (`id`),
  KEY `service_id` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='网关权限控制表';

CREATE TABLE `gateway_service_info` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `load_type` int NOT NULL DEFAULT '0' COMMENT '负载类型 0=http  1=tcp 2=grpc',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称 6-128数字字母下划线',
  `service_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称',
  `created_at` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint NOT NULL DEFAULT '0' COMMENT '更新时间',
  `is_delete` tinyint DEFAULT '0' COMMENT '删除标识，1为删除',
  PRIMARY KEY (`id`),
  KEY `service_desc` (`service_desc`),
  KEY `service_name` (`service_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='网关模块表';

CREATE TABLE `gateway_service_http_rule` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `service_id` bigint NOT NULL DEFAULT '0' COMMENT '服务ID',
  `rule_type` tinyint NOT NULL DEFAULT '0' COMMENT '匹配类型 0=url前缀，url_prefix,1=域名domain',
  `rule` varchar(255) NOT NULL DEFAULT '' COMMENT 'type=domain表示域名，type=url_prefix表示url前缀',
  `need_https` tinyint DEFAULT '0' COMMENT '支持https ,1 支持',
  `need_strip_uri` tinyint NOT NULL DEFAULT '0' COMMENT '启用strip_uri,1启用',
  `need_websocket` tinyint NOT NULL DEFAULT '0' COMMENT '支持websocket,1支持',
  `url_rewrite` varchar(5000) NOT NULL DEFAULT '' COMMENT 'url重写功能，格式:^/gatekeeper/test_service(.*)$1多个逗号间隔',
  `header_transfor` varchar(5000) NOT NULL DEFAULT '' COMMENT 'header转换支持增加(add),删除(del),修改(edit).格式 add headname headvalue 多个逗号分隔',
  PRIMARY KEY (`id`),
  KEY `service_id` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='网关路由匹配表';

CREATE TABLE `gateway_service_load_balance` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint NOT NULL DEFAULT '0' COMMENT '服务ID',
  `check_method` int NOT NULL DEFAULT '0' COMMENT '检查方法，0=tcpchk,检测握手成功',
  `check_interval` int NOT NULL DEFAULT '0' COMMENT '间隔时间，单位s',
  `round_type` tinyint NOT NULL DEFAULT '0' COMMENT '轮询方式，0=random,1=round-robin,2=weight_round-robin,3=ip_hash',
  `ip_list` varchar(2000) NOT NULL DEFAULT '' COMMENT 'ip列表',
  `weight_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '权重列表',
  `forbid_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '禁用ip列表',
  `upstream_connect_timeout` int NOT NULL DEFAULT '0' COMMENT '建立连接超时，单位s',
  `upstream_header_timeout` int NOT NULL DEFAULT '0' COMMENT '获取header超时，单位s',
  `upstream_idle_timeout` int NOT NULL DEFAULT '0' COMMENT '连接最大空闲时间，单位s',
  `upstream_max_idle` int NOT NULL DEFAULT '0' COMMENT '最大空闲连接数',
  PRIMARY KEY (`id`),
  KEY `service_id` (`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='负载均衡表';