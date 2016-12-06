package accounts

import (
    "net/http"
    "foolhttp"
)

type AuthHandler struct{}

func (self *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *AuthHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	me := Authorize(r)
    foolhttp.WriteJson(w, me)
	return nil
}

func NewAuthHandler() *AuthHandler {
	handler := new(AuthHandler)
	return handler
}
