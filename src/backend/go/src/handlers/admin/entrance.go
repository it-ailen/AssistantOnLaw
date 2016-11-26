package admin

import "github.com/gorilla/mux"

func Register(router *mux.Router) error {
	adminRouter := router.PathPrefix("/admin").Subrouter()

	channelHandler := NewChannelHandler()
	adminRouter.Handle("/channels", channelHandler)
	adminRouter.Handle("/channels/{id:[^/]+}", channelHandler)

	entryHandler := NewEntryHandler()
	adminRouter.Handle("/entries", entryHandler)
	adminRouter.Handle("/entries/{id:[^/]+}", entryHandler)

	optionHandler := NewOptionHandler()
	adminRouter.Handle("/options", optionHandler)
	adminRouter.Handle("/options/{id:[^/]+}", optionHandler)

	issueHandler := NewIssueHandler()
	adminRouter.Handle("/issues", issueHandler)
	adminRouter.Handle("/issues/{id:[^/]+}/solutions", issueHandler)

	return nil
}
