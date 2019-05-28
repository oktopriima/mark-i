-- +migrate Up

CREATE TABLE IF NOT EXISTS `mark-one`.`user_socials`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `users_id`    INT UNSIGNED NOT NULL,
    `social_name` VARCHAR(255) NOT NULL,
    `social_id`   VARCHAR(255) NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_user_social_users1_idx` (`users_id` ASC),
    CONSTRAINT `fk_user_social_users1`
        FOREIGN KEY (`users_id`)
            REFERENCES `mark-one`.`users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;

-- +migrate Down

DROP TABLE IF EXISTS `mark-one`.`user_socials`;