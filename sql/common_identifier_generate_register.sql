# ************************************************************
# Sequel Pro SQL dump
# Version 5446
#
# https://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.10)
# Database: common_micro
# Generation Time: 2023-03-01 10:31:45 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table common_identifier_generate_register
# ------------------------------------------------------------

DROP TABLE IF EXISTS `common_identifier_generate_register`;

CREATE TABLE `common_identifier_generate_register` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '节点ID 对应WorkerID',
  `ip` varchar(16) NOT NULL DEFAULT '' COMMENT '服务器IP',
  `last_timestamp` bigint(30) NOT NULL COMMENT '最后操作时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ID生成服务注册记录表';

LOCK TABLES `common_identifier_generate_register` WRITE;
/*!40000 ALTER TABLE `common_identifier_generate_register` DISABLE KEYS */;

INSERT INTO `common_identifier_generate_register` (`id`, `ip`, `last_timestamp`)
VALUES
	(1,'192.168.198.24',1677579367857687);

/*!40000 ALTER TABLE `common_identifier_generate_register` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
