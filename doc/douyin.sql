/*
 Navicat Premium Data Transfer

 Source Server         : 124.223.207.249
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : 124.223.207.249:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 02/08/2023 15:14:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
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
INSERT INTO `sys_user` VALUES (1, 'tom', 'default', 'default', 'default', 'tom', '123456');
INSERT INTO `sys_user` VALUES (2, 'jack', 'default', 'default', 'default', 'jack', '123456');
INSERT INTO `sys_user` VALUES (3, 'jerry', 'default', 'default', 'default', 'jerry', '123456');
INSERT INTO `sys_user` VALUES (4, 'linda', 'default', 'default', 'default', 'linda', '123456');
INSERT INTO `sys_user` VALUES (10, 'aa', 'aa', 'aa', 'aa', 'aa', 'aa');

-- ----------------------------
-- Table structure for user_favorite
-- ----------------------------
DROP TABLE IF EXISTS `user_favorite`;
CREATE TABLE `user_favorite`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `work_id` bigint(20) NOT NULL COMMENT '作品id',
  `user_id` bigint(20) NOT NULL COMMENT '点赞者id',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` int(11) NOT NULL DEFAULT 0 COMMENT '记录是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

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
INSERT INTO `user_favorite` VALUES (8, 52, 2, '2023-08-02 15:09:07', 1);
INSERT INTO `user_favorite` VALUES (9, 49, 2, '2023-08-02 15:09:03', 1);

-- ----------------------------
-- Table structure for user_follow
-- ----------------------------
DROP TABLE IF EXISTS `user_follow`;
CREATE TABLE `user_follow`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `follow_user_id` bigint(20) NOT NULL COMMENT '关注者(粉丝ID)',
  `user_id` bigint(20) NOT NULL COMMENT '被关注者id',
  `create_time` datetime NULL DEFAULT NULL COMMENT '关注时间',
  `deleted` int(11) NOT NULL DEFAULT 0 COMMENT '是否取关(删除记录)\r\n0 表示正常记录\r\n1 表示已取关',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

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
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) NOT NULL COMMENT '作品发布者的用户id',
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '视频的播放地址',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '视频封面地址',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '视频标题',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间',
  `deleted` int(11) NOT NULL DEFAULT 0 COMMENT '0正常显示\r\n1表示已删除该作品',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_work
-- ----------------------------
INSERT INTO `user_work` VALUES (7, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476586603057152.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476586603057152.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title0', '2023-07-31 12:00:17', 0);
INSERT INTO `user_work` VALUES (8, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476589295800320.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476589295800320.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title1', '2023-07-31 12:00:18', 0);
INSERT INTO `user_work` VALUES (9, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476590646366208.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476590646366208.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title2', '2023-07-31 12:00:19', 0);
INSERT INTO `user_work` VALUES (10, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476593670459392.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476593670459392.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title3', '2023-07-31 12:00:19', 0);
INSERT INTO `user_work` VALUES (11, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476594907779072.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476594907779072.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title4', '2023-07-31 12:00:19', 0);
INSERT INTO `user_work` VALUES (12, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476597688602624.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476597688602624.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title5', '2023-07-31 12:00:20', 0);
INSERT INTO `user_work` VALUES (13, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476599676702720.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476599676702720.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title6', '2023-07-31 12:00:20', 0);
INSERT INTO `user_work` VALUES (14, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476601975181312.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476601975181312.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title7', '2023-07-31 12:00:21', 0);
INSERT INTO `user_work` VALUES (15, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476603355107328.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476603355107328.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title8', '2023-07-31 12:00:21', 0);
INSERT INTO `user_work` VALUES (16, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476604303020032.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476604303020032.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title9', '2023-07-31 12:00:21', 0);
INSERT INTO `user_work` VALUES (17, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476605859106816.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476605859106816.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title10', '2023-07-31 12:00:23', 0);
INSERT INTO `user_work` VALUES (18, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476611311702016.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476611311702016.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title11', '2023-07-31 12:00:23', 0);
INSERT INTO `user_work` VALUES (19, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476612733571072.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476612733571072.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title12', '2023-07-31 12:00:24', 0);
INSERT INTO `user_work` VALUES (20, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476616277757952.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476616277757952.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title13', '2023-07-31 12:00:24', 0);
INSERT INTO `user_work` VALUES (21, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476617829650432.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476617829650432.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title14', '2023-08-01 21:13:45', 0);
INSERT INTO `user_work` VALUES (22, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476621445140480.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476621445140480.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title15', '2023-07-31 12:00:26', 0);
INSERT INTO `user_work` VALUES (23, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476623596818432.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476623596818432.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title16', '2023-07-31 12:00:26', 0);
INSERT INTO `user_work` VALUES (24, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476625404563456.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476625404563456.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title17', '2023-07-31 12:00:27', 0);
INSERT INTO `user_work` VALUES (25, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476628558680064.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476628558680064.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title18', '2023-07-31 12:00:27', 0);
INSERT INTO `user_work` VALUES (26, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476630701969408.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476630701969408.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title19', '2023-07-31 12:00:28', 0);
INSERT INTO `user_work` VALUES (27, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476632413245440.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476632413245440.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title20', '2023-07-31 12:00:28', 0);
INSERT INTO `user_work` VALUES (28, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476633386323968.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476633386323968.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title21', '2023-07-31 12:00:28', 0);
INSERT INTO `user_work` VALUES (29, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476634787221504.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476634787221504.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title22', '2023-07-31 12:00:29', 0);
INSERT INTO `user_work` VALUES (30, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476636288782336.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476636288782336.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title23', '2023-07-31 12:00:29', 0);
INSERT INTO `user_work` VALUES (31, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476637861646336.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476637861646336.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title24', '2023-07-31 12:00:30', 0);
INSERT INTO `user_work` VALUES (32, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476642722844672.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476642722844672.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title25', '2023-07-31 12:00:31', 0);
INSERT INTO `user_work` VALUES (33, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476645033906176.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476645033906176.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title26', '2023-07-31 12:00:31', 0);
INSERT INTO `user_work` VALUES (34, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476646300585984.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476646300585984.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title27', '2023-07-31 12:00:32', 0);
INSERT INTO `user_work` VALUES (35, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476648305463296.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476648305463296.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title28', '2023-07-31 12:00:32', 0);
INSERT INTO `user_work` VALUES (36, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476649458896896.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476649458896896.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title29', '2023-07-31 12:00:32', 0);
INSERT INTO `user_work` VALUES (37, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476650377449472.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476650377449472.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title30', '2023-07-31 12:00:32', 0);
INSERT INTO `user_work` VALUES (38, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476651786735616.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476651786735616.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title31', '2023-07-31 12:00:33', 0);
INSERT INTO `user_work` VALUES (39, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476653338628096.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476653338628096.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title32', '2023-07-31 12:00:34', 0);
INSERT INTO `user_work` VALUES (40, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476657386131456.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476657386131456.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title33', '2023-07-31 12:00:35', 0);
INSERT INTO `user_work` VALUES (41, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476662339604480.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476662339604480.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title34', '2023-07-31 12:00:35', 0);
INSERT INTO `user_work` VALUES (42, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476663799222272.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476663799222272.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title35', '2023-07-31 12:00:36', 0);
INSERT INTO `user_work` VALUES (43, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476665284005888.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476665284005888.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title36', '2023-07-31 12:00:36', 0);
INSERT INTO `user_work` VALUES (44, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476666970116096.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476666970116096.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title37', '2023-07-31 12:00:36', 0);
INSERT INTO `user_work` VALUES (45, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476668048052224.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476668048052224.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title38', '2023-07-31 12:00:37', 0);
INSERT INTO `user_work` VALUES (46, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476669218263040.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476669218263040.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title39', '2023-07-31 12:00:37', 0);
INSERT INTO `user_work` VALUES (47, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476670824681472.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476670824681472.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title40', '2023-07-31 12:00:38', 0);
INSERT INTO `user_work` VALUES (48, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476673727139840.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476673727139840.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title41', '2023-07-31 12:00:38', 0);
INSERT INTO `user_work` VALUES (49, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476675727822848.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476675727822848.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title42', '2023-07-31 12:00:38', 0);
INSERT INTO `user_work` VALUES (50, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476677481041920.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476677481041920.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title43', '2023-07-31 12:00:41', 0);
INSERT INTO `user_work` VALUES (51, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476686180028416.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476686180028416.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title44', '2023-07-31 12:00:41', 0);
INSERT INTO `user_work` VALUES (52, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476688319123456.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476688319123456.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title45', '2023-07-31 12:00:42', 0);
INSERT INTO `user_work` VALUES (53, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476690340777984.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476690340777984.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title46', '2023-07-31 12:00:42', 0);
INSERT INTO `user_work` VALUES (54, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476692983189504.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476692983189504.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title47', '2023-07-31 12:00:43', 0);
INSERT INTO `user_work` VALUES (55, 1, 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476694572830720.mp4', 'https://douyin1562.oss-cn-beijing.aliyuncs.com/video/692476694572830720.mp4?x-oss-process=video/snapshot,t_1000,f_jpg,w_800,h_600', 'title48', '2023-07-31 12:00:43', 0);

-- ----------------------------
-- Table structure for work_comment
-- ----------------------------
DROP TABLE IF EXISTS `work_comment`;
CREATE TABLE `work_comment`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `work_id` bigint(20) NOT NULL COMMENT '作品id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '评论的具体内容',
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` int(11) NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of work_comment
-- ----------------------------
INSERT INTO `work_comment` VALUES (1, 9, 3, '挺好的呀', '2023-08-01 07:45:15', 0);
INSERT INTO `work_comment` VALUES (4, 9, 4, '确实很好', '2023-08-02 14:19:03', 0);
INSERT INTO `work_comment` VALUES (5, 7, 2, '真不戳呢！', '2023-08-02 14:16:18', 0);
INSERT INTO `work_comment` VALUES (8, 54, 2, '真没呀', '2023-08-02 14:28:34', 1);
INSERT INTO `work_comment` VALUES (9, 54, 2, '来吧', '2023-08-02 14:28:44', 0);

SET FOREIGN_KEY_CHECKS = 1;
