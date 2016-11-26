package content

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"log"
	"strings"
	"sync"
	"time"
)

type Manager struct {
	conn *sql.DB
}

var inst *Manager
var once sync.Once

func GetManager() *Manager {
	return inst
}

func InitManager(conn *sql.DB) {
	once.Do(func() {
		inst = new(Manager)
		inst.conn = conn
	})
}

func (self *Manager) CurrentTimeMs() int64 {
	return time.Now().UnixNano() / 1000000
}

func (self *Manager) AllocateId(short bool) string {
	id := uuid.NewV4().String()
	if short {
		id = strings.Replace(id, "-", "", -1)
	}
	return id
}

func (self *Manager) ChannelGet(id string) (*Channel, error) {
	s := "SELECT `id`, `name`, `icon`, `deleted`, `created_time` FROM " +
		"`channel` " +
		"WHERE `id`=? "
	selectStmt, err := self.conn.Prepare(s)
	if err != nil {
		return nil, err
	}
	defer selectStmt.Close()
	rows, err := selectStmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		channel := &Channel{}
		err = rows.Scan(&channel.ID, &channel.Name, &channel.Icon,
			&channel.Deleted, &channel.CreatedTime)
		if err != nil {
			return nil, err
		}
		return channel, nil
	}
	return nil, nil
}

type ChannelFilter struct {
	Offset uint
	Count  uint
}

func (self *Manager) ChannelsGet(filter *ChannelFilter) ([]*Channel, error) {
	channels := make([]*Channel, 0, 1024)
	s := "SELECT `id`, `name`, `icon`, `deleted`, `created_time` FROM " +
		"`channel` "
	args := make([]interface{}, 0, 20)
	if filter != nil {
		s += "LIMIT %s, %s"
		args = append(args, filter.Offset, filter.Count)
	}
	selectStmt, err := self.conn.Prepare(s)
	if err != nil {
		return channels, err
	}
	defer selectStmt.Close()
	rows, err := selectStmt.Query(args...)
	if err != nil {
		return channels, err
	}
	defer rows.Close()
	for rows.Next() {
		channel := &Channel{}
		err = rows.Scan(&channel.ID, &channel.Name, &channel.Icon,
			&channel.Deleted, &channel.CreatedTime)
		log.Printf("channel: %#v", channel)
		if err != nil {
			return channels, err
		}
		channels = append(channels, channel)
	}
	return channels, nil
}

func (self *Manager) ChannelCreate(channel *Channel) error {
	tx, err := self.conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			log.Printf("Error(%s)", err.Error())
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	s := "INSERT INTO `channel`(`id`, `name`, `icon`, `created_time`) " +
		"VALUES(?, ?, ?, ?) "
	stat, err := tx.Prepare(s)
	if err != nil {
		return err
	}
	defer stat.Close()
	_, err = stat.Exec(channel.ID, channel.Name, channel.Icon, self.CurrentTimeMs())
	if err != nil {
		return err
	}
	return nil
}

func (self *Manager) ChannelUpdate(channel *Channel) error {
	s := "UPDATE `channel` SET `name`=?, `icon`=? WHERE `id`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(channel.Name, channel.Icon, channel.ID)
	return err
}

func (self *Manager) ChannelDelete(channel *Channel) error {
	s := "DELETE FROM `channel` WHERE `id`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(channel.ID)
	return err
}

type StepNode struct {
	Option
	Children []*StepNode `json:"children"`
}

type EntryTree struct {
	Entry
	Children []*StepNode `json:"children"`
}

func (self *Manager) loadChildrenNodes(parentId string) ([]*StepNode, error) {
	nodes := make([]*StepNode, 0, 100)
	filter := OptionsFilter{
		Parents: []string{parentId},
	}
	children, err := self.OptionsGet(&filter)
	if err != nil {
		return nodes, err
	}
	for _, child := range children {
		node := StepNode{
			Option:   *child,
			Children: []*StepNode{},
		}
		nodes = append(nodes, &node)
	}
	for _, node := range nodes {
		if node.Type == C_ST_option {
			node.Children, err = self.loadChildrenNodes(node.ID)
			if err != nil {
				return nodes, nil
			}
		}
	}
	return nodes, nil
}

