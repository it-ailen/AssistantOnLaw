package admin

import (
    "net/http"
    "foolhttp"
    "content"
    "encoding/json"
)

type IssueHandler struct {}

func (self *IssueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}
func (self *IssueHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    mgr := content.GetManager()
    issues, err := mgr.IssuesGet(nil)
    if err != nil {
        return foolhttp.UnknownHTTPError(err.Error())
    }
    resp := struct {
        List []*content.Issue `json:"list"`
    }{
        List: issues,
    }
    body, _ := json.Marshal(&resp)
    w.Write(body)
    return nil
}

func (self *IssueHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    args := struct {
        Solution string `json:"solution"`
        Tags []string `json:"tags"`
    }{}
    e := foolhttp.ParseJsonArgs(r, &args)
    if e != nil {
        return e
    }
    mgr := content.GetManager()
    id := foolhttp.RouteArgument(r, "id")
    issue, err := mgr.IssueGet(id)
    if err != nil {
        return foolhttp.UnknownHTTPError(err.Error())
    }
    if issue == nil {
        return foolhttp.NotFoundHTTPError("No issue")
    }
    err = mgr.IssueSolute(id, args.Solution, args.Tags)
    if err != nil {
        return foolhttp.UnknownHTTPError(err.Error())
    }
    return nil
}


func NewIssueHandler() *IssueHandler {
    return &IssueHandler{}
}
