package counsel

import (
	"content"
	"encoding/json"
	"foolhttp"
	"handlers/accounts"
	"log"
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
			"trigger_by": {
			    "type": "object",
			    "properties": {
			        "question_id": {"type": "string"},
			        "options": {
			        	"type": "array",
			        	"items": {
			        		"type": "integer",
			        		"minimum": 0
			        	}
			        }
			    }
			},
			"options": {
			    "type": "array",
			    "items": {"type": "string"},
			    "minItems": 2,
			    "uniqueItems": true
			}
		},
		"required": ["question", "type", "entry_id", "options"]
	}`
	requestArgs := struct {
		Question  string          `json:"question"`
		Type      string          `json:"type"`
		EntryId   string          `json:"entry_id"`
		Options   []string        `json:"options"`
		TriggerBy json.RawMessage `json:"trigger_by"`
	}{}
	args := make(content.SqlKV)
	foolhttp.JsonSchemaCheck(r, schemaDefine, &requestArgs)
	args["question"] = requestArgs.Question
	args["type"] = requestArgs.Type
	args["entry_id"] = requestArgs.EntryId
	args["options"] = requestArgs.Options

	mgr := content.GetManager()
	if len(requestArgs.TriggerBy) > 0 {
		triggerArgs := struct {
			QuestionId string `json:"question_id"`
			Options    []int  `json:"options"`
		}{}
		err := json.Unmarshal(requestArgs.TriggerBy, &triggerArgs)
		if err != nil {
			panic(foolhttp.BadArgHTTPError("Invalid json for trigger_by"))
		}
		q := mgr.SelectQuestion(triggerArgs.QuestionId)
		if q == nil {
			panic(foolhttp.BadArgHTTPError("Invalid question id"))
		}
		for _, index := range triggerArgs.Options {
			if index >= len(q.Options) {
				panic(foolhttp.BadArgHTTPError("Invalid option index(%s)", index))
			}
		}
		triggerMap := map[string][]int{
			triggerArgs.QuestionId: triggerArgs.Options,
		}
		args["trigger_by"] = triggerMap
	}
	id := mgr.CreateQuestion(args)
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
			    "type": "object",
			    "properties": {
			        "question_id": {"type": ["string", "null"]},
			        "options": {
			        	"type": "array",
			        	"items": {
			        		"type": "integer",
			        		"minimum": 0
			        	}
			        }
			    }
			},
			"options": {
			    "type": "array",
			    "items": {"type": "string"},
			    "minItems": 2,
			    "uniqueItems": true
			}
		}
	}`
	requestArgs := struct {
		Question  string          `json:"question"`
		Type      string          `json:"type"`
		EntryId   string          `json:"entry_id"`
		Options   []string          `json:"options"`
		TriggerBy json.RawMessage `json:"trigger_by"`
	}{}
	args := make(content.SqlKV)
	foolhttp.JsonSchemaCheck(r, schemaDefine, &requestArgs)
	if len(requestArgs.Question) > 0 {
		args["question"] = requestArgs.Question
	}
	if len(requestArgs.Type) > 0 {
		args["type"] = requestArgs.Type
	}
	if len(requestArgs.EntryId) > 0 {
		args["entry_id"] = requestArgs.EntryId
	}
	if len(requestArgs.Options) > 0 {
		args["options"] = requestArgs.Options
	}

	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	entry := mgr.SelectQuestion(id)
	if entry == nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
	}
	if len(requestArgs.TriggerBy) > 0 {
		triggerArgs := struct {
			QuestionId string `json:"question_id"`
			Options    []int  `json:"options"`
		}{}
		err := json.Unmarshal(requestArgs.TriggerBy, &triggerArgs)
		if err != nil {
			panic(foolhttp.BadArgHTTPError("Invalid json for trigger_by"))
		}
		if len(triggerArgs.QuestionId) == 0 {
			args["trigger_by"] = nil
		} else {
			q := mgr.SelectQuestion(triggerArgs.QuestionId)
			if q == nil {
				panic(foolhttp.BadArgHTTPError("Invalid question id"))
			}
			for _, index := range triggerArgs.Options {
				if index >= len(q.Options) {
					panic(foolhttp.BadArgHTTPError("Invalid option index(%s)", index))
				}
			}
			triggerMap := map[string][]int{
				triggerArgs.QuestionId: triggerArgs.Options,
			}
			args["trigger_by"] = triggerMap
		}
	}
	log.Printf("args: %v", args)
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
	if entry == nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
	}
	mgr.DeleteQuestion(id)
	return nil
}

func (self *QuestionHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	id := foolhttp.RouteArgumentWithDefault(r, "id", "nil")
	mgr := content.GetManager()
	if id == "nil" {
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
