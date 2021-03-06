-- CREATE DATABASE `hasagi`;
USE `hasagi`;

DROP TABLE IF EXISTS `location_histories`;
DROP TABLE IF EXISTS `alert_locations`;

CREATE TABLE IF NOT EXISTS `alert_locations` (
  `id`        BIGINT(20) NOT NULL AUTO_INCREMENT,
  `longitude` FLOAT NOT NULL DEFAULT 0,
  `latitude`  FLOAT NOT NULL DEFAULT 0,
  `user_id`   BIGINT(20) NOT NULL,
  `name`      VARCHAR(255),
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `alert_location_index` (`deleted_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `location_histories` (
  `id`        BIGINT(20) NOT NULL AUTO_INCREMENT,
  `longitude` FLOAT NOT NULL DEFAULT 0,
  `latitude`  FLOAT NOT NULL DEFAULT 0,
  `user_id`   BIGINT(20) NOT NULL,
  `name`      VARCHAR(255),
  `timestamp` TIMESTAMP,
  `created_at` TIMESTAMP,
  `updated_at` TIMESTAMP,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `location_histories_index` (`deleted_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
