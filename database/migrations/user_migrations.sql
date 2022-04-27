CREATE TABLE
IF NOT EXISTS `users`
(
    `id` INT (11) NOT NULL,
    `name` VARCHAR (255) NOT NULL,
    `input_status` ENUM ('Sudah Input', 'Belum Input') NOT NULL,
    `status` ENUM ('active', 'not-active', 'deleted') NOT NULL DEFAULT 'active',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime,
    `deleted_at` datetime
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
