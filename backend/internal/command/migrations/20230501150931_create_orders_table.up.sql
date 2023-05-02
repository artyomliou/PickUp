CREATE TABLE IF NOT EXISTS `orders` (
    `user_id` BIGINT UNSIGNED NOT NULL,
    `store_id` BIGINT UNSIGNED NOT NULL,
    `id` BIGINT UNSIGNED NOT NULL,
    `price` BIGINT UNSIGNED NOT NULL,
    `items` json NOT NULL,
    `status` tinyint unsigned not null,
    `created_at` datetime not null default CURRENT_TIMESTAMP,
    `updated_at` datetime not null default CURRENT_TIMESTAMP,
    `deleted_at` datetime default null,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (store_id) REFERENCES stores(id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;