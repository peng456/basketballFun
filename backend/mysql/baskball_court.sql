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

 Date: 02/15/2019 14:01:19 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `baskball_court`
-- ----------------------------
DROP TABLE IF EXISTS `baskball_court`;
CREATE TABLE `baskball_court` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `longitude` double(10,6) NOT NULL COMMENT '昵称',
  `latitude` double(10,6) unsigned NOT NULL DEFAULT '1.000000' COMMENT '性别 0 女  1 男',
  `courttype` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '球场类型（0 免费 1 付费）',
  `isverifyed` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否官方认证 0 未认证 1 认证',
  `numhoop` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '篮筐数量',
  `groundtype` tinyint(2) NOT NULL DEFAULT '0' COMMENT '篮球地面性质：0 水泥 1 橡胶 2 其他没想好以后补充',
  `isdeleted` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 1 是 0 否',
  `createor` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者(uid)',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次修改时间',
  `createtime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='球场表';

SET FOREIGN_KEY_CHECKS = 1;
