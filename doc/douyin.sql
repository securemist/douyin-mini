/*
 Navicat Premium Data Transfer

 Source Server         : 10.16.21.78
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 10.16.21.78:3306
 Source Schema         : douyin2

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 03/08/2023 22:03:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'default' COMMENT '用户昵称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'default' COMMENT '用户头像',
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'default' COMMENT '用户个人页顶部大图',
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'default' COMMENT '个人简介',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'default' COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'default' COMMENT '登录密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'tom', '/static/avatar/default.jpg', '/static/bg/default.jpg', '我的个性签名', 'tom', '123456');
INSERT INTO `sys_user` VALUES (2, 'jack', '/static/avatar/default.jpg', '/static/bg/default.jpg', '我的个性签名', 'jack', '123456');
INSERT INTO `sys_user` VALUES (3, 'jerry', '/static/avatar/default.jpg', '/static/bg/default.jpg', '我的个性签名', 'jerry', '123456');
INSERT INTO `sys_user` VALUES (4, 'linda', '/static/avatar/default.jpg', '/static/bg/default.jpg', '我的个性签名', 'linda', '123456');
INSERT INTO `sys_user` VALUES (5, 'jim', '/static/avatar/default.jpg', '/static/bg/default.jpg', '我的个性签名', 'jim', '123456');

-- ----------------------------
-- Table structure for user_favorite
-- ----------------------------
DROP TABLE IF EXISTS `user_favorite`;
CREATE TABLE `user_favorite`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `work_id` bigint NOT NULL COMMENT '作品id',
  `user_id` bigint NOT NULL COMMENT '点赞者id',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` int NOT NULL DEFAULT 0 COMMENT '记录是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_favorite
-- ----------------------------
INSERT INTO `user_favorite` VALUES (1, 9, 2, '2023-08-02 11:42:27', 0);
INSERT INTO `user_favorite` VALUES (2, 7, 3, '2023-08-01 07:43:49', 0);
INSERT INTO `user_favorite` VALUES (3, 7, 4, '2023-08-01 07:43:51', 0);
INSERT INTO `user_favorite` VALUES (4, 9, 4, '2023-08-01 07:45:04', 0);
INSERT INTO `user_favorite` VALUES (5, 8, 2, '2023-08-02 14:34:59', 0);
INSERT INTO `user_favorite` VALUES (6, 55, 2, '2023-08-02 15:09:09', 1);
INSERT INTO `user_favorite` VALUES (7, 54, 2, '2023-08-02 15:09:11', 1);
INSERT INTO `user_favorite` VALUES (8, 52, 2, '2023-08-02 15:31:44', 0);
INSERT INTO `user_favorite` VALUES (9, 49, 2, '2023-08-02 15:09:03', 1);
INSERT INTO `user_favorite` VALUES (10, 21, 2, '2023-08-02 17:43:29', 1);
INSERT INTO `user_favorite` VALUES (11, 54, 1, '2023-08-02 19:39:51', 0);
INSERT INTO `user_favorite` VALUES (12, 21, 10, '2023-08-02 21:06:56', 1);
INSERT INTO `user_favorite` VALUES (13, 55, 10, '2023-08-02 21:07:01', 0);
INSERT INTO `user_favorite` VALUES (14, 54, 10, '2023-08-02 21:07:02', 0);

-- ----------------------------
-- Table structure for user_follow
-- ----------------------------
DROP TABLE IF EXISTS `user_follow`;
CREATE TABLE `user_follow`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `follow_user_id` bigint NOT NULL COMMENT '关注者(粉丝ID)',
  `user_id` bigint NOT NULL COMMENT '被关注者id',
  `create_time` datetime NULL DEFAULT NULL COMMENT '关注时间',
  `deleted` int NOT NULL DEFAULT 0 COMMENT '是否取关(删除记录)\r\n0 表示正常记录\r\n1 表示已取关',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_follow
-- ----------------------------
INSERT INTO `user_follow` VALUES (1, 1, 2, '2023-07-30 15:11:18', 0);
INSERT INTO `user_follow` VALUES (2, 1, 3, '2023-07-30 15:11:31', 0);
INSERT INTO `user_follow` VALUES (3, 1, 4, '2023-07-30 19:13:22', 0);
INSERT INTO `user_follow` VALUES (4, 2, 4, '2023-07-30 19:13:32', 0);
INSERT INTO `user_follow` VALUES (5, 4, 2, '2023-07-30 19:13:32', 0);
INSERT INTO `user_follow` VALUES (6, 3, 2, '2023-08-02 16:37:08', 0);
INSERT INTO `user_follow` VALUES (8, 1, 10, '2023-08-02 21:06:50', 0);

-- ----------------------------
-- Table structure for user_message
-- ----------------------------
DROP TABLE IF EXISTS `user_message`;
CREATE TABLE `user_message`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `from_user_id` bigint NOT NULL DEFAULT 0 COMMENT '发送者',
  `to_user_id` bigint NOT NULL DEFAULT 0 COMMENT '接收者',
  `content` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发送时间',
  `deleted` int NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_message
-- ----------------------------
INSERT INTO `user_message` VALUES (1, 2, 1, '你好牛逼啊', '2023-08-04 16:55:58', 0);
INSERT INTO `user_message` VALUES (2, 1, 2, '你也是', '2023-08-02 16:56:14', 0);
INSERT INTO `user_message` VALUES (3, 2, 1, '卧槽', '2023-08-02 17:40:50', 0);
INSERT INTO `user_message` VALUES (4, 2, 4, '在么', '2023-08-02 19:13:56', 0);
INSERT INTO `user_message` VALUES (5, 2, 4, '你好啊', '2023-08-02 19:35:14', 0);

-- ----------------------------
-- Table structure for user_work
-- ----------------------------
DROP TABLE IF EXISTS `user_work`;
CREATE TABLE `user_work`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint NOT NULL COMMENT '作品发布者的用户id',
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '视频的播放地址',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '视频封面地址',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '视频标题',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `deleted` int NOT NULL DEFAULT 0 COMMENT '0正常显示\r\n1表示已删除该作品',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 67 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_work
-- ----------------------------

-- ----------------------------
-- Table structure for work_comment
-- ----------------------------
DROP TABLE IF EXISTS `work_comment`;
CREATE TABLE `work_comment`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `work_id` bigint NOT NULL COMMENT '作品id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '评论的具体内容',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` int NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of work_comment
-- ----------------------------
INSERT INTO `work_comment` VALUES (1, 9, 3, '挺好的呀', '2023-08-01 07:45:15', 0);
INSERT INTO `work_comment` VALUES (4, 9, 4, '确实很好', '2023-08-02 14:19:03', 0);
INSERT INTO `work_comment` VALUES (5, 7, 2, '真不戳呢！', '2023-08-02 14:16:18', 0);
INSERT INTO `work_comment` VALUES (8, 54, 2, '真没呀', '2023-08-02 14:28:34', 1);
INSERT INTO `work_comment` VALUES (9, 54, 2, '来吧', '2023-08-02 14:28:44', 0);
INSERT INTO `work_comment` VALUES (10, 52, 2, '在么啊', '2023-08-02 15:31:44', 0);
INSERT INTO `work_comment` VALUES (11, 54, 1, '挺好的', '2023-08-02 19:39:51', 0);
INSERT INTO `work_comment` VALUES (12, 21, 10, '哈哈', '2023-08-02 21:06:56', 0);

SET FOREIGN_KEY_CHECKS = 1;
