CREATE TABLE IF NOT EXISTS `users` (
    `id` MEDIUMINT NOT NULL AUTO_INCREMENT,
    `email` varchar(100) not null,
    `email_verified_at` datetime default null,
    `password` varchar(100) default null,
    `created_at` datetime not null default now(),
    `updated_at` datetime not null default now(),
    `deleted_at` datetime default null,
    PRIMARY KEY (id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;