func (self *Manager) EntryTreeGet(id string) (*EntryTree, error) {
	tree, err := func() (*EntryTree, error) {
		entry := EntryTree{
			Children: []*StepNode{},
		}
		s := "SELECT `id`, `channel_id`, `text`, `layout_type` FROM `entry` " +
			"WHERE `id`=? "
		stmt, err := self.conn.Prepare(s)
		if err != nil {
			return nil, nil
		}
		defer stmt.Close()
		rows, err := stmt.Query(id)
		if rows.Next() {
			err = rows.Scan(&entry.ID, &entry.ChannelId, &entry.Text, &entry.LayoutType)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, nil
		}
		return &entry, nil
	}()
	if err != nil || tree == nil {
		return tree, err
	}
	tree.Children, err = self.loadChildrenNodes(tree.ID)
	if err != nil {
		return nil, err
	}
	return tree, err
}

func (self *Manager) EntryGet(id string) (*Entry, error) {
	s := "SELECT `id`, `channel_id`, `text`, `layout_type` FROM `entry` " +
		"WHERE `id`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return nil, nil
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if rows.Next() {
		entry := Entry{}
		err = rows.Scan(&entry.ID, &entry.ChannelId, &entry.Text, &entry.LayoutType)
		if err != nil {
			return nil, err
		}
		return &entry, nil
	} else {
		return nil, nil
	}
}

