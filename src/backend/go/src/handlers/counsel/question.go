package counsel

import (
	"content"
	"foolhttp"
	"handlers/accounts"
	"net/http"
)

type QuestionHandler struct{}

func (self *QuestionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *QuestionHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	schemaDefine := `{
		"type": "object",
		"properties": {
			"question": {"type": "string"},
			"type": {
			    "type": "string",
			    "enum": ["single", "multiple"]
			},
			"entry_id": {"type": "string"},
			"trigger_by": {"type": "string"},
			"options": {
			    "type": "array",
			    "items": {"type": "string"},
			    "minItems": 2,
			    "uniqueItems": true
			}
		},
		"required": ["question", "type", "entry_id", "options"]
	}`
	question := content.Question{}
	foolhttp.JsonSchemaCheck(r, schemaDefine, &question)

	mgr := content.GetManager()
	id := mgr.CreateQuestion(&question)
	inst := mgr.SelectQuestion(id)
	foolhttp.WriteJson(w, inst)
	return nil
}

func (self *QuestionHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	schemaDefine := `{
		"type": "object",
		"properties": {
			"question": {"type": "string"},
			"type": {
			    "type": "string",
			    "enum": ["single", "multiple"]
			},
			"entry_id": {"type": "string"},
			"trigger_by": {
			    "type": ["string", "null"]
			},
			"options": {
			    "type": "array",
			    "items": {"type": "string"},
			    "minItems": 2,
			    "uniqueItems": true
			}
		}
	}`
	args := make(content.SqlKV)
	foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

	id := foolhttp.RouteArgument(r, "id")

	mgr := content.GetManager()
	entry := mgr.SelectQuestion(id)
	if entry != nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
	}
	mgr.UpdateQuestion(id, args)
	inst := mgr.SelectQuestion(id)
	foolhttp.WriteJson(w, inst)
	return nil
}

func (self *QuestionHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	entry := mgr.SelectQuestion(id)
	if entry != nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
	}
	mgr.DeleteQuestion(id)
	return nil
}

func (self *QuestionHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	id := foolhttp.RouteArgumentWithDefault(r, "id", "nil")
	mgr := content.GetManager()
	if id != "nil" {
		filter := content.QuestionFilter{}
		entryId := foolhttp.QueryArgumentWithDefault(r, "entry_id", "nil")
		if entryId != "nil" {
			filter.EntryIds = []string{entryId}
		}
		classes := mgr.SelectQuestions(&filter)
		foolhttp.WriteJson(w, classes)
	} else {
		class := mgr.SelectQuestion(id)
		foolhttp.WriteJson(w, class)
	}
	return nil
}

func NewQuestionHandler() *QuestionHandler {
	handler := new(QuestionHandler)
	return handler
}
