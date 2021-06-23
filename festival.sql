/*
 Navicat Premium Data Transfer

 Source Server         : 本地Mysql
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : 127.0.0.1:3306
 Source Schema         : festival

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 21/06/2021 17:08:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for mod_goods
-- ----------------------------
DROP TABLE IF EXISTS `mod_goods`;
CREATE TABLE `mod_goods` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(200) DEFAULT NULL COMMENT '标题',
  `cover` varchar(300) DEFAULT NULL COMMENT '封面',
  `num` int(11) DEFAULT '0' COMMENT '库存量',
  PRIMARY KEY (`id`),
  KEY `idx_mod_goods_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of mod_goods
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mod_member
-- ----------------------------
DROP TABLE IF EXISTS `mod_member`;
CREATE TABLE `mod_member` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `openid` varchar(280) DEFAULT NULL COMMENT '微信标识',
  `route` tinyint(1) DEFAULT '0' COMMENT '路线:0未选择,1建党,2建司',
  `name` varchar(280) DEFAULT NULL COMMENT '姓名',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `agency` varchar(280) DEFAULT NULL COMMENT '投保机构',
  `avg` int(10) DEFAULT '0' COMMENT '平均步数',
  `total` int(10) DEFAULT '0' COMMENT '累积步数',
  `address` varchar(280) DEFAULT NULL COMMENT '地址',
  PRIMARY KEY (`id`),
  KEY `idx_mod_member_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of mod_member
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mod_route
-- ----------------------------
DROP TABLE IF EXISTS `mod_route`;
CREATE TABLE `mod_route` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `type` tinyint(1) DEFAULT '1' COMMENT '类型:1建党,2建司',
  `cover` varchar(280) DEFAULT NULL COMMENT '封面图',
  `name` varchar(280) DEFAULT NULL COMMENT '姓名',
  `step` int(10) DEFAULT '0' COMMENT '需要步数',
  PRIMARY KEY (`id`),
  KEY `idx_mod_route_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of mod_route
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mod_user_goods
-- ----------------------------
DROP TABLE IF EXISTS `mod_user_goods`;
CREATE TABLE `mod_user_goods` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `userId` int(11) DEFAULT NULL COMMENT '用户ID',
  `goodsId` int(11) DEFAULT NULL COMMENT '礼品ID',
  `status` tinyint(1) DEFAULT '0' COMMENT '状态：0审核中1已发货-1驳回',
  `name` varchar(280) DEFAULT NULL COMMENT '用户名',
  `phone` varchar(11) DEFAULT NULL COMMENT '电话号',
  `address` varchar(300) DEFAULT NULL COMMENT '联系地址',
  PRIMARY KEY (`id`),
  KEY `idx_mod_user_goods_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of mod_user_goods
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for mod_user_route
-- ----------------------------
DROP TABLE IF EXISTS `mod_user_route`;
CREATE TABLE `mod_user_route` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `userId` int(11) DEFAULT NULL COMMENT '用户ID',
  `routeId` int(11) DEFAULT NULL COMMENT '线路ID',
  PRIMARY KEY (`id`),
  KEY `idx_mod_user_route_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of mod_user_route
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for s_menus
-- ----------------------------
DROP TABLE IF EXISTS `s_menus`;
CREATE TABLE `s_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(20) NOT NULL COMMENT '名称',
  `parent_id` bigint(20) DEFAULT NULL COMMENT '父ID',
  `sort` int(10) DEFAULT '0' COMMENT '排序',
  `router` varchar(100) NOT NULL COMMENT '路由地址/链接/接口',
  `permission` varchar(100) NOT NULL DEFAULT '' COMMENT '权限标识',
  `type` tinyint(1) DEFAULT '0' COMMENT '类型(0-页面 1-接口)',
  `visible` tinyint(1) DEFAULT '1' COMMENT '是否显示(0否 1是)',
  `icon` varchar(100) DEFAULT '' COMMENT '图标',
  `extra` varchar(100) DEFAULT NULL COMMENT '额外参数',
  `method` varchar(10) DEFAULT 'GET' COMMENT '类型:(GETPOST)',
  PRIMARY KEY (`id`),
  KEY `idx_s_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of s_menus
-- ----------------------------
BEGIN;
INSERT INTO `s_menus` VALUES (1, '2021-01-29 23:16:57', '2021-02-06 14:24:04', NULL, '首页', 0, 0, '/admin/index', 'home', 0, 1, 'fa-tachometer-alt', '', 'GET');
INSERT INTO `s_menus` VALUES (2, '2021-01-29 23:16:57', '2021-02-06 14:23:47', NULL, '系统管理', 0, 0, '/admin/system', 'system', 0, 1, 'fa-cogs', '', 'GET');
INSERT INTO `s_menus` VALUES (3, '2021-01-29 23:16:57', '2021-02-06 14:24:17', NULL, '用户管理', 2, 0, '/admin/system/users', 'system:user:list', 0, 1, 'fa-users', '', 'GET');
INSERT INTO `s_menus` VALUES (4, '2021-01-29 23:16:57', '2021-02-06 14:24:31', NULL, '菜单管理', 2, 0, '/admin/system/menus', 'system:menu:list', 0, 1, 'fa-bars', '', 'GET');
INSERT INTO `s_menus` VALUES (5, '2021-01-29 23:16:57', '2021-02-06 14:24:44', NULL, '角色管理', 2, 0, '/admin/system/roles', 'system:role:list', 0, 1, 'fa-cog', '', 'GET');
INSERT INTO `s_menus` VALUES (6, '2021-02-04 17:22:39', '2021-02-06 14:08:00', NULL, '编辑用户', 2, 0, '/admin/system/user/edit', 'system:user:edit', 0, 0, '', NULL, 'GET');
INSERT INTO `s_menus` VALUES (7, '2021-02-06 14:07:05', '2021-02-06 14:21:48', NULL, '编辑菜单', 2, 0, '/admin/system/menu/edit', 'system:menu:edit', 0, 0, 'fa-pencil-square-o', '', 'GET');
INSERT INTO `s_menus` VALUES (8, '2021-02-06 14:23:06', '2021-02-06 14:23:22', NULL, '编辑角色', 2, 0, '/admin/system/role/edit', 'system:role:edit', 0, 0, 'fa-cog', '', 'GET');
INSERT INTO `s_menus` VALUES (9, '2021-02-06 14:48:10', NULL, NULL, '编辑角色API', 2, 0, '/admin/system/role/save', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (10, '2021-02-06 14:48:10', NULL, NULL, '编辑菜单API', 2, 0, '/admin/system/menu/save', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (11, '2021-02-06 14:51:39', '2021-02-07 00:31:17', NULL, '编辑用户API', 2, 0, '/admin/system/user/save', '', 1, 0, '', '', 'POST');
INSERT INTO `s_menus` VALUES (12, '2021-02-06 21:37:17', '2021-02-06 22:54:38', NULL, '删除菜单API', 2, 0, '/admin/system/menu/del', '', 1, 0, '', '', 'POST');
INSERT INTO `s_menus` VALUES (13, '2021-02-06 21:37:17', '2021-02-06 22:54:38', NULL, '删除角色API', 2, 0, '/admin/system/role/del', '', 1, 0, '', '', 'POST');
INSERT INTO `s_menus` VALUES (14, '2021-02-06 21:37:17', '2021-02-06 22:54:38', NULL, '删除用户API', 2, 0, '/admin/system/user/del', '', 1, 0, '', '', 'POST');
INSERT INTO `s_menus` VALUES (20, '2021-06-21 16:54:19', '2021-06-21 17:00:43', NULL, '礼品表管理', 0, 0, '/admin/module/goods', '', 0, 1, 'fa-gift', '', 'GET');
INSERT INTO `s_menus` VALUES (21, '2021-06-21 16:54:19', '2021-06-21 17:02:27', NULL, '礼品表列表', 20, 0, '/admin/module/goodss', '', 0, 1, 'fa-bars', '', 'GET');
INSERT INTO `s_menus` VALUES (22, '2021-06-21 16:54:19', NULL, NULL, '编辑礼品表', 20, 0, '/admin/module/goods/edit', '', 0, 1, '', NULL, 'GET');
INSERT INTO `s_menus` VALUES (23, '2021-06-21 16:54:19', NULL, NULL, '编辑礼品表API', 20, 0, '/admin/module/goods/save', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (24, '2021-06-21 16:54:19', NULL, NULL, '删除礼品表API', 20, 0, '/admin/module/goods/del', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (25, '2021-06-21 16:54:27', NULL, NULL, '微信用户管理', 0, 0, '/admin/module/member', '', 0, 1, 'fa-users', NULL, 'GET');
INSERT INTO `s_menus` VALUES (26, '2021-06-21 16:54:27', '2021-06-21 17:02:08', NULL, '微信用户列表', 25, 0, '/admin/module/members', '', 0, 1, 'fa-bars', '', 'GET');
INSERT INTO `s_menus` VALUES (27, '2021-06-21 16:54:27', NULL, NULL, '编辑微信用户', 25, 0, '/admin/module/member/edit', '', 0, 1, '', NULL, 'GET');
INSERT INTO `s_menus` VALUES (28, '2021-06-21 16:54:27', NULL, NULL, '编辑微信用户API', 25, 0, '/admin/module/member/save', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (29, '2021-06-21 16:54:27', NULL, NULL, '删除微信用户API', 25, 0, '/admin/module/member/del', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (30, '2021-06-21 16:54:34', '2021-06-21 17:05:25', NULL, '线路管理', 0, 0, '/admin/module/route', '', 0, 1, 'fa-road', '', 'GET');
INSERT INTO `s_menus` VALUES (31, '2021-06-21 16:54:34', '2021-06-21 17:02:01', NULL, '线路列表', 30, 0, '/admin/module/routes', '', 0, 1, 'fa-bars', '', 'GET');
INSERT INTO `s_menus` VALUES (32, '2021-06-21 16:54:34', NULL, NULL, '编辑线路', 30, 0, '/admin/module/route/edit', '', 0, 1, '', NULL, 'GET');
INSERT INTO `s_menus` VALUES (33, '2021-06-21 16:54:34', NULL, NULL, '编辑线路API', 30, 0, '/admin/module/route/save', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (34, '2021-06-21 16:54:34', NULL, NULL, '删除线路API', 30, 0, '/admin/module/route/del', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (35, '2021-06-21 16:54:42', '2021-06-21 17:06:15', NULL, '兑换礼品表管理', 0, 0, '/admin/module/userGoods', '', 0, 1, 'fa-bars', '', 'GET');
INSERT INTO `s_menus` VALUES (36, '2021-06-21 16:54:42', '2021-06-21 17:01:53', NULL, '兑换礼品表列表', 35, 0, '/admin/module/userGoodss', '', 0, 1, 'fa-bars', '', 'GET');
INSERT INTO `s_menus` VALUES (37, '2021-06-21 16:54:42', NULL, NULL, '编辑兑换礼品表', 35, 0, '/admin/module/userGoods/edit', '', 0, 1, '', NULL, 'GET');
INSERT INTO `s_menus` VALUES (38, '2021-06-21 16:54:42', NULL, NULL, '编辑兑换礼品表API', 35, 0, '/admin/module/userGoods/save', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (39, '2021-06-21 16:54:42', NULL, NULL, '删除兑换礼品表API', 35, 0, '/admin/module/userGoods/del', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (40, '2021-06-21 16:54:50', '2021-06-21 17:08:05', NULL, '点亮线路表管理', 0, 0, '/admin/module/userRoute', '', 0, 1, 'fa-registered', '', 'GET');
INSERT INTO `s_menus` VALUES (41, '2021-06-21 16:54:50', '2021-06-21 17:01:43', NULL, '点亮线路表列表', 40, 0, '/admin/module/userRoutes', '', 0, 1, 'fa-bars', '', 'GET');
INSERT INTO `s_menus` VALUES (42, '2021-06-21 16:54:50', NULL, NULL, '编辑点亮线路表', 40, 0, '/admin/module/userRoute/edit', '', 0, 1, '', NULL, 'GET');
INSERT INTO `s_menus` VALUES (43, '2021-06-21 16:54:50', NULL, NULL, '编辑点亮线路表API', 40, 0, '/admin/module/userRoute/save', '', 1, 0, '', NULL, 'POST');
INSERT INTO `s_menus` VALUES (44, '2021-06-21 16:54:50', NULL, NULL, '删除点亮线路表API', 40, 0, '/admin/module/userRoute/del', '', 1, 0, '', NULL, 'POST');
COMMIT;

-- ----------------------------
-- Table structure for s_role_menus
-- ----------------------------
DROP TABLE IF EXISTS `s_role_menus`;
CREATE TABLE `s_role_menus` (
  `sys_role_id` bigint(20) unsigned NOT NULL,
  `sys_menu_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`sys_role_id`,`sys_menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of s_role_menus
-- ----------------------------
BEGIN;
INSERT INTO `s_role_menus` VALUES (1, 1);
INSERT INTO `s_role_menus` VALUES (1, 2);
INSERT INTO `s_role_menus` VALUES (1, 3);
INSERT INTO `s_role_menus` VALUES (1, 4);
INSERT INTO `s_role_menus` VALUES (1, 5);
INSERT INTO `s_role_menus` VALUES (1, 6);
INSERT INTO `s_role_menus` VALUES (1, 7);
INSERT INTO `s_role_menus` VALUES (1, 8);
INSERT INTO `s_role_menus` VALUES (1, 9);
INSERT INTO `s_role_menus` VALUES (1, 10);
INSERT INTO `s_role_menus` VALUES (1, 11);
INSERT INTO `s_role_menus` VALUES (1, 12);
INSERT INTO `s_role_menus` VALUES (1, 13);
INSERT INTO `s_role_menus` VALUES (1, 14);
INSERT INTO `s_role_menus` VALUES (1, 20);
INSERT INTO `s_role_menus` VALUES (1, 21);
INSERT INTO `s_role_menus` VALUES (1, 23);
INSERT INTO `s_role_menus` VALUES (1, 24);
INSERT INTO `s_role_menus` VALUES (1, 25);
INSERT INTO `s_role_menus` VALUES (1, 26);
INSERT INTO `s_role_menus` VALUES (1, 28);
INSERT INTO `s_role_menus` VALUES (1, 29);
INSERT INTO `s_role_menus` VALUES (1, 30);
INSERT INTO `s_role_menus` VALUES (1, 31);
INSERT INTO `s_role_menus` VALUES (1, 33);
INSERT INTO `s_role_menus` VALUES (1, 34);
INSERT INTO `s_role_menus` VALUES (1, 35);
INSERT INTO `s_role_menus` VALUES (1, 36);
INSERT INTO `s_role_menus` VALUES (1, 38);
INSERT INTO `s_role_menus` VALUES (1, 39);
INSERT INTO `s_role_menus` VALUES (1, 40);
INSERT INTO `s_role_menus` VALUES (1, 41);
INSERT INTO `s_role_menus` VALUES (1, 43);
INSERT INTO `s_role_menus` VALUES (1, 44);
INSERT INTO `s_role_menus` VALUES (2, 1);
INSERT INTO `s_role_menus` VALUES (2, 2);
INSERT INTO `s_role_menus` VALUES (2, 3);
COMMIT;

-- ----------------------------
-- Table structure for s_roles
-- ----------------------------
DROP TABLE IF EXISTS `s_roles`;
CREATE TABLE `s_roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(20) NOT NULL COMMENT '名称',
  `desc` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_s_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of s_roles
-- ----------------------------
BEGIN;
INSERT INTO `s_roles` VALUES (1, '2021-01-29 23:15:58', '2021-06-21 16:58:34', NULL, '超级管理员', '全部权限');
INSERT INTO `s_roles` VALUES (2, '2021-02-05 23:12:52', '2021-02-06 09:46:48', NULL, '管理员', '普通管理');
COMMIT;

-- ----------------------------
-- Table structure for s_users
-- ----------------------------
DROP TABLE IF EXISTS `s_users`;
CREATE TABLE `s_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `account` varchar(20) NOT NULL COMMENT '账号',
  `user_name` varchar(20) DEFAULT NULL COMMENT '昵称',
  `email` varchar(30) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '电话',
  `gender` tinyint(1) DEFAULT '0' COMMENT '性别(0未知 1男 2女)',
  `avatar` varchar(200) DEFAULT 'http://f.rn4.cn/logo.png' COMMENT '头像',
  `password` varchar(200) NOT NULL COMMENT '密码',
  `salt` varchar(20) DEFAULT '' COMMENT '盐值',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态(0异常 1正常)',
  `role_id` bigint(20) unsigned DEFAULT NULL COMMENT '角色id',
  PRIMARY KEY (`id`,`account`),
  UNIQUE KEY `account` (`account`),
  KEY `idx_s_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of s_users
-- ----------------------------
BEGIN;
INSERT INTO `s_users` VALUES (1, '2021-02-05 14:28:55', '2021-02-06 15:44:10', NULL, 'appjobs', '超管', '', '', 0, '', '4f10323179fe98326cce1f456822b698', 'Ai123', 0, 1);
INSERT INTO `s_users` VALUES (2, '2021-02-05 14:41:08', '2021-02-06 15:02:40', NULL, 'gravity', '引力跃迁', NULL, NULL, 0, 'http://f.rn4.cn/logo.png', '02c1e5bd5145e30006b701f65a31f4c9', 'LWivtl', 1, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
