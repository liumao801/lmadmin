/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : go.test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2018-12-03 16:05:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for lm_admin
-- ----------------------------
DROP TABLE IF EXISTS `lm_admin`;
CREATE TABLE `lm_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL COMMENT '用户名',
  `passwd` varchar(255) NOT NULL COMMENT '密码',
  `face` varchar(255) DEFAULT NULL COMMENT '头像',
  `name` varchar(255) DEFAULT NULL COMMENT '真实姓名',
  `tel` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `is_super` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否超级管理员 0否；1是',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；2已删除',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '关联用户表id',
  `created_at` int(10) unsigned DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(10) unsigned DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='管理员信息表';

-- ----------------------------
-- Records of lm_admin
-- ----------------------------
INSERT INTO `lm_admin` VALUES ('1', 'admin', '111111', null, '管理员', '13838383838', '138@gmail.com', '1', '1', '0', null, null);
INSERT INTO `lm_admin` VALUES ('2', 'test', '111111', null, '测试', '13909820984', '213@qq.com', '0', '1', '0', null, null);
INSERT INTO `lm_admin` VALUES ('3', 'customer', '111111', null, '游客', '13288743837', '123@qq.com', '0', '1', '0', null, null);

-- ----------------------------
-- Table structure for lm_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `lm_admin_log`;
CREATE TABLE `lm_admin_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int(10) unsigned NOT NULL COMMENT '用户id',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '操作用户名',
  `menu_id` int(10) unsigned NOT NULL COMMENT '用户id',
  `menu_name` varchar(255) NOT NULL DEFAULT '0' COMMENT '菜单名称',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT 'url地址',
  `params` varchar(255) NOT NULL DEFAULT '0' COMMENT '附加参数',
  `created_at` int(10) unsigned NOT NULL COMMENT '操作时间',
  `ip` varchar(255) NOT NULL COMMENT '操作IP',
  `menu` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台操作日志';

-- ----------------------------
-- Records of lm_admin_log
-- ----------------------------

-- ----------------------------
-- Table structure for lm_article
-- ----------------------------
DROP TABLE IF EXISTS `lm_article`;
CREATE TABLE `lm_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '文章标题',
  `subtitle` varchar(255) DEFAULT NULL COMMENT '副标题',
  `logo` varchar(255) DEFAULT NULL COMMENT 'logo图片',
  `keywords` varchar(255) DEFAULT NULL COMMENT '关键字',
  `desc` varchar(255) DEFAULT NULL COMMENT '简述',
  `content` text NOT NULL COMMENT '文章内容',
  `menu_id` int(11) DEFAULT NULL,
  `video` varchar(255) DEFAULT NULL COMMENT '视频地址',
  `comment_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0不允许评论；1可评论',
  `comment_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '评论数量',
  `is_back` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0回收站文章；1正常文章',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0待审核；1审核通过；2审核不通过',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lm_article
-- ----------------------------

-- ----------------------------
-- Table structure for lm_menu
-- ----------------------------
DROP TABLE IF EXISTS `lm_menu`;
CREATE TABLE `lm_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '菜单名称',
  `url` varchar(255) NOT NULL,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0菜单分类无链接；1菜单有链接；2页面菜单',
  `icon` varchar(255) DEFAULT NULL COMMENT '菜单图标',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；2已删除',
  `show` tinyint(1) unsigned DEFAULT '1' COMMENT '1显示；0隐藏',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序；越小越靠前',
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级菜单id',
  `created_at` int(10) unsigned DEFAULT NULL,
  `updated_at` int(10) unsigned DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `url_for` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='菜单资源信息表';

-- ----------------------------
-- Records of lm_menu
-- ----------------------------

