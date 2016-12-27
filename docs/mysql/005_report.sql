CREATE TABLE `report_class` (
  `id` CHAR(32) PRIMARY KEY,
  `name` CHAR(64) NOT NULL UNIQUE,
  `description` TEXT NOT NULL,
  `logo` VARCHAR(200),
  `bg` VARCHAR(200)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `report_entry` (
  `id` CHAR(32) PRIMARY KEY,
  `class_id` CHAR(32) NOT NULL,
  `name` CHAR(64) NOT NULL,
  `logo` VARCHAR(200),
  `layout_type` CHAR(40) NOT NULL, /* single|multiple */
  FOREIGN KEY `class`(`class_id`) REFERENCES `report_class`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `report_question` (
  `id` CHAR(32) PRIMARY KEY,
  `entry_id` CHAR(32) NOT NULL,
  `question` CHAR(64) NOT NULL,
  `options` JSON NOT NULL,
  `type` CHAR(40) NOT NULL, /* single|multiple */
  `trigger_by` CHAR(32), /* hide until triggered by other question */
  FOREIGN KEY `entry`(`entry_id`) REFERENCES `report_entry`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `conclusion` (
  `id` CHAR(32) PRIMARY KEY,
  `title` VARCHAR(200) NOT NULL,
  `context` TEXT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `question_conclusion_map` (
  `hash` CHAR(40) PRIMARY KEY,
  `conclusion_id` CHAR(32),
  FOREIGN KEY `conclusion`(`conclusion_id`) REFERENCES `conclusion`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
