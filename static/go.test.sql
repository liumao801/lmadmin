/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : go.test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-07-08 18:24:21
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
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0停用；1正常；-1已删除',
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
INSERT INTO `lm_admin` VALUES ('1', 'admin', 'e3ceb5881a0a1fdaad01296d7554868d', '管理员', '13838383838', '138@gmail.com', '1', '1', '/static/upload/admincenter/2019-01/04/e6BCD_1546614717.jpg', '0', '2018-11-30 19:51:51', '2018-11-30 19:51:51', '2019-06-21 13:47:41', '127.0.0.1', '', 'f109b2ee1ccab7f4bb87775af81f92a7', '1561150061');
INSERT INTO `lm_admin` VALUES ('2', 'test', '111111', '测试', '13909820984', '213@qq.com', '0', '1', null, '0', '2018-12-13 11:51:51', '2018-12-13 11:51:51', null, null, '', '0', '0');
INSERT INTO `lm_admin` VALUES ('3', 'customer', '111111', '游客', '13288743837', '123@qq.com', '0', '0', '', '0', '2018-12-13 11:51:51', '2018-12-13 11:51:51', null, null, '', '0', '0');
INSERT INTO `lm_admin` VALUES ('4', 'tester', '96e79218965eb72c92a549dd5a330112', '测试用户', '13221234543', '231@gmail.com', '0', '-1', '', '0', null, null, null, '', '', '', '0');

-- ----------------------------
-- Table structure for lm_admin_log
-- ----------------------------
DROP TABLE IF EXISTS `lm_admin_log`;
CREATE TABLE `lm_admin_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int(10) unsigned NOT NULL COMMENT '用户id',
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '操作用户名',
  `path` varchar(249) NOT NULL DEFAULT '' COMMENT 'url地址',
  `method` varchar(10) NOT NULL COMMENT 'get post put delete header option',
  `input` text NOT NULL COMMENT '附加参数',
  `created_at` datetime DEFAULT NULL COMMENT '操作时间',
  `ip` varchar(16) NOT NULL COMMENT '操作IP',
  PRIMARY KEY (`id`),
  KEY `admin_id` (`admin_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='后台操作日志';

