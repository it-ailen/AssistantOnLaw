/*
	This package aims to create a common framework for HTTP server.
 */
package foolhttp

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"runtime/debug"
	"strconv"
)

const (
	ErrorOK = "OK"
	ErrorBadArgument = "BadArgument"
	ErrorMethodNotAllowed = "MethodNotAllowed"
	ErrorNotFound = "ResourceNotFound"
	ErrorInvalidJson = "InvalidJson"
	ErrorUnknown = "Unknown"
	ErrorUnknownResource = "UnknownResource"
	ErrorForbidden = "Forbidden"
	ErrorThirdParty = "ThirdPartyError"
)


type HTTPError struct {
	StatusCode int `json:"-"`   /* Hide the status code from body. */
	ErrorCode string `json:"error"`
	Detail string `json:"detail,omitempty"`
}

func (self *HTTPError) Error() string {
	return fmt.Sprintf("%d: %s-%s", self.StatusCode, self.ErrorCode, self.Detail)
}

func (self *HTTPError) JSON() []byte {
	res, err := json.Marshal(self)
	if err != nil {
		panic(err)
	}
	return res
}

func NewHTTPError(code int, err string, detail string) (*HTTPError) {
	return &HTTPError{
		StatusCode: code,
		ErrorCode: err,
		Detail: detail,
	}
}

func NewDefaultHTTPError() (*HTTPError) {
	return &HTTPError{
		StatusCode: 200,
		ErrorCode: ErrorOK,
	}
}

func BadArgHTTPError(detail string) (*HTTPError) {
	return NewHTTPError(400, ErrorBadArgument, detail)
}

func ForbiddenHTTPError(detail string) (*HTTPError) {
	return NewHTTPError(403, ErrorForbidden, detail)
}

func NotFoundHTTPError(detail string) (*HTTPError) {
	return NewHTTPError(404, ErrorNotFound, detail)
}

func UnknownHTTPError(detail string) (*HTTPError) {
	return NewHTTPError(500, ErrorUnknown, detail)
}

type HttpGetHandler interface {
	GET(http.ResponseWriter, *http.Request) (*HTTPError)
}

type HttpPostHandler interface {
	POST(http.ResponseWriter, *http.Request) (*HTTPError)
}

type HttpDeleteHandler interface {
	DELETE(http.ResponseWriter, *http.Request) (*HTTPError)
}

type HttpPutHandler interface {
	PUT(http.ResponseWriter, *http.Request) (*HTTPError)
}

type HttpHeadHandler interface {
	HEAD(http.ResponseWriter, *http.Request) (*HTTPError)
}


//type HandlerFunc func (http.ResponseWriter, *http.Request)


type NotFoundHandler struct {}

func (self *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := HTTPError{
		StatusCode: http.StatusNotFound,
		ErrorCode: ErrorNotFound,
	}
	log.Printf("%s %d %s\n", r.Method, err.StatusCode, r.RequestURI)
	w.WriteHeader(err.StatusCode)
	w.Write(err.JSON())
}

type Handler func(w http.ResponseWriter, r *http.Request)

type Method func(w http.ResponseWriter, r *http.Request) (*HTTPError)

type GetMethod interface {
	GET(w http.ResponseWriter, r *http.Request) *HTTPError
}

type PostMethod interface {
	POST(w http.ResponseWriter, r *http.Request) *HTTPError
}

type PutMethod interface {
	PUT(w http.ResponseWriter, r *http.Request) *HTTPError
}

type DeleteMethod interface {
	DELETE(w http.ResponseWriter, r *http.Request) *HTTPError
}

/**
统一执行ServeHTTP的方法，规范接口，由具体的 http.Handler 类在 ServeHTTP 方法中调用。
 */
func DoServeHTTP(self interface{}, w http.ResponseWriter, r *http.Request) {
	if _, ok := self.(http.Handler); !ok {
		panic("The instance is invalid")
	}
	w.Header()["Content-Type"] = []string{"application/json;charset=utf-8"}
	serveHTTP(self, w, r)
}

func serveHTTP(self interface{}, w http.ResponseWriter, r *http.Request) {
	var err *HTTPError
	defer func() {
		if tmp := recover(); tmp != nil {
			if e, ok := tmp.(*HTTPError); ok {
				log.Println(e.Error())
				debug.PrintStack()
				w.WriteHeader(e.StatusCode)
				w.Write(e.JSON())
				log.Printf("%s %d %s\n", r.Method, e.StatusCode, r.RequestURI)
			} else if e, ok := tmp.(error); ok {
				log.Println(e.Error())
				debug.PrintStack()
				fe := NewHTTPError(http.StatusInternalServerError, ErrorUnknown, e.Error())
				w.WriteHeader(fe.StatusCode)
				w.Write(fe.JSON())
				log.Printf("%s %d %s\n", r.Method, fe.StatusCode, r.RequestURI)
			} else {
				debug.PrintStack()
				e := NewHTTPError(http.StatusInternalServerError, ErrorUnknown, "")
				w.WriteHeader(e.StatusCode)
				w.Write(e.JSON())
				log.Printf("%s %d %s\n", r.Method, e.StatusCode, r.RequestURI)
			}
		}
	}()
	err = &HTTPError{
		StatusCode: 405,
		ErrorCode: ErrorMethodNotAllowed,
		Detail: fmt.Sprintf("Method `%s` isn't allowed", r.Method),
	}
	switch r.Method {
	case "GET":
		if h, ok := self.(GetMethod); ok {
			err = h.GET(w, r)
		}
	case "POST":
		if h, ok := self.(PostMethod); ok {
			err = h.POST(w, r)
		}
	case "PUT":
		if h, ok := self.(PutMethod); ok {
			err = h.PUT(w, r)
		}
	case "DELETE":
		if h, ok := self.(DeleteMethod); ok {
			err = h.DELETE(w, r)
		}

	}
	if err != nil && err.StatusCode != 200 {
		log.Printf("%d %s %s: %s - %s\n", err.StatusCode, r.Method, r.RequestURI, err.ErrorCode, err.Detail)
		w.WriteHeader(err.StatusCode)
		w.Write(err.JSON())
	} else {
		log.Printf("%d %s %s\n", 200, r.Method, r.RequestURI)
	}
}

