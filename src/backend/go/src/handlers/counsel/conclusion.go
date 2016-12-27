package counsel

import (
	"content"
	"encoding/json"
	"foolhttp"
	"handlers/accounts"
	"net/http"
)

type ConclusionHandler struct{}

func (self *ConclusionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *ConclusionHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	_ = accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})
	schemaDefine := `{
		"type": "object",
		"properties": {
			"title": {"type": "string"},
			"context": {"type": "string"}
		},
		"required": ["title", "context"]
	}`
	conclusion := content.Conclusion{}
	foolhttp.JsonSchemaCheck(r, schemaDefine, &conclusion)

	mgr := content.GetManager()
	id := mgr.CreateConclusion(&conclusion)
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
	if id != "nil" {
		selectionsJson := foolhttp.QueryArgumentWithDefault(r, "selection", "nil")
		if selectionsJson != "nil" {
			var selections content.Selections
			err := json.Unmarshal([]byte(selectionsJson), &selections)
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