-- ----------------------------
-- Records of lm_admin_log
-- ----------------------------
INSERT INTO `lm_admin_log` VALUES ('1', '1', 'admin', '/admin/commonset/index', 'GET', '{}', '2019-07-08 18:23:42', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('2', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-07-08 18:23:42', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('3', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-07-08 18:23:43', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('4', '1', 'admin', '/admin/commonset/edit/10', 'GET', '{}', '2019-07-08 18:23:45', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('5', '1', 'admin', '/admin/commonset/edit/10', 'POST', '{\"Id\":[\"10\"],\"Name\":[\"using_module\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"前端使用模块|{\\\"portal\\\":\\\"门户网站模板\\\",\\\"home\\\":\\\"内容管理系统模板\\\"}\"],\"Type\":[\"home_conf\"],\"Value\":[\"portal\"]}', '2019-07-08 18:23:48', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('6', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-07-08 18:23:48', '127.0.0.1');

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
  `menu_web_id` int(11) DEFAULT NULL,
  `video` varchar(255) DEFAULT NULL COMMENT '视频地址',
  `comment_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0不允许评论；1可评论',
  `comment_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '评论数量',
  `is_back` tinyint(1) NOT NULL DEFAULT '1' COMMENT '-1回收站文章；1正常文章',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0审核中；1审核通过；-1审核失败；',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `view_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文章阅读次数',
  `author` varchar(32) NOT NULL DEFAULT 'admin' COMMENT '文章作者',
  `img` varchar(255) DEFAULT NULL COMMENT '文章头图，标题旁边图片',
  `pub_at` datetime DEFAULT NULL COMMENT '发布开始时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='文章信息';

-- ----------------------------
-- Records of lm_article
-- ----------------------------
INSERT INTO `lm_article` VALUES ('1', 'request', '热弯', '', '2,123,1234', ' 范德萨', '<p>&nbsp; &nbsp; &nbsp; &nbsp; 放给大家看联发科多斯拉克莲富大厦</p>\n\n<p>nihaoa</p>\n\n<p style=\"text-align: right;\">shide</p>\n\n<p style=\"text-align: right;\">2019年7月1日 23:35:00</p>\n', '8', '2', '0', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-06 16:22:49', '0', 'admin', '/static/upload/article/2019-07/06/62siT_1562401369.jpg', null);
INSERT INTO `lm_article` VALUES ('2', '二级分类', '范德萨', '', '范德萨,123,1234', '范德萨', '<p>范德萨范德萨</p>\n', '9', '', '1', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-01 21:00:34', '0', 'admin', null, null);
INSERT INTO `lm_article` VALUES ('3', '顶级分类', '范德萨', '', '发', '范德萨', '<p>佛挡杀佛的</p>\n', '8', '2432', '1', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-01 21:33:10', '0', 'admin', null, null);
INSERT INTO `lm_article` VALUES ('4', '顶级分类', '范德萨', '', '发', '范德萨', '<p>范德萨发达</p>\n', '9', '', '0', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-01 21:33:49', '0', 'admin', null, null);
INSERT INTO `lm_article` VALUES ('5', '顶级分类', '范德萨', '', '范德萨,123,1234', '范德萨', '<p>就罚款落地随机</p>\n', '8', '', '0', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-01 21:35:08', '0', 'admin', null, null);
INSERT INTO `lm_article` VALUES ('6', 'fds432', '范德萨', '', '范德萨,123,1234', '发', '<p>发的</p>\n', '10', '', '0', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-01 21:36:35', '0', 'admin', null, null);
INSERT INTO `lm_article` VALUES ('7', 'fds432', '范德萨', '', '范德萨,123,1234', '范德萨', '<p>范德萨范德萨范德萨范德萨范德萨范德萨ff</p>\n', '8', '', '0', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-01 21:40:05', '0', 'admin', null, null);
INSERT INTO `lm_article` VALUES ('8', '前端使用模板|{\"protal\":\"门户网站\"}', '范德萨', '', '范德萨,123,1234', '范德萨', '<p>封455</p>\n', '10', '', '0', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-06 20:14:00', '0', 'admin', '/static/upload/article/2019-07/06/258c0_1562415240.jpg', '2019-07-06 20:13:52');
INSERT INTO `lm_article` VALUES ('9', '测试图片类型', '范德萨', '', '范德萨,123,1234', '范德萨', '<p><img alt=\"\" src=\"/static/upload/upload/2019-07/08/WC53c_1562576526.png\" style=\"width: 200px; height: 200px;\" />范德萨发0000122</p>\n\n<figure class=\"easyimage easyimage-full\"><img alt=\"\" src=\"/static/upload/upload/2019-07/06/58F83_1562416556.jpg\" width=\"135\" />\n<figcaption>&nbsp;</figcaption>\n</figure>\n', '8', '', '1', '0', '0', '1', '2019-07-01 21:00:34', '2019-07-08 17:02:42', '0', 'admin', '/static/upload/article/2019-07/03/zKWk5_1562146977.png', '2019-07-06 20:30:33');
INSERT INTO `lm_article` VALUES ('10', '讲课费', '京东方', '', '发的', '就开了发', '<p>范德萨范德萨范德萨范德萨范德萨范德萨</p>\n', '10', '', '1', '0', '0', '1', '2019-07-06 19:54:58', '2019-07-06 19:54:58', '0', '', '/static/upload/article/2019-07/06/5cF3E_1562414098.jpg', null);

-- ----------------------------
-- Table structure for lm_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `lm_article_tag`;
CREATE TABLE `lm_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL COMMENT '标签名称',
  `icon` varchar(255) NOT NULL COMMENT '标签 logo',
  `status` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='文章标签';

-- ----------------------------
-- Records of lm_article_tag
-- ----------------------------
INSERT INTO `lm_article_tag` VALUES ('1', '最新', '/static/upload/articletag/2019-07/03/17hqG_1562162851.png', '1');
INSERT INTO `lm_article_tag` VALUES ('2', '热门', '/static/upload/articletag/2019-07/04/n76u7_1562228972.jpg', '1');

-- ----------------------------
-- Table structure for lm_article_tag_rel
-- ----------------------------
DROP TABLE IF EXISTS `lm_article_tag_rel`;
CREATE TABLE `lm_article_tag_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL,
  `article_tag_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 COMMENT='文章-文章标签 关联关系表';

-- ----------------------------
-- Records of lm_article_tag_rel
-- ----------------------------
INSERT INTO `lm_article_tag_rel` VALUES ('22', '9', '1', '2019-07-08 17:02:42', '2019-07-08 17:02:42');
INSERT INTO `lm_article_tag_rel` VALUES ('6', '1', '1', '2019-07-06 16:22:49', '2019-07-06 16:22:49');
INSERT INTO `lm_article_tag_rel` VALUES ('7', '1', '2', '2019-07-06 16:22:49', '2019-07-06 16:22:49');
INSERT INTO `lm_article_tag_rel` VALUES ('8', '10', '2', '2019-07-06 19:54:58', '2019-07-06 19:54:58');
INSERT INTO `lm_article_tag_rel` VALUES ('23', '9', '2', '2019-07-08 17:02:42', '2019-07-08 17:02:42');

-- ----------------------------
-- Table structure for lm_common_set
-- ----------------------------
DROP TABLE IF EXISTS `lm_common_set`;
CREATE TABLE `lm_common_set` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(32) NOT NULL COMMENT '常用设置 类型',
  `name` varchar(32) NOT NULL COMMENT '配置名称key',
  `value` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT '配置值',
  `show_type` varchar(20) NOT NULL DEFAULT 'show' COMMENT 'show 直接显示；switch 开关；select 下拉选框',
  `title` varchar(255) NOT NULL COMMENT '后台用于用户显示的名称',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态；0停用；1正常',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序；越小越靠前',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=22 DEFAULT CHARSET=utf8 COMMENT='公共配置信息';

-- ----------------------------
-- Records of lm_common_set
-- ----------------------------
INSERT INTO `lm_common_set` VALUES ('1', 'admin_conf', 'head_title', 'LM-后台管理系统', '', '后台页面顶部title', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('2', 'author_info', 'name', 'liumao801', '', '系统作者', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('3', 'author_info', 'email', 'liumao801@gmail.com', '', '系统作者邮箱', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('4', 'home_conf', 'head_title', 'LM-CMF', '', '前端页面顶部title', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('5', 'test', 'show', 'show-value0', 'show', '测试显示类型', '0', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('6', 'test', 'switch', '1', 'switch', '测试开关类型', '0', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('7', 'test', 'select', 'select4', 'select', '测试下拉选框类型|{\"select1\":\"测试-123\",\"select4\":\"测试-456\",\"select7\":\"测试-789\"}', '0', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('8', 'test', 'image', '/static/upload/commonset/2019-06/28/8kz6b_1561728853.png', 'image', '测试图片类型', '0', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('17', 'home_index_imgs', 'item2', '/static/modules/home/img/banner2.jpg', 'image', '首页轮播图2', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('16', 'home_index_imgs', 'item1', '/static/modules/home/img/banner1.jpg', 'image', '首页轮播图1', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('9', 'home_conf', 'view_tpl', 'default', 'select', '前端使用样式模板|{\"default\":\"默认内容管理系统模板\"}', '1', '100', null, '2019-07-08 18:14:55');
INSERT INTO `lm_common_set` VALUES ('18', 'home_index_imgs', 'item3', '/static/modules/home/img/banner3.jpg', 'image', '首页轮播图3', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('19', 'home_index_imgs', 'item4', '/static/modules/home/img/banner4.jpg', 'image', '首页轮播图4', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('20', 'home_index_imgs', 'item5', '/static/modules/home/img/banner5.jpg', 'image', '首页轮播图5', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('21', 'home_conf', 'logo', '/static/modules/home/img/logo.png', 'image', 'logo', '1', '100', null, null);
INSERT INTO `lm_common_set` VALUES ('10', 'home_conf', 'using_module', 'portal', 'select', '前端使用模块|{\"portal\":\"门户网站模板\",\"home\":\"内容管理系统模板\"}', '1', '100', null, '2019-07-08 18:23:48');

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
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0停用；1正常；-1已删除',
  `is_check` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否检测权限 1检测；0不检测',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `show` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=39 DEFAULT CHARSET=utf8 COMMENT='菜单资源信息表';

-- ----------------------------
-- Records of lm_menu
-- ----------------------------
INSERT INTO `lm_menu` VALUES ('1', '1', '首页', null, '1', 'fa-dashboard', 'HomeController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('2', '0', '系统管理', '0', '100', 'fa-dashboard', 'IndexController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('7', '1', '系统配置', '2', '100', 'fa-dashboard', 'IndexController.Index', 'admin/index/index', '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('8', '1', '日志管理', '2', '100', 'fa-dashboard', 'AdminLogController.Index', 'admin/index/index', '1', '1', null, null, '0');
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
INSERT INTO `lm_menu` VALUES ('33', '0', 'CMF系统', null, '100', 'fa-file-text', '', null, '1', '0', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0');
INSERT INTO `lm_menu` VALUES ('34', '1', '文章菜单分类', '33', '100', 'fa-ge', '', null, '1', '0', null, null, '0');
INSERT INTO `lm_menu` VALUES ('35', '1', '文章分类管理', '34', '100', 'fa-adn', 'MenuWebController.Index', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('36', '1', '文章管理', '34', '100', 'fa-file-text-o', 'ArticleController.Index', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('37', '1', '公共配置', '2', '100', 'fa-cogs', 'CommonSetController.Index', null, '1', '0', '0000-00-00 00:00:00', '0000-00-00 00:00:00', '0');
INSERT INTO `lm_menu` VALUES ('38', '1', '文章标签管理', '34', '99', 'fa-frown-o', 'ArticleTagController.Index', null, '1', '1', '2019-07-03 21:59:13', '2019-07-03 21:59:13', '0');

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
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '0停用；1正常；-1已删除',
  `front_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '前端菜单列表显示 0不显示；1显示',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序',
  `img` varchar(255) DEFAULT NULL COMMENT 'type=4 的缩略图',
  `seo_title` varchar(255) DEFAULT NULL,
  `seo_desc` varchar(255) DEFAULT NULL,
  `content` text COMMENT 'type=4 的页面内容',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='前端菜单/文章分类';

-- ----------------------------
-- Records of lm_menu_web
-- ----------------------------
INSERT INTO `lm_menu_web` VALUES ('1', '顶级分类1', '', '1', null, '', '1', '', '', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('2', '二级分类', '', '1', '1', '', '0', '', '', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('3', '三级跳转菜单', '', '2', '2', '', '0', '', 'https://www.baidu.com', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('6', '首页', 'fa-home', '2', null, '', null, '', '/', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('7', '公司简介', 'fa fa-file-text', '4', null, '', null, '', '', '1', '1', '100', '', '公司简介', '公司简介', '<p>公司简介</p>\n\n<p>公司简介</p>\n\n<p>公司简介</p>\n');
INSERT INTO `lm_menu_web` VALUES ('8', '公告信息', ' fa-flag', '3', '1', '', null, '', '', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('9', '新闻消息', ' fa-flag', '3', '1', '', null, '', '', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('10', '第二个三级', ' fa-flag', '3', '2', '', null, '', '', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('11', '第三级栏目页', ' fa-flag', '1', '2', '', null, '', '', '0', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('12', 'request324', 'fa-cogs', '2', '11', '', null, '', 'https://www.baidu.com?name=123', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('13', '第二顶级', '', '1', null, '', null, '', '', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('14', '第二顶级子菜单', '', '3', '13', '', null, '', '', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('15', '第三级栏目二', '', '1', '2', '', null, '', '', '1', '1', '100', '', '', '', '');
INSERT INTO `lm_menu_web` VALUES ('16', '跳转菜单32', '', '2', '15', '', null, '', 'https://www.baidu.com', '1', '1', '100', '', '', '', '');

-- ----------------------------
-- Table structure for lm_role
-- ----------------------------
DROP TABLE IF EXISTS `lm_role`;
CREATE TABLE `lm_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '角色名称',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '100' COMMENT '排序',
  `mark` varchar(255) NOT NULL COMMENT '角色标识',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0停用；1正常；-1已删除',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色信息表';

-- ----------------------------
-- Records of lm_role
-- ----------------------------
INSERT INTO `lm_role` VALUES ('1', '管理员', '1', 'admin', '1', null, null);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='角色用户关联表';

-- ----------------------------
-- Records of lm_role_admin_rel
-- ----------------------------
INSERT INTO `lm_role_admin_rel` VALUES ('4', '1', '4', '2019-06-21 14:12:58', '2019-06-21 14:12:58');

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
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8 COMMENT='角色菜单关联表';

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
INSERT INTO `lm_role_menu_rel` VALUES ('49', '2', '1', '2019-07-06 13:28:26', '2019-07-06 13:28:26', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('50', '2', '2', '2019-07-06 13:28:26', '2019-07-06 13:28:26', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('51', '2', '7', '2019-07-06 13:28:26', '2019-07-06 13:28:26', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('52', '2', '9', '2019-07-06 13:28:26', '2019-07-06 13:28:26', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('53', '2', '10', '2019-07-06 13:28:26', '2019-07-06 13:28:26', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('54', '2', '8', '2019-07-06 13:28:26', '2019-07-06 13:28:26', '0000-00-00 00:00:00');
INSERT INTO `lm_role_menu_rel` VALUES ('55', '2', '37', '2019-07-06 13:28:26', '2019-07-06 13:28:26', '0000-00-00 00:00:00');

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
  `status` tinyint(1) DEFAULT '0' COMMENT '0停用；1正常；-1锁定',
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
