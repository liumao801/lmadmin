/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : go.test

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-06-28 21:36:35
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
) ENGINE=InnoDB AUTO_INCREMENT=540 DEFAULT CHARSET=utf8 COMMENT='后台操作日志';

-- ----------------------------
-- Records of lm_admin_log
-- ----------------------------
INSERT INTO `lm_admin_log` VALUES ('99', '1', 'admin', '/admin/adminlog/datagrid', 'POST', '{}', '2019-06-27 22:15:01', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('105', '1', 'admin', '/admin/commonset/edit/14', 'POST', '{\"Id\":[\"14\"],\"Name\":[\"测试\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"测试类型|{\\\"a\\\":\\\"倒计时\\\",\\\"b\\\":\\\"减肥的\\\"}\"],\"Type\":[\"fds111\"],\"Value\":[\"a\"]}', '2019-06-27 22:15:07', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('242', '1', 'admin', '/admin/adminlog/delete', 'POST', '{\"ids\":[\"241,240,239,238,237,236,235,234,233,232,231,230,229,228,227,226,225,224,223,222,221,220,219,218,217,216,215,214,213,212,211,210,209,208,207,206,205,204,203,202,201,200,199,198,197,196,195,194,193,192\"]}', '2019-06-28 13:13:30', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('244', '1', 'admin', '/admin/adminlog/delete', 'POST', '{\"ids\":[\"243,191,190,189,188,187,186,185,184,183,182,181,180,179,178,177,176,175,174,173,172,171,170,169,168,167,166,165,164,108,107,106,104,103,102,101,100\"]}', '2019-06-28 13:13:56', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('245', '1', 'admin', '/admin/adminlog/datagrid', 'POST', '{}', '2019-06-28 13:13:56', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('378', '1', 'admin', '/admin/commonset/index', 'GET', '{}', '2019-06-28 16:51:27', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('379', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 16:51:28', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('380', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 16:51:28', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('381', '1', 'admin', '/admin/commonset/edit/0', 'GET', '{}', '2019-06-28 16:51:31', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('382', '1', 'admin', '/admin/commonset/edit/0', 'POST', '{\"Id\":[\"0\"],\"Name\":[\"view_tpl\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"前端使用模板|{\\\"protal\\\":\\\"门户网站\\\"}\"],\"Type\":[\"home_conf\"],\"Value\":[\"protal\"]}', '2019-06-28 16:52:01', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('388', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 16:53:08', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('389', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"1\"]}', '2019-06-28 16:53:08', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('395', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 16:59:21', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('396', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 16:59:21', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('397', '1', 'admin', '/admin/menuweb/edit', 'POST', '{\"Content\":[\"\"],\"Icon\":[\" fa-flag\"],\"Id\":[\"11\"],\"Parent\":[\"2\"],\"Sort\":[\"100\"],\"Title\":[\"第三级栏目页\"],\"Type\":[\"1\"]}', '2019-06-28 16:59:23', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('416', '1', 'admin', '/admin/menuweb/edit/11', 'GET', '{}', '2019-06-28 17:01:56', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('417', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:01:56', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('418', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 17:01:56', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('419', '1', 'admin', '/admin/menuweb/edit/11', 'GET', '{}', '2019-06-28 17:03:13', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('420', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:03:13', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('421', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 17:03:13', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('422', '1', 'admin', '/admin/menuweb/edit', 'POST', '{\"Content\":[\"\"],\"FrontShow\":[\"1\"],\"Icon\":[\" fa-flag\"],\"Id\":[\"11\"],\"Parent\":[\"2\"],\"Sort\":[\"100\"],\"Title\":[\"第三级栏目页\"],\"Type\":[\"1\"]}', '2019-06-28 17:03:15', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('423', '1', 'admin', '/admin/menuweb/index', 'GET', '{}', '2019-06-28 17:03:17', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('424', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:03:17', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('425', '1', 'admin', '/admin/menuweb/treegrid', 'POST', '{}', '2019-06-28 17:03:17', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('426', '1', 'admin', '/admin/menuweb/edit/11', 'GET', '{}', '2019-06-28 17:03:22', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('427', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:03:23', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('428', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 17:03:23', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('429', '1', 'admin', '/admin/menuweb/edit/11', 'GET', '{}', '2019-06-28 17:03:54', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('430', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:03:54', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('431', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 17:03:54', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('432', '1', 'admin', '/admin/menuweb/edit/11', 'GET', '{}', '2019-06-28 17:03:55', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('433', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:03:55', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('434', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 17:03:55', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('435', '1', 'admin', '/admin/menuweb/edit/11', 'GET', '{}', '2019-06-28 17:04:20', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('436', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:04:20', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('437', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 17:04:20', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('438', '1', 'admin', '/admin/menuweb/edit/11', 'GET', '{}', '2019-06-28 17:04:30', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('439', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:04:30', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('440', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 17:04:30', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('441', '1', 'admin', '/admin/menuweb/edit/11', 'GET', '{}', '2019-06-28 17:04:31', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('442', '1', 'admin', '/admin/menu/adminmenutree', 'POST', '{}', '2019-06-28 17:04:31', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('443', '1', 'admin', '/admin/menuweb/parent', 'POST', '{\"id\":[\"11\"]}', '2019-06-28 17:04:31', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('444', '1', 'admin', '/admin/commonset/edit/15', 'GET', '{}', '2019-06-28 17:10:04', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('445', '1', 'admin', '/admin/commonset/edit/15', 'POST', '{\"Id\":[\"15\"],\"Name\":[\"view_tpl\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"前端使用模板|{\\\"portal\\\":\\\"门户网站\\\"}\"],\"Type\":[\"home_conf\"],\"Value\":[\"protal\"]}', '2019-06-28 17:10:11', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('446', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 17:10:11', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('447', '1', 'admin', '/admin/commonset/edit/15', 'GET', '{}', '2019-06-28 17:10:12', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('448', '1', 'admin', '/admin/commonset/edit/15', 'POST', '{\"Id\":[\"15\"],\"Name\":[\"view_tpl\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"前端使用模板|{\\\"portal\\\":\\\"门户网站\\\"}\"],\"Type\":[\"home_conf\"],\"Value\":[\"portal\"]}', '2019-06-28 17:10:15', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('460', '1', 'admin', '/admin/commonset/edit/15', 'GET', '{}', '2019-06-28 19:30:09', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('461', '1', 'admin', '/admin/commonset/edit/15', 'POST', '{\"Id\":[\"15\"],\"Name\":[\"view_tpl\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"前端使用模板|{\\\"portal\\\":\\\"门户网站模板\\\",\\\"cms\\\":\\\"内容关系系统模板\\\"}\"],\"Type\":[\"home_conf\"],\"Value\":[\"portal\"]}', '2019-06-28 19:31:34', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('462', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 19:31:35', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('463', '1', 'admin', '/admin/commonset/edit/15', 'GET', '{}', '2019-06-28 19:31:37', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('464', '1', 'admin', '/admin/commonset/edit/15', 'POST', '{\"Id\":[\"15\"],\"Name\":[\"view_tpl\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"前端使用模板|{\\\"portal\\\":\\\"门户网站模板\\\",\\\"cms\\\":\\\"内容管理系统模板\\\"}\"],\"Type\":[\"home_conf\"],\"Value\":[\"portal\"]}', '2019-06-28 19:31:48', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('465', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 19:31:48', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('466', '1', 'admin', '/admin/commonset/edit/15', 'GET', '{}', '2019-06-28 19:31:49', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('467', '1', 'admin', '/admin/commonset/edit/15', 'POST', '{\"Id\":[\"15\"],\"Name\":[\"view_tpl\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"前端使用模板|{\\\"portal\\\":\\\"门户网站模板\\\",\\\"cms\\\":\\\"内容管理系统模板\\\"}\"],\"Type\":[\"home_conf\"],\"Value\":[\"cms\"]}', '2019-06-28 19:31:54', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('476', '1', 'admin', '/admin/commonset/edit/15', 'GET', '{}', '2019-06-28 20:41:23', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('477', '1', 'admin', '/admin/commonset/edit/15', 'POST', '{\"Id\":[\"15\"],\"Name\":[\"view_tpl\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"1\"],\"Title\":[\"前端使用模板|{\\\"portal\\\":\\\"门户网站模板\\\",\\\"cms\\\":\\\"内容管理系统模板\\\"}\"],\"Type\":[\"home_conf\"],\"Value\":[\"portal\"]}', '2019-06-28 20:41:29', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('478', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 20:41:29', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('502', '1', 'admin', '/admin/commonset/edit/5', 'GET', '{}', '2019-06-28 21:32:03', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('503', '1', 'admin', '/admin/commonset/edit/5', 'POST', '{\"Id\":[\"5\"],\"Name\":[\"show\"],\"ShowType\":[\"show\"],\"Sort\":[\"100\"],\"Status\":[\"0\"],\"Title\":[\"测试显示类型\"],\"Type\":[\"test\"],\"Value\":[\"show-value\"]}', '2019-06-28 21:32:32', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('504', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:32:32', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('505', '1', 'admin', '/admin/commonset/edit/6', 'GET', '{}', '2019-06-28 21:32:35', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('506', '1', 'admin', '/admin/commonset/edit/6', 'POST', '{\"Id\":[\"6\"],\"Name\":[\"switch\"],\"ShowType\":[\"switch\"],\"Sort\":[\"100\"],\"Status\":[\"0\"],\"Title\":[\"测试开关类型\"],\"Type\":[\"test\"],\"Value\":[\"0\"]}', '2019-06-28 21:33:02', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('507', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:33:02', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('508', '1', 'admin', '/admin/commonset/edit/7', 'GET', '{}', '2019-06-28 21:33:04', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('509', '1', 'admin', '/admin/commonset/edit/7', 'POST', '{\"Id\":[\"7\"],\"Name\":[\"select\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"0\"],\"Title\":[\"测试下拉选框类型|{\\\"select1\\\":\\\"测试-123\\\",\\\"select4\\\":\\\"测试-456\\\",\\\"select7\\\":\\\"测试-789\\\"}\"],\"Type\":[\"test\"],\"Value\":[\"select1\"]}', '2019-06-28 21:33:42', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('510', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:33:42', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('511', '1', 'admin', '/admin/commonset/edit/8', 'GET', '{}', '2019-06-28 21:33:45', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('512', '1', 'admin', '/admin/commonset/edit/8', 'POST', '{\"Id\":[\"8\"],\"Name\":[\"image\"],\"ShowType\":[\"image\"],\"Sort\":[\"100\"],\"Status\":[\"0\"],\"Title\":[\"测试图片类型\"],\"Type\":[\"test\"],\"Value\":[\"123\"]}', '2019-06-28 21:34:06', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('513', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:34:06', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('514', '1', 'admin', '/admin/commonset/edit/8', 'GET', '{}', '2019-06-28 21:34:07', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('515', '1', 'admin', '/admin/commonset/edit/123', 'GET', '{}', '2019-06-28 21:34:07', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('516', '1', 'admin', '/admin/commonset/edit/123', 'GET', '{}', '2019-06-28 21:34:11', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('517', '1', 'admin', '/admin/commonset/edit/8', 'POST', '{\"Id\":[\"8\"],\"Name\":[\"image\"],\"ShowType\":[\"image\"],\"Sort\":[\"100\"],\"Status\":[\"0\"],\"Title\":[\"测试图片类型\"],\"Type\":[\"test\"],\"oldValue\":[\"123\"]}', '2019-06-28 21:34:13', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('518', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:34:13', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('519', '1', 'admin', '/admin/commonset/edit/7', 'GET', '{}', '2019-06-28 21:34:16', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('520', '1', 'admin', '/admin/commonset/edit/7', 'POST', '{\"Id\":[\"7\"],\"Name\":[\"select\"],\"ShowType\":[\"select\"],\"Sort\":[\"100\"],\"Status\":[\"0\"],\"Title\":[\"测试下拉选框类型|{\\\"select1\\\":\\\"测试-123\\\",\\\"select4\\\":\\\"测试-456\\\",\\\"select7\\\":\\\"测试-789\\\"}\"],\"Type\":[\"test\"],\"Value\":[\"select4\"]}', '2019-06-28 21:34:19', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('521', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:34:19', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('522', '1', 'admin', '/admin/commonset/edit/6', 'GET', '{}', '2019-06-28 21:34:22', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('523', '1', 'admin', '/admin/commonset/edit/6', 'POST', '{\"Id\":[\"6\"],\"Name\":[\"switch\"],\"ShowType\":[\"switch\"],\"Sort\":[\"100\"],\"Status\":[\"0\"],\"Title\":[\"测试开关类型\"],\"Type\":[\"test\"],\"Value\":[\"1\"]}', '2019-06-28 21:34:25', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('524', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:34:25', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('525', '1', 'admin', '/admin/commonset/edit/5', 'GET', '{}', '2019-06-28 21:34:28', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('526', '1', 'admin', '/admin/commonset/edit/5', 'POST', '{\"Id\":[\"5\"],\"Name\":[\"show\"],\"ShowType\":[\"show\"],\"Sort\":[\"100\"],\"Status\":[\"0\"],\"Title\":[\"测试显示类型\"],\"Type\":[\"test\"],\"Value\":[\"show-value0\"]}', '2019-06-28 21:34:32', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('527', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:34:32', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('528', '1', 'admin', '/admin/commonset/delete', 'POST', '{\"ids\":[\"14,13,12,9\"]}', '2019-06-28 21:34:45', '127.0.0.1');
INSERT INTO `lm_admin_log` VALUES ('529', '1', 'admin', '/admin/commonset/datagrid', 'POST', '{}', '2019-06-28 21:34:45', '127.0.0.1');

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
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='文章信息';

-- ----------------------------
-- Records of lm_article
-- ----------------------------
INSERT INTO `lm_article` VALUES ('1', 'request', '热弯', '', '2', ' 范德萨', '<p>放给大家看联发科多斯拉克莲富大厦</p>\n', '8', '2', '0', '0', '0', '1', null, '0');
INSERT INTO `lm_article` VALUES ('2', '二级分类', '范德萨', '', '范德萨', '范德萨', '<p>范德萨范德萨</p>\n', '9', '', '1', '0', '0', '0', '2019', '0');

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
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='公共配置信息';

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
INSERT INTO `lm_common_set` VALUES ('9', 'home_conf', 'view_tpl', 'portal', 'select', '前端使用模板|{\"portal\":\"门户网站模板\",\"cms\":\"内容管理系统模板\"}', '1', '100', null, null);

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
  `created_at` int(10) unsigned DEFAULT NULL,
  `updated_at` int(10) unsigned DEFAULT NULL,
  `show` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 COMMENT='菜单资源信息表';

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
INSERT INTO `lm_menu` VALUES ('33', '0', 'CMF系统', null, '100', 'fa-file-text', '', null, '1', '0', '2018', '2018', '0');
INSERT INTO `lm_menu` VALUES ('34', '1', '文章菜单分类', '33', '100', 'fa-ge', '', null, '1', '0', null, null, '0');
INSERT INTO `lm_menu` VALUES ('35', '1', '文章分类管理', '34', '100', 'fa-adn', 'MenuWebController.Index', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('36', '1', '文章管理', '34', '100', 'fa-file-text-o', 'ArticleController.Index', null, '1', '1', null, null, '0');
INSERT INTO `lm_menu` VALUES ('37', '1', '公共配置', '2', '100', 'fa-cogs', 'CommonSetController.Index', null, '1', '0', '2019', '2019', '0');

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
INSERT INTO `lm_menu_web` VALUES ('6', '首页', 'fa-home', '2', null, '', null, '', '/home', '1', '1', '100', '', '', '', '');
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
