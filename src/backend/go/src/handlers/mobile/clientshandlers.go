package mobile

import (
	"content"
	"foolhttp"
	"net/http"
	"content/definition"
)

type IssueHandler struct{}

func (self *IssueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *IssueHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	type arguments struct {
		Client struct {
			Name    string `json:"name"`
			Contact string `json:"contact"`
		} `json:"client"`
		Detail struct {
			Desc        string   `json:"description"`
			Attachments []string `json:"attachments"`
		} `json:"detail"`
	}
	args := arguments{}
	e := foolhttp.ParseJsonArgs(r, &args)
	if e != nil {
		return e
	}
	issue := definition.Issue{}
	issue.Detail.Desc = args.Detail.Desc
	issue.Detail.Attachments = args.Detail.Attachments
	issue.Client = args.Client
	mgr := content.GetManager()
	err := mgr.IssueCreate(&issue)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	return nil
}

func NewIssueHandler() *IssueHandler {
	handler := IssueHandler{}
	return &handler
}
