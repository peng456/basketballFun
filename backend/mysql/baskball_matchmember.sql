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

 Date: 02/15/2019 14:01:38 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `baskball_matchmember`
-- ----------------------------
DROP TABLE IF EXISTS `baskball_matchmember`;
CREATE TABLE `baskball_matchmember` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'uid',
  `matchid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '比赛id',
  `courtId` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '球场id',
  `isnotice` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否Notice 0 否 1 提醒',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0 取消 1、报名',
  `createor` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者(uid)',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次修改时间',
  `createtime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='篮球赛报名表';

SET FOREIGN_KEY_CHECKS = 1;
