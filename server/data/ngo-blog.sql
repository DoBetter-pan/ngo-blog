-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.11 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  9.3.0.4984
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 ngo-blog 的数据库结构
CREATE DATABASE IF NOT EXISTS `ngo-blog` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `ngo-blog`;


-- 导出  表 ngo-blog.ng_blog_article 结构
CREATE TABLE IF NOT EXISTS `ng_blog_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
  `authorId` int(11) NOT NULL COMMENT '作者ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '原始标题',
  `titleHtml` varchar(255) NOT NULL DEFAULT '' COMMENT 'HTML标题',
  `content` mediumtext NOT NULL COMMENT '原始内容',
  `contentHtml` mediumtext NOT NULL COMMENT 'HTML内容',
  `excerpt` text COMMENT '原始简介',
  `excerptHtml` text COMMENT 'HTML简介',
  `sectionId` int(6) NOT NULL COMMENT '板块ID',
  `categoryId` int(6) NOT NULL COMMENT '类别ID',
  `commentsCount` int(11) NOT NULL DEFAULT '0' COMMENT '评论个数',
  `status` int(11) NOT NULL DEFAULT '2' COMMENT '状态，0：草稿，1：隐藏，2：发布',
  `posted` datetime NOT NULL COMMENT '发布时间',
  `lastMod` datetime NOT NULL COMMENT '最后修改时间',
  `expires` datetime DEFAULT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`),
  KEY `authorIdx` (`authorId`),
  KEY `sectionIdx` (`sectionId`),
  KEY `categorieIdx` (`categoryId`),
  KEY `posted` (`posted`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 COMMENT='文章表';

-- 正在导出表  ngo-blog.ng_blog_article 的数据：~8 rows (大约)
DELETE FROM `ng_blog_article`;
/*!40000 ALTER TABLE `ng_blog_article` DISABLE KEYS */;
INSERT INTO `ng_blog_article` (`id`, `authorId`, `title`, `titleHtml`, `content`, `contentHtml`, `excerpt`, `excerptHtml`, `sectionId`, `categoryId`, `commentsCount`, `status`, `posted`, `lastMod`, `expires`) VALUES
	(1, 1, 'C语言基础', 'C语言基础', '怎么写好一个好用的 hello world 呢？ 首先打开 vim，然后写入代码，最后编译运行。', '怎么写好一个好用的 hello world 呢？ 首先打开 vim，然后写入代码，最后编译运行。', NULL, NULL, 2, 2, 0, 2, '2014-01-01 00:00:00', '0000-00-00 00:00:00', NULL),
	(2, 1, 'Find命令', 'Find命令', 'linux下强大的文件查找工具，非常牛，特别牛。', 'linux下强大的文件查找工具，非常牛，特别牛。', NULL, NULL, 2, 1, 0, 2, '2014-04-04 00:00:04', '0000-00-00 00:00:00', NULL),
	(3, 1, 'Awk命令', 'Awk命令', 'linux下强大的文本处理工具，非常牛，特别牛。', 'linux下强大的文本处理工具，非常牛，特别牛。', NULL, NULL, 2, 1, 0, 2, '2015-10-10 00:00:05', '2016-01-11 18:12:05', NULL),
	(4, 1, 'Sed命令', 'Sed命令', 'linux下强大的文本编辑工具，非常牛，特别牛。', 'linux下强大的文本编辑工具，非常牛，特别牛。', NULL, NULL, 2, 1, 0, 2, '2015-11-11 00:00:06', '0000-00-00 00:00:00', NULL),
	(5, 1, 'Ls命令', 'Ls命令', 'linux下强大的目录管理工具，非常牛，特别牛。', 'linux下强大的目录管理工具，非常牛，特别牛。', NULL, NULL, 2, 1, 0, 2, '2016-01-01 00:00:07', '0000-00-00 00:00:00', NULL),
	(17, 1, '', '世界第一', '', '<p>世界第一的好书~~，真的非常值得一看的。</p>', NULL, NULL, 2, 1, 0, 2, '2016-10-15 22:02:55', '2016-10-15 22:02:55', NULL),
	(18, 1, '', 'du命令', '', '<p>du是一个统计磁盘空间的命令。</p>', NULL, NULL, 2, 1, 0, 2, '2016-10-30 12:16:53', '2016-10-30 12:16:53', NULL),
	(19, 1, '', 'cd命令', '', '<p>linux下改变当前目录的命令！</p>', NULL, NULL, 2, 1, 0, 2, '2016-10-30 19:01:37', '2016-10-30 19:01:37', NULL);
/*!40000 ALTER TABLE `ng_blog_article` ENABLE KEYS */;


-- 导出  表 ngo-blog.ng_blog_category 结构
CREATE TABLE IF NOT EXISTS `ng_blog_category` (
  `id` int(6) NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
  `sectionId` int(6) NOT NULL COMMENT '所属板块',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '类型名称',
  `url` text NOT NULL COMMENT '链接URL',
  `description` varchar(256) NOT NULL DEFAULT '' COMMENT '说明',
  `isPage` int(6) NOT NULL DEFAULT '0' COMMENT '是否在首页',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='类型表';

-- 正在导出表  ngo-blog.ng_blog_category 的数据：~2 rows (大约)
DELETE FROM `ng_blog_category`;
/*!40000 ALTER TABLE `ng_blog_category` DISABLE KEYS */;
INSERT INTO `ng_blog_category` (`id`, `sectionId`, `name`, `url`, `description`, `isPage`) VALUES
	(1, 2, 'Linux命令', '', '讲述关于linux的一些命令与相关使用方法', 1),
	(2, 2, 'C语言', '', '讲述C语言的一些故事与用法', 1);
/*!40000 ALTER TABLE `ng_blog_category` ENABLE KEYS */;


-- 导出  表 ngo-blog.ng_blog_link 结构
CREATE TABLE IF NOT EXISTS `ng_blog_link` (
  `id` int(6) NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '链接名称',
  `url` text NOT NULL COMMENT '链接URL',
  `external` int(2) NOT NULL DEFAULT '0' COMMENT '0--内部链接，1--外部链接',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='链接表';

-- 正在导出表  ngo-blog.ng_blog_link 的数据：~2 rows (大约)
DELETE FROM `ng_blog_link`;
/*!40000 ALTER TABLE `ng_blog_link` DISABLE KEYS */;
INSERT INTO `ng_blog_link` (`id`, `name`, `url`, `external`) VALUES
	(1, '登录', '/login', 0),
	(2, '注册', '', 0);
/*!40000 ALTER TABLE `ng_blog_link` ENABLE KEYS */;


-- 导出  表 ngo-blog.ng_blog_section 结构
CREATE TABLE IF NOT EXISTS `ng_blog_section` (
  `id` int(6) NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '板块名称',
  `url` text NOT NULL COMMENT '链接URl',
  `authority` int(2) NOT NULL DEFAULT '0' COMMENT '0--全部，1--仅登陆用户，2--仅朋友，3--自己',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='板块表';

-- 正在导出表  ngo-blog.ng_blog_section 的数据：~2 rows (大约)
DELETE FROM `ng_blog_section`;
/*!40000 ALTER TABLE `ng_blog_section` DISABLE KEYS */;
INSERT INTO `ng_blog_section` (`id`, `name`, `url`, `authority`) VALUES
	(1, '首页', '#/', 0),
	(2, '技术', '', 0);
/*!40000 ALTER TABLE `ng_blog_section` ENABLE KEYS */;


-- 导出  表 ngo-blog.ng_blog_user 结构
CREATE TABLE IF NOT EXISTS `ng_blog_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名称',
  `password` varchar(128) NOT NULL COMMENT '用户名称',
  `role` varchar(128) NOT NULL COMMENT '角色：admin, writer, reader',
  `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '昵称',
  `email` varchar(254) NOT NULL DEFAULT '' COMMENT '邮箱',
  `lastAccess` datetime DEFAULT NULL COMMENT '最后访问时间',
  `nonce` varchar(256) NOT NULL DEFAULT '' COMMENT '随机数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- 正在导出表  ngo-blog.ng_blog_user 的数据：~0 rows (大约)
DELETE FROM `ng_blog_user`;
/*!40000 ALTER TABLE `ng_blog_user` DISABLE KEYS */;
INSERT INTO `ng_blog_user` (`id`, `name`, `password`, `role`, `nickname`, `email`, `lastAccess`, `nonce`) VALUES
	(1, 'yingx', '123456', 'admin', '', '', '2016-10-27 22:42:53', '9898');
/*!40000 ALTER TABLE `ng_blog_user` ENABLE KEYS */;


-- 导出  表 ngo-blog.ng_counter_visitor 结构
CREATE TABLE IF NOT EXISTS `ng_counter_visitor` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自动增加id',
  `username` varchar(48) NOT NULL DEFAULT '“”' COMMENT '用户名称',
  `ip` varchar(48) NOT NULL DEFAULT '“”' COMMENT '用户ip',
  `url` varchar(512) NOT NULL DEFAULT '“”' COMMENT '访问url',
  `visitTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '访问时间',
  `new` int(11) NOT NULL DEFAULT '0' COMMENT '是否为new',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

-- 正在导出表  ngo-blog.ng_counter_visitor 的数据：15 rows
DELETE FROM `ng_counter_visitor`;
/*!40000 ALTER TABLE `ng_counter_visitor` DISABLE KEYS */;
INSERT INTO `ng_counter_visitor` (`id`, `username`, `ip`, `url`, `visitTime`, `new`) VALUES
	(1, '', '127.0.0.1:52024', '', '2016-11-01 22:54:39', 1),
	(2, '', '127.0.0.1:52027', '', '2016-11-01 22:54:41', 1),
	(3, '', '127.0.0.1:52024', '', '2016-11-01 22:54:42', 1),
	(4, '', '127.0.0.1:52023', '', '2016-11-01 22:54:43', 1),
	(5, '', '127.0.0.1:52024', '', '2016-11-01 22:54:44', 1),
	(6, '', '127.0.0.1:52026', '', '2016-11-01 22:54:45', 1),
	(7, '', '127.0.0.1:52026', '', '2016-11-01 22:54:48', 1),
	(8, '', '127.0.0.1', '', '2016-11-01 22:59:16', 1),
	(9, '', '127.0.0.1', '', '2016-11-01 23:00:45', 1),
	(10, '', '127.0.0.1', '', '2016-11-01 23:01:46', 1),
	(11, '', '127.0.0.1', '', '2016-11-01 23:01:47', 1),
	(12, '', '127.0.0.1', '', '2016-11-01 23:01:47', 1),
	(13, '', '127.0.0.1', '', '2016-11-01 23:01:48', 1),
	(14, '', '127.0.0.1', '', '2016-11-01 23:03:46', 1),
	(15, '', '127.0.0.1', '', '2016-11-01 23:05:04', 1);
/*!40000 ALTER TABLE `ng_counter_visitor` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
