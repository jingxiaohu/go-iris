/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50714
Source Host           : localhost:3306
Source Database       : casbin

Target Server Type    : MYSQL
Target Server Version : 50714
File Encoding         : 65001

Date: 2020-03-01 22:30:07
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `gender` tinyint(5) DEFAULT NULL COMMENT '性别',
  `enable` tinyint(1) DEFAULT '0' COMMENT '0：启用用户 1：禁用用户',
  `name` varchar(255) NOT NULL DEFAULT '',
  `age` int(5) DEFAULT '0',
  `phone` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `userface` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `label` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=147 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES ('76', 'yhm1', 'x04jpoIrc8/mvNRqAG59Wg==', '2', '1', '1', '0', '', '', '', '2019-01-14 11:54:11', null, null);
INSERT INTO `admin_user` VALUES ('90', 'root', 'x04jpoIrc8/mvNRqAG59Wg==', '2', '1', '超级用户', '0', '17830895301', '', '', '2019-01-16 10:42:34', '2020-02-06 13:14:50', null);
INSERT INTO `admin_user` VALUES ('92', 'zhuangzi', 'x04jpoIrc8/mvNRqAG59Wg==', '0', '1', '庄子', '18', '3213123', 'qq.com', '', '2019-01-18 10:08:12', '2020-02-05 05:14:14', null);
INSERT INTO `admin_user` VALUES ('104', 'mozi', 'CBpaVgtlw7hA/zD1hFbcdw==', '2', '0', '墨子', '0', '1250', '98089089', '', '2019-01-18 11:58:53', '2020-02-06 10:47:42', null);
INSERT INTO `admin_user` VALUES ('146', '123', 'CBpaVgtlw7hA/zD1hFbcdw==', '2', '0', '123', '123', '123', '', '', '2020-02-06 10:48:25', '2020-02-09 08:18:48', '');

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `p_type` varchar(100) DEFAULT '',
  `v0` varchar(100) DEFAULT '',
  `v1` varchar(100) DEFAULT '',
  `v2` varchar(100) DEFAULT '',
  `v3` varchar(100) DEFAULT '',
  `v4` varchar(100) DEFAULT '',
  `v5` varchar(100) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `IDX_casbin_rule_p_type` (`p_type`) USING BTREE,
  KEY `IDX_casbin_rule_v2` (`v2`) USING BTREE,
  KEY `IDX_casbin_rule_v3` (`v3`) USING BTREE,
  KEY `IDX_casbin_rule_v4` (`v4`) USING BTREE,
  KEY `IDX_casbin_rule_v5` (`v5`) USING BTREE,
  KEY `IDX_casbin_rule_v0` (`v0`) USING BTREE,
  KEY `IDX_casbin_rule_v1` (`v1`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=185 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('62', 'g', '90', 'admin', 'a', '', '', '系统管理员');
INSERT INTO `casbin_rule` VALUES ('63', 'p', 'admin', 'a', '/api/admin/user*', 'GET|POST|PUT|DELETE', '.*', 'admin-用户权限');
INSERT INTO `casbin_rule` VALUES ('64', 'p', 'admin', 'a', '/api/admin/menu*', 'GET|POST|PUT|DELETE', '.*', 'admin-菜单权限');
INSERT INTO `casbin_rule` VALUES ('65', 'p', 'admin', 'a', '/api/admin/role*', 'GET|POST|PUT|DELETE', '.*', 'admin-角色权限');
INSERT INTO `casbin_rule` VALUES ('66', 'p', 'admin', 'a', '/api/admin/policy*', 'GET|POST|PUT|DELETE', '.*', 'admin-权限模块');
INSERT INTO `casbin_rule` VALUES ('75', 'p', 'demo', 'a', '/api/admin/demo*', 'GET|POST|PUT|DELETE', '.*', 'demo-测试权限');
INSERT INTO `casbin_rule` VALUES ('168', 'g', '104', 'admin', 'a', '', '', '系统管理员');
INSERT INTO `casbin_rule` VALUES ('176', 'g', '', 'akka', 'a', '', '', 'akka');
INSERT INTO `casbin_rule` VALUES ('179', 'g', '90', 'demo', 'a', '', '', '测试员');
INSERT INTO `casbin_rule` VALUES ('181', 'p', 'akka', 'a', '/api/admin/user*', 'GET|POST', '.*', 'akka-用户权限');
INSERT INTO `casbin_rule` VALUES ('184', 'g', '146', 'demo', 'a', '', '', '测试员');

-- ----------------------------
-- Table structure for dep
-- ----------------------------
DROP TABLE IF EXISTS `dep`;
CREATE TABLE `dep` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dep_name` varchar(64) DEFAULT '' COMMENT '部门名称',
  `leader` varchar(64) NOT NULL COMMENT '部门领导人uid',
  `phone` varchar(64) DEFAULT '' COMMENT '部门电话',
  `email` varchar(64) DEFAULT '' COMMENT '部门邮箱',
  `disabled` tinyint(1) DEFAULT NULL COMMENT '0:可用 否则:不可用',
  `parent_id` int(10) DEFAULT NULL,
  `dep_desc` varchar(255) DEFAULT '',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of dep
-- ----------------------------
INSERT INTO `dep` VALUES ('1', '稷下之学', '荀子', '17830895300', '614143260@qq.com', '0', '0', '', '2019-03-27 15:19:39', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('2', '儒家', '孔子', '1000', 'laozi@163.com', '0', '1', '崇尚\"礼乐\"和\"仁义\"，主张\"德治\"和\"仁政\"，教育了很多人为人处世所遵循的准则规范。', '2019-04-04 17:13:43', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('3', '道家', '老子', '1001', 'laozi@163.com', '0', '1', '天道运行的原理。天道轮回各种哲学上的思想都能在道家找到答案。世界上很多发生的事情都有其规律，人力不可更改。', '2019-04-30 17:13:45', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('4', '法家', '李斯', '1002', 'lisi@163.com', '0', '1', '依法治国，天子犯法与庶民同罪。法律视为一种有利于社会统治的强制性工具，体现法制建设的思想，一直被沿用至今。', '2019-04-16 17:13:49', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('5', '墨家', '墨子', '1003', 'mozi@163.com', '0', '1', '兼爱非攻，反对强大的王国去攻击弱小的王国，在思想上尊天事鬼，一切以保持社会现状的稳定为主。', '2019-04-10 17:13:52', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('6', '名家', '公孙龙', '1004', 'gongsl@163.com', '0', '1', '以辩论出名，著名的白马非马也是名家的思想。以逻辑推理来证明事物的好与坏、真实与否。', '2019-04-09 17:13:55', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('7', '阴阳家', '邹衍', '1005', 'zhouyan@163.com', '0', '1', '五行学说，从天文和地理方面来判断事物的凶吉。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('8', '纵横家', '苏秦|张仪', '1006', 'sz@163.com', '0', '1', '合纵连横，捭阖之术，阴阳之变化也。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('9', '农家', '许行', '1007', 'xuxing@163.com', '0', '1', '注重生产，研究植物生长和产出的学派。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('10', '兵家', '孙膑', '1008', 'sunbing@163.com', '0', '1', '讲究利用武力，最大化的夺取敌方的利益从而赢得战争的胜利。', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('11', '医家', '扁鹊', '1009', 'bianque@163.com', '0', '1', '医者仁心', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('12', '礼乐', '孟子', '1010', 'mengzi@163.com', '0', '2', '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('13', '武当', '张三丰', '1011', 'zsf@163.com', '0', '3', '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');
INSERT INTO `dep` VALUES ('14', '庄子游', '庄子', '1012', 'zz@163.com', '0', '3', '', '2019-04-24 17:13:59', '2019-04-24 17:13:59');

-- ----------------------------
-- Table structure for dep_admin
-- ----------------------------
DROP TABLE IF EXISTS `dep_admin`;
CREATE TABLE `dep_admin` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dep_ip` int(10) DEFAULT NULL COMMENT '部门id',
  `aid` int(10) DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of dep_admin
-- ----------------------------

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `path` varchar(64) DEFAULT '',
  `name` varchar(64) DEFAULT '' COMMENT '必须是唯一的',
  `modular` varchar(64) DEFAULT '',
  `component` varchar(64) DEFAULT '',
  `sort` int(3) DEFAULT '0' COMMENT '排序',
  `meta_id` int(11) DEFAULT NULL,
  `parent_id` int(10) DEFAULT NULL,
  `is_sub` tinyint(1) NOT NULL,
  `create_time` timestamp NULL DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `key` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES ('1', 'sys_mgr', 'sys_mgr', '', '', '1', '1', null, '0', null, null);
INSERT INTO `menu` VALUES ('2', 'user', 'user', '/admin/user', '/user', '3', '2', '1', '1', null, null);
INSERT INTO `menu` VALUES ('3', 'role', 'role', '/admin/role', '/role', '2', '3', '1', '1', null, null);
INSERT INTO `menu` VALUES ('4', 'menu', 'menu', '/admin/menu', '/menu', '1', '4', '1', '1', null, null);
INSERT INTO `menu` VALUES ('5', 'dep', 'dep', '/admin/dep', '/dep', '0', '5', '1', '1', null, null);
INSERT INTO `menu` VALUES ('6', 'policy', 'policy', '/admin/policy', '/policy', '4', '26', '1', '1', '2020-02-04 18:13:23', '2020-02-09 12:55:27');
INSERT INTO `menu` VALUES ('10', 'doc', 'doc', '', '', '2', '6', '0', '0', '0000-00-00 00:00:00', '2020-02-10 02:13:49');
INSERT INTO `menu` VALUES ('15', 'components', 'components', '', '', '3', '7', null, '0', null, null);
INSERT INTO `menu` VALUES ('16', 'tree_select_page', 'tree_select_page', '/components/tree-select', '/index', '0', '8', '15', '1', null, null);
INSERT INTO `menu` VALUES ('17', 'count_to_page', 'count_to_page', '/components/count-to', '/count-to', '0', '9', '15', '1', null, null);
INSERT INTO `menu` VALUES ('18', 'drag_list_page', 'drag_list_page', '/components/drag-list', '/drag-list', '0', '10', '15', '1', null, null);
INSERT INTO `menu` VALUES ('19', 'drag_drawer_page', 'drag_drawer_page', '/components/drag-drawer', '/index', '0', '11', '15', '1', null, null);
INSERT INTO `menu` VALUES ('20', 'org_tree_page', 'org_tree_page', '/components/org-tree', '/index', '0', '12', '15', '1', null, '2020-02-03 05:18:26');
INSERT INTO `menu` VALUES ('21', 'tree_table_page', 'tree_table_page', '/components/tree-table', '/index', '0', '13', '15', '1', null, null);
INSERT INTO `menu` VALUES ('22', 'cropper_page', 'cropper_page', '/components/cropper', '/cropper', '0', '14', '15', '1', null, null);
INSERT INTO `menu` VALUES ('23', 'tables_page', 'tables_page', '/components/tables', '/tables', '0', '15', '15', '1', null, null);
INSERT INTO `menu` VALUES ('24', 'split_pane_page', 'split_pane_page', '/components/split-pane', '/split-pane', '0', '16', '15', '1', null, null);
INSERT INTO `menu` VALUES ('25', 'markdown_page', 'markdown_page', '/components/markdown', '/markdown', '0', '17', '15', '1', null, null);
INSERT INTO `menu` VALUES ('26', 'editor_page', 'editor_page', '/components/editor', '/editor', '0', '18', '15', '1', null, null);
INSERT INTO `menu` VALUES ('27', 'icons_page', 'icons_page', '/components/icons', '/icons', '0', '19', '15', '1', null, null);
INSERT INTO `menu` VALUES ('28', 'update', 'update', '', '', '3', '20', null, '0', null, null);
INSERT INTO `menu` VALUES ('29', 'update_table_page', 'update_table_page', '/update', '/update-table', '0', '21', '28', '1', null, null);
INSERT INTO `menu` VALUES ('30', 'update_paste_page', 'update_paste_page', '/update', '/update-paste', '0', '22', '28', '1', null, null);
INSERT INTO `menu` VALUES ('31', 'excel', 'excel', '', '', '4', '23', null, '0', null, null);
INSERT INTO `menu` VALUES ('32', 'upload-excel', 'upload-excel', '/excel', '/upload-excel', '0', '24', '31', '1', null, null);
INSERT INTO `menu` VALUES ('33', 'export-excel', 'export-excel', '/excel', '/export-excel', '0', '25', '31', '1', null, null);
INSERT INTO `menu` VALUES ('34', 'tools_methods', 'tools_methods', '', '', '5', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('35', 'tools_methods_page', 'tools_methods_page', '/tools-methods', '/tools-methods', '0', null, '34', '1', null, null);
INSERT INTO `menu` VALUES ('36', 'i18n', 'i18n', '', '', '6', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('37', 'i18n_page', 'i18n_page', '/i18n', '/i18n-page', '0', null, '36', '1', null, null);
INSERT INTO `menu` VALUES ('38', 'error_store', 'error_store', '', '', '7', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('39', 'error_store_page', 'error_store_page', '/error-store', '/error-store', '0', null, '38', '1', null, null);
INSERT INTO `menu` VALUES ('40', 'error_logger', 'error_logger', '', '', '8', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('41', 'error_logger_page', 'error_logger_page', '/single-page', '/error-logger', '0', null, '40', '1', null, null);
INSERT INTO `menu` VALUES ('42', 'directive', 'directive', '', '', '9', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('43', 'directive_page', 'directive_page', '/directive', '/directive', '0', null, '42', '1', null, null);
INSERT INTO `menu` VALUES ('47', '401', 'error_401', '/error-page', '/401', '10', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('48', '500', 'error_500', '/error-page', '/500', '11', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('49', '*', 'error_404', '/error-page', '/404', '12', null, null, '0', null, null);
INSERT INTO `menu` VALUES ('50', 'ws-demo', 'ws测试', '/ws-demo', '/index', '8', '27', '1', '1', '0000-00-00 00:00:00', '2020-02-19 04:34:15');

-- ----------------------------
-- Table structure for menu_meta
-- ----------------------------
DROP TABLE IF EXISTS `menu_meta`;
CREATE TABLE `menu_meta` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL COMMENT '显示在侧边栏、面包屑和标签栏的文字',
  `hide_in_bread` tinyint(1) DEFAULT '0' COMMENT '设为true后此级路由将不会出现在面包屑中',
  `hide_in_menu` tinyint(1) DEFAULT '0' COMMENT '设为true后在左侧菜单不会显示该页面选项',
  `show_always` tinyint(1) DEFAULT '0',
  `not_cache` tinyint(1) DEFAULT '0' COMMENT '设为true后页面在切换标签后不会缓存，如果需要缓存，无需设置这个字段，而且需要设置页面组件name属性和路由配置的name一致',
  `access` varchar(255) DEFAULT NULL COMMENT '可访问该页面的权限数组，当前路由设置的权限会影响子路由',
  `icon` varchar(255) DEFAULT '' COMMENT '该页面在左侧菜单、面包屑和标签导航处显示的图标，如果是自定义图标，需要在图标名称前加下划线''_''',
  `before_close_name` varchar(255) DEFAULT '' COMMENT '设置该字段，则在关闭当前tab页时会去''@/router/before-close.js''里寻找该字段名对应的方法，作为关闭前的钩子函数',
  `href` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu_meta
-- ----------------------------
INSERT INTO `menu_meta` VALUES ('1', '系统管理', '0', '0', '0', '0', null, 'md-settings', '', '');
INSERT INTO `menu_meta` VALUES ('2', '用户管理', '0', '0', '0', '0', null, 'md-person', '', '');
INSERT INTO `menu_meta` VALUES ('3', '角色管理', '0', '0', '0', '0', null, 'ios-infinite', '', '');
INSERT INTO `menu_meta` VALUES ('4', '菜单管理', '0', '0', '0', '0', null, 'md-flower', '', '');
INSERT INTO `menu_meta` VALUES ('5', '部门管理', '0', '0', '0', '0', null, 'ios-infinite', '', '');
INSERT INTO `menu_meta` VALUES ('6', '文档', '0', '0', '0', '0', '', 'ios-book', '', 'https://lison16.github.io/iview-admin-doc/#/');
INSERT INTO `menu_meta` VALUES ('7', '组件', '0', '0', '0', '0', null, 'logo-buffer', '', '');
INSERT INTO `menu_meta` VALUES ('8', '树状下拉选择器', '0', '0', '0', '0', null, 'md-arrow-dropdown-circle', '', '');
INSERT INTO `menu_meta` VALUES ('9', '数字渐变', '0', '0', '0', '0', null, 'md-trending-up', '', '');
INSERT INTO `menu_meta` VALUES ('10', '拖拽列表', '0', '0', '0', '0', null, 'ios-infinite', '', '');
INSERT INTO `menu_meta` VALUES ('11', '可拖拽抽屉', '0', '0', '0', '0', null, 'md-list', '', '');
INSERT INTO `menu_meta` VALUES ('12', '组织结构树', '0', '0', '0', '0', '', 'ios-appstore', '', '');
INSERT INTO `menu_meta` VALUES ('13', '树状表格', '0', '0', '0', '0', null, 'md-git-branch', '', '');
INSERT INTO `menu_meta` VALUES ('14', '图片裁剪', '0', '0', '0', '0', null, 'md-crop', '', '');
INSERT INTO `menu_meta` VALUES ('15', '多功能表格', '0', '0', '0', '0', null, 'md-grid', '', '');
INSERT INTO `menu_meta` VALUES ('16', '分割窗口', '0', '0', '0', '0', null, 'md-pause', '', '');
INSERT INTO `menu_meta` VALUES ('17', 'Markdown编辑器', '0', '0', '0', '0', null, 'logo-markdown', '', '');
INSERT INTO `menu_meta` VALUES ('18', '富文本编辑器', '0', '0', '0', '0', null, 'ios-create', '', '');
INSERT INTO `menu_meta` VALUES ('19', '自定义图标', '0', '0', '0', '0', null, 'ios-create', '', '');
INSERT INTO `menu_meta` VALUES ('20', '数据上传', '0', '0', '1', '0', null, 'md-cloud-upload', '', '');
INSERT INTO `menu_meta` VALUES ('21', '上传Csv', '0', '0', '0', '0', null, 'ios-document', '', '');
INSERT INTO `menu_meta` VALUES ('22', '粘贴表格数据', '0', '0', '0', '0', null, 'md-clipboard', '', '');
INSERT INTO `menu_meta` VALUES ('23', 'EXCEL导入导出', '0', '0', '0', '0', null, 'md-download', '', '');
INSERT INTO `menu_meta` VALUES ('24', '导入EXCEL', '0', '0', '0', '0', null, 'md-add', '', '');
INSERT INTO `menu_meta` VALUES ('25', '导出EXCEL', '0', '0', '0', '0', null, 'md-download', '', '');
INSERT INTO `menu_meta` VALUES ('26', '权限管理', '0', '0', '0', '0', '', 'ios-infinite', '', '');
INSERT INTO `menu_meta` VALUES ('27', 'ws测试', '0', '0', '0', '0', '', 'logo-android', '', '');

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `role_key` varchar(100) NOT NULL DEFAULT '' COMMENT '角色key',
  `mid` int(10) NOT NULL COMMENT '菜单id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES ('admin', '2');
INSERT INTO `role_menu` VALUES ('admin', '3');
INSERT INTO `role_menu` VALUES ('admin', '6');
INSERT INTO `role_menu` VALUES ('demo', '16');
INSERT INTO `role_menu` VALUES ('demo', '22');
INSERT INTO `role_menu` VALUES ('akka', '21');
INSERT INTO `role_menu` VALUES ('akka', '22');
INSERT INTO `role_menu` VALUES ('akka', '2');
INSERT INTO `role_menu` VALUES ('akka', '6');
INSERT INTO `role_menu` VALUES ('admin', '4');
INSERT INTO `role_menu` VALUES ('admin', '10');
INSERT INTO `role_menu` VALUES ('admin', '50');
