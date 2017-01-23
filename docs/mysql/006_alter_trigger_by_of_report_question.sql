ALTER TABLE `report_question` CHANGE `trigger_by` `trigger_by` JSON;
ALTER TABLE `report_question` ADD COLUMN `created_time` DATETIME DEFAULT now();