package content

import (
	"database/sql"
	"fmt"
	"strings"
)

type Class struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type Entry struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	LayoutType string `json:"layout_type"`
	ClassId    string `json:"class_id"`
}

type QuestionTriggerInfo struct {
	QuestionId string `json:"question_id,omitempty"`
	Options    []uint `json:"options,omitempty"`
}

type Question struct {
	ID        string   `json:"id"`
	Question  string   `json:"question"`
	Type      string   `json:"type"`
	Options   []string `json:"options,omitempty"`
	TriggerBy *QuestionTriggerInfo `json:"trigger_by,omitempty"`
	EntryId string `json:"entry_id"`
}

type Conclusion struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Context string `json:"context"`
}

type SqlKV map[string]interface{}

func (self SqlKV) Insert() (string, string, []interface{}) {
	colsLen := len(self)
	cols := make([]string, 0, colsLen)
	placeholders := make([]string, 0, colsLen)
	args := make([]interface{}, 0, colsLen)
	for k, v := range self {
		cols = append(cols, fmt.Sprintf("`%s`", k))
		placeholders = append(placeholders, "?")
		args = append(args, v)
	}
	return strings.Join(cols, ", "), strings.Join(placeholders, ", "), args
}

func (self SqlKV) Update() (string, []interface{}) {
	colsLen := len(self)
	cols := make([]string, 0, colsLen)
	args := make([]interface{}, 0, colsLen)
	for k, v := range self {
		cols = append(cols, fmt.Sprintf("`%s`=?", k))
		args = append(args, v)
	}
	return strings.Join(cols, ", "), args
}

func (self *Manager) CreateClass(cls *Class) string {
	id := self.AllocateId(true)
	kv := make(SqlKV)
	kv["id"] = id
	kv["name"] = cls.Name
	kv["description"] = cls.Description
	kv["logo"] = cls.Logo
	cols, placeholders, args := kv.Insert()
	s := fmt.Sprintf("INSERT INTO `report_class`(%s) VALUES(%s)",
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

func (self *Manager) UpdateClass(id string, toUpdate SqlKV) {
	cols, args := toUpdate.Update()
	s := fmt.Sprintf("UPDATE `report_class` SET %s WHERE `id`=? ", cols)
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

func (self *Manager) DeleteClass(id string) {
	s := "DELETE FROM `report_class` WHERE `id`=? "
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

type ClassFilter struct {
	IDs []string
}

func (self *Manager) SelectClasses(filter *ClassFilter) []*Class {
	s := "SELECT `id`, `name`, `description`, `logo`, `bg` FROM " +
		"`report_class` "
	args := make([]interface{}, 0, 1000)
	if filter != nil {
		cols := make([]string, 0, 10)
		if len(filter.IDs) > 0 {
			placeholders := make([]string, 0, 1000)
			for _, id := range filter.IDs {
				placeholders = append(placeholders, "?")
				args = append(args, id)
			}
			cols = append(cols, fmt.Sprintf("`id` IN (%s)",
				strings.Join(placeholders, ", ")))
		}
		if len(cols) > 0 {
			s += fmt.Sprintf("WHERE %s", strings.Join(cols, " AND "))
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

	classes := make([]*Class, 0, 1000)
	for rows.Next() {
		cls := Class{}
		logo := sql.NullString{}
		bg := sql.NullString{}
		err = rows.Scan(&cls.ID, &cls.Name, &cls.Description, &logo, &bg)
		if err != nil {
			panic(err)
		}
		if logo.Valid {
			cls.Logo = logo.String
		}
		classes = append(classes, &cls)
	}
	return classes
}

func (self *Manager) SelectClass(id string) *Class {
	classes := self.SelectClasses(&ClassFilter{
		IDs: []string{id},
	})
	if len(classes) > 0 {
		return classes[0]
	}
	return nil
}
