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

 Date: 02/15/2019 14:01:07 PM
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
  `longitude` double(10,6) NOT NULL COMMENT '昵称',
  `latitude
 
latitude` double(10,6) unsigned NOT NULL DEFAULT '1.000000' COMMENT '性别 0 女  1 男',
  `updatetime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次修改时间',
  `createtime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户收藏表';

SET FOREIGN_KEY_CHECKS = 1;
