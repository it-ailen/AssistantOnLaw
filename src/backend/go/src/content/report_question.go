package content

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"log"
)

func (self *Manager) CreateQuestion(args SqlKV) string {
	id := self.AllocateId(true)
    args["id"] = id
	if triggerBy, ok := args["trigger_by"]; ok {
        triggerByJson, _ := json.Marshal(triggerBy)
		args["trigger_by"] = string(triggerByJson)
	}
	optionsJson, _ := json.Marshal(args["options"])
	args["options"] = string(optionsJson)

	cols, placeholders, sqlArgs := args.Insert()
	s := fmt.Sprintf("INSERT INTO `report_question`(%s) VALUES(%s)",
		cols, placeholders)
    log.Printf("sql: %s args: %s", s, sqlArgs)
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(sqlArgs...)
	if err != nil {
		panic(err)
	}
	return id
}

func (self *Manager) UpdateQuestion(id string, toUpdate SqlKV) {
	if triggerBy, ok := toUpdate["trigger_by"]; ok {
        triggerByJson, _ := json.Marshal(triggerBy)
		toUpdate["trigger_by"] = string(triggerByJson)
	}
    if options, ok := toUpdate["options"]; ok {
        optionJson, _ := json.Marshal(options)
		toUpdate["options"] = optionJson
    }
	cols, args := toUpdate.Update()
	args = append(args, id)
	s := fmt.Sprintf("UPDATE `report_question` SET %s WHERE `id`=? ", cols)
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	log.Printf("sql: %s, %v", s, args)
	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err)
	}
}

func (self *Manager) DeleteQuestion(id string) {
	s := "DELETE FROM `report_question` WHERE `id`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
}

type QuestionFilter struct {
	IDs []string
	EntryIds []string
}

func (self *Manager) SelectQuestions(filter *QuestionFilter) []*Question {
	s := "SELECT `id`, `entry_id`, `question`, `options`, `type`, `trigger_by` " +
		"FROM `report_question` "
	args := make([]interface{}, 0, 1000)
	if filter != nil {
		cols := make([]string, 0, 1000)
		if len(filter.IDs) > 0 {
			placeholders := make([]string, 0, 1000)
			for _, id := range filter.IDs {
				placeholders = append(placeholders, "?")
				args = append(args, id)
			}
			cols = append(cols, fmt.Sprintf("`id` IN (%s)",
				strings.Join(placeholders, ", ")))
		}
		if len(filter.EntryIds) > 0 {
			placeholders := make([]string, 0, 1000)
			for _, id := range filter.EntryIds {
				placeholders = append(placeholders, "?")
				args = append(args, id)
			}
			cols = append(cols, fmt.Sprintf("`entry_id` IN (%s)",
				strings.Join(placeholders, ", ")))
		}
		if len(cols) > 0 {
			s += fmt.Sprintf("WHERE %s ", strings.Join(cols, " AND "))
		}
	}
	s += "ORDER BY `created_time` ASC "
	log.Printf("sql: %s, args: %v", s, args)
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	res := make([]*Question, 0, 1000)
	for rows.Next() {
		question := Question{}
		var optionsJson string
		triggerBy := sql.NullString{}
		err = rows.Scan(&question.ID, &question.EntryId, &question.Question, &optionsJson,
			&question.Type, &triggerBy)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal([]byte(optionsJson), &question.Options)
		if err != nil {
			panic(err)
		}
		if triggerBy.Valid {
			question.TriggerBy = new(QuestionTriggerInfo)
			dbView := make(map[string][]int)
			err = json.Unmarshal([]byte(triggerBy.String), &dbView)
			if err != nil {
				panic(err)
			}
			for questionId, options := range dbView {
				question.TriggerBy.QuestionId = questionId
				question.TriggerBy.Options = options
			}
		}
		res = append(res, &question)
	}
	return res
}

func (self *Manager) SelectQuestion(id string) *Question {
	items := self.SelectQuestions(&QuestionFilter{
		IDs: []string{id},
	})
	if len(items) > 0 {
		return items[0]
	}
	return nil
}
