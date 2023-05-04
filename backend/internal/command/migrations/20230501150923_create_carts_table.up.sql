CREATE TABLE IF NOT EXISTS `carts` (
    `user_id` BIGINT UNSIGNED NOT NULL,
    `store_id` BIGINT UNSIGNED NOT NULL,
    `id` BIGINT UNSIGNED NOT NULL,
    `created_at` datetime not null default CURRENT_TIMESTAMP,
    `updated_at` datetime not null default CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;