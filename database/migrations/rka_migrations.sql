DROP TABLE IF EXISTS `rkas`;
CREATE TABLE `rkas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `akuisisi` int(12) NOT NULL,
  `agen_jawara` int(12) NOT NULL,
  `agen_juragan` int(12) NOT NULL,
  `agen_bep` int(12) NOT NULL,
  `status` enum('active','not-active','deleted') NOT NULL DEFAULT 'active',
  `type` enum('target','position') DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb3;