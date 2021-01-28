CREATE TABLE `tags` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
  UNIQUE KEY `tag_name` (`name`)
);

CREATE TABLE `pets` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `tag_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `pet_name` (`name`),
  KEY `tag_id` (`tag_id`),
  CONSTRAINT `pets_ibfk_1` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
);
