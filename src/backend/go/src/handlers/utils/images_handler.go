package utils

import (
	"foolhttp"
	"net/http"
	"log"
	"os"
	"path"
	"io"
	"encoding/json"
)

type ImagesHandler struct {
	dir string
	prefix string
}

func (self *ImagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}


func (self *ImagesHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	tag := foolhttp.RouteArgument(r, "tag")
	file, header, err := r.FormFile("file")
	if err != nil {
		return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	}
	defer file.Close()
	dstDir := path.Join(self.dir, tag)
	err = os.MkdirAll(dstDir, 0777)
	if err != nil {
		return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	}
	log.Printf("mkdir: %s", dstDir)
	dstFile, err := os.OpenFile(path.Join(dstDir, header.Filename), os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	}
	defer dstFile.Close()
	io.Copy(dstFile, file)

	res := struct {
		Path string `json:"path"`
	}{
		Path: path.Join(self.prefix, tag, header.Filename),
	}
	body, _ := json.Marshal(&res)
	w.Write(body)
	return nil
}

func NewImagesHandler(dir, prefix string) *ImagesHandler {
	handler := &ImagesHandler{
		dir: dir,
		prefix: prefix,
	}
	return handler
}
