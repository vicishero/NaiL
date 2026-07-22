
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `paopao` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `paopao`;
DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=383 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES (193,'p','888','/api/createApi','POST','','',''),(196,'p','888','/api/deleteApi','POST','','',''),(199,'p','888','/api/deleteApisByIds','DELETE','','',''),(202,'p','888','/api/enterSyncApi','POST','','',''),(198,'p','888','/api/getAllApis','POST','','',''),(195,'p','888','/api/getApiById','POST','','',''),(201,'p','888','/api/getApiGroups','GET','','',''),(194,'p','888','/api/getApiList','POST','','',''),(204,'p','888','/api/getApiRoles','GET','','',''),(203,'p','888','/api/ignoreApi','POST','','',''),(205,'p','888','/api/setApiRoles','POST','','',''),(200,'p','888','/api/syncApi','GET','','',''),(197,'p','888','/api/updateApi','POST','','',''),(367,'p','888','/attachmentCategory/addCategory','POST','','',''),(368,'p','888','/attachmentCategory/deleteCategory','POST','','',''),(366,'p','888','/attachmentCategory/getCategoryList','GET','','',''),(206,'p','888','/authority/copyAuthority','POST','','',''),(208,'p','888','/authority/createAuthority','POST','','',''),(209,'p','888','/authority/deleteAuthority','POST','','',''),(210,'p','888','/authority/getAuthorityList','POST','','',''),(212,'p','888','/authority/getUsersByAuthority','GET','','',''),(211,'p','888','/authority/setDataAuthority','POST','','',''),(213,'p','888','/authority/setRoleUsers','POST','','',''),(207,'p','888','/authority/updateAuthority','PUT','','',''),(335,'p','888','/authorityBtn/canRemoveAuthorityBtn','POST','','',''),(334,'p','888','/authorityBtn/getAuthorityBtn','POST','','',''),(333,'p','888','/authorityBtn/setAuthorityBtn','POST','','',''),(293,'p','888','/autoCode/addFunc','POST','','',''),(284,'p','888','/autoCode/createPackage','POST','','',''),(288,'p','888','/autoCode/createPlug','POST','','',''),(281,'p','888','/autoCode/createTemp','POST','','',''),(304,'p','888','/autoCode/deleteAIWorkflowSession','POST','','',''),(287,'p','888','/autoCode/delPackage','POST','','',''),(282,'p','888','/autoCode/delSysHistory','POST','','',''),(305,'p','888','/autoCode/dumpAIWorkflowMarkdown','POST','','',''),(303,'p','888','/autoCode/getAIWorkflowSessionDetail','POST','','',''),(302,'p','888','/autoCode/getAIWorkflowSessionList','POST','','',''),(279,'p','888','/autoCode/getColumn','GET','','',''),(275,'p','888','/autoCode/getDB','GET','','',''),(276,'p','888','/autoCode/getMeta','POST','','',''),(286,'p','888','/autoCode/getPackage','POST','','',''),(292,'p','888','/autoCode/getPluginList','GET','','',''),(283,'p','888','/autoCode/getSysHistory','POST','','',''),(278,'p','888','/autoCode/getTables','GET','','',''),(285,'p','888','/autoCode/getTemplates','GET','','',''),(289,'p','888','/autoCode/installPlugin','POST','','',''),(294,'p','888','/autoCode/mcp','POST','','',''),(300,'p','888','/autoCode/mcpList','POST','','',''),(298,'p','888','/autoCode/mcpRoutes','POST','','',''),(296,'p','888','/autoCode/mcpStart','POST','','',''),(295,'p','888','/autoCode/mcpStatus','POST','','',''),(297,'p','888','/autoCode/mcpStop','POST','','',''),(299,'p','888','/autoCode/mcpTest','POST','','',''),(277,'p','888','/autoCode/preview','POST','','',''),(290,'p','888','/autoCode/pubPlug','POST','','',''),(291,'p','888','/autoCode/removePlugin','POST','','',''),(280,'p','888','/autoCode/rollback','POST','','',''),(301,'p','888','/autoCode/saveAIWorkflowSession','POST','','',''),(245,'p','888','/casbin/getPolicyPathByAuthorityId','POST','','',''),(244,'p','888','/casbin/updateCasbin','POST','','',''),(273,'p','888','/customer/customer','DELETE','','',''),(270,'p','888','/customer/customer','GET','','',''),(272,'p','888','/customer/customer','POST','','',''),(271,'p','888','/customer/customer','PUT','','',''),(274,'p','888','/customer/customerList','GET','','',''),(328,'p','888','/email/emailTest','POST','','',''),(329,'p','888','/email/sendEmail','POST','','',''),(237,'p','888','/fileUploadAndDownload/breakpointContinue','POST','','',''),(236,'p','888','/fileUploadAndDownload/breakpointContinueFinish','POST','','',''),(240,'p','888','/fileUploadAndDownload/deleteFile','POST','','',''),(241,'p','888','/fileUploadAndDownload/editFileName','POST','','',''),(235,'p','888','/fileUploadAndDownload/findFile','GET','','',''),(242,'p','888','/fileUploadAndDownload/getFileList','POST','','',''),(243,'p','888','/fileUploadAndDownload/importURL','POST','','',''),(238,'p','888','/fileUploadAndDownload/removeChunk','POST','','',''),(239,'p','888','/fileUploadAndDownload/upload','POST','','',''),(353,'p','888','/info/createInfo','POST','','',''),(354,'p','888','/info/deleteInfo','DELETE','','',''),(355,'p','888','/info/deleteInfoByIds','DELETE','','',''),(357,'p','888','/info/findInfo','GET','','',''),(358,'p','888','/info/getInfoList','GET','','',''),(356,'p','888','/info/updateInfo','PUT','','',''),(246,'p','888','/jwt/jsonInBlacklist','POST','','',''),(216,'p','888','/menu/addBaseMenu','POST','','',''),(218,'p','888','/menu/addMenuAuthority','POST','','',''),(222,'p','888','/menu/deleteBaseMenu','POST','','',''),(224,'p','888','/menu/getBaseMenuById','POST','','',''),(217,'p','888','/menu/getBaseMenuTree','POST','','',''),(214,'p','888','/menu/getMenu','POST','','',''),(219,'p','888','/menu/getMenuAuthority','POST','','',''),(215,'p','888','/menu/getMenuList','POST','','',''),(220,'p','888','/menu/getMenuRoles','GET','','',''),(221,'p','888','/menu/setMenuRoles','POST','','',''),(223,'p','888','/menu/updateBaseMenu','POST','','',''),(331,'p','888','/simpleUploader/checkFileMd5','GET','','',''),(332,'p','888','/simpleUploader/mergeFileMd5','GET','','',''),(330,'p','888','/simpleUploader/upload','POST','','',''),(261,'p','888','/skills/createReference','POST','','',''),(258,'p','888','/skills/createResource','POST','','',''),(255,'p','888','/skills/createScript','POST','','',''),(264,'p','888','/skills/createTemplate','POST','','',''),(254,'p','888','/skills/deleteSkill','POST','','',''),(267,'p','888','/skills/getGlobalConstraint','POST','','',''),(262,'p','888','/skills/getReference','POST','','',''),(259,'p','888','/skills/getResource','POST','','',''),(256,'p','888','/skills/getScript','POST','','',''),(252,'p','888','/skills/getSkillDetail','POST','','',''),(251,'p','888','/skills/getSkillList','POST','','',''),(265,'p','888','/skills/getTemplate','POST','','',''),(250,'p','888','/skills/getTools','GET','','',''),(269,'p','888','/skills/packageSkill','POST','','',''),(268,'p','888','/skills/saveGlobalConstraint','POST','','',''),(263,'p','888','/skills/saveReference','POST','','',''),(260,'p','888','/skills/saveResource','POST','','',''),(257,'p','888','/skills/saveScript','POST','','',''),(253,'p','888','/skills/saveSkill','POST','','',''),(266,'p','888','/skills/saveTemplate','POST','','',''),(380,'p','888','/sysApiToken/createApiToken','POST','','',''),(382,'p','888','/sysApiToken/deleteApiToken','POST','','',''),(381,'p','888','/sysApiToken/getApiTokenList','POST','','',''),(318,'p','888','/sysDictionary/createSysDictionary','POST','','',''),(319,'p','888','/sysDictionary/deleteSysDictionary','DELETE','','',''),(321,'p','888','/sysDictionary/exportSysDictionary','GET','','',''),(315,'p','888','/sysDictionary/findSysDictionary','GET','','',''),(317,'p','888','/sysDictionary/getSysDictionaryList','GET','','',''),(320,'p','888','/sysDictionary/importSysDictionary','POST','','',''),(316,'p','888','/sysDictionary/updateSysDictionary','PUT','','',''),(308,'p','888','/sysDictionaryDetail/createSysDictionaryDetail','POST','','',''),(310,'p','888','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','',''),(306,'p','888','/sysDictionaryDetail/findSysDictionaryDetail','GET','','',''),(313,'p','888','/sysDictionaryDetail/getDictionaryDetailsByParent','GET','','',''),(314,'p','888','/sysDictionaryDetail/getDictionaryPath','GET','','',''),(311,'p','888','/sysDictionaryDetail/getDictionaryTreeList','GET','','',''),(312,'p','888','/sysDictionaryDetail/getDictionaryTreeListByType','GET','','',''),(309,'p','888','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','',''),(307,'p','888','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','',''),(346,'p','888','/sysError/createSysError','POST','','',''),(347,'p','888','/sysError/deleteSysError','DELETE','','',''),(348,'p','888','/sysError/deleteSysErrorByIds','DELETE','','',''),(350,'p','888','/sysError/findSysError','GET','','',''),(351,'p','888','/sysError/getSysErrorList','GET','','',''),(352,'p','888','/sysError/getSysErrorSolution','GET','','',''),(349,'p','888','/sysError/updateSysError','PUT','','',''),(336,'p','888','/sysExportTemplate/createSysExportTemplate','POST','','',''),(337,'p','888','/sysExportTemplate/deleteSysExportTemplate','DELETE','','',''),(338,'p','888','/sysExportTemplate/deleteSysExportTemplateByIds','DELETE','','',''),(342,'p','888','/sysExportTemplate/exportExcel','GET','','',''),(343,'p','888','/sysExportTemplate/exportTemplate','GET','','',''),(340,'p','888','/sysExportTemplate/findSysExportTemplate','GET','','',''),(341,'p','888','/sysExportTemplate/getSysExportTemplateList','GET','','',''),(345,'p','888','/sysExportTemplate/importExcel','POST','','',''),(344,'p','888','/sysExportTemplate/previewSQL','GET','','',''),(339,'p','888','/sysExportTemplate/updateSysExportTemplate','PUT','','',''),(376,'p','888','/sysLoginLog/deleteLoginLog','DELETE','','',''),(377,'p','888','/sysLoginLog/deleteLoginLogByIds','DELETE','','',''),(378,'p','888','/sysLoginLog/findLoginLog','GET','','',''),(379,'p','888','/sysLoginLog/getLoginLogList','GET','','',''),(324,'p','888','/sysOperationRecord/createSysOperationRecord','POST','','',''),(326,'p','888','/sysOperationRecord/deleteSysOperationRecord','DELETE','','',''),(327,'p','888','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','',''),(322,'p','888','/sysOperationRecord/findSysOperationRecord','GET','','',''),(325,'p','888','/sysOperationRecord/getSysOperationRecordList','GET','','',''),(323,'p','888','/sysOperationRecord/updateSysOperationRecord','PUT','','',''),(359,'p','888','/sysParams/createSysParams','POST','','',''),(360,'p','888','/sysParams/deleteSysParams','DELETE','','',''),(361,'p','888','/sysParams/deleteSysParamsByIds','DELETE','','',''),(363,'p','888','/sysParams/findSysParams','GET','','',''),(365,'p','888','/sysParams/getSysParam','GET','','',''),(364,'p','888','/sysParams/getSysParamsList','GET','','',''),(362,'p','888','/sysParams/updateSysParams','PUT','','',''),(249,'p','888','/system/getServerInfo','POST','','',''),(247,'p','888','/system/getSystemConfig','POST','','',''),(248,'p','888','/system/setSystemConfig','POST','','',''),(374,'p','888','/sysVersion/deleteSysVersion','DELETE','','',''),(375,'p','888','/sysVersion/deleteSysVersionByIds','DELETE','','',''),(371,'p','888','/sysVersion/downloadVersionJson','GET','','',''),(372,'p','888','/sysVersion/exportVersion','POST','','',''),(369,'p','888','/sysVersion/findSysVersion','GET','','',''),(370,'p','888','/sysVersion/getSysVersionList','GET','','',''),(373,'p','888','/sysVersion/importVersion','POST','','',''),(192,'p','888','/user/admin_register','POST','','',''),(230,'p','888','/user/changePassword','POST','','',''),(229,'p','888','/user/deleteUser','DELETE','','',''),(225,'p','888','/user/getUserInfo','GET','','',''),(228,'p','888','/user/getUserList','POST','','',''),(233,'p','888','/user/resetPassword','POST','','',''),(227,'p','888','/user/setSelfInfo','PUT','','',''),(234,'p','888','/user/setSelfSetting','PUT','','',''),(232,'p','888','/user/setUserAuthorities','POST','','',''),(231,'p','888','/user/setUserAuthority','POST','','',''),(226,'p','888','/user/setUserInfo','PUT','','','');
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `exa_attachment_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_attachment_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '分类名称',
  `pid` bigint DEFAULT '0' COMMENT '父节点ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_attachment_category_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `exa_attachment_category` WRITE;
/*!40000 ALTER TABLE `exa_attachment_category` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_attachment_category` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `exa_customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_customers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `exa_customers` WRITE;
/*!40000 ALTER TABLE `exa_customers` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_customers` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `exa_file_chunks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_chunks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint unsigned DEFAULT NULL,
  `file_chunk_number` bigint DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `exa_file_chunks` WRITE;
/*!40000 ALTER TABLE `exa_file_chunks` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_file_chunks` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件名',
  `class_id` bigint DEFAULT '0' COMMENT '分类id',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `exa_file_upload_and_downloads` WRITE;
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_file_upload_and_downloads` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `exa_files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exa_files` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `chunk_total` bigint DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `exa_files` WRITE;
/*!40000 ALTER TABLE `exa_files` DISABLE KEYS */;
/*!40000 ALTER TABLE `exa_files` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `gva_announcements_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gva_announcements_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '公告标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '公告内容',
  `user_id` bigint DEFAULT NULL COMMENT '发布者',
  `attachments` json DEFAULT NULL COMMENT '相关附件',
  PRIMARY KEY (`id`),
  KEY `idx_gva_announcements_info_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `gva_announcements_info` WRITE;
/*!40000 ALTER TABLE `gva_announcements_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `gva_announcements_info` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `infra_evm_chain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `infra_evm_chain` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `chain_id` bigint NOT NULL,
  `name` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `rpc_url` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `block_interval_secs` bigint DEFAULT '12',
  `confirmations` bigint DEFAULT '6',
  `batch_size` bigint DEFAULT '2000',
  `catch_up_batch_size` bigint DEFAULT '5000',
  `catch_up_interval_secs` bigint DEFAULT '1',
  `start_block` bigint DEFAULT '0',
  `last_synced_block` bigint DEFAULT '0',
  `status` tinyint DEFAULT '1',
  `create_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_chain_id` (`chain_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `infra_evm_chain` WRITE;
