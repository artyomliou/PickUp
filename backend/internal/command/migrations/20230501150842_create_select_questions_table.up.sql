CREATE TABLE IF NOT EXISTS `select_questions` (
    `store_id` BIGINT UNSIGNED NOT NULL,
    `id` BIGINT UNSIGNED NOT NULL,
    `name` varchar(100) not null,
    `is_required` boolean default false,
    `at_most` SMALLINT UNSIGNED NOT NULL default 0,
    `created_at` datetime not null default CURRENT_TIMESTAMP,
    `updated_at` datetime not null default CURRENT_TIMESTAMP,
    `deleted_at` datetime default null,
    PRIMARY KEY (id),
    FOREIGN KEY (store_id) REFERENCES stores(id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
