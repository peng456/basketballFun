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

 Date: 02/15/2019 14:01:28 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `baskball_match`
-- ----------------------------
DROP TABLE IF EXISTS `baskball_match`;
CREATE TABLE `baskball_match` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `longitude` double(10,6) NOT NULL COMMENT '昵称',
  `latitude
 
latitude` double(10,6) unsigned NOT NULL DEFAULT '1.000000' COMMENT '性别 0 女  1 男',
  `courtId` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '球场id',
  `memberNum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '报名人数限制',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '球场类型（0 公开 1 私人）',
  `isverifyed` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否官方认证 0 未认证 1 认证',
  `ishavebaskball` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否有篮球',
  `isfree` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否付费 0 免费 1 付费',
  `isAA` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否AA   0  创建者请客  1 是',
  `startTime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '开始时间',
  `endTime` int(10) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `duration` float(4,0) unsigned NOT NULL DEFAULT '0' COMMENT '持续时间 但是 小时',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0 草稿 1、发布 2、取消',
  `createor` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者(uid)',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次修改时间',
  `createtime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='球赛表';

SET FOREIGN_KEY_CHECKS = 1;
