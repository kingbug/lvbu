/*
Navicat MySQL Data Transfer

Source Server         : 192.168.2.3
Source Server Version : 50505
Source Host           : 192.168.2.3:3306
Source Database       : lvbu

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2017-02-03 09:45:51
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `lvbu_config`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_config`;
CREATE TABLE `lvbu_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `dvalue` varchar(200) NOT NULL DEFAULT '',
  `tvalue` varchar(200) NOT NULL DEFAULT '',
  `ovalue` varchar(200) NOT NULL DEFAULT '',
  `dtstatus` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `tostatus` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `content` varchar(200) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_config
-- ----------------------------

-- ----------------------------
-- Table structure for `lvbu_configver`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_configver`;
CREATE TABLE `lvbu_configver` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `value` varchar(200) NOT NULL DEFAULT '',
  `env_id` int(10) unsigned NOT NULL,
  `ver` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_configver
-- ----------------------------

-- ----------------------------
-- Table structure for `lvbu_env`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_env`;
CREATE TABLE `lvbu_env` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `sign` varchar(10) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_env
-- ----------------------------
INSERT INTO `lvbu_env` VALUES ('1', '开发环境', 'DE', '2017-01-16 14:58:27', '2017-01-16 14:58:28');
INSERT INTO `lvbu_env` VALUES ('2', '测试环境', 'QE', '2017-01-16 14:58:27', '2017-01-16 14:58:28');
INSERT INTO `lvbu_env` VALUES ('3', '线上环境', 'OE', '2017-01-16 14:58:27', '2017-01-16 14:58:28');

-- ----------------------------
-- Table structure for `lvbu_machine`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_machine`;
CREATE TABLE `lvbu_machine` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `ipaddr1` varchar(50) NOT NULL DEFAULT '',
  `ipaddr2` varchar(50) NOT NULL DEFAULT '',
  `env_id` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `adminurl` varchar(100) NOT NULL DEFAULT '',
  `content` varchar(100) NOT NULL DEFAULT '',
  `dtstatus` tinyint(1) NOT NULL DEFAULT '0',
  `tostatus` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for `lvbu_mirror`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_mirror`;
CREATE TABLE `lvbu_mirror` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `hubaddress` varchar(100) NOT NULL DEFAULT '',
  `mirrorgroup_id` int(10) unsigned DEFAULT NULL,
  `updated` datetime NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;



-- ----------------------------
-- Table structure for `lvbu_mirrorgroup`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_mirrorgroup`;
CREATE TABLE `lvbu_mirrorgroup` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for `lvbu_node`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_node`;
CREATE TABLE `lvbu_node` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `sign` varchar(50) NOT NULL DEFAULT '',
  `pro_id` int(10) unsigned NOT NULL,
  `doc_id` varchar(100) NOT NULL DEFAULT '',
  `mac_id` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_node
-- ----------------------------

-- ----------------------------
-- Table structure for `lvbu_peritem`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_peritem`;
CREATE TABLE `lvbu_peritem` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `sign` varchar(10) NOT NULL DEFAULT '',
  `menu_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_peritem
-- ----------------------------
INSERT INTO `lvbu_peritem` VALUES ('1', '主页显示', 'inds', '1');
INSERT INTO `lvbu_peritem` VALUES ('2', '项目管理', 'pros', '2');
INSERT INTO `lvbu_peritem` VALUES ('3', '主机管理', 'macs', '4');
INSERT INTO `lvbu_peritem` VALUES ('4', '配置查看', 'cons', '5');
INSERT INTO `lvbu_peritem` VALUES ('5', '镜像查看', 'mirs', '6');
INSERT INTO `lvbu_peritem` VALUES ('6', '设置查看', 'sets', '7');
INSERT INTO `lvbu_peritem` VALUES ('7', '主机编辑', 'mace', '4');
INSERT INTO `lvbu_peritem` VALUES ('8', '主机添加', 'maca', '4');
INSERT INTO `lvbu_peritem` VALUES ('9', '主机删除', 'macd', '4');
INSERT INTO `lvbu_peritem` VALUES ('10', '项目编辑', 'proe', '2');
INSERT INTO `lvbu_peritem` VALUES ('11', '项目添加', 'proa', '2');
INSERT INTO `lvbu_peritem` VALUES ('12', '项目删除', 'prod', '2');
INSERT INTO `lvbu_peritem` VALUES ('13', '节点查看', 'nods', '3');
INSERT INTO `lvbu_peritem` VALUES ('14', '节点编辑', 'node', '3');
INSERT INTO `lvbu_peritem` VALUES ('15', '节点添加', 'noda', '3');
INSERT INTO `lvbu_peritem` VALUES ('16', '节点删除', 'nodd', '3');
INSERT INTO `lvbu_peritem` VALUES ('17', '镜像编辑', 'mire', '6');
INSERT INTO `lvbu_peritem` VALUES ('18', '镜像添加', 'mira', '6');
INSERT INTO `lvbu_peritem` VALUES ('19', '镜像删除', 'mird', '6');

-- ----------------------------
-- Table structure for `lvbu_permenu`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_permenu`;
CREATE TABLE `lvbu_permenu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_permenu
-- ----------------------------
INSERT INTO `lvbu_permenu` VALUES ('1', '主页');
INSERT INTO `lvbu_permenu` VALUES ('2', '项目');
INSERT INTO `lvbu_permenu` VALUES ('3', '节点');
INSERT INTO `lvbu_permenu` VALUES ('4', '主机');
INSERT INTO `lvbu_permenu` VALUES ('5', '配置');
INSERT INTO `lvbu_permenu` VALUES ('6', '镜像');
INSERT INTO `lvbu_permenu` VALUES ('7', '系统');

-- ----------------------------
-- Table structure for `lvbu_position`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_position`;
CREATE TABLE `lvbu_position` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `sign` varchar(10) NOT NULL DEFAULT '',
  `permission` varchar(200) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_position
-- ----------------------------
INSERT INTO `lvbu_position` VALUES ('1', '运维经理', 'OS', 'pros,proa,inds,macs,maca,sets,mirs,mira,cons,cona,mace,macd', '2017-01-09 16:54:17', '2017-01-09 16:54:18');
INSERT INTO `lvbu_position` VALUES ('2', '运维工程师', 'OE', 'inds,proa,macs,maca,cons,mirs,mira,sets,', '2017-01-09 16:54:17', '2017-01-09 16:54:18');
INSERT INTO `lvbu_position` VALUES ('3', '产品经理', 'PO', 'inds,pros,proe,node,macs,mace,maca,mirs,', '2017-01-09 16:54:17', '2017-01-09 16:54:18');
INSERT INTO `lvbu_position` VALUES ('4', '项目经理', 'PM', '', '2017-01-09 16:54:17', '2017-01-09 16:54:18');
INSERT INTO `lvbu_position` VALUES ('5', '研发工程师', 'DE', 'inds,macs', '2017-01-09 16:54:17', '2017-01-09 16:54:18');
INSERT INTO `lvbu_position` VALUES ('6', '测试工程师', 'QA', 'inds,pros,nods,', '2017-01-09 16:54:17', '2017-01-09 16:54:18');
INSERT INTO `lvbu_position` VALUES ('8', '运营经理', 'OM', '', '2017-01-09 16:54:17', '2017-01-09 16:54:18');
INSERT INTO `lvbu_position` VALUES ('9', '运营人员', 'OO', '', '2017-01-09 16:54:17', '2017-01-09 16:54:18');
INSERT INTO `lvbu_position` VALUES ('10', '设计师', 'DS', '', '2017-01-23 09:30:44', '2017-01-23 09:30:47');

-- ----------------------------
-- Table structure for `lvbu_project`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_project`;
CREATE TABLE `lvbu_project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `sign` varchar(50) NOT NULL DEFAULT '',
  `git` varchar(100) NOT NULL DEFAULT '',
  `gituser` varchar(50) NOT NULL DEFAULT '',
  `gitpass` varchar(50) NOT NULL DEFAULT '',
  `insfile` varchar(500) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_project
-- ----------------------------

-- ----------------------------
-- Table structure for `lvbu_user`
-- ----------------------------
DROP TABLE IF EXISTS `lvbu_user`;
CREATE TABLE `lvbu_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) NOT NULL DEFAULT '',
  `passwd` varchar(50) NOT NULL DEFAULT '',
  `nick` varchar(50) NOT NULL DEFAULT '',
  `sex` tinyint(10) unsigned NOT NULL DEFAULT '0',
  `email` varchar(50) NOT NULL DEFAULT '',
  `phone` varchar(50) NOT NULL DEFAULT '',
  `avatar` varchar(100) NOT NULL DEFAULT '',
  `position_id` int(10) unsigned NOT NULL,
  `permission` varchar(200) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `status` int(2) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lvbu_user
-- ----------------------------
INSERT INTO `lvbu_user` VALUES ('1', 'root', 'c4ca4238a0b923820dcc509a6f75849b', '超级管理员', '1', 'root@admin.cn', '11111111111', 'user4-128x128.jpg', '1', 'DE,QE,OE', '2017-01-09 14:32:22', '2018-01-09 00:00:00', '0');
INSERT INTO `lvbu_user` VALUES ('2', 'guest', 'c4ca4238a0b923820dcc509a6f75849b', '测试', '1', 'guest@admin.cn', '2222222222', 'user4-128x128.jpg', '1', 'DE,QE,OE', '2017-01-09 00:00:00', '2017-01-22 17:25:43', '0');
