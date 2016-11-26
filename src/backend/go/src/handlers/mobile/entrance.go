package mobile

import "github.com/gorilla/mux"

func Register(router *mux.Router) error {
	mobileRouter := router.PathPrefix("/mobile").Subrouter()


	layoutRouter := mobileRouter.PathPrefix("/layout").Subrouter()

	homeLayoutHandler := NewHomeLayoutHandler()
	layoutRouter.Handle("/home", homeLayoutHandler)

	channelLayoutHandler := NewChannelLayoutHandler()
	layoutRouter.Handle("/channels/{id:[^/]+}", channelLayoutHandler)

	entryLayoutHandler := NewEntryLayoutHandler()
	layoutRouter.Handle("/entries/{id:[^/]+}", entryLayoutHandler)


	//clientRouter := mobileRouter.PathPrefix("/clients").Subrouter()
	mobileRouter.Handle("/issues", NewIssueHandler())

	return nil
}
