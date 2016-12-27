package content

import (
	"fmt"
	"log"
	"strings"
    "database/sql"
)

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URI  string `json:"uri"`
}

type SuSongWenShu struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Flow        string  `json:"flow"`
	Files       []*File `json:"files"`
}

type SuSongWenShuFilter struct {
	Flow string
	ID   []string
}

func (self *Manager) LoadSuSongWenShu(filter *SuSongWenShuFilter) []*SuSongWenShu {
	items := func() []*SuSongWenShu {
		s := "SELECT `id`, `name`, `description`, `flow` " +
			"FROM `su_song_wen_shu` "
		whereClauses := make([]string, 0, 100)
		args := make([]interface{}, 0, 100)
		if filter != nil {
			if len(filter.Flow) > 0 {
				whereClauses = append(whereClauses, "`flow`=?")
				args = append(args, filter.Flow)
			}
			if len(filter.ID) > 0 {
				placeholders := make([]string, 0, len(filter.ID))
				for _, id := range filter.ID {
					placeholders = append(placeholders, "?")
					args = append(args, id)
				}
				whereClauses = append(whereClauses, fmt.Sprintf("`id` IN (%s)", strings.Join(placeholders, ", ")))
			}
		}
		if len(whereClauses) > 0 {
			s += fmt.Sprintf("WHERE %s ", strings.Join(whereClauses, " AND "))
		}

		log.Printf("SQL: %s", s)

		stmt, err := self.conn.Prepare(s)
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		result := make([]*SuSongWenShu, 0, 1000)
		rows, err := stmt.Query(args...)
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			detail := SuSongWenShu{}
			err = rows.Scan(&detail.ID, &detail.Name, &detail.Description, &detail.Flow)
			if err != nil {
				panic(err)
			}
			result = append(result, &detail)
		}
		return result
	}()
	if len(items) > 0 {
		s := "SELECT `name`, `url`, `id`, `wen_shu_id` " +
			"FROM `wen_shu_file` AS `a` " +
			"JOIN `wen_shu_file_map` AS `b` ON `b`.`file_id`=`a`.`id`  " +
			"WHERE `wen_shu_id` IN (%s) "
		placeHolders := make([]string, 0, len(items))
		args := make([]interface{}, 0, len(items))
		for _, item := range items {
			placeHolders = append(placeHolders, "?")
			args = append(args, item.ID)
		}
		s = fmt.Sprintf(s, strings.Join(placeHolders, ", "))

		log.Printf("SQL: %s", s)

		stmt, err := self.conn.Prepare(s)
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		wenShuFileMap := make(map[string][]*File)
		rows, err := stmt.Query(args...)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			f := File{}
			var wenShuId string
			err = rows.Scan(&f.Name, &f.URI, &f.ID, &wenShuId)
			if err != nil {
				panic(err)
			}
			if _, ok := wenShuFileMap[wenShuId]; !ok {
				wenShuFileMap[wenShuId] = make([]*File, 0, 1000)
			}
			wenShuFileMap[wenShuId] = append(wenShuFileMap[wenShuId], &f)
		}

		for _, item := range items {
			item.Files, _ = wenShuFileMap[item.ID]
		}
	}
	return items
}

func (self *Manager) CreateSuSongWenShu(flow, id string, args map[string]interface{}) {
	s := "INSERT INTO `su_song_wen_shu`(`id`, `name`, `description`, `flow`) " +
		"VALUES(?, ?, ?, ?) "
	tx, err := self.conn.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
	stmt0, err := tx.Prepare(s)
	if err != nil {
		return
	}
	defer stmt0.Close()
	_, err = stmt0.Exec(id, args["name"], args["description"], flow)
	if err != nil {
		return
	}
	if files, ok := args["files"]; ok {
		if fs, ok := files.([]string); ok {
			s := "INSERT INTO `wen_shu_file_map`(`wen_shu_id`, `file_id`) " +
				"VALUES(?, ?) "
			stmt1, err := tx.Prepare(s)
			if err != nil {
				return
			}
			defer stmt1.Close()

			for _, fid := range fs {
				_, err = stmt1.Exec(id, fid)
				if err != nil {
					return
				}
			}
		}
	}
}

