/*
 Navicat Premium Data Transfer

 Source Server         : 本地Mysql
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : 127.0.0.1:3306
 Source Schema         : appjobs

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 18/02/2021 17:35:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

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
INSERT INTO `s_roles` VALUES (1, '2021-01-29 23:15:58', '2021-02-18 17:20:50', NULL, '超级管理员', '全部权限');
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
