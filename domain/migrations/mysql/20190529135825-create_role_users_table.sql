-- +migrate Up

CREATE TABLE IF NOT EXISTS `mark-one`.`user_roles`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `roles_id`   INT UNSIGNED NOT NULL,
    `users_id`   INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_user_roles_role1_idx` (`roles_id` ASC),
    INDEX `fk_user_roles_users2_idx` (`users_id` ASC),
    CONSTRAINT `fk_user_roles_role1`
        FOREIGN KEY (`roles_id`)
            REFERENCES `mark-one`.`roles` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_user_roles_users2`
        FOREIGN KEY (`users_id`)
            REFERENCES `mark-one`.`users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;

-- +migrate Down

DROP TABLE IF EXISTS 'user_roles';