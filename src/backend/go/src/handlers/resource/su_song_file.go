package resource

import (
    "content"
    "foolhttp"
    "handlers/accounts"
    "net/http"
)

type SuSongFileHandler struct{}

func (self *SuSongFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    foolhttp.DoServeHTTP(self, w, r)
}

func (self *SuSongFileHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    _ = accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    type argsDefine struct {
        Name   string `json:"name"`
        URI    string `json:"uri"`
        StepId string `json:"step_id"`
    }
    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"uri": {"type": "string"},
			"step_id": {"type": "string"}
		},
		"required": ["name", "uri", "step_id"]
	}`
    args := argsDefine{}
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

    mgr := content.GetManager()
    file := mgr.CreateSuSongFile(args.Name, args.URI, args.StepId)
    foolhttp.WriteJson(w, file)
    return nil
}

func (self *SuSongFileHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    _ = accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"uri": {"type": "string"}
		}
	}`
    args := make(map[string]string)
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

    id := foolhttp.RouteArgument(r, "id")

    mgr := content.GetManager()
    file := mgr.LoadSuSongFile(id)
    if file == nil {
        panic(foolhttp.NotFoundHTTPError("Invalid id"))
    }
    foolhttp.WriteJson(w, file)
    return nil
}

func (self *SuSongFileHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    _ = accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    id := foolhttp.RouteArgument(r, "id")
    mgr := content.GetManager()
    mgr.DeleteSuSongFile(id)
    return nil
}

func NewSuSongFileHandler() *SuSongFileHandler {
    handler := new(SuSongFileHandler)
    return handler
}
