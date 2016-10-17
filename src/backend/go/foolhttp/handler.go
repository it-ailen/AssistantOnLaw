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
		StatusCode: http.StatusOK,
		ErrorCode: ErrorOK,
	}
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

type BaseHandler struct {
	HandlerMap map[string][]Handler
}

func (self *BaseHandler) GET(w http.ResponseWriter, r *http.Request) *HTTPError {
	return &HTTPError{
		StatusCode: 405,
		ErrorCode: ErrorMethodNotAllowed,
	}
}

func (self *BaseHandler) POST(w http.ResponseWriter, r *http.Request) *HTTPError {
	return &HTTPError{
		StatusCode: 405,
		ErrorCode: ErrorMethodNotAllowed,
	}
}

func (self *BaseHandler) PUT(w http.ResponseWriter, r *http.Request) *HTTPError {
	return &HTTPError{
		StatusCode: 405,
		ErrorCode: ErrorMethodNotAllowed,
	}
}

func (self *BaseHandler) DELETE(w http.ResponseWriter, r *http.Request) *HTTPError {
	return &HTTPError{
		StatusCode: 405,
		ErrorCode: ErrorMethodNotAllowed,
	}
}

func (self *BaseHandler) HEAD(w http.ResponseWriter, r *http.Request) *HTTPError {
	return &HTTPError{
		StatusCode: 405,
		ErrorCode: ErrorMethodNotAllowed,
	}
}

func (self *BaseHandler) Init() {
	self.HandlerMap = make(map[string][]Handler)
}

func (self *BaseHandler) RegisterMethod(method string, handlers ...Handler) {
	self.HandlerMap[method] = handlers
}

func (self *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err *HTTPError
	defer func() {
		if tmp := recover(); tmp != nil {
			if e, ok := tmp.(error); ok {
				log.Println(e.Error())
				debug.PrintStack()
				fe := NewHTTPError(http.StatusInternalServerError, ErrorUnknown, e.Error())
				w.WriteHeader(fe.StatusCode)
				w.Write(fe.JSON())
				log.Printf("%s %d %s\n", r.Method, fe.StatusCode, r.RequestURI)
			} else {
				log.Println(e.Error())
				debug.PrintStack()
				e := NewHTTPError(http.StatusInternalServerError, ErrorUnknown, "")
				w.WriteHeader(e.StatusCode)
				w.Write(e.JSON())
				log.Printf("%s %d %s\n", r.Method, e.StatusCode, r.RequestURI)
			}
		}
	}()
	switch r.Method {
	case "GET":
		err = self.GET(w, r)
	case "POST":
		err = self.POST(w, r)
	case "PUT":
		err = self.PUT(w, r)
	case "DELETE":
		err = self.DELETE(w, r)
	case "HEAD":
		err = self.HEAD(w, r)
	default:
		err = &HTTPError{
			StatusCode: 405,
			ErrorCode: ErrorMethodNotAllowed,
		}
	}
	if err != nil && err.StatusCode != 200 {
		log.Printf("%s %s(%d): %s - %s\n", r.Method, r.RequestURI, err.StatusCode, err.ErrorCode, err.Detail)
		w.WriteHeader(err.StatusCode)
		w.WriteHeader(err.JSON())
	} else {
		log.Printf("%s %s(%d)\n", r.Method, r.RequestURI, 200)
	}
}

func routeArgumentWithDefault(r *http.Request, key string, defaults interface{}) interface{} {
	vars := mux.Vars(r)
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
	if val == nil {
		panic(HTTPError{
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
		panic(HTTPError{
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
		panic(HTTPError{
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
		panic(HTTPError{
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
		panic(HTTPError{
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
		panic(HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorCode: ErrorBadArgument,
			Detail: fmt.Sprintf("Wrong value of boolean type: %s", s),
		})
	}
	return res
}

func formArgumentWithDefault(r *http.Request, key string, defaults interface{}) (interface{}) {
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
		panic(HTTPError{
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

func ParseJsonArgs(r *http.Request, dst interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dst)
	if err != nil {
		panic(HTTPError{
			StatusCode: 400,
			ErrorCode: ErrorInvalidJson,
		})
	}
}
