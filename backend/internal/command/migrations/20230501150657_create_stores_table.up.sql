CREATE TABLE IF NOT EXISTS `stores` (
    `id` BIGINT UNSIGNED NOT NULL,
    `name` varchar(100) not null,
    `pic` varchar(100) not null,
    `status` tinyint unsigned not null,
    `opened_at` char(5),
    `closed_at` char(5),
    `created_at` datetime not null default CURRENT_TIMESTAMP,
    `updated_at` datetime not null default CURRENT_TIMESTAMP,
    `deleted_at` datetime default null,
    PRIMARY KEY (id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;