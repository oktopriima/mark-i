-- +migrate Up

CREATE TABLE IF NOT EXISTS `mark-one`.`users`
(
    `id`              INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`            VARCHAR(255) NOT NULL,
    `username`        VARCHAR(255) NOT NULL,
    `email`           VARCHAR(255) NOT NULL,
    `phone`           VARCHAR(255) NOT NULL,
    `password`        VARCHAR(255) NOT NULL,
    `last_login`      TIMESTAMP    NOT NULL,
    `secondary_email` VARCHAR(255) NOT NULL,
    `is_verified`     TINYINT      NOT NULL,
    `avatar`          VARCHAR(255) NOT NULL,
    `created_at`      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB;

-- +migrate Down

DROP TABLE IF EXISTS `mark-one`.`users`;
