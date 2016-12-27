package content

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

func (self *Manager) CreateQuestion(entry *Question) string {
	id := self.AllocateId(true)
	kv := make(SqlKV)
	kv["id"] = id
	kv["entry_id"] = entry.EntryId
	kv["question"] = entry.Question
	kv["type"] = entry.Type
	if len(entry.TriggerBy) > 0 {
		kv["trigger_by"] = entry.TriggerBy
	}
	optionsJson, _ := json.Marshal(entry.Options)
	kv["options"] = optionsJson

	cols, placeholders, args := kv.Insert()
	s := fmt.Sprintf("INSERT INTO `report_question`(%s) VALUES(%s)",
		cols, placeholders)
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err)
	}
	return id
}

func (self *Manager) UpdateQuestion(id string, toUpdate SqlKV) {
	cols, args := toUpdate.Update()
	s := fmt.Sprintf("UPDATE `report_question` SET %s WHERE `id`=? ", cols)
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

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
			for _, id := range filter.IDs {
				placeholders = append(placeholders, "?")
				args = append(args, id)
			}
			cols = append(cols, fmt.Sprintf("`entry_id` IN (%s)",
				strings.Join(placeholders, ", ")))
		}
		if len(cols) > 0 {
			s += fmt.Sprint("WHERE %s ", strings.Join(cols, " AND "))
		}
	}
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
		if triggerBy.Valid {
			question.TriggerBy = triggerBy.String
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
