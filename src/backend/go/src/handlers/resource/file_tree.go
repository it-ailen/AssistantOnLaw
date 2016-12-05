package resource

import (
    "net/http"
    "foolhttp"
    "content"
    "content/definition"
)

type FileTreeHandler struct {}


func (self *FileTreeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    foolhttp.DoServeHTTP(self, w, r)
}


func (self *FileTreeHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    root := foolhttp.RouteArgumentWithDefault(r, "id", "/")
    mgr := content.GetManager()
    tmp, err := mgr.LoadFile(root)
    if err != nil {
        panic(err)
    }
    if file, ok := tmp.(*definition.File); ok {
        foolhttp.WriteJson(w, file)
    } else if dir, ok := tmp.(*definition.Directory); ok {
        tree, err := mgr.LoadDirectoryTree(dir)
        if err != nil {
            panic(err)
        }
        foolhttp.WriteJson(w, tree)
    }
    return nil
}


func NewFileTreeHandler() *FileTreeHandler {
    handler := new(FileTreeHandler)
    return handler
}

