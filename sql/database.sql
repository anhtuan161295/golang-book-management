CREATE DATABASE `weather` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

-- weather.weather definition

CREATE TABLE `weather` (
  `id` int NOT NULL AUTO_INCREMENT,
  `city_id` int NOT NULL,
  `data` text,
  `time` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

