package content

import (
	"content/definition"
	"database/sql"
	"fmt"
	"strings"
	"log"
)

func (self *Manager) CreateFile(author *definition.Account, name, ref string, parent *definition.Directory) (*definition.File, error) {
	tx, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	current := self.CurrentTimeMs()
	result := definition.File{
		FileNode: definition.FileNode{
			ID:          self.AllocateId(true),
			Name:        name,
			Type:        definition.C_FT_FILE,
			CreatedTime: current,
			UpdatedTime: current,
		},
		Ref: ref,
	}

	err = func() error {
		s := "INSERT INTO `file`(`id`, `name`, `type`, `owner`, `created_time`, `updated_time`, `reference_uri`) " +
			"VALUES(?, ?, ?, ?, ?, ?, ?) "
		stmt, e := tx.Prepare(s)
		if e != nil {
			return e
		}
		defer stmt.Close()

		_, e = stmt.Exec(result.ID, result.Name, result.Type, author.ID,
			result.CreatedTime, result.UpdatedTime, result.Ref)
		if e != nil {
			return e
		}
		return nil
	}()
	if err != nil {
		return nil, err
	}

	err = func() error {
		s := "INSERT INTO `directory`(`directory_id`, `child_id`) " +
			"VALUES(?, ?) "
		stmt, e := tx.Prepare(s)
		if e != nil {
			return e
		}
		defer stmt.Close()

		_, e = stmt.Exec(parent.ID, result.ID)
		return e
	}()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (self *Manager) CreateDirectory(author *definition.Account, name string, parent *definition.Directory) (*definition.Directory, error) {
	tx, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	current := self.CurrentTimeMs()
	result := definition.Directory{
		FileNode: definition.FileNode{
			ID:          self.AllocateId(true),
			Name:        name,
			Type:        definition.C_FT_DIR,
			CreatedTime: current,
			UpdatedTime: current,
		},
	}

	err = func() error {
		s := "INSERT INTO `file`(`id`, `name`, `type`, `owner`, `created_time`, `updated_time`) " +
			"VALUES(?, ?, ?, ?, ?, ?) "
		stmt, e := tx.Prepare(s)
		if e != nil {
			return e
		}
		defer stmt.Close()

		_, e = stmt.Exec(result.ID, result.Name, result.Type, author.ID,
			result.CreatedTime, result.UpdatedTime)
		if e != nil {
			return e
		}
		return nil
	}()
	if err != nil {
		return nil, err
	}

	err = func() error {
		s := "INSERT INTO `directory`(`directory_id`, `child_id`) " +
			"VALUES(?, ?) "
		stmt, e := tx.Prepare(s)
		if e != nil {
			return e
		}
		defer stmt.Close()

		_, e = stmt.Exec(parent.ID, result.ID)
		return e
	}()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (self *Manager) loadChildren(id string) ([]string, error) {
	ids := make([]string, 0, 1000)
	s1 := "SELECT `child_id` " +
		"FROM `directory` " +
		"WHERE `directory_id`=? "
	stmt1, err := self.conn.Prepare(s1)
	if err != nil {
		return nil, err
	}
	defer stmt1.Close()

	rows, err := stmt1.Query(id)
	if err != nil {
		return ids, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (self *Manager) LoadFile(id string) (interface{}, error) {
	s0 := "SELECT `id`, `name`, `type`, `owner`, `created_time`, `updated_time`, `reference_uri` " +
		"FROM `file` WHERE `id`=? "
	stmt0, err := self.conn.Prepare(s0)
	if err != nil {
		return nil, err
	}
	defer stmt0.Close()

	rows, err := stmt0.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	log.Printf("sql: %s", s0)

	if rows.Next() {
		node := definition.FileNode{}
		uri := sql.NullString{}
		err = rows.Scan(&node.ID, &node.Name, &node.Type, &node.Owner, &node.CreatedTime, &node.UpdatedTime, &uri)
		if err != nil {
			return nil, err
		}
		switch node.Type {
		case definition.C_FT_DIR:
			dir := definition.Directory{
				FileNode: node,
			}
			dir.Children, err = self.loadChildren(node.ID)
			if err != nil {
				return nil, err
			}
			return &dir, nil
		case definition.C_FT_FILE:
			file := definition.File{
				FileNode: node,
				Ref:      uri.String,
			}
			return &file, nil
		}
	}
	return nil, nil
}

func (self *Manager) ListChildren(parent *definition.Directory) ([]interface{}, error) {
	children := make([]interface{}, 0, 1000)
	err := func() error {
		s := "SELECT `id`, `name`, `type`, `owner`, `created_time`, `updated_time`, `reference_uri` " +
			"FROM `directory` LEFT JOIN `file` ON `directory`.`child_id` = `file`.`id` " +
			"WHERE `directory_id`=? "
		stmt, err := self.conn.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()

		rows, err := stmt.Query(parent.ID)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			node := definition.FileNode{}
			uri := sql.NullString{}
			err = rows.Scan(&node.ID, &node.Name, &node.Type, &node.Owner, &node.CreatedTime, &node.UpdatedTime, &uri)
			if err != nil {
				return err
			}
			switch node.Type {
			case definition.C_FT_DIR:
				dir := definition.Directory{
					FileNode: node,
				}
				children = append(children, &dir)
			case definition.C_FT_FILE:
				file := definition.File{
					FileNode: node,
					Ref:      uri.String,
				}
				children = append(children, &file)
			}
		}
		return nil
	}()
	if err != nil {
		return children, err
	}

	for _, child := range children {
		if dir, ok := child.(*definition.Directory); ok {
			dir.Children, err = self.loadChildren(dir.ID)
			if err != nil {
				return children, err
			}
		}
	}
	return children, nil
}

func (self *Manager) LoadDirectoryTree(dir *definition.Directory) (*definition.FileTree, error) {
	root := definition.FileTree{}
	root.Current = dir

	children, err := self.ListChildren(dir)
	if err != nil {
		panic(err)
	}
	root.Children = make([]*definition.FileTree, 0, 1000)

	for _, child := range children {
		var tree *definition.FileTree
		if childDir, ok := child.(*definition.Directory); ok {
			tree, err = self.LoadDirectoryTree(childDir)
			if err != nil {
				panic(err)
			}
		} else {
			tree = &definition.FileTree{
				Current: child,
			}
		}
		root.Children = append(root.Children, tree)
	}
	return &root, nil
}

func (self *Manager) UpdateFile(file interface{}, args map[string]string) error {
	s := "UPDATE `file` SET "
	columns := []string{"`updated_time`=? "}
	values := []interface{}{self.CurrentTimeMs()}
	for k, v := range args {
		columns = append(columns, fmt.Sprintf("`%s`=? ", k))
		values = append(values, v)
	}
	s += strings.Join(columns, ", ")
	s += "WHERE `id`=? "
	if f, ok := file.(*definition.File); ok {
		values = append(values, f.ID)
		if name, ok := args["name"]; ok {
			f.Name = name
		}
		if uri, ok := args["reference_uri"]; ok {
			f.Ref = uri
		}
	} else if dir, ok := file.(*definition.Directory); ok {
		values = append(values, dir.ID)
		if name, ok := args["name"]; ok {
			dir.Name = name
		}
	}
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(values...)
	return err
}
