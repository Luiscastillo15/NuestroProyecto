/*
SQLyog Ultimate v11.11 (64 bit)
MySQL - 5.5.5-10.4.32-MariaDB : Database - proyecto
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`proyecto` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci */;

USE `proyecto`;

/*Table structure for table `telefonos` */

DROP TABLE IF EXISTS `telefonos`;

CREATE TABLE `telefonos` (
  `ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Id del telefono',
  `marca` varchar(30) NOT NULL COMMENT 'Marca del telefono',
  `modelo` varchar(20) NOT NULL COMMENT 'Modelo del telefono',
  `precio` float NOT NULL COMMENT 'precio del telefono',
  KEY `ID` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `telefonos` */

insert  into `telefonos`(`ID`,`marca`,`modelo`,`precio`) values (1,'Samsung','A54',230),(2,'Xiaomi','Poco X6 Pro',310);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
