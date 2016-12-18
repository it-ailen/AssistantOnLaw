package resource

import (
	"content"
	"foolhttp"
	"handlers/accounts"
	"net/http"
)

type MinShiSuSongHandler struct{}

func (self *MinShiSuSongHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *MinShiSuSongHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})

	type argsDefine struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Files       []string `json:"files"`
	}
	schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"description": {"type": "string"},
			"files": {
			    "type": "array",
			    "items": {
			        "type": "string"
			    }
			}
		},
		"required": ["name", "description", "files"]
	}`
	args := argsDefine{}
	foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

	mgr := content.GetManager()
	class, err := mgr.CreateFaLvWenDaClass(args.Name)
	if err != nil {
		panic(err)
	}
	foolhttp.WriteJson(w, class)
	return nil
}

func (self *MinShiSuSongHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	accounts.CheckAccessibility(r, &accounts.AccessControl{
		SuperOnly: true,
	})

	schemaDefineForCreation := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"description": {"type": "string"},
			"files": {
				"type": "array",
				"items": {"type": "string"}
			}
		},
		"required": ["name", "description"]
	}`
	schemaDefineForUpdating := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"description": {"type": "string"},
			"files": {
			    "oneOf": [
			        {
			            "type": "array",
			            "items": {"type": "string"}
			        },
			        { "type": "null" }
			    ]
			}
		}
	}`

	args := make(map[string]interface{})
	//foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	items := mgr.LoadSuSongWenShu(&content.SuSongWenShuFilter{
		ID: []string{id},
	})
	if len(items) == 0 { /* new */
		foolhttp.JsonSchemaCheck(r, schemaDefineForCreation, &args)
		mgr.CreateSuSongWenShu("min_shi_su_song", id, args)
	} else {
		foolhttp.JsonSchemaCheck(r, schemaDefineForUpdating, &args)
		mgr.UpdateSuSongWenShu(id, args)
	}

	items = mgr.LoadSuSongWenShu(&content.SuSongWenShuFilter{
		ID: []string{id},
	})
	foolhttp.WriteJson(w, items[0])
	return nil
}

func (self *MinShiSuSongHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	mgr := content.GetManager()
	id := foolhttp.RouteArgumentWithDefault(r, "id", "")
	filter := content.SuSongWenShuFilter{
		Flow: "min_shi_su_song",
	}
	if id != "" {
		filter.ID = []string{id}
	}
	items := mgr.LoadSuSongWenShu(&filter)
	if id != "" {
		if len(items) == 0 {
			defaultItem := content.SuSongWenShu{
				ID: id,
				Flow: "min_shi_su_song",
			}
			foolhttp.WriteJson(w, &defaultItem)
		} else {
			foolhttp.WriteJson(w, items[0])
		}
	} else {
		foolhttp.WriteJson(w, items)
	}
	return nil
}

func NewMinShiSuSongHandler() *MinShiSuSongHandler {
	handler := new(MinShiSuSongHandler)
	return handler
}
