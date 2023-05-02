CREATE TABLE IF NOT EXISTS `custom_question_product` (
    `product_id` BIGINT UNSIGNED NOT NULL,
    `custom_question_id` varchar(100) not null,
    PRIMARY KEY (product_id, custom_question_id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;