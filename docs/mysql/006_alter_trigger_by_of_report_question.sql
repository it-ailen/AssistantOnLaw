ALTER TABLE `report_question` CHANGE `trigger_by` `trigger_by` JSON;
/*
  like mongo, example here
  {
    "q1": [0, 1]
    "$or": [
      {"q2": [0]},
      {"q2": [1]},
      {"$and": {
          "q3": [0],
          "q4": [0]
        }
      }
    ]
  } ==> q1:(0|1) && (q2:0 || q2:1 || (q3:0 && q4:0))
 */
ALTER TABLE `report_question` ADD COLUMN `created_time` DATETIME DEFAULT now();