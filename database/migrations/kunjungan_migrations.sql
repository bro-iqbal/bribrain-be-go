DROP TABLE IF EXISTS `kunjungans`;
CREATE TABLE `kunjungans` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `akuisisi` int(12) NOT NULL,
  `agen_jawara` int(12) NOT NULL,
  `agen_juragan` int(12) NOT NULL,
  `agen_bep` int(12) NOT NULL,
  `status` enum('active','not-active','deleted') NOT NULL DEFAULT 'active',
  `user_id` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;