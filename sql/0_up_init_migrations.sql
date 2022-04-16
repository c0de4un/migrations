CREATE TABLE `migrations` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `index` INT NOT NULL UNIQUE,
    `name` VARCHAR(128) NOT NULL,
    `created_at` DATETIME NOT NULL
)