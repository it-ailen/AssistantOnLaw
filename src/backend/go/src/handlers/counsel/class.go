package counsel

import (
	"content"
	"foolhttp"
	"handlers/accounts"
	"net/http"
)

type ClassHandler struct{}

func (self *ClassHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *ClassHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"description": {"type": "string"},
			"logo": {"type": "string"}
		},
		"required": ["name", "description", "logo"]
	}`
	class := content.Class{}
	foolhttp.JsonSchemaCheck(r, schemaDefine, &class)

	mgr := content.GetManager()
	id := mgr.CreateClass(&class)
	cls := mgr.SelectClass(id)

	foolhttp.WriteJson(w, cls)
	return nil
}

func (self *ClassHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"description": {"type": "string"},
			"logo": {"type": "string"}
		}
	}`
	var args content.SqlKV
	foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

	id := foolhttp.RouteArgument(r, "id")

	mgr := content.GetManager()
	class := mgr.SelectClass(id)
	if class == nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
	}
    mgr.UpdateClass(id, args)
    cls := mgr.SelectClass(id)
	foolhttp.WriteJson(w, cls)
	return nil
}

func (self *ClassHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	mgr.DeleteClass(id)
	return nil
}

func (self *ClassHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	id := foolhttp.RouteArgumentWithDefault(r, "id", "nil")
	mgr := content.GetManager()
    if id != "nil" {
        classes := mgr.SelectClasses(nil)
        foolhttp.WriteJson(w, classes)
    } else {
        class := mgr.SelectClass(id)
        foolhttp.WriteJson(w, class)
    }
    return nil
}

func NewClassHandler() *ClassHandler {
	handler := new(ClassHandler)
	return handler
}
