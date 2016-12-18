CREATE TABLE `su_song_wen_shu` (
    `id` CHAR(32) PRIMARY KEY,
    `name` VARCHAR(200) NOT NULL,
    `description` TEXT NOT NULL,
    `flow` CHAR(40) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `wen_shu_file` (
    `id` CHAR(32) PRIMARY KEY,
    `name` VARCHAR(200) NOT NULL,
    `url` VARCHAR(250) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `wen_shu_file_map` (
    `wen_shu_id` CHAR(32),
    `file_id` CHAR(32),
    PRIMARY KEY(`wen_shu_id`, `file_id`),
    FOREIGN KEY `wen_shu`(`wen_shu_id`) REFERENCES `su_song_wen_shu`(`id`) ON DELETE CASCADE,
    FOREIGN KEY `wen_shu_file`(`file_id`) REFERENCES `wen_shu_file`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

