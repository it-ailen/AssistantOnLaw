package content

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func (self *Manager) CreateConclusion(con *Conclusion) string {
	id := self.AllocateId(true)
	kv := make(SqlKV)
	kv["id"] = id
	kv["title"] = con.Title
	kv["context"] = con.Context

	cols, placeholders, args := kv.Insert()
	s := fmt.Sprintf("INSERT INTO `conclusion`(%s) VALUES(%s)",
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

func (self *Manager) UpdateConclusion(id string, toUpdate SqlKV) {
	cols, args := toUpdate.Update()
	s := fmt.Sprintf("UPDATE `conclusion` SET %s WHERE `id`=? ", cols)
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

func (self *Manager) DeleteConclusion(id string) {
	s := "DELETE FROM `conclusion` WHERE `id`=? "
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

type ConclusionFilter struct {
	IDs []string
}

func (self *Manager) SelectConclusions(filter *ConclusionFilter) []*Conclusion {
	s := "SELECT `id`, `title`, `context` " +
		"FROM `conclusion` "
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

	res := make([]*Conclusion, 0, 1000)
	for rows.Next() {
		conclusion := Conclusion{}
		err = rows.Scan(&conclusion.ID, &conclusion.Title, &conclusion.Context)
		if err != nil {
			panic(err)
		}
		res = append(res, &conclusion)
	}
	return res
}

func (self *Manager) SelectConclusion(id string) *Conclusion {
	items := self.SelectConclusions(&ConclusionFilter{
		IDs: []string{id},
	})
	if len(items) > 0 {
		return items[0]
	}
	return nil
}

type Selection struct {
	Ques       *Question
	Selections []int
}

type Selections []Selection

func (slice Selections) Len() int {
	return len(slice)
}

func (slice Selections) Less(i, j int) bool {
	return slice[i].Ques.ID < slice[j].Ques.ID
}

func (slice Selections) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (self *Manager) CalculateSelectionHash(selections Selections) string {
	sort.Sort(selections)
	rows := make([]string, 0, len(selections))
	for _, each := range selections {
		var selectionsString string
		selectionsLen := len(each.Selections)
		if selectionsLen > 0 {
			indexes := make([]string, selectionsLen)
			for _, selection := range each.Selections {
				indexes = append(indexes, strconv.Itoa(selection))
			}
			selectionsString = strings.Join(indexes, ":")
		}
		if len(selectionsString) > 0 {
			rows = append(rows, fmt.Sprintf("[%s:%s]", each.Ques.ID, selectionsString))
		} else {
			rows = append(rows, fmt.Sprintf("[%s]", each.Ques.ID))
		}
	}
	hash := strings.Join(rows, "")
	log.Printf("raw: %s", hash)
	return fmt.Sprintf("%x", md5.Sum([]byte(hash)))
}

func (self *Manager) CalculateConclusion(selections Selections) *Conclusion {
	hashKey := self.CalculateSelectionHash(selections)
	s := "SELECT `id`, `title`, `context` " +
		"FROM `question_conclusion_map` AS `a` " +
		"JOIN `conclusion` AS `b` ON `a`.`conclusion`=`b`.`id` " +
		"WHERE `hash`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	conclusion := Conclusion{}
	err = stmt.QueryRow(hashKey).Scan(&conclusion.ID, &conclusion.Title, &conclusion.Context)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		panic(err)
	}
	return &conclusion
}
