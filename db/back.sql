/*
 Navicat Premium Data Transfer

 Source Server         : localdb
 Source Server Type    : MySQL
 Source Server Version : 90000 (9.0.0)
 Source Host           : 127.0.0.1:3306
 Source Schema         : dfes_web

 Target Server Type    : MySQL
 Target Server Version : 90000 (9.0.0)
 File Encoding         : 65001

 Date: 04/07/2024 14:26:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for r_user_file_root
-- ----------------------------
DROP TABLE IF EXISTS `r_user_file_root`;
CREATE TABLE `r_user_file_root` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `tree_root_id` bigint unsigned NOT NULL,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `r_user_file_root_nk_user_id` (`user_id`),
  UNIQUE KEY `r_user_file_root_uindex_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='map user with file root id';

-- ----------------------------
-- Table structure for t_file_tree
-- ----------------------------
DROP TABLE IF EXISTS `t_file_tree`;
CREATE TABLE `t_file_tree` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(1024) NOT NULL,
  `data_id` varchar(255) DEFAULT NULL,
  `parent` bigint unsigned DEFAULT NULL COMMENT '0 root / else other',
  `kind` tinyint NOT NULL DEFAULT '1' COMMENT '0 file 1 directory',
  `share_url` varchar(512) DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_at` timestamp NULL DEFAULT NULL,
  `file_size` int(10) unsigned zerofill DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `t_file_tree_nk_data_id` (`data_id`),
  UNIQUE KEY `t_file_tree_nk_parent` (`parent`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `ID` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` char(50) DEFAULT NULL,
  `password` char(100) DEFAULT NULL,
  `email` varchar(1024) DEFAULT NULL,
  `phone` varchar(30) DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_at` timestamp NULL DEFAULT NULL,
  `enable` tinyint DEFAULT '0' COMMENT '0 启用 1 冻结',
  PRIMARY KEY (`ID`),
  UNIQUE KEY `t_user_k_delete_time` (`delete_at`),
  UNIQUE KEY `t_user_k_phone` (`phone`),
  UNIQUE KEY `t_user_username_index` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