func routeArgumentWithDefault(r *http.Request, key string, defaults interface{}) interface{} {
	vars := mux.Vars(r)
	log.Printf("vars: %#v -- key: %s", vars, key)
	if val, ok := vars[key]; ok {
		return val
	} else {
		return defaults
	}
}

func RouteArgumentWithDefault(r *http.Request, key string, defaults string) string {
	val := routeArgumentWithDefault(r, key, nil)
	if val == nil {
		return defaults
	}
	s, _ := val.(string)
	return s
}

func RouteArgument(r *http.Request, key string) string {
	val := routeArgumentWithDefault(r, key, nil)
	log.Printf("val:::: %v", val)
	if val == nil {
		panic(&HTTPError{
			StatusCode: http.StatusNotFound,
			ErrorCode: ErrorNotFound,
		})
	}
	s, _ := val.(string)
	return s
}

func queryArgumentWithDefault(r *http.Request, key string, defaults interface{}) interface{} {
	queryMaps := r.URL.Query()
	if values, ok := queryMaps[key]; ok {
		if len(values) > 0 {
			return string(values[0])
		}
	}
	return defaults
}

func QueryArgument(r *http.Request, key string) (string) {
	val := queryArgumentWithDefault(r, key, nil)
	if val == nil {
		panic(&HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorCode: ErrorBadArgument,
			Detail: fmt.Sprintf("Missing args: %s", key),
		})
	}
	s, _ := val.(string)
	return s
}

func QueryLongArgument(r *http.Request, key string) (int64) {
	s := QueryArgument(r, key)
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(&HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorCode: ErrorBadArgument,
			Detail: fmt.Sprintf("Wrong value of long type: %s", s),
		})
	}
	return res
}

func QueryBoolArgument(r *http.Request, key string) (bool) {
	s := QueryArgument(r, key)
	res, err := strconv.ParseBool(s)
	if err != nil {
		panic(&HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorCode: ErrorBadArgument,
			Detail: fmt.Sprintf("Wrong value of boolean type: %s", s),
		})
	}
	return res
}

func QueryArgumentWithDefault(r *http.Request, key string, defaults string) (string) {
	val := queryArgumentWithDefault(r, key, interface{}(defaults))
	s, _ := val.(string)
	return s
}

func QueryLongArgumentWithDefault(r *http.Request, key string, defaults int64) (int64) {
	s := QueryArgumentWithDefault(r, key, strconv.FormatInt(defaults, 10))
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(&HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorCode: ErrorBadArgument,
			Detail: fmt.Sprintf("Wrong value of long type: %s", s),
		})
	}
	return res
}

func QueryBoolArgumentWithDefault(r *http.Request, key string, defaults bool) (bool) {
	s := QueryArgumentWithDefault(r, key, strconv.FormatBool(defaults))
	res, err := strconv.ParseBool(s)
	if err != nil {
		panic(&HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorCode: ErrorBadArgument,
			Detail: fmt.Sprintf("Wrong value of boolean type: %s", s),
		})
	}
	return res
}

func formArgumentWithDefault(r *http.Request, key string, defaults interface{}) (interface{}) {
	log.Printf("form: %#v", r.Form)
	if values, ok := r.Form[key]; ok {
		if len(values) > 0 {
			return interface{}(values[0])
		}
	}
	return defaults
}

func FormArgument(r *http.Request, key string) (string) {
	val := formArgumentWithDefault(r, key, nil)
	if val == nil {
		panic(&HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorCode: ErrorBadArgument,
			Detail: fmt.Sprintf("Missing args: %s", key),
		})
	}
	s, _ := val.(string)
	return s
}

func FormArgumentWithDefault(r *http.Request, key string, defaults string) string {
	val := formArgumentWithDefault(r, key, interface{}(defaults))
	s, _ := val.(string)
	return s
}

func ParseJsonArgs(r *http.Request, dst interface{}) (*HTTPError) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dst)
	if err != nil {
		return &HTTPError{
			StatusCode: 400,
			ErrorCode: ErrorInvalidJson,
			Detail: err.Error(),
		}
	}
	return nil
}

func WriteJson(w http.ResponseWriter, src interface{}) {
	headers := w.Header()
	headers["Content-Type"] = []string{"application/json;charset=utf-8"}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	log.Printf("write: %#v %s", src, string(data))
	w.Write(data)
}
