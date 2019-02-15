/*
 Navicat Premium Data Transfer

 Source Server         : localhost3306
 Source Server Type    : MySQL
 Source Server Version : 80000
 Source Host           : localhost
 Source Database       : baskball2019

 Target Server Type    : MySQL
 Target Server Version : 80000
 File Encoding         : utf-8

 Date: 02/15/2019 14:01:45 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `baskball_user`
-- ----------------------------
DROP TABLE IF EXISTS `baskball_user`;
CREATE TABLE `baskball_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(20) CHARACTER SET latin1 NOT NULL DEFAULT '' COMMENT '名称',
  `nick` varchar(40) CHARACTER SET latin1 NOT NULL DEFAULT '' COMMENT '昵称',
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '性别 0 女  1 男',
  `weixinid` varchar(120) CHARACTER SET latin1 NOT NULL DEFAULT '' COMMENT '微信账号',
  `qq` varchar(20) CHARACTER SET latin1 NOT NULL DEFAULT '' COMMENT 'QQ号',
  `phone` varchar(11) CHARACTER SET latin1 NOT NULL DEFAULT '' COMMENT '手机号',
  `email` varchar(120) CHARACTER SET latin1 NOT NULL DEFAULT '' COMMENT '邮箱',
  `birthday` date NOT NULL COMMENT '生日',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次修改时间',
  `createtime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='用户表（基础表）';

SET FOREIGN_KEY_CHECKS = 1;