func (self *Manager) EntriesGet(channelId string) ([]*Entry, error) {
	s := "SELECT `id`, `channel_id`, `text`, `layout_type` FROM `entry` " +
		"WHERE `channel_id`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(channelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := make([]*Entry, 0, 100)
	for rows.Next() {
		entry := &Entry{}
		err = rows.Scan(&entry.ID, &entry.ChannelId, &entry.Text, &entry.LayoutType)
		if err != nil {
			return entries, nil
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (self *Manager) EntryCreate(entry *Entry) error {
	s := "INSERT INTO `entry`(`id`, `channel_id`, `text`, `layout_type`) " +
		"VALUES(?, ?, ?, ?) "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(entry.ID, entry.ChannelId, entry.Text, entry.LayoutType)
	if err != nil {
		return err
	}
	return nil
}

func (self *Manager) EntryUpdate(entry *Entry) error {
	s := "UPDATE `entry` SET " +
		"`text`=?, `layout_type`=? " +
		"WHERE `id`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(entry.Text, entry.LayoutType, entry.ID)
	return err
}

func (self *Manager) EntryDelete(entry *Entry) error {
	s := "DELETE FROM `entry` " +
		"WHERE `id`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(entry.ID)
	return err
}

type ReportFilter struct {
	Ids []string
}

func (self *Manager) ReportsGet(filter *ReportFilter) ([]*Report, error) {
	reports := make([]*Report, 0, 1000)
	reportCasesMap := make(map[string][]string)
	reportDecreesMap := make(map[string][]string)
	err := func() error {
		s := "SELECT  `id`, `title`, `conclusion`, `cases`, `decrees` " +
			"FROM `report` "
		if filter != nil {
			whereClauses := make([]string, 0, 10)
			if len(filter.Ids) > 0 {
				wrappedIds := make([]string, len(filter.Ids))
				for index, id := range filter.Ids {
					wrappedIds[index] = fmt.Sprintf("'%s'", id)
				}
				whereClauses = append(whereClauses, "`id` IN ("+strings.Join(wrappedIds, ", ")+")")
			}
			if len(whereClauses) > 0 {
				s += "WHERE " + strings.Join(whereClauses, " AND ")
			}
		}
		log.Printf("SQL: %s", s)
		stmt, err := self.conn.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()
		rows, err := stmt.Query()
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			report := Report{}
			var casesJson sql.NullString
			var decreesJson sql.NullString
			err = rows.Scan(&report.ID, &report.Title, &report.Conclusion,
				&casesJson, &decreesJson)
			if err != nil {
				panic(err)
			}
			if casesJson.Valid && len(casesJson.String) > 0 {
				var cases []string
				err = json.Unmarshal([]byte(casesJson.String), &cases)
				if err != nil {
					panic(err)
				}
				reportCasesMap[report.ID] = cases
			}
			if decreesJson.Valid && len(decreesJson.String) > 0 {
				var decrees []string
				err = json.Unmarshal([]byte(decreesJson.String), &decrees)
				if err != nil {
					panic(err)
				}
				reportDecreesMap[report.ID] = decrees
			}
			reports = append(reports, &report)
		}
		return nil
	}()
	if err != nil {
		return reports, err
	}
	err = func() error {
		todoList := make([]string, 0, 1000)
		for _, caseIds := range reportCasesMap {
			for _, id := range caseIds {
				todoList = append(todoList, fmt.Sprintf("'%s'", id))
			}
		}
		if len(todoList) > 0 {
			s := "SELECT `id`, `content`, `link` " +
				"FROM `case` " +
				"WHERE `id` IN (" + strings.Join(todoList, ", ") + ") "
			stmt, e := self.conn.Prepare(s)
			if e != nil {
				panic(e)
			}
			defer stmt.Close()
			rows, e := stmt.Query()
			if e != nil {
				panic(e)
			}
			defer rows.Close()
			caseMap := make(map[string]*event)
			for rows.Next() {
				c := event{}
				e = rows.Scan(&c.ID, &c.Content, &c.Link)
				if e != nil {
					panic(e)
				}
				caseMap[c.ID] = &c
			}
			for _, report := range reports {
				if caseIds, ok := reportCasesMap[report.ID]; ok {
					cases := make([]*event, len(caseIds))
					for index, caseId := range caseIds {
						cases[index] = caseMap[caseId]
					}
					report.Cases = cases
				}
			}
		}
		return nil
	}()
	if err != nil {
		return reports, err
	}
	err = func() error {
		todoList := make([]string, 0, 1000)
		for _, decreeIds := range reportDecreesMap {
			for _, id := range decreeIds {
				todoList = append(todoList, fmt.Sprintf("'%s'", id))
			}
		}
		if len(todoList) > 0 {
			s := "SELECT `id`, `source`, `content`, `link` " +
				"FROM `decree` " +
				"WHERE `id` IN (" + strings.Join(todoList, ", ") + ") "
			stmt, e := self.conn.Prepare(s)
			if e != nil {
				panic(e)
			}
			defer stmt.Close()
			rows, e := stmt.Query()
			if e != nil {
				panic(e)
			}
			defer rows.Close()
			decreeMap := make(map[string]*decree)
			for rows.Next() {
				d := decree{}
				e = rows.Scan(&d.ID, &d.Source, &d.Content, &d.Link)
				if e != nil {
					panic(e)
				}
				decreeMap[d.ID] = &d
			}
			for _, report := range reports {
				if decreeIds, ok := reportDecreesMap[report.ID]; ok {
					decrees := make([]*decree, len(decreeIds))
					for index, decreeId := range decreeIds {
						decrees[index] = decreeMap[decreeId]
					}
					report.Decrees = decrees
				}
			}
		}
		return nil
	}()
	if err != nil {
		return reports, err
	}
	return reports, nil
}

type OptionsFilter struct {
	Parents []string
	Ids     []string
}

func (self *Manager) OptionsGet(filter *OptionsFilter) ([]*Option, error) {
	options := make([]*Option, 0, 100)
	optionReportMap := make(map[string]string)
	toLoadReportIds := make([]string, 0, 100)
	err := func() error {
		s := "SELECT `id`, `parent_id`, `text`, `type`, `ref_id` " +
			"FROM `option` "
		if filter != nil {
			whereClause := make([]string, 0, 100)
			if len(filter.Parents) > 0 {
				wrappedIds := make([]string, len(filter.Parents))
				for index, id := range filter.Parents {
					wrappedIds[index] = fmt.Sprintf("'%s'", id)
				}
				whereClause = append(whereClause, fmt.Sprintf("`parent_id` IN (%s)", strings.Join(wrappedIds, ", ")))
			}
			if len(filter.Ids) > 0 {
				wrappedIds := make([]string, len(filter.Ids))
				for index, id := range filter.Ids {
					wrappedIds[index] = fmt.Sprintf("'%s'", id)
				}
				whereClause = append(whereClause, fmt.Sprintf("`id` IN (%s)", strings.Join(wrappedIds, ", ")))
			}
			if len(whereClause) > 0 {
				s += fmt.Sprintf("WHERE %s ", strings.Join(whereClause, " AND "))
			}
		}
		stmt, err := self.conn.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()
		log.Println(s)
		rows, err := stmt.Query()
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			option := Option{}
			var reportId sql.NullString
			err = rows.Scan(&option.ID, &option.ParentId, &option.Text, &option.Type, &reportId)
			if err != nil {
				return err
			}
			log.Printf("%#v %#v", &option, reportId)
			if option.Type == C_ST_report {
				optionReportMap[option.ID] = reportId.String
				toLoadReportIds = append(toLoadReportIds, reportId.String)
			}
			options = append(options, &option)
		}
		return nil
	}()
	if err != nil {
		return options, err
	}
	if len(toLoadReportIds) > 0 {
		filter := ReportFilter{
			Ids: toLoadReportIds,
		}
		reports, err := self.ReportsGet(&filter)
		if err != nil {
			return options, err
		}
		reportMap := make(map[string]*Report)
		for _, report := range reports {
			reportMap[report.ID] = report
		}
		for _, option := range options {
			if option.Type == "report" {
				option.Report = reportMap[optionReportMap[option.ID]]
			}
		}
	}
	return options, nil
}

func (self *Manager) OptionGet(id string) (*Option, error) {
	filter := OptionsFilter{
		Ids: []string{id},
	}
	options, err := self.OptionsGet(&filter)
	if err != nil {
		return nil, err
	}
	if len(options) > 0 {
		return options[0], nil
	}
	return nil, nil
}

func (self *Manager) OptionCreate(option *Option) error {
	switch option.Type {
	case "option":
		s := "INSERT INTO `option`(`id`, `parent_id`, `text`, `type`) " +
			"VALUES(?, ?, ?, ?) "
		stmt, err := self.conn.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(option.ID, option.ParentId, option.Text, option.Type)
		if err != nil {
			return err
		}
	case "report":
		tx, err := self.conn.Begin()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()
		s := "INSERT INTO `option`(`id`, `parent_id`, `text`, `type`, `ref_id`) " +
			"VALUES(?, ?, ?, ?, ?) "
		stmt, err := tx.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(option.ID, option.ParentId, option.Text, option.Type, option.Report.ID)
		if err != nil {
			return err
		}
		var decreesJson []byte
		if len(option.Report.Decrees) > 0 {
			s = "INSERT INTO `decree`(`id`, `source`, `content`, `link`) " +
				"VALUES(?, ?, ?, ?) "
			decreeStmt, err := tx.Prepare(s)
			if err != nil {
				return err
			}
			defer decreeStmt.Close()
			ids := make([]string, 0, 100)
			for _, decree := range option.Report.Decrees {
				_, err = decreeStmt.Exec(decree.ID, decree.Source, decree.Content, decree.Link)
				if err != nil {
					return err
				}
				ids = append(ids, decree.ID)
			}
			decreesJson, _ = json.Marshal(ids)
		}
		var casesJson []byte
		if len(option.Report.Cases) > 0 {
			s = "INSERT INTO `case`(`id`, `content`, `link`) " +
				"VALUES(?, ?, ?) "
			caseStmt, err := tx.Prepare(s)
			if err != nil {
				return err
			}
			defer caseStmt.Close()
			ids := make([]string, 0, 100)
			for _, event := range option.Report.Cases {
				_, err = caseStmt.Exec(event.ID, event.Content, event.Link)
				if err != nil {
					return err
				}
				ids = append(ids, event.ID)
			}
			casesJson, _ = json.Marshal(ids)
		}
		s = "INSERT INTO `report`(`id`, `title`, `conclusion`,  `cases`, `decrees`) " +
			"VALUES(?, ?, ?, ?, ?) "
		reportStmt, err := tx.Prepare(s)
		if err != nil {
			return err
		}
		defer reportStmt.Close()
		_, err = reportStmt.Exec(option.Report.ID, option.Report.Title, option.Report.Conclusion,
			casesJson, decreesJson)
		return err
	}
	return nil
}

func (self *Manager) OptionDelete(option *Option) error {
	s := "DELETE FROM `option` " +
		"WHERE `id`=? "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(option.ID)
	return err
}

func (self *Manager) IssueCreate(issue *Issue) error {
	issue.ID = self.AllocateId(true)
	issue.CreatedTime = self.CurrentTimeMs()
	s := "INSERT INTO `client_issue`(`id`, `created_time`, `client`, `description`, `attachments`) " +
		"VALUES(?, ?, ?, ?, ?) "
	stmt, err := self.conn.Prepare(s)
	if err != nil {
		return err
	}
	defer stmt.Close()
	clientJson, err := json.Marshal(&issue.Client)
	if err != nil {
		return err
	}
	attachmentsJson, err := json.Marshal(&issue.Detail.Attachments)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(issue.ID, issue.CreatedTime, string(clientJson), issue.Detail.Desc, attachmentsJson)
	if err != nil {
		return err
	}
	return nil
}

func (self *Manager) IssueGet(id string) (*Issue, error) {
	filter := IssuesFilter{
		ID: id,
	}
	issues, err := self.IssuesGet(&filter)
	if err != nil {
		return nil, err
	}
	if len(issues) == 0 {
		return nil, nil
	}
	return issues[0], nil
}

type IssuesFilter struct {
	ID string
}

func (self *Manager) IssuesGet(filter *IssuesFilter) ([]*Issue, error) {
	issues, err := func()([]*Issue, error) {
		issues := make([]*Issue, 0, 100)
		s := "SELECT `id`, `created_time`, `client`, `description`, `attachments`, `status`, `solution` " +
			"FROM `client_issue` "
		args := []interface{}{}
		if filter != nil {
			if len(filter.ID) > 0 {
				s += "WHERE `id`=? "
				args = append(args, filter.ID)
			}
		}
		stmt, err := self.conn.Prepare(s)
		if err != nil {
			return issues, err
		}
		defer stmt.Close()

		rows, err := stmt.Query(args...)
		if err != nil {
			return issues, err
		}
		defer rows.Close()

		for rows.Next() {
			issue := Issue{}
			var clientJson string
			var attachmentsJson sql.NullString
			var solution sql.NullString
			err = rows.Scan(&issue.ID, &issue.CreatedTime, &clientJson, &issue.Detail.Desc,
				&attachmentsJson, &issue.Status, &solution)
			if err != nil {
				return issues, err
			}
			err = json.Unmarshal([]byte(clientJson), &issue.Client)
			if err != nil {
				return issues, err
			}
			if attachmentsJson.Valid {
				err = json.Unmarshal([]byte(attachmentsJson.String), &issue.Detail.Attachments)
				if err != nil {
					return issues, err
				}
			}
			if solution.Valid {
				issue.Solution = solution.String
			}
			issues = append(issues, &issue)
		}
		return issues, nil
	}()
	if err != nil {
		return issues, err
	}
	if len(issues) > 0 {
		wrappedIds := make([]string, 0, len(issues))
		for _, issue := range issues {
			wrappedIds = append(wrappedIds, fmt.Sprintf("'%s'", issue.ID))
		}
		s := "SELECT `issue_id`, `tag` FROM `issue_tag` " +
			"WHERE `issue_id` IN (" + strings.Join(wrappedIds, ", ") + ") "
		stmt, err := self.conn.Prepare(s)
		if err != nil {
			return issues, err
		}
		defer stmt.Close()

		rows, err := stmt.Query()
		if err != nil {
			return issues, err
		}
		defer rows.Close()
		tagMap := make(map[string][]string)
		for rows.Next() {
			var issueId string
			var tag string
			err = rows.Scan(&issueId, &tag)
			if err != nil {
				return issues, err
			}
			if _, ok := tagMap[issueId]; !ok {
				tagMap[issueId] = []string{}
			}
			tagMap[issueId] = append(tagMap[issueId], tag)
		}
		for _, issue := range issues {
			if tags, ok := tagMap[issue.ID]; ok {
				issue.Tags = tags
			} else {
				issue.Tags = []string{}
			}
		}
	}

	return issues, nil
}

func (self *Manager) IssueSolute(id, solution string, tags []string) (error) {
	tx, err := self.conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			log.Printf("Error(%s)", err.Error())
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	{
		s := "DELETE FROM `issue_tag` " +
			"WHERE `issue_id`=? "
		stmt, err := tx.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
	}
	{
		s := "INSERT INTO `issue_tag`(`issue_id`, `tag`) " +
			"VALUES(?, ?) "
		stmt, err := tx.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()
		for _, tag := range tags {
			_, err = stmt.Exec(id, tag)
			if err != nil {
				return err
			}
		}
	}
	{
		s := "UPDATE `client_issue` " +
			"SET `solution`=? WHERE `id`=? "
		stmt, err := tx.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(solution, id)
		if err != nil {
			return err
		}
	}
	return nil
}

