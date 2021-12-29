/*
 Navicat Premium Data Transfer

 Source Server         : localhost3306
 Source Server Type    : MySQL
 Source Server Version : 80000
 Source Host           : localhost
 Source Database       : baskball_2020

 Target Server Type    : MySQL
 Target Server Version : 80000
 File Encoding         : utf-8

 Date: 04/01/2020 00:33:35 AM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `baskball_collection`
-- ----------------------------
DROP TABLE IF EXISTS `baskball_collection`;
CREATE TABLE `baskball_collection` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `courtid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '球场id',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否收藏 1 收藏 0 取消',
  `longitude` double(10,6) unsigned NOT NULL COMMENT '经纬经度',
  `latitude` double(10,6) unsigned NOT NULL COMMENT '经纬纬度',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次修改时间',
  `createtime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户收藏表';

-- ----------------------------
--  Table structure for `baskball_court`
-- ----------------------------
DROP TABLE IF EXISTS `baskball_court`;
CREATE TABLE `baskball_court` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='球场表';

-- ----------------------------
--  Table structure for `baskball_match`
-- ----------------------------
DROP TABLE IF EXISTS `baskball_match`;
CREATE TABLE `baskball_match` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '名称',
  `longitude` double(10,6) NOT NULL COMMENT '昵称',
  `latitude` double(10,6) unsigned NOT NULL DEFAULT '1.000000' COMMENT '性别 0 女  1 男',
  `courtId` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '球场id',
  `describe` varchar(255) NOT NULL DEFAULT '' COMMENT '描述 活动描述',
  `memberNum` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '报名人数限制',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '球场类型（0 公开 1 私人）',
  `isverifyed` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否官方认证 0 未认证 1 认证',
  `ishavebaskball` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否有篮球',
  `cost` float unsigned NOT NULL DEFAULT '0' COMMENT '活动费用',
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

-- ----------------------------
--  Table structure for `baskball_matchmember`
-- ----------------------------
DROP TABLE IF EXISTS `baskball_matchmember`;
CREATE TABLE `baskball_matchmember` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'uid 用户ID',
  `matchid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '比赛id',
  `courtId` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '球场id',
  `isnotice` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否Notice 0 否 1 提醒',
  `status` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0 取消 1、报名',
  `createor` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '创建者(uid)',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次修改时间',
  `createtime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='篮球赛报名表';

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
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户表（基础表）';

SET FOREIGN_KEY_CHECKS = 1;
