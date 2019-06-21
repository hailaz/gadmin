/*
 Navicat Premium Data Transfer

 Source Server         : home.xiao3.top
 Source Server Type    : MySQL
 Source Server Version : 50643
 Source Host           : home.xiao3.top:3306
 Source Schema         : gadmin

 Target Server Type    : MySQL
 Target Server Version : 50643
 File Encoding         : 65001

 Date: 21/06/2019 14:12:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单路径',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '页面模块',
  `redirect` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '重定向地址',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '唯一关键名',
  `hidden` tinyint(1) NULL DEFAULT 0 COMMENT '是否隐藏',
  `alwaysshow` tinyint(1) NULL DEFAULT 0 COMMENT '是否常显示',
  `sort` tinyint(2) NULL DEFAULT 0 COMMENT '排序',
  `parent_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '父菜级关键名',
  `auto_create` tinyint(1) NULL DEFAULT 0 COMMENT '是否自动生成',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `key_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
