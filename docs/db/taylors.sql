/*
 Navicat MySQL Data Transfer

 Source Server         : taylors
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : 127.0.0.1:3306
 Source Schema         : taylors

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 28/05/2020 19:30:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/base/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/base/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/casbinTest/:pathParam', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/stock/top/list', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/stock/monitor/list', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/stock/monitor/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/stock/monitor/one', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/stock/monitor/del', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/stock/monitor/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/base/register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/stock/top/list', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/stock/monitor/list', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/stock/monitor/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/stock/monitor/one', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/stock/monitor/del', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'stock', '/stock/monitor/update', 'POST', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `customer_name` varchar(255) DEFAULT NULL,
  `customer_phone_data` varchar(255) DEFAULT NULL,
  `sys_user_id` int(10) unsigned DEFAULT NULL,
  `sys_user_authority_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_customers_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------
BEGIN;
INSERT INTO `exa_customers` VALUES (1, '2020-02-25 18:01:48', '2020-04-10 12:29:29', NULL, '测试客户', '1761111111', 10, '888');
INSERT INTO `exa_customers` VALUES (2, '2020-04-10 12:25:53', '2020-04-10 12:25:53', '2020-04-10 13:43:56', 'test', '123123123', 10, '888');
INSERT INTO `exa_customers` VALUES (3, '2020-04-10 13:44:12', '2020-04-10 13:44:12', '2020-04-10 13:44:13', '123123', '123123', 10, '888');
INSERT INTO `exa_customers` VALUES (4, '2020-04-10 13:47:10', '2020-04-10 13:47:10', '2020-04-10 13:47:12', '22222222', '222222222222222', 10, '888');
COMMIT;

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `exa_file_id` int(10) unsigned DEFAULT NULL,
  `file_chunk_path` varchar(255) DEFAULT NULL,
  `file_chunk_number` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `tag` varchar(255) DEFAULT NULL,
  `key` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_file_upload_and_downloads_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
BEGIN;
INSERT INTO `exa_file_upload_and_downloads` VALUES (17, '2020-04-26 11:51:39', '2020-04-26 11:51:39', NULL, '10.png', 'http://qmplusimg.henrongyi.top/158787308910.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (19, '2020-04-27 15:48:38', '2020-04-27 15:48:38', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1587973709logo.png', 'png', '1587973709logo.png');
COMMIT;

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `file_md5` varchar(255) DEFAULT NULL,
  `file_path` varchar(255) DEFAULT NULL,
  `chunk_total` int(11) DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_exa_files_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of exa_files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `jwt` text,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
BEGIN;
INSERT INTO `jwt_blacklists` VALUES (3, '2019-12-28 18:29:05', '2019-12-28 18:29:05', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MTMzNzM2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc1Mjc5MzZ9.T7ikGw-lgAAQlfMne7zPIF-PlfQMg37uBCYJ24Y_B38');
INSERT INTO `jwt_blacklists` VALUES (4, '2019-12-28 18:31:02', '2019-12-28 18:31:02', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MTMzODUzLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc1MjgwNTN9.tDzUm4KNFeJCErNfZGfuF2tcuolga2f_2dE0nTl_UZU');
INSERT INTO `jwt_blacklists` VALUES (5, '2019-12-28 18:31:25', '2019-12-28 18:31:25', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MTMzODcwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc1MjgwNzB9.mspXy9sqQO_5PusPReLalodo_ybWRKxb3Ownf2r2HxE');
INSERT INTO `jwt_blacklists` VALUES (6, '2019-12-30 14:20:10', '2019-12-30 14:20:10', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkxNTc2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODU3NzZ9.AR2KYShboFKsHTjwohxEkA3lytttfZqRH849sl2fNdw');
INSERT INTO `jwt_blacklists` VALUES (7, '2019-12-30 14:21:14', '2019-12-30 14:21:14', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkxNjE2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODU4MTZ9.h8zbDVHM_QbBI-ejGXeQpw0S9oYHJyP4U-TwsVFus9Q');
INSERT INTO `jwt_blacklists` VALUES (8, '2019-12-30 14:21:57', '2019-12-30 14:21:57', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkxNjgxLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODU4ODF9.CSjolDGVpU0g7YG6TaPAlWAMdhtvnBhAi-XYYWZ6RLo');
INSERT INTO `jwt_blacklists` VALUES (9, '2019-12-30 14:25:01', '2019-12-30 14:25:01', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkxODIyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODYwMjJ9.Y_s22Vh5J2ah6Kh1nZQQ8XIQspbT4I7tzc_YJqWrRWM');
INSERT INTO `jwt_blacklists` VALUES (10, '2019-12-30 14:29:26', '2019-12-30 14:29:26', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkyMTU0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODYzNTR9.4HJdx-sfYE5TUUefdwi3yZ6dY_jG7WwEC_55WuGawY8');
INSERT INTO `jwt_blacklists` VALUES (11, '2019-12-30 14:43:43', '2019-12-30 14:43:43', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkyMTcwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODYzNzB9.YEhupQVwjMVBB2eAcAoGG-vJczoxuUyn6KR-tDWU86I');
INSERT INTO `jwt_blacklists` VALUES (12, '2019-12-30 14:55:13', '2019-12-30 14:55:13', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkzMDI3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODcyMjd9.r_sE_Z31cFdS2nCf3iyQjuiZe0Z3HPR07wKBGlUHsnk');
INSERT INTO `jwt_blacklists` VALUES (13, '2019-12-30 14:58:31', '2019-12-30 14:58:31', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkzNzY2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODc5NjZ9.dYFlmyIKQZjzTCKu56wCmxXiW6zOayN_YgygCcvCyLk');
INSERT INTO `jwt_blacklists` VALUES (14, '2019-12-30 14:58:38', '2019-12-30 14:58:38', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkzOTEwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODgxMTB9.pPmzsHU4UceZuPFT_G-SDdxe6FD3MuL47HkovpI-_0c');
INSERT INTO `jwt_blacklists` VALUES (15, '2019-12-30 14:58:58', '2019-12-30 14:58:58', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkzOTE4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODgxMTh9.irf98R0belbXtb8x9SxsvuhiYsbHMPbHbFDxaaH0z6Q');
INSERT INTO `jwt_blacklists` VALUES (16, '2020-01-06 16:32:31', '2020-01-06 16:32:31', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTAzMjk5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTc0OTl9.jgLfjvek7sQyuZ2TABQvLOyu_ifNw_KYzfY3VTLL4fw');
INSERT INTO `jwt_blacklists` VALUES (17, '2020-01-06 16:33:08', '2020-01-06 16:33:08', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA0MzU4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTg1NTh9.89r6xHZUBDjfmNpmF02RjQXYTBGUiJvOEDP8pydNt-A');
INSERT INTO `jwt_blacklists` VALUES (18, '2020-01-06 16:33:18', '2020-01-06 16:33:18', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA0MzkyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTg1OTJ9.6Yv9ZYhN-TH9H4SoZEAkjevKVX0vLHL1lVQGFpfBr2U');
INSERT INTO `jwt_blacklists` VALUES (19, '2020-01-06 16:36:06', '2020-01-06 16:36:06', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA0NDA5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTg2MDl9._9zRRK76XH_KgrW1X9P5GTLW9dwfIixB4QUsC7M3RHA');
INSERT INTO `jwt_blacklists` VALUES (20, '2020-01-06 16:44:06', '2020-01-06 16:44:06', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA0NTcxLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTg3NzF9.5ki0TZooCorK81xWpYa-OO3RR-Bpp5am_uNCNPh4250');
INSERT INTO `jwt_blacklists` VALUES (21, '2020-01-06 16:45:50', '2020-01-06 16:45:50', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MDUwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTkyNTB9.A0n5faE0X0TyRb_1RvAQBLooY-peapPTD0LnJD03Ul0');
INSERT INTO `jwt_blacklists` VALUES (22, '2020-01-06 16:46:24', '2020-01-06 16:46:24', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MTU0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTkzNTR9.VtqTOJ-MQY2K3w4tM7HgT0z73CEOd3CDqmYqKCjXxnc');
INSERT INTO `jwt_blacklists` VALUES (23, '2020-01-06 16:47:20', '2020-01-06 16:47:20', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MTg3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTkzODd9.fwL1QakF30SHSaGDkPo3weIg0l7kiAGwNq_fKsFxquc');
INSERT INTO `jwt_blacklists` VALUES (24, '2020-01-06 16:47:57', '2020-01-06 16:47:57', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MjQ0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk0NDR9.VoKdA0-brmUlQ5bYufIdMWrS-cCQ2ARm7_jeVtfvCpc');
INSERT INTO `jwt_blacklists` VALUES (25, '2020-01-06 16:49:08', '2020-01-06 16:49:08', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1Mjg1LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk0ODV9.a8-zmyIlJJGdonhXAzNvNH9C-nMa-Voq4bhTbiVKJzE');
INSERT INTO `jwt_blacklists` VALUES (26, '2020-01-06 16:49:32', '2020-01-06 16:49:32', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MzUyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk1NTJ9.l4e3rjtrDgRsqnQwizJ-ZXVUVM8ywSJcNJkkEVYbdzU');
INSERT INTO `jwt_blacklists` VALUES (27, '2020-01-06 16:49:58', '2020-01-06 16:49:58', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1Mzc3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk1Nzd9.mXUPYvmXbntrdywpBNM0j9sP991cwfhc9b0KvUM4dG4');
INSERT INTO `jwt_blacklists` VALUES (28, '2020-01-06 16:50:56', '2020-01-06 16:50:56', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NDExLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk2MTF9.Z21e8nWHKV5XvYg61CZCz3nMK25m_FmlxncxGMpMS0k');
INSERT INTO `jwt_blacklists` VALUES (29, '2020-01-06 16:52:03', '2020-01-06 16:52:03', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NDY0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk2NjR9.qzptIyCcL_SPm6TGwXML8Rih3qYqj9GLUpWzTpSPPuI');
INSERT INTO `jwt_blacklists` VALUES (30, '2020-01-06 16:52:36', '2020-01-06 16:52:36', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NTI3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk3Mjd9.D9e8qbx44CLX0ZInwNlIqTGS_sSE069TRIDkQAk7tVY');
INSERT INTO `jwt_blacklists` VALUES (31, '2020-01-06 16:54:35', '2020-01-06 16:54:35', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NTY1LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk3NjV9.D4EZmVAJ96kxcyIfWkT_LA81t1JCuQZcYmQkkoNhtPo');
INSERT INTO `jwt_blacklists` VALUES (32, '2020-01-06 16:55:40', '2020-01-06 16:55:40', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NjgzLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk4ODN9.SJL2fFMbe5VL2YWBzMlrhxbBIJhIHTUeodkEpgH1Xgo');
INSERT INTO `jwt_blacklists` VALUES (33, '2020-01-06 16:57:28', '2020-01-06 16:57:28', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NzU4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk5NTh9.6y12UkOeW7vz7gGTcYaN3Y-2Ut2QmjgU9WEuy_pneGM');
INSERT INTO `jwt_blacklists` VALUES (34, '2020-01-06 16:59:02', '2020-01-06 16:59:02', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1ODU1LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgzMDAwNTV9.G0q9X7Ld3cN_BO-K219b7tFAHgtpiAwqLPoxVNKsEl8');
INSERT INTO `jwt_blacklists` VALUES (35, '2020-01-06 16:59:26', '2020-01-06 16:59:26', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1OTQ2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgzMDAxNDZ9.cmBgWiztsnh7zF3OUNIDQKv8wzGJF7fllUv-4LlYxu8');
INSERT INTO `jwt_blacklists` VALUES (36, '2020-03-21 14:46:14', '2020-03-21 14:46:14', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg1Mzc3ODY3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODQ3NzIwNjd9.DLhWhD1FdcWLyFLcXQynKJnenbVHrSiKhlDGFRzgo5k');
INSERT INTO `jwt_blacklists` VALUES (37, '2020-03-31 14:24:35', '2020-03-31 14:24:35', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg2MTM4MTA4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODU1MzIzMDh9.Ro2F2dZLfOk2Z_OPRbweOuCpchr6HlHfQIF5qjfc8y4');
INSERT INTO `jwt_blacklists` VALUES (38, '2020-04-01 16:07:57', '2020-04-01 16:07:57', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg2MjQwNzQyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODU2MzQ5NDJ9.9qaOFu7D5cq4vxTfLi4pyO_JGcKjVAEJIcoStJWJlYg');
INSERT INTO `jwt_blacklists` VALUES (39, '2020-04-15 16:30:41', '2020-04-15 16:30:41', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3MDk1Njg5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODY0ODk4ODl9.-cNmRAyqhylZlzakwoFY08x7RnjI3CiWTiQc_Iabb-c');
INSERT INTO `jwt_blacklists` VALUES (40, '2020-04-15 16:39:26', '2020-04-15 16:39:26', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3NTQ0MjUwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODY5Mzg0NTB9.smVP-Rl1EkAuUVqXW7z0mpxA5O86vXj0oH4FukG-NVA');
INSERT INTO `jwt_blacklists` VALUES (41, '2020-04-15 17:08:06', '2020-04-15 17:08:06', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3NTQ0NzgxLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODY5Mzg5ODF9.OMZ08Y8aPuj40-NGEQ402LyRFBpkLWzzaqD3_tvj1h8');
INSERT INTO `jwt_blacklists` VALUES (42, '2020-04-15 17:08:28', '2020-04-15 17:08:28', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3NTQ2NDk0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODY5NDA2OTR9.9lsoTbZrwhZ8kMXiH-Ta3A4h_yp7SwLj57mo_u5mrk4');
INSERT INTO `jwt_blacklists` VALUES (43, '2020-04-15 17:10:24', '2020-04-15 17:10:24', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3NTQ2NTE1LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODY5NDA3MTV9.5SrUrUmd4YhzlGmSpA9xJW_wbjV6yI6ty_NriIceOQo');
INSERT INTO `jwt_blacklists` VALUES (44, '2020-04-15 17:11:43', '2020-04-15 17:11:43', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3NTQ2NjI5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODY5NDA4Mjl9.SFiomEpOshboOe0JGDa1HlJt5aQIF7IeyOsoDwl1o8E');
INSERT INTO `jwt_blacklists` VALUES (45, '2020-04-15 17:12:54', '2020-04-15 17:12:54', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3NTQ2NzE4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODY5NDA5MTh9.xrwEknZQN2J3poarMTQvb7mX1Icicz2_f60kw36g9og');
INSERT INTO `jwt_blacklists` VALUES (46, '2020-04-15 17:14:47', '2020-04-15 17:14:47', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3NTQ2Nzg5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODY5NDA5ODl9.3jbdl1N0KA8ExFMWXHi3ha4aESKq8yDKDgpSH4Xdsnk');
INSERT INTO `jwt_blacklists` VALUES (47, '2020-04-22 12:04:20', '2020-04-22 12:04:20', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg3NjE2MTYwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODcwMTAzNjB9.jRHlnfXuJhp4hBE-QqCZ-lodzwK67IBkDI2xteB0OQw');
INSERT INTO `jwt_blacklists` VALUES (48, '2020-04-22 12:12:17', '2020-04-22 12:12:17', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg4MTMzMjQyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODc1Mjc0NDJ9.WJ59uRUxXJ7-rUH07mE6jCfnwgfvQnpPaLU5vJ_VhWM');
INSERT INTO `jwt_blacklists` VALUES (49, '2020-05-26 13:47:57', '2020-05-26 13:47:57', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTkwNzM5NzA2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1OTAxMzM5MDZ9.qLOiJZNi0BzUsfVf-RXlxKuSvKdiCzujj4U1MjBLIMo');
INSERT INTO `jwt_blacklists` VALUES (50, '2020-05-26 13:50:17', '2020-05-26 13:50:17', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTkxMDc2OTYyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1OTA0NzExNjJ9.DzDQLh3rcR2Q1SDFhgyldwxvmiM5pfWQXEJQznwKKCo');
INSERT INTO `jwt_blacklists` VALUES (51, '2020-05-26 13:56:21', '2020-05-26 13:56:21', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTkxMDc3MzA3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1OTA0NzE1MDd9.EN12QfW7pyQgSXH3hEyqK0xE2RG5gq-h-gb88ryI8GE');
INSERT INTO `jwt_blacklists` VALUES (52, '2020-05-26 13:57:50', '2020-05-26 13:57:50', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTkxMDc3MzkxLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1OTA0NzE1OTF9.yws3hTfYo1I28yPTbGs3kJunDSwL8onFWQ3XRmHTSo8');
INSERT INTO `jwt_blacklists` VALUES (53, '2020-05-26 14:01:10', '2020-05-26 14:01:10', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTkxMDc3NjQ5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1OTA0NzE4NDl9.bk_5o2it5DrUDowF-gcoT9wN4Lj4mC5PMf-VfCeQjkk');
INSERT INTO `jwt_blacklists` VALUES (54, '2020-05-26 14:01:33', '2020-05-26 14:01:33', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYTRiMGY2NjAtM2NhMi00MGViLWJkMjYtNzI4MmEzMGYxMmVjIiwiSUQiOjE0LCJOaWNrTmFtZSI6InpoYW5nanVuIiwiQXV0aG9yaXR5SWQiOiJzdG9jayIsImV4cCI6MTU5MTA3NzY4OCwiaXNzIjoicW1QbHVzIiwibmJmIjoxNTkwNDcxODg4fQ.n_Kqd8fXuKC-09Pk6ArGN2n_QDTXx1sKY3PcjQkCJqQ');
COMMIT;

-- ----------------------------
-- Table structure for stock_monitor
-- ----------------------------
DROP TABLE IF EXISTS `stock_monitor`;
CREATE TABLE `stock_monitor` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `symbol` varchar(255) DEFAULT NULL,
  `monitor_high` double DEFAULT NULL,
  `monitor_low` double DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `del_status` bigint(20) DEFAULT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  `is_day` tinyint(1) DEFAULT NULL,
  `day` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `base_index` (`symbol`,`del_status`,`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of stock_monitor
-- ----------------------------
BEGIN;
INSERT INTO `stock_monitor` VALUES (1, 'SH603950', 1, 50, 10, 1590471852, 1590462766, 0, NULL);
INSERT INTO `stock_monitor` VALUES (2, 'SH688003', 60, 22, 10, 1590470752, 1590462813, 0, NULL);
INSERT INTO `stock_monitor` VALUES (3, 'SZ002400', 90, 22, 10, 1590470747, 1590465041, 0, NULL);
INSERT INTO `stock_monitor` VALUES (4, 'SZ300792', 400, 360, 14, 0, 1590471772, 0, NULL);
INSERT INTO `stock_monitor` VALUES (5, 'SH603605', 170, 150, 14, 0, 1590471819, 0, NULL);
INSERT INTO `stock_monitor` VALUES (6, 'SZ000869', 30, 25, 14, 0, 1590471898, 0, NULL);
INSERT INTO `stock_monitor` VALUES (7, 'SH603969', 2, 0, 10, 0, 1590550438, 0, NULL);
INSERT INTO `stock_monitor` VALUES (8, 'SZ300836', 0, 0, 10, 0, 1590662702, 1, 1590595200);
INSERT INTO `stock_monitor` VALUES (9, 'SZ300061', 0, 0, 10, 1590662867, 1590662860, 1, 1590595200);
INSERT INTO `stock_monitor` VALUES (10, 'SH600069', 0, 0, 10, 1590664745, 1590662861, 1, 1590595200);
INSERT INTO `stock_monitor` VALUES (11, 'SH688010', 0, 0, 10, 0, 1590662862, 1, 1590595200);
INSERT INTO `stock_monitor` VALUES (12, 'SH688126', 0, 0, 10, 0, 1590664740, 1, 1590595200);
COMMIT;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `authority_id` int(10) unsigned DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `api_group` varchar(255) DEFAULT NULL,
  `method` varchar(255) DEFAULT 'POST',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_apis_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_apis_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
BEGIN;
INSERT INTO `sys_apis` VALUES (1, '2019-09-28 11:23:49', '2019-09-28 17:06:16', NULL, NULL, '/base/login', '用户登录', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2019-09-28 11:32:46', '2019-09-28 17:06:11', NULL, NULL, '/base/register', '用户注册', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (3, '2019-09-28 11:33:41', '2020-05-09 17:43:15', NULL, NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (4, '2019-09-28 14:09:04', '2019-09-28 17:05:59', NULL, NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2019-09-28 14:15:50', '2019-09-28 17:05:53', NULL, NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (7, '2019-09-28 14:19:26', '2019-09-28 17:05:44', NULL, NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (8, '2019-09-28 14:19:48', '2019-09-28 17:05:39', NULL, NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2019-09-30 15:05:38', '2019-09-30 15:05:38', NULL, NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2019-09-30 15:23:09', '2019-09-30 15:23:09', NULL, NULL, '/authority/createAuthority', '创建角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2019-09-30 15:23:33', '2019-09-30 15:23:33', NULL, NULL, '/authority/deleteAuthority', '删除角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2019-09-30 15:23:57', '2019-09-30 15:23:57', NULL, NULL, '/authority/getAuthorityList', '获取角色列表', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2019-09-30 15:24:20', '2019-09-30 15:24:20', NULL, NULL, '/menu/getMenu', '获取菜单树', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2019-09-30 15:24:50', '2019-09-30 15:24:50', NULL, NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2019-09-30 15:25:07', '2019-09-30 15:25:07', NULL, NULL, '/menu/addBaseMenu', '新增菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2019-09-30 15:25:25', '2019-09-30 15:25:25', NULL, NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2019-09-30 15:25:53', '2019-09-30 15:25:53', NULL, NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2019-09-30 15:26:20', '2019-09-30 15:26:20', NULL, NULL, '/menu/getMenuAuthority', '获取指定角色menu', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (20, '2019-09-30 15:26:43', '2019-09-30 15:26:43', NULL, NULL, '/menu/deleteBaseMenu', '删除菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2019-09-30 15:28:05', '2019-09-30 15:28:05', NULL, NULL, '/menu/updateBaseMenu', '更新菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (22, '2019-09-30 15:28:21', '2019-09-30 15:28:21', NULL, NULL, '/menu/getBaseMenuById', '根据id获取菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2019-09-30 15:29:19', '2019-09-30 15:29:19', NULL, NULL, '/user/changePassword', '修改密码', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (24, '2019-09-30 15:29:33', '2019-09-30 15:29:33', NULL, NULL, '/user/uploadHeaderImg', '上传头像', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2019-09-30 15:30:00', '2020-05-06 16:03:47', NULL, NULL, '/user/deleteUser', '删除用户', 'user', 'DELETE');
INSERT INTO `sys_apis` VALUES (28, '2019-10-09 15:15:17', '2019-10-09 15:17:07', NULL, NULL, '/user/getUserList', '获取用户列表', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2019-10-09 23:01:40', '2019-10-09 23:01:40', NULL, NULL, '/user/setUserAuthority', '修改用户角色', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2019-10-26 20:14:38', '2019-10-26 20:14:38', NULL, NULL, '/fileUploadAndDownload/upload', '文件上传示例', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2019-10-26 20:14:59', '2019-10-26 20:14:59', NULL, NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2019-12-12 13:28:47', '2019-12-12 13:28:47', NULL, NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2019-12-12 13:28:59', '2019-12-12 13:28:59', NULL, NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2019-12-12 17:02:15', '2019-12-12 17:02:15', NULL, NULL, '/fileUploadAndDownload/deleteFile', '删除文件', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2019-12-28 18:18:07', '2019-12-28 18:18:07', NULL, NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2020-01-06 17:56:36', '2020-01-06 17:56:36', NULL, NULL, '/authority/setDataAuthority', '设置角色资源权限', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (37, '2020-01-13 14:04:05', '2020-01-13 14:04:05', NULL, NULL, '/system/getSystemConfig', '获取配置文件内容', 'system', 'POST');
INSERT INTO `sys_apis` VALUES (38, '2020-01-13 15:02:06', '2020-01-13 15:02:06', NULL, NULL, '/system/setSystemConfig', '设置配置文件内容', 'system', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2020-02-25 15:32:39', '2020-02-25 15:32:39', NULL, NULL, '/customer/customer', '创建客户', 'customer', 'POST');
INSERT INTO `sys_apis` VALUES (40, '2020-02-25 15:32:51', '2020-02-25 15:34:56', NULL, NULL, '/customer/customer', '更新客户', 'customer', 'PUT');
INSERT INTO `sys_apis` VALUES (41, '2020-02-25 15:33:57', '2020-02-25 15:33:57', NULL, NULL, '/customer/customer', '删除客户', 'customer', 'DELETE');
INSERT INTO `sys_apis` VALUES (42, '2020-02-25 15:36:48', '2020-02-25 15:37:16', NULL, NULL, '/customer/customer', '获取单一客户', 'customer', 'GET');
INSERT INTO `sys_apis` VALUES (43, '2020-02-25 15:37:06', '2020-02-25 15:37:06', NULL, NULL, '/customer/customerList', '获取客户列表', 'customer', 'GET');
INSERT INTO `sys_apis` VALUES (44, '2020-03-12 14:36:54', '2020-03-12 14:56:50', NULL, NULL, '/casbin/casbinTest/:pathParam', 'RESTFUL模式测试', 'casbin', 'GET');
INSERT INTO `sys_apis` VALUES (45, '2020-03-29 23:01:28', '2020-03-29 23:01:28', NULL, NULL, '/autoCode/createTemp', '自动化代码', 'autoCode', 'POST');
INSERT INTO `sys_apis` VALUES (46, '2020-04-15 12:46:58', '2020-04-15 12:46:58', NULL, NULL, '/authority/updateAuthority', '更新角色信息', 'authority', 'PUT');
INSERT INTO `sys_apis` VALUES (47, '2020-04-20 15:14:25', '2020-04-20 15:14:25', NULL, NULL, '/authority/copyAuthority', '拷贝角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (48, '2020-05-24 16:33:25', '2020-05-24 16:43:26', NULL, NULL, '/stock/top/list', '获取Top列表', 'stock', 'POST');
INSERT INTO `sys_apis` VALUES (49, '2020-05-26 10:48:29', '2020-05-26 10:48:29', NULL, NULL, '/stock/monitor/list', '监控列表', 'stock', 'POST');
INSERT INTO `sys_apis` VALUES (50, '2020-05-26 11:11:36', '2020-05-26 11:11:36', NULL, NULL, '/stock/monitor/add', '新增监控', 'stock', 'POST');
INSERT INTO `sys_apis` VALUES (51, '2020-05-26 11:46:33', '2020-05-26 11:46:33', NULL, NULL, '/stock/monitor/one', '查询单个监控', 'stock', 'POST');
INSERT INTO `sys_apis` VALUES (52, '2020-05-26 11:47:01', '2020-05-26 11:47:01', NULL, NULL, '/stock/monitor/del', '删除监控', 'stock', 'POST');
INSERT INTO `sys_apis` VALUES (53, '2020-05-26 11:47:20', '2020-05-26 11:47:20', NULL, NULL, '/stock/monitor/update', '更新监控', 'stock', 'POST');
COMMIT;

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities` (
  `authority_id` varchar(255) NOT NULL,
  `authority_name` varchar(255) DEFAULT NULL,
  `parent_id` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE KEY `authority_id` (`authority_id`) USING BTREE,
  KEY `idx_sys_authorities_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
BEGIN;
INSERT INTO `sys_authorities` VALUES ('888', '普通用户', '0', '2020-04-04 11:44:56', '2020-05-09 17:41:29', NULL);
INSERT INTO `sys_authorities` VALUES ('8881', '普通用户子角色', '888', '2020-04-04 11:44:56', '2020-05-09 17:41:29', NULL);
INSERT INTO `sys_authorities` VALUES ('9528', '测试角色', '0', '2020-04-04 11:44:56', '2020-05-09 17:41:29', NULL);
INSERT INTO `sys_authorities` VALUES ('stock', '股票', '0', '2020-05-26 13:46:01', '2020-05-26 13:46:01', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_authority_authority_id` varchar(255) NOT NULL,
  `sys_base_menu_id` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`sys_authority_authority_id`,`sys_base_menu_id`) USING BTREE,
  KEY `sys_authority_authority_id` (`sys_authority_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_authority_menus` VALUES ('888', 1);
INSERT INTO `sys_authority_menus` VALUES ('888', 2);
INSERT INTO `sys_authority_menus` VALUES ('888', 3);
INSERT INTO `sys_authority_menus` VALUES ('888', 4);
INSERT INTO `sys_authority_menus` VALUES ('888', 5);
INSERT INTO `sys_authority_menus` VALUES ('888', 6);
INSERT INTO `sys_authority_menus` VALUES ('888', 17);
INSERT INTO `sys_authority_menus` VALUES ('888', 18);
INSERT INTO `sys_authority_menus` VALUES ('888', 19);
INSERT INTO `sys_authority_menus` VALUES ('888', 20);
INSERT INTO `sys_authority_menus` VALUES ('888', 21);
INSERT INTO `sys_authority_menus` VALUES ('888', 22);
INSERT INTO `sys_authority_menus` VALUES ('888', 23);
INSERT INTO `sys_authority_menus` VALUES ('888', 26);
INSERT INTO `sys_authority_menus` VALUES ('888', 33);
INSERT INTO `sys_authority_menus` VALUES ('888', 34);
INSERT INTO `sys_authority_menus` VALUES ('888', 38);
INSERT INTO `sys_authority_menus` VALUES ('888', 40);
INSERT INTO `sys_authority_menus` VALUES ('888', 41);
INSERT INTO `sys_authority_menus` VALUES ('888', 42);
INSERT INTO `sys_authority_menus` VALUES ('888', 45);
INSERT INTO `sys_authority_menus` VALUES ('888', 47);
INSERT INTO `sys_authority_menus` VALUES ('888', 48);
INSERT INTO `sys_authority_menus` VALUES ('888', 49);
INSERT INTO `sys_authority_menus` VALUES ('888', 50);
INSERT INTO `sys_authority_menus` VALUES ('8881', 1);
INSERT INTO `sys_authority_menus` VALUES ('8881', 2);
INSERT INTO `sys_authority_menus` VALUES ('8881', 18);
INSERT INTO `sys_authority_menus` VALUES ('8881', 38);
INSERT INTO `sys_authority_menus` VALUES ('8881', 40);
INSERT INTO `sys_authority_menus` VALUES ('8881', 41);
INSERT INTO `sys_authority_menus` VALUES ('8881', 42);
INSERT INTO `sys_authority_menus` VALUES ('9528', 1);
INSERT INTO `sys_authority_menus` VALUES ('9528', 2);
INSERT INTO `sys_authority_menus` VALUES ('9528', 3);
INSERT INTO `sys_authority_menus` VALUES ('9528', 4);
INSERT INTO `sys_authority_menus` VALUES ('9528', 5);
INSERT INTO `sys_authority_menus` VALUES ('9528', 6);
INSERT INTO `sys_authority_menus` VALUES ('9528', 17);
INSERT INTO `sys_authority_menus` VALUES ('9528', 18);
INSERT INTO `sys_authority_menus` VALUES ('9528', 19);
INSERT INTO `sys_authority_menus` VALUES ('9528', 20);
INSERT INTO `sys_authority_menus` VALUES ('9528', 21);
INSERT INTO `sys_authority_menus` VALUES ('9528', 22);
INSERT INTO `sys_authority_menus` VALUES ('9528', 23);
INSERT INTO `sys_authority_menus` VALUES ('9528', 26);
INSERT INTO `sys_authority_menus` VALUES ('9528', 33);
INSERT INTO `sys_authority_menus` VALUES ('9528', 34);
INSERT INTO `sys_authority_menus` VALUES ('9528', 38);
INSERT INTO `sys_authority_menus` VALUES ('9528', 40);
INSERT INTO `sys_authority_menus` VALUES ('9528', 41);
INSERT INTO `sys_authority_menus` VALUES ('9528', 42);
INSERT INTO `sys_authority_menus` VALUES ('stock', 1);
INSERT INTO `sys_authority_menus` VALUES ('stock', 18);
INSERT INTO `sys_authority_menus` VALUES ('stock', 47);
INSERT INTO `sys_authority_menus` VALUES ('stock', 48);
INSERT INTO `sys_authority_menus` VALUES ('stock', 49);
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `menu_level` int(10) unsigned DEFAULT NULL,
  `parent_id` int(10) unsigned DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `hidden` tinyint(1) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  `sort` int(255) DEFAULT NULL,
  `keep_alive` tinyint(1) DEFAULT NULL,
  `default_menu` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_base_menus_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menus` VALUES (1, '2019-09-19 22:05:18', '2020-05-09 17:41:51', NULL, 0, 0, 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', '仪表盘', 'setting', '仪表盘', 1, 0, 0);
INSERT INTO `sys_base_menus` VALUES (2, '2019-09-19 22:06:17', '2020-05-24 16:38:41', NULL, 0, 0, 'about', 'about', 1, 'view/about/index.vue', '关于我们', 'info', '测试菜单', 7, 0, 0);
INSERT INTO `sys_base_menus` VALUES (3, '2019-09-19 22:06:38', '2020-05-09 17:41:29', NULL, 0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', '超级管理员', 3, 0, 0);
INSERT INTO `sys_base_menus` VALUES (4, '2019-09-19 22:11:53', '2020-05-09 17:41:29', NULL, 0, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', '角色管理', 1, 0, 0);
INSERT INTO `sys_base_menus` VALUES (5, '2019-09-19 22:13:18', '2020-05-09 17:41:29', NULL, 0, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', '菜单管理', 2, 1, 0);
INSERT INTO `sys_base_menus` VALUES (6, '2019-09-19 22:13:36', '2020-05-09 17:41:29', NULL, 0, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 'api管理', 3, 1, 0);
INSERT INTO `sys_base_menus` VALUES (17, '2019-10-09 15:12:29', '2020-05-09 17:41:29', NULL, 0, 3, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', '用户管理', 4, 0, 0);
INSERT INTO `sys_base_menus` VALUES (18, '2019-10-15 22:27:22', '2020-05-09 17:41:29', NULL, 0, 0, 'person', 'person', 1, 'view/person/person.vue', '个人信息', 'user-solid', '个人信息', 4, 0, 0);
INSERT INTO `sys_base_menus` VALUES (19, '2019-10-20 11:14:42', '2020-05-24 16:38:49', NULL, 0, 0, 'example', 'example', 1, 'view/example/index.vue', '示例文件', 's-management', '示例文件', 6, 0, 0);
INSERT INTO `sys_base_menus` VALUES (20, '2019-10-20 11:18:11', '2020-05-09 17:41:29', NULL, 0, 19, 'table', 'table', 0, 'view/example/table/table.vue', '表格示例', 's-order', '表格示例', 1, 0, 0);
INSERT INTO `sys_base_menus` VALUES (21, '2019-10-20 11:19:52', '2020-05-09 17:41:29', NULL, 0, 19, 'form', 'form', 0, 'view/example/form/form.vue', '表单示例', 'document', '表单示例', 2, 0, 0);
INSERT INTO `sys_base_menus` VALUES (22, '2019-10-20 11:22:19', '2020-05-09 17:41:29', NULL, 0, 19, 'rte', 'rte', 0, 'view/example/rte/rte.vue', '富文本编辑器', 'reading', '富文本编辑器', 3, 0, 0);
INSERT INTO `sys_base_menus` VALUES (23, '2019-10-20 11:23:39', '2020-05-09 17:41:29', NULL, 0, 19, 'excel', 'excel', 0, 'view/example/excel/excel.vue', 'excel导入导出', 's-marketing', 'excel导入导出', 4, 0, 0);
INSERT INTO `sys_base_menus` VALUES (26, '2019-10-20 11:27:02', '2020-05-09 17:41:29', NULL, 0, 19, 'upload', 'upload', 0, 'view/example/upload/upload.vue', '上传下载', 'upload', '上传下载', 5, 0, 0);
INSERT INTO `sys_base_menus` VALUES (33, '2020-02-17 16:20:47', '2020-05-09 17:41:29', NULL, 0, 19, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', '断点续传', 'upload', '断点续传', 6, 0, 0);
INSERT INTO `sys_base_menus` VALUES (34, '2020-02-24 19:48:37', '2020-05-09 17:41:29', NULL, 0, 19, 'customer', 'customer', 0, 'view/example/customer/customer.vue', '客户列表（资源示例）', 's-custom', '客户列表（资源示例）', 7, 0, 0);
INSERT INTO `sys_base_menus` VALUES (38, '2020-03-29 21:31:03', '2020-05-09 17:41:29', NULL, 0, 0, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', '系统工具', 's-cooperation', '系统工具', 5, 0, 0);
INSERT INTO `sys_base_menus` VALUES (40, '2020-03-29 21:35:10', '2020-05-09 17:41:29', NULL, 0, 38, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', '代码生成器', 'cpu', '代码生成器', 1, 0, 0);
INSERT INTO `sys_base_menus` VALUES (41, '2020-03-29 21:36:26', '2020-05-09 17:41:29', NULL, 0, 38, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', '表单生成器', 'magic-stick', '表单生成器', 2, 0, 0);
INSERT INTO `sys_base_menus` VALUES (42, '2020-04-02 14:19:36', '2020-05-09 17:41:29', NULL, 0, 38, 'system', 'system', 0, 'view/systemTools/system/system.vue', '系统配置', 's-operation', '系统配置', 3, 0, 0);
INSERT INTO `sys_base_menus` VALUES (45, '2020-04-29 17:19:34', '2020-05-24 16:38:32', NULL, 0, 0, 'iconList', 'iconList', 1, 'view/iconList/index.vue', '图标集合', 'star-on', NULL, 2, 0, 0);
INSERT INTO `sys_base_menus` VALUES (47, '2020-05-24 16:34:55', '2020-05-24 16:36:07', NULL, 0, 0, 'stock', 'stock', 0, 'view/stock/index.vue', '股票', 's-marketing', NULL, 5, 0, 0);
INSERT INTO `sys_base_menus` VALUES (48, '2020-05-24 16:37:34', '2020-05-24 16:39:28', NULL, 0, 47, 'top', 'top', 0, 'view/stock/top/top.vue', 'Top列表', 'data-line', NULL, 1, 0, 0);
INSERT INTO `sys_base_menus` VALUES (49, '2020-05-26 10:45:09', '2020-05-28 18:45:42', NULL, 0, 47, 'monitor', 'monitor', 0, 'view/stock/monitor/monitor.vue', '监控列表', 's-custom', NULL, 3, 0, 0);
INSERT INTO `sys_base_menus` VALUES (50, '2020-05-28 18:29:56', '2020-05-28 18:45:49', NULL, 0, 47, 'day', 'day', 0, 'view/stock/monitor/day.vue', '监控当天', 'star-on', NULL, 2, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` varchar(255) NOT NULL,
  `data_authority_id` varchar(255) NOT NULL,
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id`) USING BTREE,
  KEY `sys_authority_authority_id` (`sys_authority_authority_id`) USING BTREE,
  KEY `data_authority_id` (`data_authority_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
BEGIN;
INSERT INTO `sys_data_authority_id` VALUES ('888', '888');
INSERT INTO `sys_data_authority_id` VALUES ('888', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('888', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('88822', '888');
INSERT INTO `sys_data_authority_id` VALUES ('88822', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('88822', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('888222', '888');
INSERT INTO `sys_data_authority_id` VALUES ('888222', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('888222', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('8883', '888');
INSERT INTO `sys_data_authority_id` VALUES ('8883', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('8883', '9528');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '8881');
INSERT INTO `sys_data_authority_id` VALUES ('9528', '9528');
COMMIT;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `uuid` varbinary(255) DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `pass_word` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT 'QMPlusUser',
  `header_img` varchar(255) DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg',
  `authority_id` varchar(255) DEFAULT '888',
  `authority_name` varchar(255) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phone_data` varchar(255) DEFAULT NULL,
  `manager` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_users_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_users_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` VALUES (10, '2019-09-13 17:23:46', '2020-05-24 16:42:48', NULL, 0x63653064363638352D633135662D343132362D613562342D383930626339643233353664, NULL, NULL, '超级管理员', 'http://qmplusimg.henrongyi.top/15887525450B978439-F04A-4a09-A8D3-DE7DE2677142.png', '888', NULL, 'admin', 'e10adc3949ba59abbe56e057f20f883e', NULL, NULL);
INSERT INTO `sys_users` VALUES (11, '2019-09-13 17:27:29', '2020-05-09 17:43:44', NULL, 0x66643665663739622D393434632D343838382D383337372D616265326432363038383538, NULL, NULL, 'QMPlusUser', 'http://qmplusimg.henrongyi.top/1572075907logo.png', '9528', NULL, 'a303176530', '3ec063004a6f31642261936a379fde3d', NULL, NULL);
INSERT INTO `sys_users` VALUES (13, '2020-05-24 16:42:35', '2020-05-26 13:46:37', '2020-05-26 13:46:45', 0x34353464643334612D353066652D346338382D386461342D303561336666633737393033, NULL, NULL, 'zj', 'http://qmplusimg.henrongyi.top/1590309743timg.jpg', 'stock', NULL, 'zj', 'e10adc3949ba59abbe56e057f20f883e', NULL, NULL);
INSERT INTO `sys_users` VALUES (14, '2020-05-26 13:47:54', '2020-05-26 13:47:54', NULL, 0x61346230663636302D336361322D343065622D626432362D373238326133306631326563, NULL, NULL, 'zhangjun', 'http://qmplusimg.henrongyi.top/1590472067timg.jpg', 'stock', NULL, 'zhangjun', 'cfb70ae5d9a8e90529b4a27a89041d85', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_workflow_step_infos
-- ----------------------------
DROP TABLE IF EXISTS `sys_workflow_step_infos`;
CREATE TABLE `sys_workflow_step_infos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `workflow_id` int(10) unsigned DEFAULT NULL,
  `is_strat` tinyint(1) DEFAULT NULL,
  `step_name` varchar(255) DEFAULT NULL,
  `step_no` double DEFAULT NULL,
  `step_authority_id` varchar(255) DEFAULT NULL,
  `is_end` tinyint(1) DEFAULT NULL,
  `sys_workflow_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_workflow_step_infos_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_workflow_step_infos_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of sys_workflow_step_infos
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_workflows
-- ----------------------------
DROP TABLE IF EXISTS `sys_workflows`;
CREATE TABLE `sys_workflows` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `workflow_nick_name` varchar(255) DEFAULT NULL,
  `workflow_name` varchar(255) DEFAULT NULL,
  `workflow_description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_workflows_deleted_at` (`deleted_at`) USING BTREE,
  KEY `idx_sys_workflows_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of sys_workflows
-- ----------------------------
BEGIN;
INSERT INTO `sys_workflows` VALUES (8, '2019-12-09 15:20:21', '2019-12-09 15:20:21', NULL, '测试改版1', 'test', '123123');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
