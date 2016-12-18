package resource

import (
    "content"
    "content/definition"
    "foolhttp"
    "handlers/accounts"
    "net/http"
)

type FileHandler struct{}

func (self *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    foolhttp.DoServeHTTP(self, w, r)
}

func (self *FileHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    me := accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    type argsDefine struct {
        Type   string `json:"type"`
        Name   string `json:"name"`
        URI    string `json:"reference_uri"`
        Parent string `json:"parent"`
    }
    schemaDefine := `{
		"type": "object",
		"properties": {
			"type": {
			    "type": "string",
			    "enum": ["directory", "file"]
			},
			"name": {"type": "string"},
			"reference_uri": {"type": "string"},
			"parent": {
				"type": "string"
			}
		},
		"required": ["type", "name", "parent"]
	}`
    args := argsDefine{}
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

    mgr := content.GetManager()
    parent, e0 := mgr.LoadFile(args.Parent)
    if e0 != nil {
        panic(foolhttp.UnknownHTTPError(e0.Error()))
    }
    if parent == nil {
        panic(foolhttp.BadArgHTTPError("Invalid parent"))
    }
    if dir, ok := parent.(*definition.Directory); ok {
        if args.Type == definition.C_FT_DIR {
            newDir, err := mgr.CreateDirectory(me, args.Name, dir)
            if err != nil {
                panic(err)
            }
            foolhttp.WriteJson(w, newDir)
        } else {
            newFile, err := mgr.CreateFile(me, args.Name, args.URI, dir)
            if err != nil {
                panic(err)
            }
            foolhttp.WriteJson(w, newFile)
        }
    } else {
        panic(foolhttp.BadArgHTTPError("Invalid parent"))
    }
    return nil
}

func (self *FileHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"reference_uri": {"type": "string"}
		}
	}`
    args := make(map[string]string)
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)
    id := foolhttp.RouteArgument(r, "id")

    mgr := content.GetManager()

    file, e0 := mgr.LoadFile(id)
    if e0 != nil {
        panic(e0)
    }
    if file == nil {
        panic(foolhttp.NotFoundHTTPError("Unknown id"))
    }
    e0 = mgr.UpdateFile(file, args)
    if e0 != nil {
        panic(e0)
    }
    foolhttp.WriteJson(w, file)
    return nil
}

func (self *FileHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })
    id := foolhttp.RouteArgument(r, "id")
    mgr := content.GetManager()
    file, err := mgr.LoadFile(id)
    if err != nil {
        panic(err)
    }
    if file == nil {
        panic(foolhttp.NotFoundHTTPError("Unknown id"))
    }
    if dir, ok := file.(*definition.Directory); ok {
        if len(dir.Children) > 0 {
            panic(foolhttp.ForbiddenHTTPError("Delete an unempty directory"))
        }
    }
    mgr.DeleteFile(id)
    return nil
}


//func (self *FileHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
//    root := foolhttp.RouteArgumentWithDefault(r, "root", "/")
//    mgr := content.GetManager()
//    tmp, err := mgr.LoadFile(root)
//    if err != nil {
//        panic(err)
//    }
//    if file, ok := tmp.(*definition.File); ok {
//        foolhttp.WriteJson(w, file)
//    } else if dir, ok := tmp.(*definition.Directory); ok {
//        tree, err := mgr.LoadDirectoryTree(dir)
//        if err != nil {
//            panic(err)
//        }
//        foolhttp.WriteJson(w, tree)
//    }
//    return nil
//}

func NewFileHandler() *FileHandler {
    handler := new(FileHandler)
    return handler
}
