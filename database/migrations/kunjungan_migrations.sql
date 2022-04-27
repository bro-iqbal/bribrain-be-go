CREATE TABLE
IF NOT EXISTS `kunjungans`
(
    `id` INT (11) NOT NULL,
    `akuisisi` INT (12) NOT NULL,
    `agen_jawara` INT (12) NOT NULL,
    `agen_juragan` INT (12) NOT NULL,
    `agen_bep` INT (12) NOT NULL,
    `status` ENUM ('active', 'not-active', 'deleted') NOT NULL DEFAULT 'active',
    `user_id` INT (11) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP,
    `deleted_at` TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

ALTER TABLE `kunjungans`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `kunjungans`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
