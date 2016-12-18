package content

import (
    "content/definition"
    "fmt"
    "strings"
    "log"
)

type FaLvWenDaArticle struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    ClassId     string `json:"class_id"`
    Content     string `json:"content"`

    CreatedTime int64 `json:"created_time"`
    UpdatedTime int64 `json:"updated_time"`
}

type FaLvWenDaClass struct {
    ID          string              `json:"id"`
    Name        string              `json:"name"`
    Articles    []*FaLvWenDaArticle `json:"articles,omitempty"`

    CreatedTime int64 `json:"created_time"`
}

func (self *Manager) CreateFaLvWenDaClass(name string) (*FaLvWenDaClass, *definition.DefinedError) {
    class := FaLvWenDaClass{
        Name:        name,
        ID:          self.AllocateId(true),
        CreatedTime: self.CurrentTimeMs(),
    }
    s := "INSERT INTO `fa_lv_wen_da_class`(`name`, `id`, `created_time`) " +
        "VALUES(?, ?, ?) "
    stmt, err := self.conn.Prepare(s)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(class.Name, class.ID, class.CreatedTime)
    switch {
    case err != nil:
        panic(err)
    }
    return &class, nil
}

func (self *Manager) UpdateFaLvWenDaClass(src *FaLvWenDaClass, name string) *definition.DefinedError {
    s := "UPDATE `fa_lv_wen_da_class` SET `name`=? " +
        "WHERE `id`=? "
    stmt, err := self.conn.Prepare(s)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(name, src.ID)
    switch {
    case err != nil:
        panic(err)
    }
    src.Name = name
    return nil
}

type FaLvWenDaClassesFilter struct {
    ID []string
}

func (self *Manager) LoadFaLvWenDaClasses(filter *FaLvWenDaClassesFilter) ([]*FaLvWenDaClass, *definition.DefinedError) {
    classes := make([]*FaLvWenDaClass, 0, 100)
    func() {
        s := "SELECT `id`, `name`, `created_time` " +
            "FROM `fa_lv_wen_da_class` "
        whereClauses := make([]string, 0, 100)
        args := make([]interface{}, 0, 100)
        if filter != nil {
            if len(filter.ID) > 0 {
                placeHolders := make([]string, 0, 100)
                for _, id := range filter.ID {
                    placeHolders = append(placeHolders, "?")
                    args = append(args, id)
                }
                whereClauses = append(whereClauses,
                    fmt.Sprintf("`id` IN (%s)", strings.Join(placeHolders, ", ")))
            }
        }
        if len(whereClauses) > 0 {
            s += fmt.Sprintf("WHERE %s ", strings.Join(whereClauses, " AND "))
        }
        s += "ORDER BY `created_time` ASC "
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
        for rows.Next() {
            class := FaLvWenDaClass{}
            err = rows.Scan(&class.ID, &class.Name, &class.CreatedTime)
            if err != nil {
                panic(err)
            }
            classes = append(classes, &class)
        }
    }()
    for _, class := range classes {
        articles, err := self.LoadFaLvWenDaArticles(&FaLvWenDaArticlesFilter{
            Class: []string{class.ID},
        })
        if err != nil {
            panic(err)
        }
        class.Articles = articles
    }
    return classes, nil
}

func (self *Manager) DeleteFaLvWenDaClasses(id string) {
    s := "DELETE FROM `fa_lv_wen_da_class` WHERE `id`=? "
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

func (self *Manager) CreateFaLvWenDaArticle(name, classId, content string) (*FaLvWenDaArticle, *definition.DefinedError) {
    now := self.CurrentTimeMs()
    article := FaLvWenDaArticle{
        ID: self.AllocateId(true),
        Name: name,
        ClassId: classId,
        Content: content,
        CreatedTime: now,
        UpdatedTime: now,
    }
    s := "INSERT INTO `fa_lv_wen_da_article`(`id`, `class_id`, `title`, `content`, `created_time`, `updated_time`) " +
        "VALUES(?, ?, ?, ?, ?, ?) "
    stmt, err := self.conn.Prepare(s)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(article.ID, article.ClassId, article.Name, article.Content,
        article.CreatedTime, article.UpdatedTime)
    if err != nil {
        panic(err);
    }
    return &article, nil
}


func (self *Manager) UpdateFaLvWenDaArticle(article *FaLvWenDaArticle, data map[string]interface{}) (error) {
    now := self.CurrentTimeMs()
    s := "UPDATE `fa_lv_wen_da_article` SET "
    updateCols := make([]string, 0, 100)
    args := make([]interface{}, 0, 100)
    for key, value := range data {
        common := true
        switch key {
        case "name":
            common = false
            updateCols = append(updateCols, "`title`=?")
            args = append(args, value)
        case "content":
        default:
            panic(fmt.Sprintf("Invalid column(%s) to update", key))
        }
        if common {
            updateCols = append(updateCols, fmt.Sprintf("`%s`=?", key))
            args = append(args, value)
        }
    }
    if len(updateCols) == 0 {
        log.Println("No columns to update")
        return nil
    }
    updateCols = append(updateCols, "`updated_time`=?")
    args = append(args, now)
    s += fmt.Sprintf("%s ", strings.Join(updateCols, ", "))
    s += "WHERE `id`=? "
    args = append(args, article.ID)

    stmt, err := self.conn.Prepare(s)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        panic(err);
    }
    for key, value := range data {
        switch key {
        case "name":
            article.Name = fmt.Sprintf("%s", value)
        case "content":
            article.Content = fmt.Sprintf("%s", value)
        }
    }
    article.UpdatedTime = now
    return nil
}

type FaLvWenDaArticlesFilter struct {
    Class []string /* class ids */
    ID []string
}

func (self *Manager) LoadFaLvWenDaArticles(filter *FaLvWenDaArticlesFilter) ([]*FaLvWenDaArticle, error) {
    s := "SELECT `id`, `class_id`, `title`, `content`, `created_time`, `updated_time` " +
        "FROM `fa_lv_wen_da_article` "
    whereClauses := make([]string, 0, 100)
    args := make([]interface{}, 0, 100)
    if (filter != nil) {
        if len(filter.Class) > 0 {
            wrappedIds := make([]string, 0, len(filter.Class))
            for _, class := range filter.Class {
                wrappedIds = append(wrappedIds, fmt.Sprintf("'%s'", class))
            }
            whereClauses = append(whereClauses, fmt.Sprintf("`class_id` IN (%s)", strings.Join(wrappedIds, ", ")))
        }
        if len(filter.ID) > 0 {
            wrappedIds := make([]string, 0, len(filter.ID))
            for _, id := range filter.ID {
                wrappedIds = append(wrappedIds, fmt.Sprintf("'%s'", id))
            }
            whereClauses = append(whereClauses, fmt.Sprintf("`id` IN (%s)", strings.Join(wrappedIds, ", ")))
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

    articles := make([]*FaLvWenDaArticle, 0, 1000)
    rows, err := stmt.Query(args...)
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        article := FaLvWenDaArticle{}
        err = rows.Scan(&article.ID, &article.ClassId, &article.Name, &article.Content,
            &article.CreatedTime, &article.UpdatedTime)
        if err != nil {
            panic(err)
        }
        articles = append(articles, &article)
    }
    return articles, nil
}

func (self *Manager) DeleteFaLvWenDaArticle(id string) {
    s := "DELETE FROM `fa_lv_wen_da_article` WHERE `id`=? "
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