func (self *Manager) UpdateSuSongWenShu(id string, args map[string]interface{}) {
	tx, err := self.conn.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
	err = func() error {
		updatingCols := make([]string, 0, 10)
		sqlArgs := make([]interface{}, 0, 10)
		for k, v := range args {
			switch {
			case k == "name":
				updatingCols = append(updatingCols, "`name`=?")
				if name, ok := v.(string); ok {
					sqlArgs = append(sqlArgs, name)
				} else {
					err = fmt.Errorf("Invalid args for `name`(%#v)", name)
				}
			case k == "description":
				updatingCols = append(updatingCols, "`description`=?")
				if description, ok := v.(string); ok {
					sqlArgs = append(sqlArgs, description)
				} else {
					err = fmt.Errorf("Invalid args for `description`(%#v)", description)
				}
			}
		}
		if len(updatingCols) > 0 {
			s := fmt.Sprintf("UPDATE `su_song_wen_shu` SET %s WHERE `id`=? ", strings.Join(updatingCols, ", "))
			sqlArgs = append(sqlArgs, id)
			stmt, e := tx.Prepare(s)
			if e != nil {
				return e
			}
			defer stmt.Close()

			_, e = stmt.Exec(sqlArgs...)
			if e != nil {
				return e
			}
		}
		return nil
	}()
	if err != nil {
		return
	}

	if v, ok := args["files"]; ok {
		if files, ok := v.([]string); ok {
			s := "DELETE FROM `wen_shu_file_map` WHERE `wen_shu_id`=? "
			stmt0, err := tx.Prepare(s)
			if err != nil {
				return
			}
			defer stmt0.Close()

			_, err = stmt0.Exec(id)
			if err != nil {
				return
			}

			s = "INSERT INTO `wen_shu_file_map`(`wen_shu_id`, `file_id`) " +
				"VALUES(?, ?) "
			stmt1, err := tx.Prepare(s)
			if err != nil {
				return
			}
			defer stmt1.Close()

			for _, fileId := range files {
				_, err = stmt1.Exec(id, fileId)
				if err != nil {
					return
				}
			}
		} else {
			err = fmt.Errorf("Invalid value of `files`(%#v)", v)
		}
	}
}

func (self *Manager) CreateSuSongFile(name, uri, stepId string) *File {
	file := File{
		ID:   self.AllocateId(true),
		Name: name,
		URI:  uri,
	}
	tx, err := self.conn.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	err = func() error {
		s := "INSERT INTO `wen_shu_file`(`id`, `name`, `url`) " +
			"VALUES(?, ?, ?) "
		stmt, e := tx.Prepare(s)
		if e != nil {
			return e
		}
		defer stmt.Close()

		_, e = stmt.Exec(file.ID, file.Name, file.URI)
		if e != nil {
			return e
		}
		return nil
	}()
	if err != nil {
		return nil
	}
	err = func() error {
		s := "INSERT INTO `wen_shu_file_map`(`wen_shu_id`, `file_id`) " +
			"VALUES(?, ?) "
		stmt, e := tx.Prepare(s)
		if e != nil {
			return e
		}
		defer stmt.Close()

		_, e = stmt.Exec(stepId, file.ID)
		if e != nil {
			return e
		}
		return nil
	}()
	if err != nil {
		return nil
	}

	return &file
}

func (self *Manager) DeleteSuSongFile(id string) {
	s := "DELETE FROM `wen_shu_file` " +
		"WHERE `id`=? "
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

func (self *Manager) UpdateSuSongFile(file *File, args map[string]string) {
	tx, err := self.conn.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()

	err = func() error {
		updatingCols := make([]string, 0, 100)
		updatingArgs := make([]interface{}, 0, 100)
		for k, v := range args {
			switch {
			case k == "name":
				updatingCols = append(updatingCols, "`name`=?")
				updatingArgs = append(updatingArgs, v)
			case k == "uri":
				updatingCols = append(updatingCols, "`url`=?")
				updatingArgs = append(updatingArgs, v)
			}
		}
        updatingArgs = append(updatingArgs, file.ID)
		if len(updatingCols) > 0 {
            s := fmt.Sprintf("UPDATE `wen_shu_file` SET %s WHERE `id`=? ", strings.Join(updatingCols, ", "))
            stmt, e := tx.Prepare(s)
            if e != nil {
                return e
            }
            defer stmt.Close()

            _, e = stmt.Exec(updatingArgs...)
            if e != nil {
                return e
            }
		}
		return nil
	}()
	if err != nil {
		panic(err)
	}
}

func (self *Manager) LoadSuSongFile(id string) *File {
    s := "SELECT `id`, `name`, `url` FROM `wen_shu_file` " +
        "WHERE `id`=? "
    stmt, err := self.conn.Prepare(s)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    row := stmt.QueryRow(id)
    file := File{}
    err = row.Scan(&file.ID, &file.Name, &file.URI)
    switch {
    case err == sql.ErrNoRows:
        return nil
    case err != nil:
        panic(err)
    }
    return &file
}
