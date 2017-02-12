package counsel

import (
    "content"
    "foolhttp"
    "handlers/accounts"
    "net/http"
)

type EntryHandler struct{}

func (self *EntryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    foolhttp.DoServeHTTP(self, w, r)
}

func (self *EntryHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    _ = accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"logo": {"type": "string"},
			"layout_type": {
			    "type": "string",
			    "enum": ["single", "multiple"]
			},
			"class_id": {"type": "string"}
		},
		"required": ["name", "logo", "layout_type", "class_id"]
	}`
    entry := content.Entry{}
    foolhttp.JsonSchemaCheck(r, schemaDefine, &entry)

    mgr := content.GetManager()
    id := mgr.CreateEntry(&entry)
    inst := mgr.SelectEntry(id)
    foolhttp.WriteJson(w, inst)
    return nil
}

func (self *EntryHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    _ = accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"logo": {"type": "string"},
			"layout_type": {
			    "type": "string",
			    "enum": ["single", "multiple"]
			}
		}
	}`
    args := make(content.SqlKV)
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

    id := foolhttp.RouteArgument(r, "id")

    mgr := content.GetManager()
    entry := mgr.SelectEntry(id)
    if entry == nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
    }
    mgr.UpdateEntry(id, args)
    inst := mgr.SelectEntry(id)
    foolhttp.WriteJson(w, inst)
    return nil
}

func (self *EntryHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    _ = accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    id := foolhttp.RouteArgument(r, "id")
    mgr := content.GetManager()
    entry := mgr.SelectEntry(id)
    if entry == nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
    }
    mgr.DeleteEntry(id)
    return nil
}

func (self *EntryHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	id := foolhttp.RouteArgumentWithDefault(r, "id", "nil")
	mgr := content.GetManager()
    if id == "nil" {
        filter := content.EntryFilter{}
        class_id := foolhttp.QueryArgumentWithDefault(r, "class_id", "nil")
        if class_id != "nil" {
            filter.ClassIds = []string{class_id}
        }
        classes := mgr.SelectEntries(&filter)
        foolhttp.WriteJson(w, classes)
    } else {
        class := mgr.SelectEntry(id)
        foolhttp.WriteJson(w, class)
    }
    return nil
}

func NewEntryHandler() *EntryHandler {
    handler := new(EntryHandler)
    return handler
}
