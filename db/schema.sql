CREATE TABLE user (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARBINARY(32) NOT NULL,
    `password`  VARBINARY(255) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY(id),
    UNIQUE  KEY(name),

    KEY(created_at),
    KEY(updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE entry (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,

    `url`  VARBINARY(2083) NOT NULL,
    `titile` VARCHAR(512) NOT NULL,

    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY (url(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE bookmark (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,

    `user_id` BIGINT UNSIGNED NOT NULL,
    `entry_id` BIGINT UNSIGNED NOT NULL,
    `comment` VARCHAR(256) NOT NULL,

    FOREIGN KEY (user_id)
      REFERENCES user(id),

    FOREIGN KEY (entry_id)
      REFERENCES entry(id),

    PRIMARY KEY(id),
    UNIQUE KEY(user_id, entry_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