/*!40000 ALTER TABLE `infra_evm_chain` DISABLE KEYS */;
INSERT INTO `infra_evm_chain` VALUES (1,137,'Polygon','https://polygon-rpc.com',2,12,2000,5000,1,0,0,1,NULL,NULL),(2,1,'Ethereum','https://eth.llamarpc.com',12,6,2000,5000,1,0,0,1,NULL,NULL);
/*!40000 ALTER TABLE `infra_evm_chain` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `infra_evm_contract_event`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `infra_evm_contract_event` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `chain_id` bigint NOT NULL,
  `contract_address` varchar(42) COLLATE utf8mb4_unicode_ci NOT NULL,
  `event_signature` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `event_name` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `alias` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT '',
  `topic0` varchar(66) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `start_block` bigint DEFAULT '0',
  `last_synced_block` bigint DEFAULT '0',
  `status` tinyint DEFAULT '1',
  `create_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_chain_contract_topic0` (`chain_id`,`contract_address`,`topic0`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `infra_evm_contract_event` WRITE;
/*!40000 ALTER TABLE `infra_evm_contract_event` DISABLE KEYS */;
INSERT INTO `infra_evm_contract_event` VALUES (1,137,'0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174','Transfer(address indexed from, address indexed to, uint256 value)','Transfer','nftTransfer','0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef',0,0,1,NULL,NULL);
/*!40000 ALTER TABLE `infra_evm_contract_event` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `infra_evm_event_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `infra_evm_event_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `chain_id` bigint NOT NULL,
  `contract_address` varchar(42) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `event_name` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `alias` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `block_number` bigint DEFAULT NULL,
  `block_hash` varchar(66) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tx_hash` varchar(66) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `log_index` bigint DEFAULT NULL,
  `topics` json DEFAULT NULL,
  `raw_data` text COLLATE utf8mb4_unicode_ci,
  `decoded_data` json DEFAULT NULL,
  `decode_status` smallint DEFAULT '0',
  `status` tinyint DEFAULT '0',
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_event_unique` (`chain_id`,`tx_hash`,`log_index`),
  KEY `idx_status` (`status`),
  KEY `idx_block_number` (`chain_id`,`block_number`),
  KEY `idx_contract` (`chain_id`,`contract_address`,`event_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `infra_evm_event_log` WRITE;
/*!40000 ALTER TABLE `infra_evm_event_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `infra_evm_event_log` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `infra_job`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `infra_job` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `handler_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `handler_param` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cron_expression` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint DEFAULT '1',
  `deleted` bit(1) DEFAULT NULL,
  `create_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `infra_job` WRITE;
/*!40000 ALTER TABLE `infra_job` DISABLE KEYS */;
INSERT INTO `infra_job` VALUES (1,'扫链[Polygon]','scanEvmChain','137','*/2 * * * * *',1,NULL,'2026-06-05 14:36:01.634','2026-06-05 16:14:40.986'),(2,'扫链[Ethereum]','scanEvmChain','1','*/12 * * * * *',1,NULL,'2026-06-05 14:36:01.636','2026-06-05 16:14:40.988'),(3,'事件消费处理','processScanEvent','','*/5 * * * * *',1,NULL,'2026-06-05 14:36:01.637','2026-06-05 16:14:40.989');
/*!40000 ALTER TABLE `infra_job` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `infra_job_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `infra_job_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `job_id` bigint NOT NULL,
  `status` tinyint DEFAULT '0',
  `message` text COLLATE utf8mb4_unicode_ci,
  `start_time` datetime(3) DEFAULT NULL,
  `end_time` datetime(3) DEFAULT NULL,
  `create_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_infra_job_log_job_id` (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `infra_job_log` WRITE;
/*!40000 ALTER TABLE `infra_job_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `infra_job_log` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `jwt_blacklists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `jwt_blacklists` WRITE;
/*!40000 ALTER TABLE `jwt_blacklists` DISABLE KEYS */;
/*!40000 ALTER TABLE `jwt_blacklists` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_attachment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_attachment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL DEFAULT '0',
  `file_size` bigint NOT NULL,
  `img_width` bigint NOT NULL DEFAULT '0',
  `img_height` bigint NOT NULL DEFAULT '0',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '1图片，2视频，3其他附件',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_attachment_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192582846009835521 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='附件';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_attachment` WRITE;
/*!40000 ALTER TABLE `p_attachment` DISABLE KEYS */;
INSERT INTO `p_attachment` VALUES (100041,100058,235201,0,0,1,'http://127.0.0.1:8008/oss/paopao/public/image/7f/75/bd/89/9626-4a2e-9163-7041bcd2b529.jpeg',1780455529,1780455529,0,0),(100042,100058,235201,0,0,1,'http://127.0.0.1:8008/oss/paopao/public/image/e7/b0/e3/a4/b80a-436a-8e21-02dc800c76ec.jpeg',1780455559,1780455559,0,0),(100043,100059,208717,0,0,1,'http://127.0.0.1:8008/oss/paopao/public/image/85/5a/f5/a2/84a6-49f1-bf44-2812df724f33.jpeg',1780457404,1780457404,0,0),(100044,100059,63029,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/f0/73/35/c2/2845-4b7d-8827-098eb1f49749.jpeg',1780459936,1780459936,0,0),(100045,100059,62710,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/8c/d2/f4/9f/a44a-4e97-9af2-6679ac6cb8c0.jpeg',1780459936,1780459936,0,0),(100046,100059,62377,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/66/cf/96/3e/483e-4721-8a80-0cbccbf819d3.jpeg',1780459936,1780459936,0,0),(100047,100059,208717,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/23/ad/bf/5a/19e5-441b-9a2b-4d234b4ed3a1.jpeg',1780459943,1780459943,0,0),(187902645048967168,100058,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/ac/b2/a5/c5/f5ef-4eb7-97d6-ce9d105fdd9a.jpeg',1780489081,1780489081,0,0),(187905413142806528,100058,872,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/d2/fb/21/89/4dec-4a7b-bdb1-e4e01a3c55ad.png',1780489741,1780489741,0,0),(187925185259634688,100058,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/4b/6c/a5/9a/8b9f-482f-9a64-7c53ba6bb595.jpeg',1780494455,1780494455,0,0),(187964075806818304,100058,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/db/02/4d/35/24d0-4ecc-bbcf-9b51a3c420a4.jpeg',1780503727,1780503727,0,0),(187964346423312384,100058,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/5a/32/f9/5f/ead5-4ce1-95f6-acf38f273ddd.jpeg',1780503792,1780503792,0,0),(188149094978420736,100058,872,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/avatar/7a/e3/4e/46/934a-4fe4-8aff-3e1968c23b09.png',1780547839,1780547839,0,0),(188190705036820480,100058,84442,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/avatar/42/08/5f/18/7dab-418c-b152-ed52a01b0f80.png',1780557760,1780557760,0,0),(188193530294829056,100058,130548,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/avatar/73/4f/f2/39/1ef4-4c41-a3ba-f315caf59b11.png',1780558434,1780558434,0,0),(188203706523582464,100058,125799,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/avatar/49/09/f5/a1/be7a-4511-853d-ef235a01114e.png',1780560860,1780560860,0,0),(188306185663807488,100058,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/c2/24/fc/56/f1a0-4d1e-be40-5305c062b31a.jpeg',1780585293,1780585293,0,0),(188306654591188992,100058,872,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/31/2d/c1/39/9d82-4158-b385-9fca32a7ef64.png',1780585405,1780585405,0,0),(190131276328468480,190122191415672832,105887,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/avatar/24/db/d7/82/cd87-4319-808e-2a6f12de6b1b.png',1781020428,1781020428,0,0),(192453242808958976,189717769330098176,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/0e/de/b0/14/093d-4dff-8301-68e3841a14ab.jpeg',1781574028,1781574028,0,0),(192453271216979968,189717769330098176,49048,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/d2/0e/d5/d1/7603-4301-a2d1-ac02c4ba0ea9.png',1781574035,1781574035,0,0),(192453945614925824,189717769330098176,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/ff/a2/e9/4a/9e9d-40e3-bd78-6984b090aff3.jpeg',1781574196,1781574196,0,0),(192454671653142528,189717769330098176,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/76/f8/44/b9/0e18-465e-a9ef-18b728c2f5bf.jpeg',1781574369,1781574369,0,0),(192454736383836160,189717769330098176,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/bb/4c/4f/36/b9c0-4a74-8281-a51bceabc9de.jpeg',1781574384,1781574384,0,0),(192454791534739456,189717769330098176,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/98/8c/7a/de/f897-46dd-8be5-a0da91ad9f00.jpeg',1781574397,1781574397,0,0),(192455053661962240,189717769330098176,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/6c/f9/65/78/e86e-4678-bdba-d58cd34ecc5e.jpeg',1781574460,1781574460,0,0),(192572795165605889,192570609316659200,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/d2/14/f1/fb/98c8-48f5-8904-c146d828675a.jpeg',1781602532,1781602532,0,0),(192572912182493184,192570609316659200,235201,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/image/53/f0/05/53/1e5e-4d37-a707-b9f308c9a4ac.jpeg',1781602560,1781602560,0,0),(192575934174330880,192570609316659200,87019,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/avatar/3b/2a/3c/69/0e4e-4970-a2aa-376ce6d41ccc.jpeg',1781603280,1781603280,0,0),(192582846009835520,192570609316659200,79974,0,0,1,'http://192.168.30.44:8008/oss/paopao/public/avatar/63/78/f5/c1/43fc-45de-8fcd-3b2950924ac3.jpeg',1781604928,1781604928,0,0);
/*!40000 ALTER TABLE `p_attachment` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_captcha`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_captcha` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '验证码ID',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `captcha` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '验证码',
  `use_times` int NOT NULL DEFAULT '0' COMMENT '使用次数',
  `expired_on` bigint NOT NULL DEFAULT '0' COMMENT '过期时间',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_captcha_phone` (`phone`) USING BTREE,
  KEY `idx_captcha_expired_on` (`expired_on`) USING BTREE,
  KEY `idx_captcha_use_times` (`use_times`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1021 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='手机验证码';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_captcha` WRITE;
/*!40000 ALTER TABLE `p_captcha` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_captcha` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_chat_conversation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_chat_conversation` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `kol_user_id` bigint NOT NULL DEFAULT '0' COMMENT 'KOL用户ID, 0=默认助手',
  `dify_conversation_id` varchar(64) NOT NULL DEFAULT '' COMMENT 'Dify会话ID',
  `title` varchar(128) DEFAULT '' COMMENT '会话标题',
  `created_on` bigint NOT NULL,
  `modified_on` bigint NOT NULL,
  `deleted_on` bigint NOT NULL DEFAULT '0',
  `is_del` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_kol` (`user_id`,`kol_user_id`),
  KEY `idx_user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户聊天会话映射';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_chat_conversation` WRITE;
/*!40000 ALTER TABLE `p_chat_conversation` DISABLE KEYS */;
INSERT INTO `p_chat_conversation` VALUES (1,189717769330098176,0,'83c68994-1da7-4e51-9cbe-241223e3e053','',1781596405,1781596441,0,0);
/*!40000 ALTER TABLE `p_chat_conversation` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_comment` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `post_id` bigint NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `is_essence` tinyint NOT NULL DEFAULT '0' COMMENT '是否精选',
  `reply_count` int NOT NULL DEFAULT '0' COMMENT '回复数',
  `thumbs_up_count` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  `thumbs_down_count` int NOT NULL DEFAULT '0' COMMENT '点踩数',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_post_id` (`post_id`) USING BTREE,
  KEY `idx_comment_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192625375283511297 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_comment` WRITE;
/*!40000 ALTER TABLE `p_comment` DISABLE KEYS */;
INSERT INTO `p_comment` VALUES (6001736,1080017990,100059,'192.168.30.1','局域网',0,0,0,0,1780459555,1780459555,0,0),(6001737,1080017994,100058,'192.168.30.44','局域网',0,0,0,0,1780460046,1780460046,0,0),(6001738,1080017994,100058,'192.168.30.44','局域网',0,0,0,0,1780460057,1780460057,0,0),(187963967778324480,187962543409463296,100058,'192.168.30.44','局域网',0,0,0,0,1780503702,1780503702,0,0),(188304094555799552,187968258895249408,100058,'192.168.30.44','局域网',0,0,0,0,1780584794,1780584794,0,0),(188310675905314816,188310514168758272,100058,'192.168.30.44','局域网',0,0,0,0,1780586363,1780586363,0,0),(188310720339771392,188308277799419904,100058,'192.168.30.44','局域网',0,0,0,0,1780586374,1780586374,0,0),(190124725194719232,187942239173869568,190122191415672832,'192.168.30.10','局域网',0,0,0,0,1781018866,1781018866,0,0),(192625375283511296,192572795165605888,192570609316659200,'127.0.0.1','本机地址',0,0,0,0,1781615068,1781615084,0,0);
/*!40000 ALTER TABLE `p_comment` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_comment_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_comment_content` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '内容ID',
  `comment_id` bigint NOT NULL DEFAULT '0' COMMENT '评论ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `content` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `type` tinyint NOT NULL DEFAULT '2' COMMENT '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址',
  `sort` bigint NOT NULL DEFAULT '100' COMMENT '排序，越小越靠前',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_content_comment_id` (`comment_id`) USING BTREE,
  KEY `idx_comment_content_user_id` (`user_id`) USING BTREE,
  KEY `idx_comment_content_type` (`type`) USING BTREE,
  KEY `idx_comment_content_sort` (`sort`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192625375296094209 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论内容';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_comment_content` WRITE;
/*!40000 ALTER TABLE `p_comment_content` DISABLE KEYS */;
INSERT INTO `p_comment_content` VALUES (11001738,6001736,100059,'123456',2,100,1780459555,1780459555,0,0),(11001739,6001737,100058,'23',2,100,1780460046,1780460046,0,0),(11001740,6001738,100058,'2r3r',2,100,1780460057,1780460057,0,0),(187963967786713088,187963967778324480,100058,'21212',2,100,1780503702,1780503702,0,0),(188304094572576768,188304094555799552,100058,'22',2,100,1780584794,1780584794,0,0),(188310675922092032,188310675905314816,100058,'23244',2,100,1780586363,1780586363,0,0),(188310720343965696,188310720339771392,100058,'23124',2,100,1780586374,1780586374,0,0),(190124725211496448,190124725194719232,190122191415672832,'具体',2,100,1781018866,1781018866,0,0),(192625375296094208,192625375283511296,192570609316659200,'白来了\n',2,100,1781615068,1781615068,0,0);
/*!40000 ALTER TABLE `p_comment_content` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_comment_metric`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_comment_metric` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comment_id` bigint NOT NULL,
  `rank_score` bigint NOT NULL DEFAULT '0',
  `incentive_score` int NOT NULL DEFAULT '0',
  `decay_factor` int NOT NULL DEFAULT '0',
  `motivation_factor` int NOT NULL DEFAULT '0',
  `is_del` tinyint NOT NULL DEFAULT '0',
  `created_on` bigint NOT NULL DEFAULT '0',
  `modified_on` bigint NOT NULL DEFAULT '0',
  `deleted_on` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_metric_comment_id_rank_score` (`comment_id`,`rank_score`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192625375321260033 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_comment_metric` WRITE;
/*!40000 ALTER TABLE `p_comment_metric` DISABLE KEYS */;
INSERT INTO `p_comment_metric` VALUES (1,6001736,0,0,0,0,0,1780459555,1780459555,0),(2,6001737,0,0,0,0,0,1780460046,1780460046,0),(3,6001738,0,0,0,0,0,1780460057,1780460057,0),(187963967811878912,187963967778324480,0,0,0,0,0,1780503702,1780503702,0),(188304094597742593,188304094555799552,0,0,0,0,0,1780584794,1780584794,0),(188310675943063552,188310675905314816,0,0,0,0,0,1780586363,1780586363,0),(188310720369131520,188310720339771392,0,0,0,0,0,1780586374,1780586374,0),(190124725232467968,190124725194719232,0,0,0,0,0,1781018866,1781018866,0),(192625375321260032,192625375283511296,0,0,0,0,0,1781615068,1781615084,0);
/*!40000 ALTER TABLE `p_comment_metric` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_comment_reply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_comment_reply` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '回复ID',
  `comment_id` bigint NOT NULL DEFAULT '0' COMMENT '评论ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `at_user_id` bigint NOT NULL DEFAULT '0' COMMENT '@用户ID',
  `content` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `thumbs_up_count` int NOT NULL DEFAULT '0' COMMENT '点赞数',
  `thumbs_down_count` int NOT NULL DEFAULT '0' COMMENT '点踩数',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment_reply_comment_id` (`comment_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12000015 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论回复';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_comment_reply` WRITE;
/*!40000 ALTER TABLE `p_comment_reply` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_comment_reply` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_contact`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_contact` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `friend_id` bigint NOT NULL COMMENT '好友ID',
  `group_id` bigint NOT NULL DEFAULT '0' COMMENT '好友分组ID:默认为0无分组',
  `remark` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '好友备注',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '好友状态: 1请求好友, 2已好友, 3拒绝好友, 4已删好友',
  `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶, 0否, 1是',
  `is_black` tinyint NOT NULL DEFAULT '0' COMMENT '是否为黑名单, 0否, 1是',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除好友, 0否, 1是',
  `notice_enable` tinyint NOT NULL DEFAULT '0' COMMENT '是否有消息提醒, 0否, 1是',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_contact_user_friend` (`user_id`,`friend_id`) USING BTREE,
  KEY `idx_contact_user_friend_status` (`user_id`,`friend_id`,`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=190132703385878529 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联系人';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_contact` WRITE;
/*!40000 ALTER TABLE `p_contact` DISABLE KEYS */;
INSERT INTO `p_contact` VALUES (1,100058,100059,0,'',2,0,0,0,0,1780459603,1780459603,0),(2,100059,100058,0,'',2,0,0,0,0,1780459655,1780459655,0),(187968546007941120,187968194357493760,100058,0,'',2,0,0,0,0,1780504793,1780505020,0),(187969437243015168,100058,187968194357493760,0,'',2,0,0,0,0,1780505006,1780505006,0),(187973580653330432,187973080667127808,187968194357493760,0,'',2,0,0,0,0,1780505994,1780505994,0),(187973630192254976,187968194357493760,187973080667127808,0,'',2,0,0,0,0,1780506005,1780506005,0),(190113595139293184,190112139342512128,189717769330098176,0,'',2,0,0,0,0,1781016213,1781016213,0),(190115540176470016,189717769330098176,190112139342512128,0,'',2,0,0,0,0,1781016676,1781016676,0),(190132272173678592,190112139342512128,190122191415672832,0,'',2,0,0,0,0,1781020666,1781020666,0),(190132703385878528,190122191415672832,190112139342512128,0,'',2,0,0,0,0,1781020768,1781020768,0);
/*!40000 ALTER TABLE `p_contact` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_contact_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_contact_group` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `name` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分组名称',
  `is_del` tinyint NOT NULL DEFAULT '1' COMMENT '是否删除, 0否, 1是',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联系人分组';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_contact_group` WRITE;
/*!40000 ALTER TABLE `p_contact_group` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_contact_group` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_following`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_following` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `follow_id` bigint NOT NULL,
  `is_del` tinyint NOT NULL DEFAULT '0',
  `created_on` bigint NOT NULL DEFAULT '0',
  `modified_on` bigint NOT NULL DEFAULT '0',
  `deleted_on` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_following_user_follow` (`user_id`,`follow_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192548695902781441 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_following` WRITE;
/*!40000 ALTER TABLE `p_following` DISABLE KEYS */;
INSERT INTO `p_following` VALUES (1,100059,100058,0,1780459535,1780459535,0),(2,100058,100059,0,1780459720,1780459720,0),(187968567470194688,187968194357493760,100058,0,1780504798,1780504798,0),(187969058988097536,100058,187968194357493760,0,1780504915,1780504915,0),(192511443596214272,189717769330098176,190119095696359424,0,1781587904,1781587904,0),(192511491168010240,189717769330098176,190122191415672832,0,1781587916,1781587916,0),(192512169978363904,189717769330098176,190319267848126464,0,1781588078,1781588078,0),(192548695902781440,189717769330098176,100058,0,1781596786,1781596786,0);
/*!40000 ALTER TABLE `p_following` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_kol_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_kol_category` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='KOL分类';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_kol_category` WRITE;
/*!40000 ALTER TABLE `p_kol_category` DISABLE KEYS */;
INSERT INTO `p_kol_category` VALUES (1,'初恋学生',1,1781001811,0,0,0),(2,'都市御姐',2,1781001811,0,0,0),(3,'温柔少妇',3,1781001811,0,0,0),(4,'运动少女',4,1781001811,0,0,0),(5,'暗黑病娇',5,1781001811,0,0,0),(6,'慵懒纯欲',6,1781001811,0,0,0),(7,'精灵奇幻',7,1781001811,0,0,0),(8,'街头辣妹',8,1781001811,0,0,0),(9,'文艺知性',9,1781001811,0,0,0),(10,'邻家治愈',10,1781001811,0,0,0);
/*!40000 ALTER TABLE `p_kol_category` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_kol_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_kol_profile` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `category_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'KOL分类ID',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序(越大越前)',
  `system_prompt` text COLLATE utf8mb4_general_ci COMMENT '系统提示词',
  `api_key` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'Dify API私钥',
  `height` varchar(16) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '160cm' COMMENT '身高',
  `weight` varchar(16) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '44kg' COMMENT '体重',
  `measurements` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '84/58/84' COMMENT '三围',
  `skin_tone` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '冷白病态肌' COMMENT '肤色',
  `eye_color` varchar(16) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '酒红' COMMENT '瞳色',
  `orientation` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '偏双性恋（情感依赖向）' COMMENT '性向',
  `preferences` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '独占欲、暗调氛围、偏执温柔' COMMENT '喜好',
  `favorite_foods` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '黑森林、红酒、冷食' COMMENT '喜欢食物',
  `clothing_style` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '黑裙、蕾丝、丝带、哥特风' COMMENT '服装风格',
  `makeup_style` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '苍白底妆、下垂眼、暗红眼影、冷唇' COMMENT '妆风格',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0为未删除、1为已删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='KOL人物属性';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_kol_profile` WRITE;
/*!40000 ALTER TABLE `p_kol_profile` DISABLE KEYS */;
INSERT INTO `p_kol_profile` VALUES (1,189718285166575616,1,3,'','','160cm','44kg','84/58/84','冷白病态肌','酒红','偏双性恋（情感依赖向）','独占欲、暗调氛围、偏执温柔','黑森林、红酒、冷食','黑裙、蕾丝、丝带、哥特风','苍白底妆、下垂眼、暗红眼影、冷唇',1780999120,1781613970,0,0),(2,189717769330098176,1,2,'','','160cm','44kg','84/58/84','冷白病态肌','酒红','偏双性恋（情感依赖向）','独占欲、暗调氛围、偏执温柔','黑森林、红酒、冷食','黑裙、蕾丝、丝带、哥特风','苍白底妆、下垂眼、暗红眼影、冷唇',1781003022,1781613739,0,0),(3,100059,1,0,NULL,'','160cm','44kg','84/58/84','冷白病态肌','酒红','偏双性恋（情感依赖向）','独占欲、暗调氛围、偏执温柔','黑森林、红酒、冷食','黑裙、蕾丝、丝带、哥特风','苍白底妆、下垂眼、暗红眼影、冷唇',1781003126,1781003126,0,0),(4,190338733612990464,4,0,'地方试点改革22','app-hwL3g2syuknNoj4SjBnsEhbz','160cm','44kg','84/58/84','冷白病态肌','酒红','偏双性恋（情感依赖向）','独占欲、暗调氛围、偏执温柔','黑森林、红酒、冷食','黑裙、蕾丝、丝带、哥特风','苍白底妆、下垂眼、暗红眼影、冷唇',1781585746,1781612151,0,0),(5,192570609316659200,1,4,'','','160cm','44kg','84/58/84','冷白病态肌','酒红','偏双性恋（情感依赖向）','独占欲、暗调氛围、偏执温柔','黑森林、红酒、冷食','黑裙、蕾丝、丝带、哥特风','苍白底妆、下垂眼、暗红眼影、冷唇',1781607709,1781613723,0,0);
/*!40000 ALTER TABLE `p_kol_profile` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_message` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '消息通知ID',
  `sender_user_id` bigint NOT NULL DEFAULT '0' COMMENT '发送方用户ID',
  `receiver_user_id` bigint NOT NULL DEFAULT '0' COMMENT '接收方用户ID',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '通知类型，1动态，2评论，3回复，4私信，99系统通知',
  `brief` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '摘要说明',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细内容',
  `post_id` bigint NOT NULL DEFAULT '0' COMMENT '动态ID',
  `comment_id` bigint NOT NULL DEFAULT '0' COMMENT '评论ID',
  `reply_id` bigint NOT NULL DEFAULT '0' COMMENT '回复ID',
  `is_read` tinyint NOT NULL DEFAULT '0' COMMENT '是否已读',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_message_receiver_user_id` (`receiver_user_id`) USING BTREE,
  KEY `idx_message_is_read` (`is_read`) USING BTREE,
  KEY `idx_message_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=190132272173678594 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='消息通知';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_message` WRITE;
/*!40000 ALTER TABLE `p_message` DISABLE KEYS */;
INSERT INTO `p_message` VALUES (16000033,100059,100058,2,'在泡泡中评论了你','',1080017990,6001736,0,1,1780459555,1780459555,0,0),(16000034,100058,100059,5,'请求添加好友，并附言:','22',0,0,2,0,1780459603,1780459603,0,0),(16000035,100058,100059,4,'给你发送新私信了','223',0,0,0,0,1780459707,1780459707,0,0),(16000036,100058,100059,2,'在泡泡中评论了你','',1080017994,6001737,0,0,1780460046,1780460046,0,0),(16000037,100058,100059,2,'在泡泡中评论了你','',1080017994,6001738,0,0,1780460057,1780460057,0,0),(187968546007941121,187968194357493760,100058,5,'请求添加好友，并附言:','11',0,0,2,1,1780504793,1780504793,0,0),(187969437247209472,100058,187968194357493760,5,'请求添加好友，并附言:','1232',0,0,1,1,1780505006,1780505006,0,0),(187973580653330433,187973080667127808,187968194357493760,5,'请求添加好友，并附言:','232',0,0,2,1,1780505994,1780505994,0,0),(188304094597742592,100058,187968194357493760,2,'在泡泡中评论了你','',187968258895249408,188304094555799552,0,0,1780584794,1780584794,0,0),(188304094597742594,0,0,99,'2e',' 阿道夫士大夫',0,0,0,1,1781005584,0,0,0),(190113298048352256,190112139342512128,189717769330098176,4,'给你发送新私信了','123',0,0,0,1,1781016142,1781016142,0,0),(190113419175657472,190112139342512128,189717769330098176,4,'给你发送新私信了','12323',0,0,0,1,1781016171,1781016171,0,0),(190113595139293185,190112139342512128,189717769330098176,5,'请求添加好友，并附言:','122432',0,0,2,1,1781016213,1781016213,0,0),(190124725228273664,190122191415672832,100058,2,'在泡泡中评论了你','',187942239173869568,190124725194719232,0,0,1781018866,1781018866,0,0),(190132081462870016,190112139342512128,190122191415672832,4,'给你发送新私信了','12323',0,0,0,1,1781020620,1781020620,0,0),(190132272173678593,190112139342512128,190122191415672832,5,'请求添加好友，并附言:','77',0,0,2,1,1781020666,1781020666,0,0);
/*!40000 ALTER TABLE `p_message` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_notice` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '消息通知ID',
  `sender_user_id` bigint NOT NULL DEFAULT '0' COMMENT '发送方用户ID',
  `receiver_user_id` bigint NOT NULL DEFAULT '0' COMMENT '接收方用户ID',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '通知类型，1动态，2评论，3回复，4私信，99系统通知',
  `brief` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '摘要说明',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细内容',
  `post_id` bigint NOT NULL DEFAULT '0' COMMENT '动态ID',
  `comment_id` bigint NOT NULL DEFAULT '0' COMMENT '评论ID',
  `reply_id` bigint NOT NULL DEFAULT '0' COMMENT '回复ID',
  `is_read` tinyint NOT NULL DEFAULT '0' COMMENT '是否已读',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_message_receiver_user_id` (`receiver_user_id`) USING BTREE,
  KEY `idx_message_is_read` (`is_read`) USING BTREE,
  KEY `idx_message_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=188304094597742598 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统通知';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_notice` WRITE;
/*!40000 ALTER TABLE `p_notice` DISABLE KEYS */;
INSERT INTO `p_notice` VALUES (188304094597742593,0,0,99,'122','22',0,0,0,0,1781003691,0,0,0),(188304094597742594,0,0,1,'士大夫','似懂非懂',0,0,0,0,1781015924,0,0,0),(188304094597742595,0,190122191415672830,1,'22','反反复复方法',0,0,0,0,1781061319,0,0,0),(188304094597742596,0,0,1,'d f','发的是对方是个 ',0,0,0,0,1781064758,0,0,0),(188304094597742597,0,190122191415672832,1,'的发',' 上的撒旦发  ',0,0,0,0,1781064769,0,0,0);
/*!40000 ALTER TABLE `p_notice` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_notice_read`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_notice_read` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `msg_id` bigint NOT NULL COMMENT '关联p_notice.id',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `read_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '阅读时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_msg_user` (`msg_id`,`user_id`),
  KEY `idx_user` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统消息阅读记录';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_notice_read` WRITE;
/*!40000 ALTER TABLE `p_notice_read` DISABLE KEYS */;
INSERT INTO `p_notice_read` VALUES (1,188304094597742596,189717769330098176,'2026-06-16 08:53:39'),(2,188304094597742594,189717769330098176,'2026-06-16 08:53:40'),(3,188304094597742593,189717769330098176,'2026-06-16 08:53:42'),(4,188304094597742596,192570609316659200,'2026-06-16 17:26:59'),(5,188304094597742594,192570609316659200,'2026-06-16 17:27:03'),(6,188304094597742593,192570609316659200,'2026-06-16 17:27:03'),(7,188304094597742596,190112139342512128,'2026-06-17 10:48:59'),(9,188304094597742594,190112139342512128,'2026-06-17 10:49:01'),(10,188304094597742593,190112139342512128,'2026-06-17 10:49:03');
/*!40000 ALTER TABLE `p_notice_read` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_post` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主题ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `comment_count` bigint NOT NULL DEFAULT '0' COMMENT '评论数',
  `collection_count` bigint NOT NULL DEFAULT '0' COMMENT '收藏数',
  `upvote_count` bigint NOT NULL DEFAULT '0' COMMENT '点赞数',
  `share_count` bigint NOT NULL DEFAULT '0' COMMENT '分享数',
  `visibility` tinyint NOT NULL DEFAULT '0' COMMENT '可见性: 0私密 10充电可见 20订阅可见 30保留 40保留 50好友可见 60关注可见 70保留 80保留 90公开',
  `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶',
  `is_essence` tinyint NOT NULL DEFAULT '0' COMMENT '是否精华',
  `is_lock` tinyint NOT NULL DEFAULT '0' COMMENT '是否锁定',
  `latest_replied_on` bigint NOT NULL DEFAULT '0' COMMENT '最新回复时间',
  `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签',
  `attachment_price` bigint NOT NULL DEFAULT '0' COMMENT '附件价格(分)',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_user_id` (`user_id`) USING BTREE,
  KEY `idx_post_visibility` (`visibility`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192572928192151553 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_post` WRITE;
/*!40000 ALTER TABLE `p_post` DISABLE KEYS */;
INSERT INTO `p_post` VALUES (1080017989,100058,0,0,0,0,50,0,0,0,1780455529,'',0,'127.0.0.1','本机地址',1780455529,1780455529,0,0),(1080017990,100058,1,1,1,0,90,0,0,0,1780459555,'',0,'127.0.0.1','本机地址',1780455561,1780455561,0,0),(1080017991,100059,0,0,0,0,50,0,0,0,1780457417,'',0,'192.168.30.1','局域网',1780457417,1780457417,0,0),(1080017992,100058,0,0,0,0,90,0,0,0,1780459384,'CEO',0,'192.168.30.44','局域网',1780459384,1780459384,0,0),(1080017993,100058,0,1,0,0,90,0,0,0,1780459446,'CEO',0,'192.168.30.44','局域网',1780459446,1780459446,0,0),(1080017994,100059,2,0,2,0,50,0,0,0,1780460057,'',0,'192.168.30.1','局域网',1780459953,1780459953,0,0),(187902656608468992,100058,0,0,0,0,90,0,0,0,1780489084,'',0,'192.168.30.44','局域网',1780489084,1780489084,0,0),(187903825783291904,187901818129350656,0,0,0,0,0,0,0,0,1780489363,'test',0,'127.0.0.1','本机地址',1780489363,1780489363,0,0),(187905447682899968,100058,0,0,0,0,90,0,0,0,1780489749,'',0,'192.168.30.44','局域网',1780489749,1780489749,0,0),(187912355961634816,187901818129350656,0,0,0,0,90,0,0,0,1780491396,'x',0,'127.0.0.1','本机地址',1780491396,1780491396,0,0),(187912693439528960,100058,0,0,0,0,50,0,0,0,1780491477,'',0,'192.168.30.44','局域网',1780491477,1780491477,0,0),(187916958144921600,100058,0,0,0,0,50,0,0,0,1780492494,'',0,'192.168.30.44','局域网',1780492494,1780492494,0,0),(187917096942829568,100058,0,0,0,0,50,0,0,0,1780492527,'',0,'192.168.30.44','局域网',1780492527,1780492527,0,0),(187917149526818816,100058,0,0,0,0,90,0,0,0,1780492539,'',0,'192.168.30.44','局域网',1780492539,1780492539,0,0),(187917786427686912,187901818129350656,0,0,0,0,90,0,0,0,1780492691,'t',0,'127.0.0.1','本机地址',1780492691,1780492691,0,0),(187919141263376384,187901818129350656,0,0,0,0,90,0,0,0,1780493014,'x',0,'127.0.0.1','本机地址',1780493014,1780493014,0,0),(187919527143538688,187901818129350656,0,0,0,0,90,0,0,0,1780493106,'x',0,'127.0.0.1','本机地址',1780493106,1780493106,0,0),(187919577437437952,187901818129350656,0,0,0,0,90,0,0,0,1780493118,'x',0,'127.0.0.1','本机地址',1780493118,1780493118,0,0),(187921515851808768,100058,0,0,0,0,90,0,0,0,1780493580,'',0,'192.168.30.44','局域网',1780493580,1780493580,0,0),(187922349511671808,100058,0,0,0,0,50,0,0,0,1780493779,'',0,'192.168.30.44','局域网',1780493779,1780493779,0,0),(187922402921938944,100058,0,0,0,0,90,0,0,0,1780493792,'',0,'192.168.30.44','局域网',1780493792,1780493792,0,0),(187923490404302848,100058,0,0,0,0,50,0,0,0,1780494051,'',0,'192.168.30.44','局域网',1780494051,1780494051,0,0),(187923537984487424,100058,0,0,0,0,90,0,0,0,1780494062,'',0,'192.168.30.44','局域网',1780494062,1780494062,0,0),(187924409904791552,187901818129350656,0,0,0,0,90,0,0,0,1780494270,'x',0,'127.0.0.1','本机地址',1780494270,1780494270,0,0),(187925033207726080,100058,0,0,0,0,90,0,0,0,1780494419,'',0,'192.168.30.44','局域网',1780494419,1780494419,0,0),(187925186769584128,100058,0,0,0,0,90,0,0,0,1780494456,'',0,'192.168.30.44','局域网',1780494456,1780494456,0,0),(187940231306018816,100058,0,0,0,0,90,0,0,0,1780498042,'OKX',0,'192.168.30.44','局域网',1780498042,1780498042,0,0),(187942239173869568,100058,1,0,1,0,90,0,0,0,1781018866,'六一',0,'192.168.30.44','局域网',1780498521,1780498521,0,0),(187962543409463296,100058,1,0,0,0,90,0,0,0,1780503702,'',0,'192.168.30.44','局域网',1780503362,1780503362,0,0),(187964129628127232,100058,0,0,1,0,90,0,0,0,1780503740,'',0,'192.168.30.44','局域网',1780503740,1780503740,0,0),(187964381126983680,100058,0,0,0,0,90,0,0,0,1780503800,'',0,'192.168.30.44','局域网',1780503800,1780503800,0,0),(187968258895249408,187968194357493760,1,0,0,0,90,0,0,0,1780584794,'',0,'192.168.30.44','局域网',1780504725,1780504725,0,0),(188306214596116480,100058,0,0,0,0,90,0,0,0,1780585300,'',0,'192.168.30.44','局域网',1780585300,1780585300,0,0),(188306298826129408,100058,0,0,0,0,90,0,0,0,1780585320,'',0,'192.168.30.44','局域网',1780585320,1780585320,0,0),(188306668067487744,100058,0,0,0,0,90,0,0,0,1780585408,'',0,'192.168.30.44','局域网',1780585408,1780585408,0,0),(188307666135678976,100058,0,0,1,0,90,0,0,0,1780585646,'',0,'192.168.30.44','局域网',1780585646,1780585646,0,0),(188308277799419904,100058,1,0,0,0,90,0,0,0,1780586374,'',0,'192.168.30.44','局域网',1780585792,1780585792,0,0),(188310514168758272,100058,1,0,1,0,90,0,0,0,1780586363,'',0,'192.168.30.44','局域网',1780586325,1780586325,0,0),(188310555746893824,100058,0,0,0,0,90,0,0,0,1780586335,'',0,'192.168.30.44','局域网',1780586335,1780586335,0,0),(189718039749459968,189717769330098176,0,0,0,0,90,0,0,0,1780921905,'',0,'192.168.30.44','局域网',1780921905,1780921905,0,0),(190119319424729088,190119095696359424,0,0,0,0,90,0,0,0,1781017577,'',0,'192.168.30.10','局域网',1781017577,1781017577,0,0),(190131339897339904,190122191415672832,0,0,0,0,90,0,0,0,1781020443,'',0,'192.168.30.10','局域网',1781020443,1781020443,0,0),(190319463197835264,190319267848126464,0,0,1,0,90,0,0,0,1781065295,'',0,'192.168.30.10','局域网',1781065295,1781065295,0,0),(192453341018587136,189717769330098176,0,0,0,0,50,0,0,0,1781574052,'',0,'127.0.0.1','本机地址',1781574052,1781574052,0,0),(192455707159691264,189717769330098176,0,0,0,0,50,0,0,0,1781574616,'',0,'127.0.0.1','本机地址',1781574616,1781574616,0,0),(192460204531515392,189717769330098176,0,1,1,0,90,0,0,0,1781575688,'',0,'127.0.0.1','本机地址',1781575688,1781575688,0,0),(192572795165605888,192570609316659200,1,0,1,0,90,0,0,0,1781615068,'',0,'127.0.0.1','本机地址',1781602532,1781602532,0,0),(192572928192151552,192570609316659200,0,0,1,0,90,0,0,0,1781602563,'',0,'127.0.0.1','本机地址',1781602563,1781602563,0,0);
/*!40000 ALTER TABLE `p_post` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_post_attachment_bill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_post_attachment_bill` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '购买记录ID',
  `post_id` bigint NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `paid_amount` bigint NOT NULL DEFAULT '0' COMMENT '支付金额',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_attachment_bill_post_id` (`post_id`) USING BTREE,
  KEY `idx_post_attachment_bill_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5000002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章附件账单';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_post_attachment_bill` WRITE;
/*!40000 ALTER TABLE `p_post_attachment_bill` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_post_attachment_bill` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_post_by_comment`;
/*!50001 DROP VIEW IF EXISTS `p_post_by_comment`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `p_post_by_comment` AS SELECT 
 1 AS `id`,
 1 AS `user_id`,
 1 AS `comment_count`,
 1 AS `collection_count`,
 1 AS `upvote_count`,
 1 AS `share_count`,
 1 AS `visibility`,
 1 AS `is_top`,
 1 AS `is_essence`,
 1 AS `is_lock`,
 1 AS `latest_replied_on`,
 1 AS `tags`,
 1 AS `attachment_price`,
 1 AS `ip`,
 1 AS `ip_loc`,
 1 AS `created_on`,
 1 AS `modified_on`,
 1 AS `deleted_on`,
 1 AS `is_del`,
 1 AS `comment_user_id`*/;
