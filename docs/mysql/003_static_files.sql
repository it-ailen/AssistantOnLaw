CREATE TABLE `file` (
  `id` CHAR(32) PRIMARY KEY,
  `name` VARCHAR(200) NOT NULL,
  `type` CHAR(40) NOT NULL, /* file|directory */
  `owner` CHAR(32) NOT NULL,
  `created_time` LONG DEFAULT 0,
  `updated_time` LONG DEFAULT 0,
  `reference_uri` VARCHAR(400) NULL,  /* 普通文件需要 */
  `download_count` INT DEFAULT 0,
  `etc` JSON
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `directory` (
  `directory_id` CHAR(32),
  `child_id` CHAR(32),
  PRIMARY KEY(`directory_id`, `child_id`),
  FOREIGN KEY `directory_parent`(`directory_id`) REFERENCES `file`(`id`),
  FOREIGN KEY `directory_child`(`child_id`) REFERENCES `file`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `file`(`id`, `name`, `type`, `owner`)
  VALUES('root', 'root', 'directory', 'builtin_root');

INSERT INTO `file`(`id`, `name`, `type`, `owner`)
  VALUES
  ('xie_yi_fan_ben', '协议范本', 'directory', 'builtin_root')
;

INSERT INTO `directory`(`directory_id`, `child_id`)
  VALUES
  ('root', 'xie_yi_fan_ben')
;

CREATE TABLE `fa_lv_wen_da_class` (
  `id` CHAR(32) PRIMARY KEY,
  `name` VARCHAR (200) UNIQUE NOT NULL,
  `created_time` LONG NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `fa_lv_wen_da_article` (
  `id` CHAR(32) PRIMARY KEY,
  `class_id` CHAR(32) NOT NULL,
  `title` VARCHAR (200) NOT NULL,
  `content` TEXT NOT NULL,
  `created_time` LONG NOT NULL,
  `updated_time` LONG NOT NULL,
  INDEX `title_index`(`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;