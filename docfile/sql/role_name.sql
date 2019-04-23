/*
 Navicat Premium Data Transfer

 Source Server         : home.xiao3.top
 Source Server Type    : MySQL
 Source Server Version : 50643
 Source Host           : home.xiao3.top:3306
 Source Schema         : mixweb

 Target Server Type    : MySQL
 Target Server Version : 50643
 File Encoding         : 65001

 Date: 23/04/2019 09:04:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for role_name
-- ----------------------------
DROP TABLE IF EXISTS `role_name`;
CREATE TABLE `role_name`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL DEFAULT 0,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