SET character_set_client = @saved_cs_client;
DROP TABLE IF EXISTS `p_post_by_media`;
/*!50001 DROP VIEW IF EXISTS `p_post_by_media`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `p_post_by_media` AS SELECT 
 1 AS `id`,
 1 AS `user_id`,
 1 AS `comment_count`,
 1 AS `collection_count`,
 1 AS `upvote_count`,
 1 AS `share_count`,
 1 AS `visibility`,
 1 AS `is_top`,
 1 AS `is_essence`,
 1 AS `is_lock`,
 1 AS `latest_replied_on`,
 1 AS `tags`,
 1 AS `attachment_price`,
 1 AS `ip`,
 1 AS `ip_loc`,
 1 AS `created_on`,
 1 AS `modified_on`,
 1 AS `deleted_on`,
 1 AS `is_del`*/;
SET character_set_client = @saved_cs_client;
DROP TABLE IF EXISTS `p_post_collection`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_post_collection` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
  `post_id` bigint NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_collection_post_id` (`post_id`) USING BTREE,
  KEY `idx_post_collection_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=193386889506455553 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章收藏';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_post_collection` WRITE;
/*!40000 ALTER TABLE `p_post_collection` DISABLE KEYS */;
INSERT INTO `p_post_collection` VALUES (6000012,1080017990,100059,1780459545,1780459547,1780459547,1),(6000013,1080017990,100058,1780459751,1780459756,1780459756,1),(6000014,1080017993,100058,1780459754,1780459754,0,0),(6000015,1080017990,100058,1780459757,1780459757,0,0),(187963976900935680,187962543409463296,100058,1780503704,1780503705,1780503705,1),(188302314296049664,187968258895249408,100058,1780584370,1780584371,1780584371,1),(193386889506455552,192460204531515392,189717769330098176,1781796627,1781796627,0,0);
/*!40000 ALTER TABLE `p_post_collection` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_post_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_post_content` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '内容ID',
  `post_id` bigint NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `content` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `type` tinyint NOT NULL DEFAULT '2' COMMENT '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源',
  `sort` int NOT NULL DEFAULT '100' COMMENT '排序，越小越靠前',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_content_post_id` (`post_id`) USING BTREE,
  KEY `idx_post_content_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192572928221511681 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章内容';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_post_content` WRITE;
/*!40000 ALTER TABLE `p_post_content` DISABLE KEYS */;
INSERT INTO `p_post_content` VALUES (180022546,1080017989,100058,'erer',2,100,1780455529,1780455529,0,0),(180022547,1080017990,100058,'qwqwaf',2,100,1780455561,1780455561,0,0),(180022548,1080017990,100058,'http://192.168.30.44:8008/oss/paopao/public/image/e7/b0/e3/a4/b80a-436a-8e21-02dc800c76ec.jpeg',3,101,1780455561,1780455561,0,0),(180022549,1080017991,100059,'今天好热',2,100,1780457417,1780457417,0,0),(180022550,1080017991,100059,'http://192.168.30.44:8008/oss/paopao/public/image/85/5a/f5/a2/84a6-49f1-bf44-2812df724f33.jpeg',3,101,1780457417,1780457417,0,0),(180022551,1080017992,100058,'#CEO dk',2,100,1780459384,1780459384,0,0),(180022552,1080017993,100058,'#CEO dfs',2,100,1780459446,1780459446,0,0),(180022553,1080017994,100059,'123456',2,100,1780459953,1780459953,0,0),(180022554,1080017994,100059,'http://192.168.30.44:8008/oss/paopao/public/image/f0/73/35/c2/2845-4b7d-8827-098eb1f49749.jpeg',3,101,1780459953,1780459953,0,0),(180022555,1080017994,100059,'http://192.168.30.44:8008/oss/paopao/public/image/66/cf/96/3e/483e-4721-8a80-0cbccbf819d3.jpeg',3,102,1780459953,1780459953,0,0),(180022556,1080017994,100059,'http://192.168.30.44:8008/oss/paopao/public/image/8c/d2/f4/9f/a44a-4e97-9af2-6679ac6cb8c0.jpeg',3,103,1780459953,1780459953,0,0),(180022557,1080017994,100059,'http://192.168.30.44:8008/oss/paopao/public/image/23/ad/bf/5a/19e5-441b-9a2b-4d234b4ed3a1.jpeg',3,104,1780459953,1780459953,0,0),(187902656629440512,187902656608468992,100058,'2ef',2,100,1780489084,1780489084,0,0),(187902656637829120,187902656608468992,100058,'http://192.168.30.44:8008/oss/paopao/public/image/ac/b2/a5/c5/f5ef-4eb7-97d6-ce9d105fdd9a.jpeg',3,101,1780489084,1780489084,0,0),(187903825800069120,187903825783291904,187901818129350656,'',1,0,1780489363,1780489363,0,0),(187905447695482880,187905447682899968,100058,'sdfsd',2,100,1780489749,1780489749,0,0),(187905447699677184,187905447682899968,100058,'http://192.168.30.44:8008/oss/paopao/public/image/d2/fb/21/89/4dec-4a7b-bdb1-e4e01a3c55ad.png',3,101,1780489749,1780489749,0,0),(187912355986800640,187912355961634816,187901818129350656,'',1,0,1780491396,1780491396,0,0),(187912693452111872,187912693439528960,100058,'232e',2,100,1780491477,1780491477,0,0),(187916958165893120,187916958144921600,100058,'sdf',2,100,1780492494,1780492494,0,0),(187917096955412480,187917096942829568,100058,'dfsdgasgdsf',2,100,1780492527,1780492527,0,0),(187917149543596032,187917149526818816,100058,'ssdfsdgsdf\n1111111111111',2,100,1780492539,1780492539,0,0),(187917786448658432,187917786427686912,187901818129350656,'',1,0,1780492691,1780492691,0,0),(187919141284347904,187919141263376384,187901818129350656,'',1,0,1780493014,1780493014,0,0),(187919527164510208,187919527143538688,187901818129350656,'',1,0,1780493106,1780493106,0,0),(187919577458409472,187919577437437952,187901818129350656,'',1,0,1780493118,1780493118,0,0),(187921515868585984,187921515851808768,100058,'qwqeqd',2,100,1780493580,1780493580,0,0),(187922349541031936,187922349511671808,100058,'222223242313113',2,100,1780493779,1780493779,0,0),(187922402938716160,187922402921938944,100058,'qwrwewq11111111111111111',2,100,1780493792,1780493792,0,0),(187923490416885760,187923490404302848,100058,'2131413213',2,100,1780494051,1780494051,0,0),(187923538005458944,187923537984487424,100058,'212222',2,100,1780494062,1780494062,0,0),(187924409925763072,187924409904791552,187901818129350656,'',1,0,1780494270,1780494270,0,0),(187925033224503296,187925033207726080,100058,'232323',2,100,1780494419,1780494419,0,0),(187925186790555648,187925186769584128,100058,'sdfdf',2,100,1780494456,1780494456,0,0),(187925186794749952,187925186769584128,100058,'http://192.168.30.44:8008/oss/paopao/public/image/4b/6c/a5/9a/8b9f-482f-9a64-7c53ba6bb595.jpeg',3,101,1780494456,1780494456,0,0),(187940231326990336,187940231306018816,100058,'#OKX 完美\n111',2,100,1780498042,1780498042,0,0),(187942239199035392,187942239173869568,100058,'#六一 在“六一”国际儿童节之际，习近平总书记给中共一大纪念馆、南湖革命纪念馆少先队红领巾讲解员回信，对他们予以亲切勉励，“希望你们高举队旗跟党走，传承红色基因，增长知识本领，磨练意志品质，做党和人民的红孩子，在新征程上跑好历史接力赛。”‌‌',2,100,1780498521,1780498521,0,0),(187962543426240512,187962543409463296,100058,'士大夫 六一 水电工 ',2,100,1780503362,1780503362,0,0),(187964129640710144,187964129628127232,100058,'1242321',2,100,1780503740,1780503740,0,0),(187964129649098752,187964129628127232,100058,'http://192.168.30.44:8008/oss/paopao/public/image/db/02/4d/35/24d0-4ecc-bbcf-9b51a3c420a4.jpeg',3,101,1780503740,1780503740,0,0),(187964381143760896,187964381126983680,100058,'23423523',2,100,1780503800,1780503800,0,0),(187964381152149504,187964381126983680,100058,'http://192.168.30.44:8008/oss/paopao/public/image/5a/32/f9/5f/ead5-4ce1-95f6-acf38f273ddd.jpeg',3,101,1780503800,1780503800,0,0),(187968258916220928,187968258895249408,187968194357493760,'士大夫 ',2,100,1780504725,1780504725,0,0),(188306214617088000,188306214596116480,100058,'132132',2,100,1780585300,1780585300,0,0),(188306214625476608,188306214596116480,100058,'http://192.168.30.44:8008/oss/paopao/public/image/c2/24/fc/56/f1a0-4d1e-be40-5305c062b31a.jpeg',3,101,1780585300,1780585300,0,0),(188306298847100928,188306298826129408,100058,'132132',2,100,1780585320,1780585320,0,0),(188306298851295232,188306298826129408,100058,'http://192.168.30.44:8008/oss/paopao/public/image/c2/24/fc/56/f1a0-4d1e-be40-5305c062b31a.jpeg',3,101,1780585320,1780585320,0,0),(188306668084264960,188306668067487744,100058,'说的方法 ',2,100,1780585408,1780585408,0,0),(188306668088459264,188306668067487744,100058,'http://192.168.30.44:8008/oss/paopao/public/image/31/2d/c1/39/9d82-4158-b385-9fca32a7ef64.png',3,101,1780585408,1780585408,0,0),(188307666152456192,188307666135678976,100058,'似懂非懂',2,100,1780585646,1780585646,0,0),(188308277816197120,188308277799419904,100058,'23124',2,100,1780585792,1780585792,0,0),(188310514189729792,188310514168758272,100058,'2424',2,100,1780586325,1780586325,0,0),(188310555767865344,188310555746893824,100058,'千万千万人',2,100,1780586335,1780586335,0,0),(189718039774625792,189718039749459968,189717769330098176,'轻微而士大夫 ',2,100,1780921905,1780921905,0,0),(190119319441506304,190119319424729088,190119095696359424,'你们',2,100,1781017578,1781017578,0,0),(190131339914117120,190131339897339904,190122191415672832,'策略',2,100,1781020443,1781020443,0,0),(190319463223001088,190319463197835264,190319267848126464,'出来',2,100,1781065295,1781065295,0,0),(192453341035364352,192453341018587136,189717769330098176,'112',2,100,1781574052,1781574052,0,0),(192453341043752960,192453341018587136,189717769330098176,'http://192.168.30.44:8008/oss/paopao/public/image/0e/de/b0/14/093d-4dff-8301-68e3841a14ab.jpeg',3,101,1781574052,1781574052,0,0),(192453341047947264,192453341018587136,189717769330098176,'http://192.168.30.44:8008/oss/paopao/public/image/d2/0e/d5/d1/7603-4301-a2d1-ac02c4ba0ea9.png',3,102,1781574052,1781574052,0,0),(192455707197440000,192455707159691264,189717769330098176,'2324',2,100,1781574616,1781574616,0,0),(192455707205828608,192455707159691264,189717769330098176,'http://192.168.30.44:8008/oss/paopao/public/image/6c/f9/65/78/e86e-4678-bdba-d58cd34ecc5e.jpeg',3,101,1781574616,1781574616,0,0),(192455707210022912,192455707159691264,189717769330098176,'https://www.shenfendaquan.com/#google_vignette',6,102,1781574616,1781574616,0,0),(192460204560875520,192460204531515392,189717769330098176,'通天塔',2,100,1781575688,1781575688,0,0),(192572795186577408,192572795165605888,192570609316659200,'222',2,100,1781602532,1781602532,0,0),(192572928217317376,192572928192151552,192570609316659200,'3232',2,100,1781602563,1781602563,0,0),(192572928221511680,192572928192151552,192570609316659200,'http://192.168.30.44:8008/oss/paopao/public/image/53/f0/05/53/1e5e-4d37-a707-b9f308c9a4ac.jpeg',3,101,1781602563,1781602563,0,0);
/*!40000 ALTER TABLE `p_post_content` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_post_metric`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_post_metric` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint NOT NULL,
  `rank_score` bigint NOT NULL DEFAULT '0',
  `incentive_score` int NOT NULL DEFAULT '0',
  `decay_factor` int NOT NULL DEFAULT '0',
  `motivation_factor` int NOT NULL DEFAULT '0',
  `is_del` tinyint NOT NULL DEFAULT '0',
  `created_on` bigint NOT NULL DEFAULT '0',
  `modified_on` bigint NOT NULL DEFAULT '0',
  `deleted_on` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_metric_post_id_rank_score` (`post_id`,`rank_score`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192572928208928769 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_post_metric` WRITE;
/*!40000 ALTER TABLE `p_post_metric` DISABLE KEYS */;
INSERT INTO `p_post_metric` VALUES (1,1080017989,0,0,0,0,0,1780455529,1780455529,0),(2,1080017990,7,0,0,0,0,1780455561,1780459757,0),(3,1080017991,0,0,0,0,0,1780457417,1780457417,0),(4,1080017992,0,0,0,0,0,1780459384,1780459384,0),(5,1080017993,4,0,0,0,0,1780459446,1780459754,0),(6,1080017994,6,0,0,0,0,1780459953,1780460057,0),(187902656621051904,187902656608468992,0,0,0,0,0,1780489084,1780489084,0),(187903825791680512,187903825783291904,0,0,0,0,0,1780489363,1780489363,0),(187905447687094272,187905447682899968,0,0,0,0,0,1780489749,1780489749,0),(187912355978412032,187912355961634816,0,0,0,0,0,1780491396,1780491396,0),(187912693447917568,187912693439528960,0,0,0,0,0,1780491477,1780491477,0),(187916958157504512,187916958144921600,0,0,0,0,0,1780492494,1780492494,0),(187917096951218176,187917096942829568,0,0,0,0,0,1780492527,1780492527,0),(187917149535207424,187917149526818816,0,0,0,0,0,1780492539,1780492539,0),(187917786440269824,187917786427686912,0,0,0,0,0,1780492691,1780492691,0),(187919141275959296,187919141263376384,0,0,0,0,0,1780493014,1780493014,0),(187919527156121600,187919527143538688,0,0,0,0,0,1780493106,1780493106,0),(187919577454215168,187919577437437952,0,0,0,0,0,1780493118,1780493118,0),(187921515860197376,187921515851808768,0,0,0,0,0,1780493580,1780493580,0),(187922349532643328,187922349511671808,0,0,0,0,0,1780493779,1780493779,0),(187922402934521856,187922402921938944,0,0,0,0,0,1780493792,1780493792,0),(187923490412691456,187923490404302848,0,0,0,0,0,1780494051,1780494051,0),(187923537997070336,187923537984487424,0,0,0,0,0,1780494062,1780494062,0),(187924409917374464,187924409904791552,0,0,0,0,0,1780494270,1780494270,0),(187925033216114688,187925033207726080,0,0,0,0,0,1780494419,1780494419,0),(187925186782167040,187925186769584128,0,0,0,0,0,1780494456,1780494456,0),(187940231318601728,187940231306018816,0,0,0,0,0,1780498042,1780498042,0),(187942239190646784,187942239173869568,3,0,0,0,0,1780498521,1781615120,0),(187962543422046208,187962543409463296,1,0,0,0,0,1780503362,1780503707,0),(187964129636515840,187964129628127232,2,0,0,0,0,1780503740,1780503827,0),(187964381135372288,187964381126983680,0,0,0,0,0,1780503800,1780578231,0),(187968258907832320,187968258895249408,1,0,0,0,0,1780504725,1780595048,0),(188306214608699392,188306214596116480,0,0,0,0,0,1780585300,1780585300,0),(188306298838712320,188306298826129408,0,0,0,0,0,1780585320,1780585320,0),(188306668075876352,188306668067487744,0,0,0,0,0,1780585408,1780585408,0),(188307666144067584,188307666135678976,2,0,0,0,0,1780585646,1780585783,0),(188308277807808512,188308277799419904,1,0,0,0,0,1780585792,1780586374,0),(188310514181341184,188310514168758272,3,0,0,0,0,1780586325,1780586363,0),(188310555759476736,188310555746893824,0,0,0,0,0,1780586335,1780586335,0),(189718039766237184,189718039749459968,0,0,0,0,0,1780921905,1780921905,0),(190119319437312000,190119319424729088,0,0,0,0,0,1781017578,1781017578,0),(190131339905728512,190131339897339904,0,0,0,0,0,1781020443,1781020443,0),(190319463214612480,190319463197835264,2,0,0,0,0,1781065295,1781615106,0),(192453341031170048,192453341018587136,0,0,0,0,0,1781574052,1781574052,0),(192455707189051392,192455707159691264,0,0,0,0,0,1781574616,1781574616,0),(192460204552486912,192460204531515392,6,0,0,0,0,1781575688,1781796627,0),(192572795178188800,192572795165605888,3,0,0,0,0,1781602532,1781615068,0),(192572928208928768,192572928192151552,2,0,0,0,0,1781602563,1781615103,0);
/*!40000 ALTER TABLE `p_post_metric` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_post_star`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_post_star` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
  `post_id` bigint NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post_star_post_id` (`post_id`) USING BTREE,
  KEY `idx_post_star_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192625594687553537 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章点赞';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_post_star` WRITE;
/*!40000 ALTER TABLE `p_post_star` DISABLE KEYS */;
INSERT INTO `p_post_star` VALUES (6000028,1080017990,100059,1780459539,1780459562,1780459562,1),(6000029,1080017990,100059,1780459564,1780459564,0,0),(6000030,1080017994,100059,1780459958,1780459958,0,0),(6000031,1080017994,100058,1780460036,1780460036,0,0),(6000032,1080017994,100058,1780460036,1780460036,0,0),(187963925533294592,187962543409463296,100058,1780503692,1780503707,1780503707,1),(187964457215852544,187964381126983680,100058,1780503818,1780503827,1780503827,1),(187964492347342848,187964129628127232,100058,1780503827,1780503827,0,0),(187964492351537152,187964129628127232,100058,1780503827,1780503827,0,0),(187964504682790912,187964381126983680,100058,1780503830,1780503830,1780503830,1),(187964565324038144,187964381126983680,100058,1780503844,1780503852,1780503852,1),(187964600031903744,187964381126983680,100058,1780503852,1780578228,1780578228,1),(187964600040292352,187964381126983680,100058,1780503852,1780578229,1780578229,1),(188276528906240000,187968258895249408,100058,1780578222,1780584367,1780584367,1),(188276565644148736,187964381126983680,100058,1780578231,1780578231,0,0),(188302308788928512,187968258895249408,100058,1780584368,1780584786,1780584786,1),(188304065212448768,187968258895249408,100058,1780584787,1780595048,1780595048,1),(188308243963969536,188307666135678976,100058,1780585783,1780585783,0,0),(188310651242807296,188310514168758272,100058,1780586357,1780586357,0,0),(192621865800302592,192572795165605888,192570609316659200,1781614231,1781614231,0,0),(192621939796213760,192460204531515392,192570609316659200,1781614249,1781615099,1781615099,1),(192625510398820352,192460204531515392,192570609316659200,1781615100,1781615100,0,0),(192625523183058944,192572928192151552,192570609316659200,1781615103,1781615103,0,0),(192625535015190528,190319463197835264,192570609316659200,1781615106,1781615106,0,0),(192625594687553536,187942239173869568,192570609316659200,1781615120,1781615120,0,0);
/*!40000 ALTER TABLE `p_post_star` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_site_settings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_site_settings` (
  `key` varchar(191) COLLATE utf8mb4_general_ci NOT NULL,
  `value` text COLLATE utf8mb4_general_ci NOT NULL,
  `is_encrypted` tinyint NOT NULL DEFAULT '0',
  `created_on` bigint NOT NULL DEFAULT '0',
  `modified_on` bigint NOT NULL DEFAULT '0',
  `deleted_on` bigint NOT NULL DEFAULT '0',
  `is_del` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Admin settings overrides';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_site_settings` WRITE;
/*!40000 ALTER TABLE `p_site_settings` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_site_settings` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_tag` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签名',
  `quote_num` bigint NOT NULL DEFAULT '0' COMMENT '引用数',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_tag_tag` (`tag`) USING BTREE,
  KEY `idx_tag_user_id` (`user_id`) USING BTREE,
  KEY `idx_tag_quote_num` (`quote_num`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=187942239207424001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='标签';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_tag` WRITE;
/*!40000 ALTER TABLE `p_tag` DISABLE KEYS */;
INSERT INTO `p_tag` VALUES (9000065,100058,'CEO',2,1780459384,1780459384,0,0),(187912355995189248,187901818129350656,'x',5,1780491396,1780491396,0,0),(187917786457047040,187901818129350656,'t',1,1780492691,1780492691,0,0),(187940231335378944,100058,'OKX',1,1780498042,1780498042,0,0),(187942239207424000,100058,'六一',1,1780498521,1780498521,0,0);
/*!40000 ALTER TABLE `p_tag` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_topic_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_topic_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `topic_id` bigint NOT NULL COMMENT '标签ID',
  `user_id` bigint NOT NULL COMMENT '创建者ID',
  `alias_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '别名',
  `remark` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `quote_num` bigint DEFAULT NULL COMMENT '引用数',
  `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶 0 为未置顶、1 为已置顶',
  `is_pin` tinyint NOT NULL DEFAULT '0' COMMENT '是否钉住 0 为未钉住、1 为已钉住',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  `reserve_a` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '保留字段a',
  `reserve_b` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '保留字段b',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_topic_user_uid_tid` (`topic_id`,`user_id`) USING BTREE,
  KEY `idx_topic_user_uid_ispin` (`user_id`,`is_pin`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户话题';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_topic_user` WRITE;
/*!40000 ALTER TABLE `p_topic_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_topic_user` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_tweet_comment_thumbs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_tweet_comment_thumbs` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'thumbs ID',
  `user_id` bigint NOT NULL,
  `tweet_id` bigint NOT NULL COMMENT '推文ID',
  `comment_id` bigint NOT NULL COMMENT '评论ID',
  `reply_id` bigint DEFAULT NULL COMMENT '评论回复ID',
  `comment_type` tinyint NOT NULL DEFAULT '0' COMMENT '评论类型 0为推文评论、1为评论回复',
  `is_thumbs_up` tinyint NOT NULL DEFAULT '0' COMMENT '是否点赞',
  `is_thumbs_down` tinyint NOT NULL DEFAULT '0' COMMENT '是否点踩',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tweet_comment_thumbs_uid_tid` (`user_id`,`tweet_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=192625389766443009 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='推文评论点赞';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_tweet_comment_thumbs` WRITE;
/*!40000 ALTER TABLE `p_tweet_comment_thumbs` DISABLE KEYS */;
INSERT INTO `p_tweet_comment_thumbs` VALUES (192625389766443008,192570609316659200,192572795165605888,192625375283511296,0,0,0,0,1781615071,1781615084,0,0);
/*!40000 ALTER TABLE `p_tweet_comment_thumbs` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'MD5密码',
  `salt` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态，1正常，2停用',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `cover_image` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '展示图',
  `balance` bigint NOT NULL COMMENT '用户余额（分）',
  `is_admin` tinyint NOT NULL DEFAULT '0' COMMENT '是否管理员',
  `is_kol` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为KOL',
  `chat_enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否启用AI聊天',
  `address` varchar(64) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '钱包地址',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user_username` (`username`) USING BTREE,
  KEY `idx_user_phone` (`phone`) USING BTREE,
  KEY `idx_address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=193383319155507201 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_user` WRITE;
/*!40000 ALTER TABLE `p_user` DISABLE KEYS */;
INSERT INTO `p_user` VALUES (100058,'abc1230','abc123','','20fea643006e976dd2b01c5bae0deab2','9e883100',1,'http://192.168.30.44:8008/oss/paopao/public/avatar/49/09/f5/a1/be7a-4511-853d-ef235a01114e.png','',0,1,0,0,'0x1234567890123456789012345678901234567890',1780455458,1780455458,0,0),(100059,'wesley','wesley','','901396f1b1b16a502df9abcacfc76601','a0e19d47',1,'https://paopao-demo.vercel.app/avatar/default/ryan.png','',0,0,1,0,'',1780457198,1780457198,1780482023,0),(100060,'sf77','sf77','','d9e611aa912412386dfcd974869173f9','b5bd3644',1,'https://paopao-demo.vercel.app/avatar/default/naomi.png','',0,0,0,0,'',1780488837,1780488837,0,0),(100061,'sf88','sf88','','174f52cdedce61ba48f02dc8859dc575','5172c5d5',1,'https://paopao-demo.vercel.app/avatar/default/audrey.png','',0,0,0,0,'',1780488844,1780488844,0,0),(187901818129350656,'snow1','snow1','','afa6c10a6ad868c8ca31881349adb2bb','dd086c4d',1,'https://paopao-demo.vercel.app/avatar/default/emma.png','',0,0,0,0,'',1780488884,1780488884,0,0),(187901895635894272,'snow2','snow2','','abf47d11ea2a7d521f0a3c21dadec493','347b8c2a',1,'https://paopao-demo.vercel.app/avatar/default/zoe.png','',0,0,0,0,'',1780488902,1780488902,0,0),(187968194357493760,'abc1234','abc1234','','7948a3d4d0be1d7a6451bded38632005','0db37e81',1,'https://paopao-demo.vercel.app/avatar/default/joshua.png','',0,0,1,0,'',1780504709,1780504709,0,0),(187973080667127808,'abc1235','abc1235','','a78ac26b7a44e45748edace5c76954ec','bdc4ac8c',1,'https://paopao-demo.vercel.app/avatar/default/clara.png','',0,0,1,0,'',1780505874,1780505874,0,0),(189717769330098176,'user_D3a5479686','user_D3a5479686','','','',1,'https://paopao-demo.vercel.app/avatar/default/ruby.png','',0,1,1,0,'0xEA9b0FeE68fa8C42351005f7Bb9800D3a5479686',1780921840,1780921840,0,0),(189718285166575616,'user_8E26272005','user_8E26272005','','','',1,'https://paopao-demo.vercel.app/avatar/default/zoe.png','',0,0,1,0,'0x2770ac4E800769c753965c71A2ac718E26272005',1780921963,1780921963,0,0),(190112139342512128,'user_f9Bb7d37Ef','user_f9Bb7d37Ef','','','',1,'https://paopao-demo.vercel.app/avatar/default/miley.png','',0,0,0,0,'0x44eAfC89A0Df82B714e36d153Ad147f9Bb7d37Ef',1781015866,1781015866,0,0),(190119095696359424,'user_7455146cD2','user_7455146cD2','','','',1,'https://paopao-demo.vercel.app/avatar/default/hanna.png','',0,0,0,0,'0xA7fec6e194637072eaaa0e8151A5747455146cD2',1781017524,1781017524,0,0),(190122191415672832,'user_9BFDaCC2ab','user_9BFDaCC2ab','','','',1,'http://192.168.30.44:8008/oss/paopao/public/avatar/24/db/d7/82/cd87-4319-808e-2a6f12de6b1b.png','',0,0,0,0,'0xD438DC158150f5d700F8110236f95A9BFDaCC2ab',1781018262,1781018262,0,0),(190319267848126464,'user_cA244c8D2A','user_cA244c8D2A','','','',1,'https://paopao-demo.vercel.app/avatar/default/jane.png','',0,0,0,0,'0x25EdF0156130240E35526B70013ECecA244c8D2A',1781065249,1781065249,0,0),(190326398693408768,'user_8e6f31A8F8','user_8e6f31A8F8','','','',1,'https://paopao-demo.vercel.app/avatar/default/ryan.png','',0,0,0,0,'0x51cF387F793155963B274E5D702e3F8e6f31A8F8',1781066949,1781066949,0,0),(190330337203585024,'user_7391201694','user_7391201694','','','',1,'https://paopao-demo.vercel.app/avatar/default/miley.png','',0,0,0,0,'0xa4eb1EAcC6c2067FB1Cc7A822c2E737391201694',1781067888,1781067888,0,0),(190330674861834240,'user_EA60F60756','user_EA60F60756','','','',1,'https://paopao-demo.vercel.app/avatar/default/emily.png','',0,0,0,0,'0xF5276420Fb1736E199dBa125395AEdEA60F60756',1781067969,1781067969,0,0),(190334927470854144,'user_3D1A573095','user_3D1A573095','','','',1,'https://paopao-demo.vercel.app/avatar/default/alexa.png','',0,0,1,1,'0x1D97B92788E9bB4D163c5b02Ef68de3D1A573095',1781068982,1781068982,0,0),(190336026986676224,'user_A93ba7b701','user_A93ba7b701','','','',1,'https://paopao-demo.vercel.app/avatar/default/claire.png','',0,0,1,1,'0x7235c147c54F157D59FB66ABccF8CcA93ba7b701',1781069245,1781069245,0,0),(190336279810932736,'user_2b498e261C','user_2b498e261C','','','',1,'https://paopao-demo.vercel.app/avatar/default/finn.png','',0,0,1,1,'0xA2eB23D7A8Cdd3D67e262D5165a90B2b498e261C',1781069305,1781069305,0,0),(190337599783567360,'user_b58043AbF2','user_b58043AbF2','','','',1,'https://paopao-demo.vercel.app/avatar/default/george.png','',0,0,1,1,'0xaFDf2C2d2F6cB29605F3ECE56165C1b58043AbF2',1781069620,1781069620,0,0),(190338733612990464,'user_1026A62D6E','user_1026A62D6E','','','',1,'https://paopao-demo.vercel.app/avatar/default/john.png','',0,0,1,1,'0x7f31Cf3999ec936716f9d3FDdA76901026A62D6E',1781069890,1781069890,0,0),(192570609316659200,'user_bd82E35029','user_bd82E35029','','','',1,'https://paopao-demo.vercel.app/avatar/default/ryan.png','http://192.168.30.44:8008/oss/paopao/public/avatar/63/78/f5/c1/43fc-45de-8fcd-3b2950924ac3.jpeg',0,0,1,0,'0x35997Bfcae6c5a7811B85e837dB844bd82E35029',1781602011,1781602011,0,0),(193383319155507200,'user_F6bC0cb5b3','user_F6bC0cb5b3','','','',1,'https://paopao-demo.vercel.app/avatar/default/arthur.png','',0,0,0,0,'0x0D745B5061DabC52641a9f29f8dFe1F6bC0cb5b3',1781795776,1781795776,0,0);
/*!40000 ALTER TABLE `p_user` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_user_metric`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_user_metric` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `tweets_count` int NOT NULL DEFAULT '0',
  `latest_trends_on` bigint NOT NULL DEFAULT '0' COMMENT '最新动态时间',
  `is_del` tinyint NOT NULL DEFAULT '0',
  `created_on` bigint NOT NULL DEFAULT '0',
  `modified_on` bigint NOT NULL DEFAULT '0',
  `deleted_on` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_metric_user_id_tweets_count_trends` (`user_id`,`tweets_count`,`latest_trends_on`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=193383319180673025 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_user_metric` WRITE;
/*!40000 ALTER TABLE `p_user_metric` DISABLE KEYS */;
INSERT INTO `p_user_metric` VALUES (1,100058,29,1780586335,0,1780455458,1780586335,0),(2,100059,2,1780459953,0,1780457198,1780459953,0),(3,100060,0,0,0,1780488837,1780488837,0),(4,100061,0,0,0,1780488844,1780488844,0),(187901818141933568,187901818129350656,7,1780494270,0,1780488884,1780494270,0),(187901895648477184,187901895635894272,0,0,0,1780488902,1780488902,0),(187968194365882368,187968194357493760,1,1780504725,0,1780504709,1780504725,0),(187973080679710720,187973080667127808,0,0,0,1780505874,1780505874,0),(189717769342681088,189717769330098176,4,1781575688,0,1780921840,1781575688,0),(189718285179158528,189718285166575616,0,0,0,1780921963,1780921963,0),(190112139350900736,190112139342512128,0,0,0,1781015866,1781015866,0),(190119095713136640,190119095696359424,1,1781017578,0,1781017524,1781017578,0),(190122191428255744,190122191415672832,1,1781020443,0,1781018262,1781020443,0),(190319267864903680,190319267848126464,1,1781065295,0,1781065249,1781065296,0),(190326398701797376,190326398693408768,0,0,0,1781066949,1781066949,0),(190330337220362240,190330337203585024,0,0,0,1781067888,1781067888,0),(190330674874417152,190330674861834240,0,0,0,1781067969,1781067969,0),(190334927487631360,190334927470854144,0,0,0,1781068982,1781068982,0),(190336026999259136,190336026986676224,0,0,0,1781069245,1781069245,0),(190336279827709952,190336279810932736,0,0,0,1781069305,1781069305,0),(190337599800344576,190337599783567360,0,0,0,1781069620,1781069620,0),(190338733625573376,190338733612990464,0,0,0,1781069890,1781069890,0),(192570609329242112,192570609316659200,2,1781602563,0,1781602011,1781602563,0),(193383319180673024,193383319155507200,0,0,0,1781795776,1781795776,0);
/*!40000 ALTER TABLE `p_user_metric` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_user_relation`;
/*!50001 DROP VIEW IF EXISTS `p_user_relation`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `p_user_relation` AS SELECT 
 1 AS `user_id`,
 1 AS `he_uid`,
 1 AS `style`*/;
