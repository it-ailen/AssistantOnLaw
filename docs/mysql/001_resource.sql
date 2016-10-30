CREATE TABLE `channel` (
  `id` CHAR(32) PRIMARY KEY,
  `name` CHAR(48) UNIQUE NOT NULL,
  `icon` CHAR(100) NOT NULL,
  `deleted` INT(1) DEFAULT 0 NOT NULL,
  `created_time` LONG NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `entry` (
  `id` CHAR(32) PRIMARY KEY,
  `channel_id` CHAR(32) NOT NULL,
  `text` CHAR(48) NOT NULL,
  `layout_type` CHAR(20) NOT NULL DEFAULT 'single-page' /* 'single'|'multiple'*/
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `option` (
  `id` CHAR(32) PRIMARY KEY,
  `parent_id` CHAR(32) NOT NULL, /* id of the parent entry or option */
  `text` CHAR(48),  /* null for report */
  `type` CHAR(10) NOT NULL DEFAULT 'option', /* option|report */
  `ref_id` CHAR(32), /* fk to report */
  INDEX `parent_index` (`parent_id`),
  INDEX `report_index` (`ref_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `decree` (
  `id` CHAR(32) PRIMARY KEY,
  `source` TEXT NOT NULL,
  `content` TEXT NOT NULL,
  `link` TEXT /* Link to the reference page, optional */,
  `deleted` INT(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `case` (
  `id` CHAR(32) PRIMARY KEY,
  `content` TEXT NOT NULL,
  `link` TEXT /* Link to the reference page, optional */,
  `deleted` INT(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `report` (
  `id` CHAR(32) PRIMARY KEY,
  `title` CHAR(64) NOT NULL,
  `conclusion` TEXT NOT NULL,
  `deleted` INT(1) DEFAULT 0,
  `cases` JSON,
  `decrees` JSON
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
