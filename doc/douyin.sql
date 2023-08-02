/*
 Navicat Premium Data Transfer

 Source Server         : 10.16.21.78
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 10.16.21.78:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 01/08/2023 18:22:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'default' COMMENT '用户昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'default' COMMENT '用户头像',
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'default' COMMENT '用户个人页顶部大图',
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'default' COMMENT '个人简介',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'default' COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'default' COMMENT '登录密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'tom', 'default', 'default', 'default', 'tom', '12345');
INSERT INTO `sys_user` VALUES (2, 'jack', 'default', 'default', 'default', 'jack', '12345');
INSERT INTO `sys_user` VALUES (3, 'jerry', 'default', 'default', 'default', 'jerry', '12345');
INSERT INTO `sys_user` VALUES (4, 'linda', 'default', 'default', 'default', 'linda', '12345');
INSERT INTO `sys_user` VALUES (8, 'default', 'default', 'default', 'default', 'tom', '12345');
INSERT INTO `sys_user` VALUES (9, 'default', 'default', 'default', 'default', 'tom1', '12345');

-- ----------------------------
-- Table structure for user_favorite
-- ----------------------------
DROP TABLE IF EXISTS `user_favorite`;
CREATE TABLE `user_favorite`  (
  `id` bigint NOT NULL,
  `work_id` bigint NOT NULL COMMENT '作品id',
  `user_id` bigint NOT NULL COMMENT '点赞者id',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` int NOT NULL DEFAULT 0 COMMENT '记录是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_favorite
-- ----------------------------
INSERT INTO `user_favorite` VALUES (1, 9, 2, '2023-08-01 07:45:54', 0);
INSERT INTO `user_favorite` VALUES (2, 7, 3, '2023-08-01 07:43:49', 0);
INSERT INTO `user_favorite` VALUES (3, 7, 4, '2023-08-01 07:43:51', 0);
INSERT INTO `user_favorite` VALUES (4, 9, 4, '2023-08-01 07:45:04', 0);

-- ----------------------------
-- Table structure for user_follow
-- ----------------------------
DROP TABLE IF EXISTS `user_follow`;
CREATE TABLE `user_follow`  (
  `id` bigint NOT NULL,
  `follow_user_id` bigint NOT NULL COMMENT '关注者(粉丝ID)',
  `user_id` bigint NOT NULL COMMENT '被关注者id',
  `create_time` datetime NULL DEFAULT NULL COMMENT '关注时间',
  `deleted` int NOT NULL DEFAULT 0 COMMENT '是否取关(删除记录)\r\n0 表示正常记录\r\n1 表示已取关',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_follow
-- ----------------------------
INSERT INTO `user_follow` VALUES (1, 1, 2, '2023-07-30 15:11:18', 0);
INSERT INTO `user_follow` VALUES (2, 1, 3, '2023-07-30 15:11:31', 0);
INSERT INTO `user_follow` VALUES (3, 1, 4, '2023-07-30 19:13:22', 0);
INSERT INTO `user_follow` VALUES (4, 2, 4, '2023-07-30 19:13:32', 0);

-- ----------------------------
-- Table structure for user_work
-- ----------------------------
DROP TABLE IF EXISTS `user_work`;
CREATE TABLE `user_work`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint NOT NULL COMMENT '作品发布者的用户id',
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'default' COMMENT '视频的播放地址',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'default' COMMENT '视频封面地址',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'default' COMMENT '视频标题',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `deleted` int NOT NULL DEFAULT 0 COMMENT '0正常显示\r\n1表示已删除该作品',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_work
-- ----------------------------

-- ----------------------------
-- Table structure for work_comment
-- ----------------------------
DROP TABLE IF EXISTS `work_comment`;
CREATE TABLE `work_comment`  (
  `id` bigint NOT NULL,
  `work_id` bigint NOT NULL COMMENT '作品id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'default' COMMENT '评论的具体内容',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` int NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of work_comment
-- ----------------------------
INSERT INTO `work_comment` VALUES (1, 9, 3, '挺好的呀', '2023-08-01 07:45:15', 0);
INSERT INTO `work_comment` VALUES (4, 9, 4, '确实很好', '2023-08-01 07:45:17', 0);

SET FOREIGN_KEY_CHECKS = 1;
