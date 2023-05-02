CREATE TABLE IF NOT EXISTS `category_product` (
    `category_id` BIGINT UNSIGNED NOT NULL,
    `product_id` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (category_id, product_id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;