CREATE TABLE IF NOT EXISTS `select_options` (
    `select_question_id` BIGINT UNSIGNED NOT NULL,
    `id` BIGINT UNSIGNED NOT NULL,
    `name` varchar(100) not null,
    `price` MEDIUMINT UNSIGNED NOT NULL default 0,
    `created_at` datetime not null default CURRENT_TIMESTAMP,
    `updated_at` datetime not null default CURRENT_TIMESTAMP,
    `deleted_at` datetime default null,
    PRIMARY KEY (id),
    FOREIGN KEY (select_question_id) REFERENCES select_questions(id)
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;