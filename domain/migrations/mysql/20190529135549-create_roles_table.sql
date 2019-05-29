-- +migrate Up

CREATE TABLE IF NOT EXISTS `mark-one`.`roles`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `value`       VARCHAR(45)  NOT NULL,
    `description` TEXT         NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB;

-- +migrate Down

DROP TABLE IF EXISTS `mark-one`.roles;

