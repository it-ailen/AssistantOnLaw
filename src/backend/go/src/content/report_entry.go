package content

import (
	"fmt"
	"strings"
    "database/sql"
)

func (self *Manager) CreateEntry(entry *Entry) string {
    id := self.AllocateId(true)
	kv := make(SqlKV)
	kv["id"] = id
	kv["name"] = entry.Name
	kv["logo"] = entry.Logo
	kv["layout_type"] = entry.LayoutType
	kv["class_id"] = entry.ClassId
	cols, placeholders, args := kv.Insert()
	s := fmt.Sprintf("INSERT INTO `report_entry`(%s) VALUES(%s)",
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

func (self *Manager) UpdateEntry(id string, toUpdate SqlKV) {
	cols, args := toUpdate.Update()
	s := fmt.Sprintf("UPDATE `report_entry` SET %s WHERE `id`=? ", cols)
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

func (self *Manager) DeleteEntry(id string) {
	s := "DELETE FROM `report_entry` WHERE `id`=? "
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

type EntryFilter struct {
	IDs      []string
	ClassIds []string
}

func (self *Manager) SelectEntries(filter *EntryFilter) []*Entry {
	s := "SELECT `id`, `class_id`, `name`, `logo`, `layout_type` " +
		"FROM `report_entry` "
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
		if len(filter.ClassIds) > 0 {
			placeholders := make([]string, 0, 1000)
			for _, id := range filter.ClassIds {
				placeholders = append(placeholders, "?")
				args = append(args, id)
			}
			cols = append(cols, fmt.Sprintf("`class_id` IN (%s)",
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

    res := make([]*Entry, 0, 1000)
    for rows.Next() {
        entry := Entry{}
        logo := sql.NullString{}
        err = rows.Scan(&entry.ID, &entry.ClassId, &entry.Name, &logo, &entry.LayoutType)
        if err != nil {
            panic(err)
        }
        if logo.Valid {
            entry.Logo = logo.String
        }
        res = append(res, &entry)
    }
    return res
}

func (self *Manager) SelectEntry(id string) *Entry {
    entries := self.SelectEntries(&EntryFilter{
        IDs: []string{id},
    })
    if len(entries) > 0 {
        return entries[0]
    }
    return nil
}
