package accounts

import (
	"content"
	"content/definition"
	"encoding/json"
	"foolhttp"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
	"net/http"
)

type LoginHandler struct{}

func (self *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *LoginHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	type argsDefine struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	schemaDefine := `{
		"type": "object",
		"properties": {
			"account": {"type": "string"},
			"password": {"type": "string"}
		},
		"required": ["account", "password"]
	}`
	schemaLoader := gojsonschema.NewStringLoader(schemaDefine)
	data, e := ioutil.ReadAll(r.Body)
	log.Printf("body: %s", data)
	if e != nil {
		return foolhttp.UnknownHTTPError(e.Error())
	}
	sourceLoader := gojsonschema.NewBytesLoader(data)
	result, e := gojsonschema.Validate(schemaLoader, sourceLoader)
	if e != nil {
		return foolhttp.UnknownHTTPError(e.Error())
	}
	if !result.Valid() {
		for _, e := range result.Errors() {
			log.Printf("error: %#v", e)
			return foolhttp.BadArgHTTPError(e.String())
		}
	}

	mgr := content.GetManager()
	args := argsDefine{}
	//err := foolhttp.ParseJsonArgs(r, &args)
	e0 := json.Unmarshal([]byte(data), &args)
	if e0 != nil {
		return foolhttp.UnknownHTTPError(e0.Error())
	}
	session, e1 := mgr.AccountsLogin(args.Account, args.Password)
	if e1 != nil {
		switch {
		case e1.Err == definition.C_ERR_ACCOUNT_MISSING || e1.Err == definition.C_ERR_PASSWORD_WRONG:
			return foolhttp.NewHTTPError(401, e1.Err, e1.Detail)
		default:
			return foolhttp.UnknownHTTPError(e1.Error())
		}
	}
	cookie := http.Cookie{
		Name:  "session",
		Value: session,
		Path: "/",
	}
	http.SetCookie(w, &cookie)
	account, e2 := mgr.AccountsAuthSession(session)
	if e2 != nil {
		return foolhttp.UnknownHTTPError(e2.Error())
	}
	foolhttp.WriteJson(w, account)
	return nil
}

func NewLoginHandler() *LoginHandler {
	handler := new(LoginHandler)
	return handler
}

type LogoutHandler struct{}

func (self *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *LogoutHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	me := Authorize(r)
	mgr := content.GetManager()
	mgr.AccountsLogout(me.Session)
	return nil
}

func NewLogoutHandler() *LogoutHandler {
	handler := new(LogoutHandler)
	return handler
}