SET character_set_client = @saved_cs_client;
DROP TABLE IF EXISTS `p_wallet_recharge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_wallet_recharge` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '充值ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `amount` bigint NOT NULL DEFAULT '0' COMMENT '充值金额',
  `trade_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付宝订单号',
  `trade_status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易状态',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_wallet_recharge_user_id` (`user_id`) USING BTREE,
  KEY `idx_wallet_recharge_trade_no` (`trade_no`) USING BTREE,
  KEY `idx_wallet_recharge_trade_status` (`trade_status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10023 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包流水';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_wallet_recharge` WRITE;
/*!40000 ALTER TABLE `p_wallet_recharge` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_wallet_recharge` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `p_wallet_statement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `p_wallet_statement` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '账单ID',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `change_amount` bigint NOT NULL DEFAULT '0' COMMENT '变动金额',
  `balance_snapshot` bigint NOT NULL DEFAULT '0' COMMENT '资金快照',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '变动原因',
  `post_id` bigint NOT NULL DEFAULT '0' COMMENT '关联动态',
  `created_on` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_wallet_statement_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10010 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包流水';
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `p_wallet_statement` WRITE;
/*!40000 ALTER TABLE `p_wallet_statement` DISABLE KEYS */;
/*!40000 ALTER TABLE `p_wallet_statement` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_ai_workflow_sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_ai_workflow_sessions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `tab` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '会话类型',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '会话标题',
  `summary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '摘要',
  `conversation_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Dify会话ID',
  `message_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'Dify消息ID',
  `current_node_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '当前选中节点ID',
  `settings` json DEFAULT NULL COMMENT '页面设置',
  `form_data` json DEFAULT NULL COMMENT '表单数据',
  `result_data` json DEFAULT NULL COMMENT '当前展示结果',
  `messages` json DEFAULT NULL COMMENT '会话消息',
  PRIMARY KEY (`id`),
  KEY `idx_sys_ai_workflow_sessions_deleted_at` (`deleted_at`),
  KEY `idx_sys_ai_workflow_sessions_user_id` (`user_id`),
  KEY `idx_sys_ai_workflow_sessions_tab` (`tab`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_ai_workflow_sessions` WRITE;
/*!40000 ALTER TABLE `sys_ai_workflow_sessions` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_ai_workflow_sessions` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_api_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_api_tokens` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `token` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT 'Token',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `expires_at` datetime(3) DEFAULT NULL COMMENT '过期时间',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_tokens_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_api_tokens` WRITE;
/*!40000 ALTER TABLE `sys_api_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_api_tokens` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_apis` WRITE;
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` VALUES (1,'2026-05-25 17:32:16.982','2026-05-25 17:32:16.982',NULL,'/info/createInfo','新建公告','公告','POST'),(2,'2026-05-25 17:32:16.983','2026-05-25 17:32:16.983',NULL,'/info/deleteInfo','删除公告','公告','DELETE'),(3,'2026-05-25 17:32:16.983','2026-05-25 17:32:16.983',NULL,'/info/deleteInfoByIds','批量删除公告','公告','DELETE'),(4,'2026-05-25 17:32:16.984','2026-05-25 17:32:16.984',NULL,'/info/updateInfo','更新公告','公告','PUT'),(5,'2026-05-25 17:32:16.984','2026-05-25 17:32:16.984',NULL,'/info/findInfo','根据ID获取公告','公告','GET'),(6,'2026-05-25 17:32:16.985','2026-05-25 17:32:16.985',NULL,'/info/getInfoList','获取公告列表','公告','GET'),(7,'2026-05-25 17:32:17.001','2026-05-25 17:32:17.001',NULL,'/autoCode/getDB','获取数据库列表','代码生成器','GET'),(8,'2026-05-25 17:32:17.002','2026-05-25 17:32:17.002',NULL,'/autoCode/getTables','获取数据表列表','代码生成器','GET'),(9,'2026-05-25 17:32:17.002','2026-05-25 17:32:17.002',NULL,'/autoCode/getColumn','获取字段列表','代码生成器','GET'),(10,'2026-05-25 17:32:17.002','2026-05-25 17:32:17.002',NULL,'/autoCode/preview','预览自动代码','代码生成器','POST'),(11,'2026-05-25 17:32:17.003','2026-05-25 17:32:17.003',NULL,'/autoCode/createTemp','生成自动代码','代码生成器','POST'),(12,'2026-05-25 17:32:17.003','2026-05-25 17:32:17.003',NULL,'/autoCode/mcp','生成 MCP 工具','代码生成器','POST'),(13,'2026-05-25 17:32:17.003','2026-05-25 17:32:17.003',NULL,'/autoCode/mcpStatus','获取 MCP 状态','代码生成器','POST'),(14,'2026-05-25 17:32:17.003','2026-05-25 17:32:17.003',NULL,'/autoCode/mcpStart','启动 MCP','代码生成器','POST'),(15,'2026-05-25 17:32:17.004','2026-05-25 17:32:17.004',NULL,'/autoCode/mcpStop','停止 MCP','代码生成器','POST'),(16,'2026-05-25 17:32:17.004','2026-05-25 17:32:17.004',NULL,'/autoCode/mcpList','获取 MCP 工具列表','代码生成器','POST'),(17,'2026-05-25 17:32:17.004','2026-05-25 17:32:17.004',NULL,'/autoCode/mcpRoutes','获取 MCP 路由','代码生成器','POST'),(18,'2026-05-25 17:32:17.004','2026-05-25 17:32:17.004',NULL,'/autoCode/mcpTest','测试 MCP 调用','代码生成器','POST'),(19,'2026-05-25 17:32:17.004','2026-05-25 17:32:17.004',NULL,'/autoCode/pubPlug','打包插件','代码生成器','POST'),(20,'2026-05-25 17:32:17.005','2026-05-25 17:32:17.005',NULL,'/autoCode/installPlugin','安装插件','代码生成器','POST'),(21,'2026-05-25 17:32:17.005','2026-05-25 17:32:17.005',NULL,'/autoCode/removePlugin','移除插件','代码生成器','POST'),(22,'2026-05-25 17:32:17.005','2026-05-25 17:32:17.005',NULL,'/autoCode/getPluginList','获取插件列表','代码生成器','GET'),(23,'2026-05-25 17:32:17.005','2026-05-25 17:32:17.005',NULL,'/autoCode/saveAIWorkflowSession','保存 AI 工作流会话','代码生成器','POST'),(24,'2026-05-25 17:32:17.005','2026-05-25 17:32:17.005',NULL,'/autoCode/getAIWorkflowSessionList','获取 AI 工作流列表','代码生成器','POST'),(25,'2026-05-25 17:32:17.006','2026-05-25 17:32:17.006',NULL,'/autoCode/getAIWorkflowSessionDetail','获取 AI 工作流详情','代码生成器','POST'),(26,'2026-05-25 17:32:17.006','2026-05-25 17:32:17.006',NULL,'/autoCode/deleteAIWorkflowSession','删除 AI 工作流会话','代码生成器','POST'),(27,'2026-05-25 17:32:17.006','2026-05-25 17:32:17.006',NULL,'/autoCode/dumpAIWorkflowMarkdown','导出 AI 工作流 Markdown','代码生成器','POST'),(28,'2026-05-25 17:32:17.006','2026-05-25 17:32:17.006',NULL,'/autoCode/getPackage','获取自动化包列表','模板配置','POST'),(29,'2026-05-25 17:32:17.006','2026-05-25 17:32:17.006',NULL,'/autoCode/delPackage','删除自动化包','模板配置','POST'),(30,'2026-05-25 17:32:17.007','2026-05-25 17:32:17.007',NULL,'/autoCode/createPackage','创建自动化包','模板配置','POST'),(31,'2026-05-25 17:32:17.007','2026-05-25 17:32:17.007',NULL,'/autoCode/getTemplates','获取模板列表','模板配置','GET'),(32,'2026-05-25 17:32:17.007','2026-05-25 17:32:17.007',NULL,'/autoCode/getMeta','获取自动代码历史元数据','代码生成器历史','POST'),(33,'2026-05-25 17:32:17.007','2026-05-25 17:32:17.007',NULL,'/autoCode/rollback','回滚自动代码历史','代码生成器历史','POST'),(34,'2026-05-25 17:32:17.007','2026-05-25 17:32:17.007',NULL,'/autoCode/delSysHistory','删除自动代码历史','代码生成器历史','POST'),(35,'2026-05-25 17:32:17.008','2026-05-25 17:32:17.008',NULL,'/autoCode/getSysHistory','获取自动代码历史列表','代码生成器历史','POST'),(36,'2026-05-25 17:32:17.008','2026-05-25 17:32:17.008',NULL,'/autoCode/addFunc','追加自动代码方法','代码生成器历史','POST'),(37,'2026-05-25 17:32:17.008','2026-05-25 17:32:17.008',NULL,'/skills/getTools','获取 AI 工具列表','skills','GET'),(38,'2026-05-25 17:32:17.009','2026-05-25 17:32:17.009',NULL,'/skills/getSkillList','获取技能列表','skills','POST'),(39,'2026-05-25 17:32:17.009','2026-05-25 17:32:17.009',NULL,'/skills/getSkillDetail','获取技能详情','skills','POST'),(40,'2026-05-25 17:32:17.009','2026-05-25 17:32:17.009',NULL,'/skills/saveSkill','保存技能','skills','POST'),(41,'2026-05-25 17:32:17.009','2026-05-25 17:32:17.009',NULL,'/skills/deleteSkill','删除技能','skills','POST'),(42,'2026-05-25 17:32:17.010','2026-05-25 17:32:17.010',NULL,'/skills/createScript','创建脚本','skills','POST'),(43,'2026-05-25 17:32:17.010','2026-05-25 17:32:17.010',NULL,'/skills/getScript','获取脚本','skills','POST'),(44,'2026-05-25 17:32:17.010','2026-05-25 17:32:17.010',NULL,'/skills/saveScript','保存脚本','skills','POST'),(45,'2026-05-25 17:32:17.010','2026-05-25 17:32:17.010',NULL,'/skills/createResource','创建资源','skills','POST'),(46,'2026-05-25 17:32:17.010','2026-05-25 17:32:17.010',NULL,'/skills/getResource','获取资源','skills','POST'),(47,'2026-05-25 17:32:17.011','2026-05-25 17:32:17.011',NULL,'/skills/saveResource','保存资源','skills','POST'),(48,'2026-05-25 17:32:17.011','2026-05-25 17:32:17.011',NULL,'/skills/createReference','创建参考资料','skills','POST'),(49,'2026-05-25 17:32:17.011','2026-05-25 17:32:17.011',NULL,'/skills/getReference','获取参考资料','skills','POST'),(50,'2026-05-25 17:32:17.011','2026-05-25 17:32:17.011',NULL,'/skills/saveReference','保存参考资料','skills','POST'),(51,'2026-05-25 17:32:17.012','2026-05-25 17:32:17.012',NULL,'/skills/createTemplate','创建模板','skills','POST'),(52,'2026-05-25 17:32:17.013','2026-05-25 17:32:17.013',NULL,'/skills/getTemplate','获取模板','skills','POST'),(53,'2026-05-25 17:32:17.013','2026-05-25 17:32:17.013',NULL,'/skills/saveTemplate','保存模板','skills','POST'),(54,'2026-05-25 17:32:17.014','2026-05-25 17:32:17.014',NULL,'/skills/getGlobalConstraint','获取全局约束','skills','POST'),(55,'2026-05-25 17:32:17.014','2026-05-25 17:32:17.014',NULL,'/skills/saveGlobalConstraint','保存全局约束','skills','POST'),(56,'2026-05-25 17:32:17.015','2026-05-25 17:32:17.015',NULL,'/skills/packageSkill','打包技能','skills','POST'),(57,'2026-05-25 17:32:17.015','2026-05-25 17:32:17.015',NULL,'/skills/downloadOnlineSkill','下载在线技能','skills','POST'),(58,'2026-06-09 18:53:39.000','2026-06-09 18:53:39.000',NULL,'/api/v1/h5Admin/kolCategory','保存KOL分类','KOL管理','PUT'),(59,'2026-06-09 18:53:39.000','2026-06-09 18:53:39.000',NULL,'/api/v1/h5Admin/kolCategory','删除KOL分类','KOL管理','DELETE'),(60,'2026-06-09 18:53:39.000','2026-06-09 18:53:39.000',NULL,'/api/v1/h5Admin/kolManageList','KOL用户管理列表','KOL管理','GET'),(61,'2026-06-09 18:53:39.000','2026-06-09 18:53:39.000',NULL,'/api/v1/h5Admin/kolAssignCategory','分配KOL分类','KOL管理','PUT');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_authorities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `uni_sys_authorities_authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9529 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_authorities` WRITE;
/*!40000 ALTER TABLE `sys_authorities` DISABLE KEYS */;
INSERT INTO `sys_authorities` VALUES ('2026-05-25 17:53:02.094','2026-05-25 17:53:02.108',NULL,888,'普通用户',0,'dashboard'),('2026-05-25 17:53:02.094','2026-05-25 17:53:02.112',NULL,8881,'普通用户子角色',888,'dashboard'),('2026-05-25 17:53:02.094','2026-05-25 17:53:02.110',NULL,9528,'测试角色',0,'dashboard');
/*!40000 ALTER TABLE `sys_authorities` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_authority_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_authority_btns` WRITE;
/*!40000 ALTER TABLE `sys_authority_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_authority_btns` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_authority_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_authority_menus` WRITE;
/*!40000 ALTER TABLE `sys_authority_menus` DISABLE KEYS */;
INSERT INTO `sys_authority_menus` VALUES (16,888),(16,8881),(16,9528),(18,888),(18,8881),(19,888),(19,8881),(19,9528),(20,888),(20,8881),(23,888),(23,8881),(23,9528),(25,888),(26,888),(27,888),(28,888),(29,888),(30,888),(31,888),(32,888),(33,888),(34,888),(35,888),(36,888),(37,888),(37,8881),(38,888),(38,8881),(39,888),(39,8881),(40,888),(41,888),(42,888),(43,888),(59,888),(60,888);
/*!40000 ALTER TABLE `sys_authority_menus` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_auto_code_histories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名',
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模块名或插件名',
  `request` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '前端传入的结构化信息',
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '结构体名称',
  `abbreviation` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '结构体简称',
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '业务库',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '结构体中文名',
  `templates` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '模板信息',
  `injections` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '注入信息',
  `flag` bigint DEFAULT NULL COMMENT '[0:创建,1:回滚]',
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联API ID',
  `menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  `export_template_id` bigint unsigned DEFAULT NULL COMMENT '导出模板ID',
  `package_id` bigint unsigned DEFAULT NULL COMMENT '包ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_auto_code_histories` WRITE;
/*!40000 ALTER TABLE `sys_auto_code_histories` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_auto_code_histories` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_auto_code_packages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_auto_code_packages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '显示名称',
  `template` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板',
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '包名',
  `module` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_packages_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_auto_code_packages` WRITE;
/*!40000 ALTER TABLE `sys_auto_code_packages` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_auto_code_packages` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_base_menu_btns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_base_menu_btns` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_btns` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_btns` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint unsigned DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_base_menu_parameters` WRITE;
/*!40000 ALTER TABLE `sys_base_menu_parameters` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_base_menu_parameters` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_base_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_base_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint unsigned DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '高亮菜单',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '是否是基础路由（开发中）',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单名',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单图标',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '自动关闭tab',
  `transition_type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由切换动画',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_base_menus` WRITE;
/*!40000 ALTER TABLE `sys_base_menus` DISABLE KEYS */;
INSERT INTO `sys_base_menus` VALUES (16,'2026-05-25 17:53:02.104','2026-05-25 17:53:02.104',NULL,0,0,'dashboard','dashboard',0,'view/dashboard/index.vue',1,'',0,0,'仪表盘','odometer',0,''),(18,'2026-05-25 17:53:02.104','2026-05-25 17:53:02.104',NULL,0,0,'admin','superAdmin',0,'view/superAdmin/index.vue',3,'',0,0,'超级管理员','user',0,''),(19,'2026-05-25 17:53:02.104','2026-05-25 17:53:02.104',NULL,0,0,'person','person',1,'view/person/person.vue',4,'',0,0,'个人信息','message',0,''),(20,'2026-05-25 17:53:02.104','2026-05-25 17:53:02.104',NULL,0,0,'example','example',0,'view/example/index.vue',7,'',0,0,'示例文件','management',0,''),(23,'2026-05-25 17:53:02.104','2026-05-25 17:53:02.104',NULL,0,0,'state','state',0,'view/system/state.vue',8,'',0,0,'服务器状态','cloudy',0,''),(25,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'authority','authority',0,'view/superAdmin/authority/authority.vue',1,'',0,0,'角色管理','avatar',0,''),(26,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'menu','menu',0,'view/superAdmin/menu/menu.vue',2,'',1,0,'菜单管理','tickets',0,''),(27,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'api','api',0,'view/superAdmin/api/api.vue',3,'',1,0,'api管理','platform',0,''),(28,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'user','user',0,'view/superAdmin/user/user.vue',4,'',0,0,'用户管理','coordinate',0,''),(29,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'dictionary','dictionary',0,'view/superAdmin/dictionary/sysDictionary.vue',5,'',0,0,'字典管理','notebook',0,''),(30,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'operation','operation',0,'view/superAdmin/operation/sysOperationRecord.vue',6,'',0,0,'操作历史','pie-chart',0,''),(31,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'sysParams','sysParams',0,'view/superAdmin/params/sysParams.vue',7,'',0,0,'参数管理','compass',0,''),(32,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'system','system',0,'view/systemTools/system/system.vue',8,'',0,0,'系统配置','operation',0,''),(33,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'apiToken','apiToken',0,'view/systemTools/apiToken/index.vue',9,'',0,0,'API Token','key',0,''),(34,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'loginLog','loginLog',0,'view/systemTools/loginLog/index.vue',10,'',0,0,'登录日志','monitor',0,''),(35,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'sysVersion','sysVersion',0,'view/systemTools/version/version.vue',11,'',0,0,'版本管理','server',0,''),(36,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,18,'sysError','sysError',0,'view/systemTools/sysError/sysError.vue',12,'',0,0,'错误日志','warn',0,''),(37,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,20,'upload','upload',0,'view/example/upload/upload.vue',5,'',0,0,'媒体库（上传下载）','upload',0,''),(38,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,20,'breakpoint','breakpoint',0,'view/example/breakpoint/breakpoint.vue',6,'',0,0,'断点续传','upload-filled',0,''),(39,'2026-05-25 17:53:02.105','2026-05-25 17:53:02.105',NULL,1,20,'customer','customer',0,'view/example/customer/customer.vue',7,'',0,0,'客户列表（资源示例）','avatar',0,''),(40,'2026-06-03 18:13:31.000','2026-06-03 18:13:31.000',NULL,NULL,0,'ops','ops',0,'view/superAdmin/index.vue',5,NULL,0,NULL,'运维管理','setting',NULL,NULL),(41,'2026-06-03 18:13:31.000','2026-06-03 18:13:31.000',NULL,NULL,40,'h5User','h5User',0,'view/ops/h5Users/h5Users.vue',1,NULL,0,NULL,'用户管理','user',NULL,NULL),(42,'2026-06-03 18:13:31.000','2026-06-03 18:13:31.000',NULL,NULL,40,'h5Post','h5Post',0,'view/ops/h5Posts/h5Posts.vue',2,NULL,0,NULL,'贴文管理','edit',NULL,NULL),(43,'2026-06-03 18:13:31.000','2026-06-03 18:13:31.000',NULL,NULL,40,'h5Tag','h5Tag',0,'view/ops/h5Tags/h5Tags.vue',3,NULL,0,NULL,'话题管理','collection-tag',NULL,NULL),(59,'2026-06-09 18:55:38.000','2026-06-09 18:55:38.000',NULL,NULL,40,'kolCategory','kolCategory',0,'view/ops/kolCategories/kolCategories.vue',4,NULL,NULL,NULL,'KOL分类','collection',NULL,NULL),(60,'2026-06-09 19:12:22.000','2026-06-09 19:12:22.000',NULL,NULL,40,'sysMsg','sysMsg',0,'view/ops/sysMsg/sysMsg.vue',5,NULL,NULL,NULL,'系统消息','message',NULL,NULL);
/*!40000 ALTER TABLE `sys_base_menus` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_data_authority_id`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_data_authority_id` WRITE;
/*!40000 ALTER TABLE `sys_data_authority_id` DISABLE KEYS */;
INSERT INTO `sys_data_authority_id` VALUES (888,888),(888,8881),(888,9528),(9528,8881),(9528,9528);
/*!40000 ALTER TABLE `sys_data_authority_id` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_dictionaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionaries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父级字典ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_dictionaries` WRITE;
/*!40000 ALTER TABLE `sys_dictionaries` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_dictionaries` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_dictionary_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint unsigned DEFAULT NULL COMMENT '关联标记',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父级字典详情ID',
  `level` bigint DEFAULT NULL COMMENT '层级深度',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '层级路径',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_dictionary_details` WRITE;
/*!40000 ALTER TABLE `sys_dictionary_details` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_dictionary_details` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_error`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_error` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `form` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '错误来源',
  `info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '错误内容',
  `level` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '日志等级',
  `solution` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '解决方案',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '未处理' COMMENT '处理状态',
  PRIMARY KEY (`id`),
  KEY `idx_sys_error_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_error` WRITE;
/*!40000 ALTER TABLE `sys_error` DISABLE KEYS */;
INSERT INTO `sys_error` VALUES (1,'2026-05-25 17:32:17.077','2026-05-25 17:32:17.077',NULL,'后端','server启动失败 | 错误: listen tcp :8888: bind: address already in use \n 源文件:/home/v3/workspace/gin-vue-admin/server/core/server_run.go:36 \n 调用栈：github.com/flipped-aurora/gin-vue-admin/server/core.initServer.func1\n	/home/v3/workspace/gin-vue-admin/server/core/server_run.go:36','error',NULL,'未处理'),(2,'2026-05-25 17:37:20.765','2026-05-25 17:37:20.765',NULL,'前端','错误信息: TypeError: can\'t access property \"forEach\", asyncRouters.value[0].children is undefined\nStack: 调用栈: useRouterStore</<@http://localhost:8080/src/pinia/modules/router.js:145:5\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\ndoWatch/baseWatchOptions.call@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2538:72\ngetter@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1209:18\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\njob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1256:17\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:31\nflushJobs@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2225:26\n','error',NULL,'未处理'),(3,'2026-05-25 17:37:29.003','2026-05-25 17:37:29.003',NULL,'前端','错误信息: TypeError: can\'t access property \"forEach\", asyncRouters.value[0].children is undefined\nStack: 调用栈: useRouterStore</<@http://localhost:8080/src/pinia/modules/router.js:145:5\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\ndoWatch/baseWatchOptions.call@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2538:72\ngetter@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1209:18\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\njob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1256:17\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:31\nflushJobs@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2225:26\npromise callback*queueFlush@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2168:66\nqueueJob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2164:3\ndoWatch/baseWatchOptions.scheduler@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2547:17\nwatch$1/effect.scheduler<@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1260:48\ntrigger@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1496:34\nendBatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:543:7\nnotify@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1587:5\ntrigger@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1579:9\nset value@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1817:14\nSetAsyncRouter@http://localhost:8080/src/pinia/modules/router.js:191:5\nasync*wrappedAction@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:1029:14\npatchActionForGrouping/store[actionName]@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:716:40\nsetupRouter@http://localhost:8080/src/permission.js:120:36\n@http://localhost:8080/src/permission.js:176:15\nasync*guardToPromiseFn/</</guardReturn<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:700:50\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:682:83\nguardToPromiseFn/</<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:700:37\nguardToPromiseFn/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:684:15\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4094:13\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2455:64\ncreateRouter/runGuardQueue/</<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:77\npromise callback*createRouter/runGuardQueue/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:52\nrunGuardQueue@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:17\ncreateRouter/navigate/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2470:11\npromise callback*navigate@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2466:32\npushWithRedirect@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2427:48\npushWithRedirect@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2412:30\npush@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2380:10\ninstall@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2627:9\nuse@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4020:13\n@http://localhost:8080/src/main.js:35:4\nsetTimeout handler*@http://localhost:8080/__uno.css:19:42\n@http://localhost:8080/__uno.css:19:9\n','error',NULL,'未处理'),(4,'2026-05-25 17:37:34.345','2026-05-25 17:37:34.345',NULL,'前端','错误信息: TypeError: can\'t access property \"forEach\", asyncRouters.value[0].children is undefined\nStack: 调用栈: useRouterStore</<@http://localhost:8080/src/pinia/modules/router.js:145:5\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\ndoWatch/baseWatchOptions.call@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2538:72\ngetter@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1209:18\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\njob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1256:17\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:31\nflushJobs@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2225:26\n','error',NULL,'未处理'),(5,'2026-05-25 17:37:40.695','2026-05-25 17:37:40.695',NULL,'前端','错误信息: TypeError: can\'t access property \"forEach\", asyncRouters.value[0].children is undefined\nStack: 调用栈: useRouterStore</<@http://localhost:8080/src/pinia/modules/router.js:145:5\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\ndoWatch/baseWatchOptions.call@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2538:72\ngetter@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1209:18\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\njob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1256:17\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:31\nflushJobs@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2225:26\npromise callback*queueFlush@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2168:66\nqueueJob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2164:3\ndoWatch/baseWatchOptions.scheduler@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2547:17\nwatch$1/effect.scheduler<@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1260:48\ntrigger@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1496:34\nendBatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:543:7\nnotify@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1587:5\ntrigger@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1579:9\nset value@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1817:14\nSetAsyncRouter@http://localhost:8080/src/pinia/modules/router.js:191:5\nasync*wrappedAction@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:1029:14\npatchActionForGrouping/store[actionName]@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:716:40\nsetupRouter@http://localhost:8080/src/permission.js:120:36\n@http://localhost:8080/src/permission.js:176:15\nasync*guardToPromiseFn/</</guardReturn<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:700:50\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:682:83\nguardToPromiseFn/</<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:700:37\nguardToPromiseFn/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:684:15\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4094:13\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2455:64\ncreateRouter/runGuardQueue/</<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:77\npromise callback*createRouter/runGuardQueue/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:52\nrunGuardQueue@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:17\ncreateRouter/navigate/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2470:11\npromise callback*navigate@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2466:32\npushWithRedirect@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2427:48\npush@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2380:10\ninstall@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2627:9\nuse@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4020:13\n@http://localhost:8080/src/main.js:35:4\nsetTimeout handler*@http://localhost:8080/__uno.css:19:42\n@http://localhost:8080/__uno.css:19:9\n','error',NULL,'未处理'),(6,'2026-05-25 17:39:12.461','2026-05-25 17:39:12.461',NULL,'后端','重新写入cookie token失败,未能成功解析token,请检查请求头是否存在x-token且claims是否为规定结构 \n 源文件:/home/v3/workspace/gin-vue-admin/server/utils/claims.go:49 \n 调用栈：github.com/flipped-aurora/gin-vue-admin/server/utils.GetToken\n	/home/v3/workspace/gin-vue-admin/server/utils/claims.go:49\ngithub.com/flipped-aurora/gin-vue-admin/server/initialize.Routers.JWTAuth.func3\n	/home/v3/workspace/gin-vue-admin/server/middleware/jwt.go:19\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/flipped-aurora/gin-vue-admin/server/initialize.Routers.GinRecovery.func2\n	/home/v3/workspace/gin-vue-admin/server/middleware/error.go:78\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:633\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:589\nnet/http.serverHandler.ServeHTTP\n	/snap/go/11127/src/net/http/server.go:3311\nnet/http.(*conn).serve\n	/snap/go/11127/src/net/http/server.go:2073 \n 最终调用方法:/home/v3/workspace/gin-vue-admin/server/utils/claims.go:49 (GetToken lines 42-55)\n----- 产生日志的方法代码如下 -----\nfunc GetToken(c *gin.Context) string {\n	token := c.Request.Header.Get(\"x-token\")\n	if token == \"\" {\n		j := NewJWT()\n		token, _ = c.Cookie(\"x-token\")\n		claims, err := j.ParseToken(token)\n		if err != nil {\n			global.GVA_LOG.Error(\"重新写入cookie token失败,未能成功解析token,请检查请求头是否存在x-token且claims是否为规定结构\")\n			return token\n		}\n		SetToken(c, token, int(claims.ExpiresAt.Unix()-time.Now().Unix()))\n	}\n	return token\n}','error',NULL,'未处理'),(7,'2026-05-25 17:46:22.776','2026-05-25 17:46:22.776',NULL,'前端','错误信息: TypeError: can\'t access property \"forEach\", asyncRouters.value[0].children is undefined\nStack: 调用栈: useRouterStore</<@http://localhost:8080/src/pinia/modules/router.js:145:5\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\ndoWatch/baseWatchOptions.call@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2538:72\ngetter@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1209:18\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\njob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1256:17\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:31\nflushJobs@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2225:26\npromise callback*queueFlush@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2168:66\nqueueJob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2164:3\ndoWatch/baseWatchOptions.scheduler@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2547:17\nwatch$1/effect.scheduler<@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1260:48\ntrigger@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1496:34\nendBatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:543:7\nnotify@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1587:5\ntrigger@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1579:9\nset value@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1817:14\nSetAsyncRouter@http://localhost:8080/src/pinia/modules/router.js:191:5\nasync*wrappedAction@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:1029:14\npatchActionForGrouping/store[actionName]@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:716:40\nsetupRouter@http://localhost:8080/src/permission.js:120:36\n@http://localhost:8080/src/permission.js:176:15\nasync*guardToPromiseFn/</</guardReturn<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:700:50\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:682:83\nguardToPromiseFn/</<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:700:37\nguardToPromiseFn/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:684:15\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4094:13\nrunWithContext@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2455:64\ncreateRouter/runGuardQueue/</<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:77\npromise callback*createRouter/runGuardQueue/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:52\nrunGuardQueue@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2657:17\ncreateRouter/navigate/<@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2470:11\npromise callback*navigate@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2466:32\npushWithRedirect@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2427:48\npush@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2380:10\ninstall@http://localhost:8080/node_modules/.vite/deps/vue-router.js?v=78c52a94:2627:9\nuse@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4020:13\n@http://localhost:8080/src/main.js:35:4\nsetTimeout handler*@http://localhost:8080/__uno.css:19:42\n@http://localhost:8080/__uno.css:19:9\n','error',NULL,'未处理'),(8,'2026-05-25 17:46:41.680','2026-05-25 17:46:41.680',NULL,'前端','错误信息: TypeError: can\'t access property \"forEach\", asyncRouters.value[0].children is undefined\nStack: 调用栈: useRouterStore</<@http://localhost:8080/src/pinia/modules/router.js:145:5\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\ndoWatch/baseWatchOptions.call@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2538:72\ngetter@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1209:18\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\njob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1256:17\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:31\nflushJobs@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2225:26\npromise callback*queueFlush@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2168:66\nqueueJob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2164:3\ndoWatch/baseWatchOptions.scheduler@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2547:17\nwatch$1/effect.scheduler<@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1260:48\ntrigger@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1496:34\nendBatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:543:7\nnotify@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1587:5\ntrigger@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1579:9\nset value@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1817:14\nSetAsyncRouter@http://localhost:8080/src/pinia/modules/router.js:191:5\nasync*wrappedAction@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:1029:14\npatchActionForGrouping/store[actionName]@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:716:40\nLoginIn@http://localhost:8080/src/pinia/modules/user.js:81:25\nasync*wrappedAction@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:1029:14\npatchActionForGrouping/store[actionName]@http://localhost:8080/node_modules/.vite/deps/pinia.js?v=44ac4499:716:40\nlogin@http://localhost:8080/src/view/login/index.vue:103:28\nsetup/submitForm/<@http://localhost:8080/src/view/login/index.vue:118:26\nvalidateField@http://localhost:8080/node_modules/.vite/deps/es-Dg0EwXc_.js?v=da719756:10752:32\nasync*validate@http://localhost:8080/node_modules/.vite/deps/es-Dg0EwXc_.js?v=da719756:10729:53\nsubmitForm@http://localhost:8080/src/view/login/index.vue:106:21\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\nemit@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4188:41\ncreateSetupContext/get emit/<@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:6186:41\nhandleClick@http://localhost:8080/node_modules/.vite/deps/es-Dg0EwXc_.js?v=da719756:16145:7\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\ninvoker@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:7939:29\nEventListener.handleEvent*addEventListener@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:7905:5\npatchEvent@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:7916:34\npatchProp@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:8374:41\nmountElement@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4796:87\nprocessElement@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4775:31\npatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4721:46\ncomponentUpdateFn@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4995:11\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\nsetupRenderEffect@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:5063:3\nmountComponent@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4947:27\nprocessComponent@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4927:22\npatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4722:44\nmountChildren@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4834:54\nprocessFragment@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4918:17\npatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4719:20\nmountChildren@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4834:54\nmountElement@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4792:41\nprocessElement@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4775:31\npatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4721:46\nmountChildren@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4834:54\nmountElement@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4792:41\nprocessElement@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4775:31\npatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4721:46\ncomponentUpdateFn@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4995:11\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\nsetupRenderEffect@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:5063:3\nmountComponent@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4947:27\nprocessComponent@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4927:22\npatch@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4722:44\nmountChildren@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:4834:54\n','error',NULL,'未处理'),(9,'2026-05-25 17:47:37.015','2026-05-25 17:47:37.015',NULL,'前端','错误信息: TypeError: can\'t access property \"forEach\", asyncRouters.value[0].children is undefined\nStack: 调用栈: useRouterStore</<@http://localhost:8080/src/pinia/modules/router.js:145:5\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:17\ncallWithAsyncErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2092:36\ndoWatch/baseWatchOptions.call@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2538:72\ngetter@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1209:18\nrun@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1476:17\njob@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:1256:17\ncallWithErrorHandling@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2085:31\nflushJobs@http://localhost:8080/node_modules/.vite/deps/runtime-dom.esm-bundler-CtfAi8Ci.js?v=da719756:2225:26\n','error',NULL,'未处理'),(10,'2026-05-25 17:53:44.492','2026-05-25 17:53:44.492',NULL,'后端','获取失败! | 错误: record not found \n 源文件:/home/v3/workspace/gin-vue-admin/server/api/v1/system/sys_user.go:487 \n 调用栈：github.com/flipped-aurora/gin-vue-admin/server/api/v1/system.(*BaseApi).GetUserInfo\n	/home/v3/workspace/gin-vue-admin/server/api/v1/system/sys_user.go:487\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/flipped-aurora/gin-vue-admin/server/initialize.Routers.CasbinHandler.func4\n	/home/v3/workspace/gin-vue-admin/server/middleware/casbin_rbac.go:30\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/flipped-aurora/gin-vue-admin/server/initialize.Routers.JWTAuth.func3\n	/home/v3/workspace/gin-vue-admin/server/middleware/jwt.go:69\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/flipped-aurora/gin-vue-admin/server/initialize.Routers.GinRecovery.func2\n	/home/v3/workspace/gin-vue-admin/server/middleware/error.go:78\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:633\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:589\nnet/http.serverHandler.ServeHTTP\n	/snap/go/11127/src/net/http/server.go:3311\nnet/http.(*conn).serve\n	/snap/go/11127/src/net/http/server.go:2073 \n 最终调用方法:/home/v3/workspace/gin-vue-admin/server/api/v1/system/sys_user.go:487 (GetUserInfo lines 483-492)\n----- 产生日志的方法代码如下 -----\nfunc (b *BaseApi) GetUserInfo(c *gin.Context) {\n	uuid := utils.GetUserUuid(c)\n	ReqUser, err := userService.GetUserInfo(uuid)\n	if err != nil {\n		global.GVA_LOG.Error(\"获取失败!\", zap.Error(err))\n		response.FailWithMessage(\"获取失败\", c)\n		return\n	}\n	response.OkWithDetailed(gin.H{\"userInfo\": ReqUser}, \"获取成功\", c)\n}','error',NULL,'未处理'),(11,'2026-05-25 17:53:47.425','2026-05-25 17:53:47.425',NULL,'后端','获取失败! | 错误: record not found \n 源文件:/home/v3/workspace/gin-vue-admin/server/api/v1/system/sys_user.go:487 \n 调用栈：github.com/flipped-aurora/gin-vue-admin/server/api/v1/system.(*BaseApi).GetUserInfo\n	/home/v3/workspace/gin-vue-admin/server/api/v1/system/sys_user.go:487\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/flipped-aurora/gin-vue-admin/server/initialize.Routers.CasbinHandler.func4\n	/home/v3/workspace/gin-vue-admin/server/middleware/casbin_rbac.go:30\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/flipped-aurora/gin-vue-admin/server/initialize.Routers.JWTAuth.func3\n	/home/v3/workspace/gin-vue-admin/server/middleware/jwt.go:69\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.LoggerWithConfig.func1\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/logger.go:249\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/flipped-aurora/gin-vue-admin/server/initialize.Routers.GinRecovery.func2\n	/home/v3/workspace/gin-vue-admin/server/middleware/error.go:78\ngithub.com/gin-gonic/gin.(*Context).Next\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/context.go:185\ngithub.com/gin-gonic/gin.(*Engine).handleHTTPRequest\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:633\ngithub.com/gin-gonic/gin.(*Engine).ServeHTTP\n	/home/v3/go/pkg/mod/github.com/gin-gonic/gin@v1.10.0/gin.go:589\nnet/http.serverHandler.ServeHTTP\n	/snap/go/11127/src/net/http/server.go:3311\nnet/http.(*conn).serve\n	/snap/go/11127/src/net/http/server.go:2073 \n 最终调用方法:/home/v3/workspace/gin-vue-admin/server/api/v1/system/sys_user.go:487 (GetUserInfo lines 483-492)\n----- 产生日志的方法代码如下 -----\nfunc (b *BaseApi) GetUserInfo(c *gin.Context) {\n	uuid := utils.GetUserUuid(c)\n	ReqUser, err := userService.GetUserInfo(uuid)\n	if err != nil {\n		global.GVA_LOG.Error(\"获取失败!\", zap.Error(err))\n		response.FailWithMessage(\"获取失败\", c)\n		return\n	}\n	response.OkWithDetailed(gin.H{\"userInfo\": ReqUser}, \"获取成功\", c)\n}','error',NULL,'未处理');
/*!40000 ALTER TABLE `sys_error` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_export_template_condition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_template_condition` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_condition_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_export_template_condition` WRITE;
/*!40000 ALTER TABLE `sys_export_template_condition` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_template_condition` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_export_template_join`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_template_join` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `joins` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联',
  `table` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联表',
  `on` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '关联条件',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_join_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_export_template_join` WRITE;
/*!40000 ALTER TABLE `sys_export_template_join` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_template_join` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_export_templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_export_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `db_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '模板标识',
  `template_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `sql` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '自定义导出SQL',
  `import_sql` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '自定义导入SQL',
  `limit` bigint DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_export_templates` WRITE;
/*!40000 ALTER TABLE `sys_export_templates` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_export_templates` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_ignore_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_ignore_apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_ignore_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_ignore_apis` WRITE;
/*!40000 ALTER TABLE `sys_ignore_apis` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_ignore_apis` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_login_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_login_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名',
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `status` tinyint(1) DEFAULT NULL COMMENT '登录状态',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '代理',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_login_logs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_login_logs` WRITE;
/*!40000 ALTER TABLE `sys_login_logs` DISABLE KEYS */;
INSERT INTO `sys_login_logs` VALUES (1,'2026-05-25 17:45:16.383','2026-05-25 17:45:16.383',NULL,'admin','::1',0,'验证码错误','curl/8.5.0',0),(2,'2026-05-25 17:46:41.363','2026-05-25 17:46:41.363',NULL,'admin','127.0.0.1',1,'登录成功','Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:150.0) Gecko/20100101 Firefox/150.0',1),(3,'2026-05-25 17:49:37.363','2026-05-25 17:49:37.363',NULL,'admin','127.0.0.1',0,'验证码错误','Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:150.0) Gecko/20100101 Firefox/150.0',0),(4,'2026-05-25 17:49:44.824','2026-05-25 17:49:44.824',NULL,'admin','127.0.0.1',1,'登录成功','Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:150.0) Gecko/20100101 Firefox/150.0',1),(5,'2026-05-25 17:53:59.632','2026-05-25 17:53:59.632',NULL,'admin','127.0.0.1',1,'登录成功','Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:150.0) Gecko/20100101 Firefox/150.0',3);
/*!40000 ALTER TABLE `sys_login_logs` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_operation_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_operation_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint DEFAULT NULL COMMENT '请求状态',
  `latency` bigint DEFAULT NULL COMMENT '延迟',
  `agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_operation_records` WRITE;
/*!40000 ALTER TABLE `sys_operation_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_operation_records` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_params`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_params` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数名称',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数键',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数值',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '参数说明',
  PRIMARY KEY (`id`),
  KEY `idx_sys_params_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_params` WRITE;
/*!40000 ALTER TABLE `sys_params` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_params` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_user_authority`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint unsigned NOT NULL,
  `sys_authority_authority_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_user_authority` WRITE;
/*!40000 ALTER TABLE `sys_user_authority` DISABLE KEYS */;
INSERT INTO `sys_user_authority` VALUES (3,888),(3,8881),(3,9528),(4,888);
/*!40000 ALTER TABLE `sys_user_authority` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `authority_id` bigint unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  `origin_setting` json DEFAULT NULL COMMENT '配置',
  `mfa_secret` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'MFA密钥(base32编码)',
  `mfa_bound` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'MFA是否已绑定',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (3,'2026-05-25 17:53:02.100','2026-05-25 17:53:02.101',NULL,'b94ed65c-c1c0-4ea6-b8eb-18323c3abb02','admin','$2a$10$MRlsLRotONzJqBRpjlIfw.2hWP6JhtIZeDfZFe4RvrO/0bDI/ZDKO','Mr.奇淼','https://qmplusimg.henrongyi.top/gva_header.jpg',888,'17611111111','333333333@qq.com',1,NULL,'',0),(4,'2026-05-25 17:53:02.100','2026-05-25 17:53:02.103',NULL,'a6593696-b22c-4ed2-9da0-1be09aa38c1e','a303176530','$2a$10$MRlsLRotONzJqBRpjlIfw.2hWP6JhtIZeDfZFe4RvrO/0bDI/ZDKO','用户1','https://qmplusimg.henrongyi.top/1572075907logo.png',9528,'17611111111','333333333@qq.com',1,NULL,'',0);
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `sys_versions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_versions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `version_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '版本名称',
  `version_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '版本号',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '版本描述',
  `version_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '版本数据JSON',
  PRIMARY KEY (`id`),
  KEY `idx_sys_versions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `sys_versions` WRITE;
/*!40000 ALTER TABLE `sys_versions` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_versions` ENABLE KEYS */;
UNLOCK TABLES;

USE `paopao`;
/*!50001 DROP VIEW IF EXISTS `p_post_by_comment`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `p_post_by_comment` AS select `P`.`id` AS `id`,`P`.`user_id` AS `user_id`,`P`.`comment_count` AS `comment_count`,`P`.`collection_count` AS `collection_count`,`P`.`upvote_count` AS `upvote_count`,`P`.`share_count` AS `share_count`,`P`.`visibility` AS `visibility`,`P`.`is_top` AS `is_top`,`P`.`is_essence` AS `is_essence`,`P`.`is_lock` AS `is_lock`,`P`.`latest_replied_on` AS `latest_replied_on`,`P`.`tags` AS `tags`,`P`.`attachment_price` AS `attachment_price`,`P`.`ip` AS `ip`,`P`.`ip_loc` AS `ip_loc`,`P`.`created_on` AS `created_on`,`P`.`modified_on` AS `modified_on`,`P`.`deleted_on` AS `deleted_on`,`P`.`is_del` AS `is_del`,`C`.`user_id` AS `comment_user_id` from ((select `p_comment`.`post_id` AS `post_id`,`p_comment`.`user_id` AS `user_id` from `p_comment` where (`p_comment`.`is_del` = 0) union select `COMMENT`.`post_id` AS `post_id`,`reply`.`user_id` AS `user_id` from (`p_comment_reply` `reply` join `p_comment` `COMMENT` on((`reply`.`comment_id` = `COMMENT`.`id`))) where ((`reply`.`is_del` = 0) and (`COMMENT`.`is_del` = 0))) `C` join `p_post` `P` on((`C`.`post_id` = `P`.`id`))) where (`P`.`is_del` = 0) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!50001 DROP VIEW IF EXISTS `p_post_by_media`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `p_post_by_media` AS select `post`.`id` AS `id`,`post`.`user_id` AS `user_id`,`post`.`comment_count` AS `comment_count`,`post`.`collection_count` AS `collection_count`,`post`.`upvote_count` AS `upvote_count`,`post`.`share_count` AS `share_count`,`post`.`visibility` AS `visibility`,`post`.`is_top` AS `is_top`,`post`.`is_essence` AS `is_essence`,`post`.`is_lock` AS `is_lock`,`post`.`latest_replied_on` AS `latest_replied_on`,`post`.`tags` AS `tags`,`post`.`attachment_price` AS `attachment_price`,`post`.`ip` AS `ip`,`post`.`ip_loc` AS `ip_loc`,`post`.`created_on` AS `created_on`,`post`.`modified_on` AS `modified_on`,`post`.`deleted_on` AS `deleted_on`,`post`.`is_del` AS `is_del` from ((select distinct `p_post_content`.`post_id` AS `post_id` from `p_post_content` where (((`p_post_content`.`type` = 3) or (`p_post_content`.`type` = 4) or (`p_post_content`.`type` = 7) or (`p_post_content`.`type` = 8)) and (`p_post_content`.`is_del` = 0))) `media` join `p_post` `post` on((`media`.`post_id` = `post`.`id`))) where (`post`.`is_del` = 0) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!50001 DROP VIEW IF EXISTS `p_user_relation`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `p_user_relation` AS select `p_contact`.`user_id` AS `user_id`,`p_contact`.`friend_id` AS `he_uid`,5 AS `style` from `p_contact` where ((`p_contact`.`status` = 2) and (`p_contact`.`is_del` = 0)) union select `p_following`.`user_id` AS `user_id`,`p_following`.`follow_id` AS `he_uid`,10 AS `style` from `p_following` where (`p_following`.`is_del` = 0) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

