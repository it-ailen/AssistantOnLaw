package main

import (
	"github.com/gorilla/mux"
	"foolhttp"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = &foolhttp.NotFoundHandler{}

	return router
}
