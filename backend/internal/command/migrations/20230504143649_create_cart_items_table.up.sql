CREATE TABLE IF NOT EXISTS `cart_items` (
    `cart_id` BIGINT UNSIGNED NOT NULL,
    `product_id` BIGINT UNSIGNED NOT NULL,
    `id` BIGINT UNSIGNED NOT NULL,
    `amount` SMALLINT UNSIGNED NOT NULL DEFAULT 0,
    `selects` json not null,
    `customs` json not null,
    `created_at` datetime not null default CURRENT_TIMESTAMP,
    `updated_at` datetime not null default CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;