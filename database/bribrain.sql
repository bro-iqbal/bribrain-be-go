/*
 Navicat Premium Data Transfer

 Source Server         : Local DB
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : bribrain

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 27/04/2022 18:31:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for kunjungans
-- ----------------------------
DROP TABLE IF EXISTS `kunjungans`;
CREATE TABLE `kunjungans` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `akuisisi` int(12) NOT NULL,
  `agen_jawara` int(12) NOT NULL,
  `agen_juragan` int(12) NOT NULL,
  `agen_bep` int(12) NOT NULL,
  `status` enum('active','not-active','deleted') NOT NULL DEFAULT 'active',
  `user_id` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of kunjungans
-- ----------------------------
BEGIN;
INSERT INTO `kunjungans` VALUES (1, 10, 8, 5, 9, 'active', 1, '2022-04-27 17:26:19', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for rkas
-- ----------------------------
DROP TABLE IF EXISTS `rkas`;
CREATE TABLE `rkas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `akuisisi` int(12) NOT NULL,
  `agen_jawara` int(12) NOT NULL,
  `agen_juragan` int(12) NOT NULL,
  `agen_bep` int(12) NOT NULL,
  `status` enum('active','not-active','deleted') NOT NULL DEFAULT 'active',
  `type` enum('target','position') DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of rkas
-- ----------------------------
BEGIN;
INSERT INTO `rkas` VALUES (1, 10, 3, 5, 9, 'active', 'target', 1, '2021-04-27 16:04:58', NULL, NULL);
INSERT INTO `rkas` VALUES (2, 10, 3, 13, 9, 'active', 'position', 1, '2021-04-27 16:04:58', NULL, NULL);
INSERT INTO `rkas` VALUES (3, 800, 870, 450, 200, 'active', 'target', 1, '2022-04-27 17:26:57', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `input_status_kunjungan` enum('Sudah Input','Belum Input') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `input_status_rka` enum('Sudah Input','Belum Input') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` enum('active','not-active','deleted') NOT NULL DEFAULT 'active',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'Erick Hermawan', 'Sudah Input', 'Sudah Input', 'active', '2022-04-27 15:13:34', NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