-- ----------------------------
-- Table structure for lm_menu_web
-- ----------------------------
DROP TABLE IF EXISTS `lm_menu_web`;
CREATE TABLE `lm_menu_web` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '菜单名称',
  `type` tinyint(3) unsigned NOT NULL COMMENT '1频道页；2跳转页；3栏目页；4单页',
  `model_id` int(11) DEFAULT NULL COMMENT '模型ID',
  `par_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级Id',
  `list_tpl` varchar(255) DEFAULT NULL COMMENT '列表页模板',
  `article_tpl` varchar(255) DEFAULT NULL COMMENT '文章页模板',
  `url` varchar(255) DEFAULT NULL COMMENT 'type=2 时的跳转地址',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；',
  `sort` tinyint(3) unsigned DEFAULT '100' COMMENT '排序',
  `img` varchar(255) DEFAULT NULL COMMENT 'type=4 的缩略图',
  `seo_title` varchar(255) DEFAULT NULL,
  `seo_desc` varchar(255) DEFAULT NULL,
  `content` text COMMENT 'type=4 的页面内容',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='前端菜单/文章分类';

-- ----------------------------
-- Records of lm_menu_web
-- ----------------------------

-- ----------------------------
-- Table structure for lm_role
-- ----------------------------
DROP TABLE IF EXISTS `lm_role`;
CREATE TABLE `lm_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '角色名称',
  `mark` varchar(255) NOT NULL COMMENT '角色标识',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；2已删除',
  `created_at` int(10) unsigned DEFAULT NULL,
  `updated_at` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='角色信息表';

-- ----------------------------
-- Records of lm_role
-- ----------------------------

-- ----------------------------
-- Table structure for lm_role_admin_rel
-- ----------------------------
DROP TABLE IF EXISTS `lm_role_admin_rel`;
CREATE TABLE `lm_role_admin_rel` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL,
  `admin_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色用户关联表';

-- ----------------------------
-- Records of lm_role_admin_rel
-- ----------------------------

-- ----------------------------
-- Table structure for lm_role_menu_rel
-- ----------------------------
DROP TABLE IF EXISTS `lm_role_menu_rel`;
CREATE TABLE `lm_role_menu_rel` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL,
  `menu_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色菜单关联表';

-- ----------------------------
-- Records of lm_role_menu_rel
-- ----------------------------

-- ----------------------------
-- Table structure for lm_test
-- ----------------------------
DROP TABLE IF EXISTS `lm_test`;
CREATE TABLE `lm_test` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lm_test
-- ----------------------------
INSERT INTO `lm_test` VALUES ('1', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('2', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('3', 'Name 3', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('4', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('5', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('6', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('7', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('8', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('9', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('10', 'Name 1', 'Test Data 1');
INSERT INTO `lm_test` VALUES ('11', 'Name 1', 'Test Data 1');

-- ----------------------------
-- Table structure for lm_user
-- ----------------------------
DROP TABLE IF EXISTS `lm_user`;
CREATE TABLE `lm_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `passwd` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL COMMENT '姓名',
  `status` tinyint(1) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of lm_user
-- ----------------------------
INSERT INTO `lm_user` VALUES ('1', 'admin', '111111', null, '1', null, null);
INSERT INTO `lm_user` VALUES ('2', 'test', '111111', 'test', '1', '1543725789', '1543725789');
INSERT INTO `lm_user` VALUES ('3', 'test2', '111111', 'test2', '1', '1543725853', '1543725853');
INSERT INTO `lm_user` VALUES ('4', 'lm', 'lm', '刘茂良', '1', '1543794534', '1543794534');

-- ----------------------------
-- Table structure for session
-- ----------------------------
DROP TABLE IF EXISTS `session`;
CREATE TABLE `session` (
  `session_key` char(64) NOT NULL,
  `session_data` blob,
  `session_expiry` int(11) unsigned NOT NULL,
  PRIMARY KEY (`session_key`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='mysql session 存储表';

-- ----------------------------
-- Records of session
-- ----------------------------

-- ----------------------------
-- Table structure for test
-- ----------------------------
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of test
-- ----------------------------
INSERT INTO `test` VALUES ('1', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('2', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('3', 'Name 3', 'Test Data 3');
INSERT INTO `test` VALUES ('4', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('5', 'Updated Name 5-2', 'Updated Name');
INSERT INTO `test` VALUES ('7', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('8', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('9', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('10', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('11', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('12', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('13', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('14', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('15', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('16', 'Name 1', 'Test Data 1');
INSERT INTO `test` VALUES ('17', 'Name 1', 'Test Data 1');
