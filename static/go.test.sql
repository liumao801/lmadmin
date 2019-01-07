/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : go.test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-01-07 22:44:23
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
  `real_name` varchar(255) DEFAULT NULL COMMENT '真实姓名',
  `tel` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `is_super` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否超级管理员 0否；1是',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；2已删除',
  `face` varchar(255) DEFAULT NULL COMMENT '头像',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '关联用户表id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '最后更新时间',
  `login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `login_ip` varchar(255) DEFAULT NULL,
  `name` varchar(32) NOT NULL DEFAULT '',
  `remember_passwd` varchar(255) NOT NULL DEFAULT '' COMMENT '记住密码的密码',
  `remember_out` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '记住密码的过期时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='管理员信息表';

-- ----------------------------
-- Records of lm_admin
-- ----------------------------
INSERT INTO `lm_admin` VALUES ('1', 'admin', 'e3ceb5881a0a1fdaad01296d7554868d', '管理员', '13838383838', '138@gmail.com', '1', '1', '/static/upload/admincenter/2019-01/04/e6BCD_1546614717.jpg', '0', '2018-12-01 03:51:51', '2018-12-01 03:51:51', '2019-01-07 12:59:53', '::1', '', 'a68738287709c96abb8e15eef94fd910', '1546865993');
INSERT INTO `lm_admin` VALUES ('2', 'test', '111111', '测试', '13909820984', '213@qq.com', '0', '1', null, '0', '2018-12-13 11:51:51', '2018-12-13 11:51:51', null, null, '', '0', '0');
INSERT INTO `lm_admin` VALUES ('3', 'customer', '111111', '游客', '13288743837', '123@qq.com', '0', '0', '', '0', '2018-12-13 11:51:51', '2018-12-13 11:51:51', null, null, '', '0', '0');
INSERT INTO `lm_admin` VALUES ('4', 'tester', '96e79218965eb72c92a549dd5a330112', '测试用户', '13221234543', '231@gmail.com', '0', '1', '', '0', '2018-12-13 11:51:51', '2018-12-13 11:51:51', null, null, '', '0', '0');

-- ----------------------------
-- Table structure for lm_admin_copy
-- ----------------------------
DROP TABLE IF EXISTS `lm_admin_copy`;
CREATE TABLE `lm_admin_copy` (
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
-- Records of lm_admin_copy
-- ----------------------------
INSERT INTO `lm_admin_copy` VALUES ('1', 'admin', '222222', null, '管理员', '13838383838', '138@gmail.com', '1', '1', '0', null, null);
INSERT INTO `lm_admin_copy` VALUES ('2', 'test', '111111', null, '测试', '13909820984', '213@qq.com', '0', '1', '0', null, null);
INSERT INTO `lm_admin_copy` VALUES ('3', 'customer', '111111', null, '游客', '13288743837', '123@qq.com', '0', '1', '0', null, null);

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
  `created_at` datetime NOT NULL COMMENT '操作时间',
  `ip` varchar(255) NOT NULL COMMENT '操作IP',
  `menu` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='后台操作日志';

-- ----------------------------
-- Records of lm_admin_log
-- ----------------------------
INSERT INTO `lm_admin_log` VALUES ('1', '1', 'admin', '14', '0', '/admin/role/index', '{}', '2018-12-15 12:50:10', '::1', '');
INSERT INTO `lm_admin_log` VALUES ('2', '1', 'admin', '14', '0', '/admin/role/index', '{}', '2018-12-15 12:50:25', '::1', '');
INSERT INTO `lm_admin_log` VALUES ('3', '1', 'admin', '14', '0', '/admin/role/index', '{}', '2018-12-15 12:50:43', '::1', '');

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
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='文章信息';

-- ----------------------------
-- Records of lm_article
-- ----------------------------

-- ----------------------------
-- Table structure for lm_menu
-- ----------------------------
DROP TABLE IF EXISTS `lm_menu`;
CREATE TABLE `lm_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0菜单分类无链接；1菜单有链接；2页面菜单',
  `name` varchar(255) DEFAULT NULL COMMENT '菜单名称',
  `parent_id` int(11) unsigned DEFAULT '0',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序；越小越靠前',
  `icon` varchar(255) DEFAULT 'fa-circle-o' COMMENT '菜单图标',
  `url_for` varchar(256) NOT NULL DEFAULT '' COMMENT 'beego URLFor',
  `url` varchar(255) DEFAULT NULL COMMENT 'url地址',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；2已删除',
  `is_check` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否检测权限 1检测；0不检测',
  `created_at` int(10) unsigned DEFAULT NULL,
  `updated_at` int(10) unsigned DEFAULT NULL,
  `show` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 COMMENT='菜单资源信息表';

