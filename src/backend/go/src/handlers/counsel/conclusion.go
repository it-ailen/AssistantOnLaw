package counsel

import (
	"content"
	"foolhttp"
	"handlers/accounts"
	"net/http"
    "fmt"
)

type ConclusionHandler struct{}

const C_SELECTIONS_SCHEMA = `
			{
				"type": "array",
				"items": {
					"type": "object",
					"properties": {
						"question_id": {"type": "string"},
						"selections": {
							"type": "array",
							"items": {
								"type": "integer",
								"minimum": 0
							}
						}
					},
					"required": ["question_id", "selections"]
				},
				"minItems": 1
			}
			`

func (self *ConclusionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *ConclusionHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
    type argsDefine struct {
        Title string `json:"title"`
        Context string `json:"context"`
        Selections content.Selections `json:"selections"`
    }
	schemaDefine := fmt.Sprintf(`{
		"type": "object",
		"properties": {
			"title": {"type": "string"},
			"context": {"type": "string"},
			"selections": %s
		},
		"required": ["title", "context"]
	}`, C_SELECTIONS_SCHEMA)
    args := argsDefine{}
	foolhttp.JsonSchemaCheck(r, schemaDefine, &args)
	conclusion := content.Conclusion{
        Title: args.Title,
        Context: args.Context,
    }

	mgr := content.GetManager()
	id := mgr.CreateConclusion(&conclusion, args.Selections)
	inst := mgr.SelectConclusion(id)
	foolhttp.WriteJson(w, inst)
	return nil
}

func (self *ConclusionHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	schemaDefine := `{
		"type": "object",
		"properties": {
			"title": {"type": "string"},
			"context": {"type": "string"}
		}
	}`
	args := make(content.SqlKV)
	foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

	id := foolhttp.RouteArgument(r, "id")

	mgr := content.GetManager()
	item := mgr.SelectConclusion(id)
	if item != nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
	}
	mgr.UpdateConclusion(id, args)
	inst := mgr.SelectConclusion(id)
	foolhttp.WriteJson(w, inst)
	return nil
}

func (self *ConclusionHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	item := mgr.SelectConclusion(id)
	if item != nil {
		panic(foolhttp.NotFoundHTTPError("Invalid id"))
	}
	mgr.DeleteConclusion(id)
	return nil
}

func (self *ConclusionHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	id := foolhttp.RouteArgumentWithDefault(r, "id", "nil")
	mgr := content.GetManager()
	if id == "nil" {
		selectionsJson := foolhttp.QueryArgumentWithDefault(r, "selections", "nil")
		if selectionsJson != "nil" {
			var selections content.Selections
			err := foolhttp.JsonStringCheck(selectionsJson, C_SELECTIONS_SCHEMA, &selections)
			//err := json.Unmarshal([]byte(selectionsJson), &selections)
			if err != nil {
				panic(foolhttp.BadArgHTTPError("Invalid json: %s", err.Error()))
			}
			item := mgr.CalculateConclusion(selections)
			items := []*content.Conclusion{}
			if item != nil {
				items = append(items, item)
			}
			foolhttp.WriteJson(w, items)
		} else {
			items := mgr.SelectConclusions(nil)
			foolhttp.WriteJson(w, items)
		}
	} else {
		item := mgr.SelectConclusion(id)
		foolhttp.WriteJson(w, item)
	}
	return nil
}

func NewConclusionHandler() *ConclusionHandler {
	handler := new(ConclusionHandler)
	return handler
}
