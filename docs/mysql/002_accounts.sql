CREATE TABLE `accounts` (
  `id` CHAR(32) PRIMARY KEY,
  `account` CHAR(20) NOT NULL UNIQUE,
  `type` CHAR(20) NOT NULL, /* customer|super */
  `password` CHAR(60),
  `nick` CHAR(40) NULL, /*  */
  `contact` CHAR(40) NULL,
  `etc` JSON
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `accounts`(`id`, `account`, `type`)
VALUES('builtin_root', 'root', 'super');