-- ----------------------------
-- Records of lm_menu
-- ----------------------------
INSERT INTO `lm_menu` VALUES ('1', '1', '首页', null, '1', 'fa-dashboard', 'HomeController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('2', '0', '系统管理', '0', '100', 'fa-dashboard', 'IndexController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('7', '1', '系统配置', '2', '100', 'fa-dashboard', 'IndexController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('8', '1', '日志管理', '2', '100', 'fa-dashboard', 'IndexController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('9', '1', '网站配置', '7', '100', 'fa-dashboard', 'IndexController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('10', '1', '注册管理', '7', '100', 'fa-dashboard', 'IndexController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('11', '1', '权限管理', '12', '100', 'fa fa-balance-scale', '', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('12', '0', '系统设置', null, '100', '', '', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('13', '1', '菜单管理', '11', '100', 'fa-pinterest', 'MenuController.Index', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('14', '1', '角色管理', '11', '100', 'fa-users', 'RoleController.Index', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('15', '1', '用户管理', '11', '100', 'fa-user', 'AdminController.Index', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('17', '0', '业务菜单', null, '170', '', '', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('18', '1', '课程资源(空)', '17', '100', 'fa fa-book', '', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('20', '2', '编辑', '13', '100', 'fa fa-pencil', 'MenuController.Edit', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('21', '2', '编辑', '15', '100', 'fa fa-pencil', 'AdminController.Edit', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('22', '2', '删除', '13', '100', 'fa fa-trash', 'MenuController.Delete', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('23', '2', '删除', '15', '100', 'fa fa-trash', 'AdminController.Delete', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('24', '2', '编辑', '14', '100', 'fa fa-pencil', 'RoleController.Edit', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('25', '2', '删除', '14', '100', 'fa fa-trash', 'RoleController.Delete', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('26', '2', '分配资源', '14', '100', 'fa fa-th', 'RoleController.Allocate', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('33', '0', 'CMF系统', null, '100', 'fa-file-text', '', null, '1', '0', '2018', '2018', '0');
INSERT INTO `lm_menu` VALUES ('34', '1', '文章菜单分类', '33', '100', 'fa-ge', '', null, '1', '0', null, null, '0');
INSERT INTO `lm_menu` VALUES ('35', '1', '文章分类管理', '34', '100', 'fa-adn', 'MenuWebController.Index', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('36', '1', '文章管理', '34', '100', 'fa-file-text-o', 'MenuWebController.Index', null, '1', '1', '2019', '2019', '0');

-- ----------------------------
-- Table structure for lm_menu_copy
-- ----------------------------
DROP TABLE IF EXISTS `lm_menu_copy`;
CREATE TABLE `lm_menu_copy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '菜单名称',
  `url_for` varchar(256) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0菜单分类无链接；1菜单有链接；2页面菜单',
  `icon` varchar(255) DEFAULT 'fa-circle-o' COMMENT '菜单图标',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；2已删除',
  `show` tinyint(1) unsigned DEFAULT '1' COMMENT '1显示；0隐藏',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序；越小越靠前',
  `parent_id` int(11) unsigned DEFAULT '0',
  `created_at` int(10) unsigned DEFAULT NULL,
  `updated_at` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='菜单资源信息表';

-- ----------------------------
-- Records of lm_menu_copy
-- ----------------------------
INSERT INTO `lm_menu_copy` VALUES ('1', '首页', 'admin/IndexController.index', 'admin/index/index', '1', 'fa-dashboard', '1', '1', '100', '0', null, null);
INSERT INTO `lm_menu_copy` VALUES ('2', '系统管理', 'admin/IndexController.index', 'admin/index/index', '0', 'fa-dashboard', '1', '1', '100', '0', null, null);
INSERT INTO `lm_menu_copy` VALUES ('3', '权限管理', 'admin/IndexController.index', 'admin/index/index', '0', 'fa-dashboard', '1', '1', '100', '0', null, null);
INSERT INTO `lm_menu_copy` VALUES ('4', '菜单管理', 'admin/IndexController.index', 'admin/index/index', '1', 'fa-dashboard', '1', '1', '100', '3', null, null);
INSERT INTO `lm_menu_copy` VALUES ('5', '角色管理', 'admin/IndexController.index', 'admin/index/index', '1', 'fa-dashboard', '1', '1', '100', '3', null, null);
INSERT INTO `lm_menu_copy` VALUES ('6', '管理员', 'admin/IndexController.index', 'admin/index/index', '1', 'fa-dashboard', '1', '1', '100', '3', null, null);
INSERT INTO `lm_menu_copy` VALUES ('7', '系统设置', 'admin/IndexController.index', 'admin/index/index', '0', 'fa-dashboard', '1', '1', '100', '2', null, null);
INSERT INTO `lm_menu_copy` VALUES ('8', '日志管理', 'admin/IndexController.index', 'admin/index/index', '1', 'fa-dashboard', '1', '1', '100', '2', null, null);
INSERT INTO `lm_menu_copy` VALUES ('9', '网站配置', 'admin/IndexController.index', 'admin/index/index', '1', 'fa-dashboard', '1', '1', '100', '7', null, null);
INSERT INTO `lm_menu_copy` VALUES ('10', '注册管理', 'admin/IndexController.index', 'admin/index/index', '1', 'fa-dashboard', '1', '1', '100', '7', null, null);

-- ----------------------------
-- Table structure for lm_menu_web
-- ----------------------------
DROP TABLE IF EXISTS `lm_menu_web`;
CREATE TABLE `lm_menu_web` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '菜单名称',
  `icon` varchar(255) DEFAULT NULL COMMENT '图标',
  `type` tinyint(3) unsigned NOT NULL COMMENT '1频道页；2跳转页；3栏目页；4单页',
  `parent_id` int(10) unsigned DEFAULT '0' COMMENT '父级Id',
  `list_tpl` varchar(255) DEFAULT NULL COMMENT '列表页模板',
  `page_tpl` varchar(255) DEFAULT NULL COMMENT '页面模型',
  `article_tpl` varchar(255) DEFAULT NULL COMMENT '文章页模板',
  `url` varchar(255) DEFAULT NULL COMMENT 'type=2 时的跳转地址',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序',
  `img` varchar(255) DEFAULT NULL COMMENT 'type=4 的缩略图',
  `seo_title` varchar(255) DEFAULT NULL,
  `seo_desc` varchar(255) DEFAULT NULL,
  `content` text COMMENT 'type=4 的页面内容',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='前端菜单/文章分类';

-- ----------------------------
-- Records of lm_menu_web
-- ----------------------------
INSERT INTO `lm_menu_web` VALUES ('1', '就发大水', null, '1', '0', 'jk', '1', '；l', null, '1', '100', null, null, null, null);
INSERT INTO `lm_menu_web` VALUES ('2', '范德萨', null, '1', '1', '', '0', '', '', '1', '100', '', '123456', '', '');
INSERT INTO `lm_menu_web` VALUES ('3', '测试二', null, '2', '2', '', '0', '', '', '1', '100', '', '21', '21', '');
INSERT INTO `lm_menu_web` VALUES ('6', '首页', 'fa-home', '2', null, '', null, '', '/home', '1', '100', '', '', '', '');

-- ----------------------------
-- Table structure for lm_role
-- ----------------------------
DROP TABLE IF EXISTS `lm_role`;
CREATE TABLE `lm_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '角色名称',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序',
  `mark` varchar(255) NOT NULL COMMENT '角色标识',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0停用；1正常；2已删除',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色信息表';

-- ----------------------------
-- Records of lm_role
-- ----------------------------
INSERT INTO `lm_role` VALUES ('1', '管理员', '0', 'admin', '1', null, null);
INSERT INTO `lm_role` VALUES ('2', '游客', '100', 'tester', '0', '2018-12-13 11:51:51', '2018-12-13 11:51:51');

-- ----------------------------
-- Table structure for lm_role_admin_rel
-- ----------------------------
DROP TABLE IF EXISTS `lm_role_admin_rel`;
CREATE TABLE `lm_role_admin_rel` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL,
  `admin_id` int(10) unsigned NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='角色用户关联表';

-- ----------------------------
-- Records of lm_role_admin_rel
-- ----------------------------
INSERT INTO `lm_role_admin_rel` VALUES ('3', '1', '4', '2018-12-13 11:30:41', '2018-12-13 11:30:41');

-- ----------------------------
-- Table structure for lm_role_admin_rel_copy
-- ----------------------------
DROP TABLE IF EXISTS `lm_role_admin_rel_copy`;
CREATE TABLE `lm_role_admin_rel_copy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL,
  `admin_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色用户关联表';

-- ----------------------------
-- Records of lm_role_admin_rel_copy
-- ----------------------------

-- ----------------------------
-- Table structure for lm_role_copy
-- ----------------------------
DROP TABLE IF EXISTS `lm_role_copy`;
CREATE TABLE `lm_role_copy` (
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
-- Records of lm_role_copy
-- ----------------------------

-- ----------------------------
-- Table structure for lm_role_menu_rel
-- ----------------------------
DROP TABLE IF EXISTS `lm_role_menu_rel`;
CREATE TABLE `lm_role_menu_rel` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL,
  `menu_id` int(10) unsigned NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `created_at` datetime NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8 COMMENT='角色菜单关联表';

-- ----------------------------
-- Records of lm_role_menu_rel
-- ----------------------------
INSERT INTO `lm_role_menu_rel` VALUES ('15', '1', '1', '2018-12-13 12:26:31', '2018-12-13 12:26:31', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('16', '1', '2', '2018-12-13 12:26:31', '2018-12-13 12:26:31', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('17', '1', '7', '2018-12-13 12:26:31', '2018-12-13 12:26:31', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('18', '1', '9', '2018-12-13 12:26:31', '2018-12-13 12:26:31', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('19', '1', '10', '2018-12-13 12:26:31', '2018-12-13 12:26:31', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('20', '1', '8', '2018-12-13 12:26:31', '2018-12-13 12:26:31', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('21', '1', '18', '2018-12-13 12:26:31', '2018-12-13 12:26:31', '0000-00-00 00:00:00');

-- ----------------------------
-- Table structure for lm_role_menu_rel_copy
-- ----------------------------
DROP TABLE IF EXISTS `lm_role_menu_rel_copy`;
CREATE TABLE `lm_role_menu_rel_copy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(10) unsigned NOT NULL,
  `menu_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `created` datetime NOT NULL,
  `updated_at` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色菜单关联表';

-- ----------------------------
-- Records of lm_role_menu_rel_copy
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
INSERT INTO `session` VALUES ('1bb97604d8de95aa30bdaab1df251760', 0x0EFF81040102FF8200011001100000FE0118FF82000106737472696E670C07000561646D696E1E6C69756D616F3830312F6C6D61646D696E2F6D6F64656C732E41646D696EFF8F0301010541646D696E01FF900001100102496401040001085265616C4E616D65010C000108557365726E616D65010C000106506173737764010C000107497353757065720102000106537461747573010400010E52656D656D626572506173737764010C00010B52656D656D6265724F7574010400010354656C010C000105456D61696C010C00010446616365010C000107526F6C6549647301FF9200010C526F6C6541646D696E52656C01FF9400010E4D656E7555726C466F724C69737401FF9600010943726561746564417401FF8C00010955706461746564417401FF8C00000013FF91020101055B5D696E7401FF92000104000025FF93020101165B5D2A6D6F64656C732E526F6C6541646D696E52656C01FF940001FF8E000046FF8D030102FF8E000105010249640104000104526F6C6501FF8A00010541646D696E01FF9000010943726561746564417401FF8C00010955706461746564417401FF8C000000FF80FF8903010104526F6C6501FF8A0001090102496401040001044E616D65010C0001044D61726B010C000104536F72740106000106537461747573010400010943726561746564417401FF8C00010955706461746564417401FF8C00010C526F6C6541646D696E52656C01FF9400010B526F6C654D656E7552656C01FF9800000010FF8B0501010454696D6501FF8C00000024FF97020101155B5D2A6D6F64656C732E526F6C654D656E7552656C01FF980001FF88000045FF87030102FF88000105010249640104000104526F6C6501FF8A0001044D656E7501FF8400010943726561746564417401FF8C00010955706461746564417401FF8C000000FFD0FF83030102FF840001110102496401040001044E616D65010C000106506172656E7401FF84000104547970650104000104536F72740106000104536F6E7301FF86000106536F6E4E756D010400010449636F6E010C0001074C696E6B55726C010C00010655726C466F72010C0001054C6576656C010400010C48746D6C44697361626C6564010400010B526F6C654D656E7552656C01FF9800010653746174757301040001074973436865636B010400010943726561746564417401FF8C00010955706461746564417401FF8C0000000DFF85020102FF860001FF84000016FF95020101085B5D737472696E6701FF9600010C0000FE023DFF90FE023801020109E7AEA1E79086E59198010561646D696E01206533636562353838316130613166646161643031323936643735353438363864010101020120613638373338323837373039633936616262386531356565663934666439313001FCB82ABC6C010B3133383338333833383338010D31333840676D61696C2E636F6D01342F7374617469632F75706C6F61642F61646D696E63656E7465722F323031382D31322F75736572332D313238783132382E6A7067031414486F6D65436F6E74726F6C6C65722E496E64657815496E646578436F6E74726F6C6C65722E496E64657815496E646578436F6E74726F6C6C65722E496E64657815496E646578436F6E74726F6C6C65722E496E64657815496E646578436F6E74726F6C6C65722E496E64657815496E646578436F6E74726F6C6C65722E496E6465780000144D656E75436F6E74726F6C6C65722E496E646578134D656E75436F6E74726F6C6C65722E45646974154D656E75436F6E74726F6C6C65722E44656C65746514526F6C65436F6E74726F6C6C65722E496E64657813526F6C65436F6E74726F6C6C65722E4564697415526F6C65436F6E74726F6C6C65722E44656C65746517526F6C65436F6E74726F6C6C65722E416C6C6F636174651541646D696E436F6E74726F6C6C65722E496E6465781441646D696E436F6E74726F6C6C65722E456469741641646D696E436F6E74726F6C6C65722E44656C6574650000010F010000000ED39CC6570000000001E0010F010000000ED39CC6570000000001E000, '1544879089');

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
