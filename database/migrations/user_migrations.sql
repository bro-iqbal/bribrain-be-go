DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `input_status_kunjungan` enum('Sudah Input','Belum Input') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `input_status_rka` enum('Sudah Input','Belum Input') CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `status` enum('active','not-active','deleted') NOT NULL DEFAULT 'active',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;