package counsel

import "github.com/gorilla/mux"

func Register(router *mux.Router) error {
	counselRouter := router.PathPrefix("/counsels").Subrouter()

	classHandler := NewClassHandler()
	counselRouter.Handle("/classes", classHandler)
	counselRouter.Handle("/classes/{id:[^/]*}", classHandler)

	entryHandler := NewEntryHandler()
	counselRouter.Handle("/entries", entryHandler)
	counselRouter.Handle("/entries/{id:[^/]*}", entryHandler)

	questionHandler := NewQuestionHandler()
	counselRouter.Handle("/questions", questionHandler)
	counselRouter.Handle("/questions/{id:[^/]*}", questionHandler)

	conclusionHandler := NewConclusionHandler()
	counselRouter.Handle("/conclusions", conclusionHandler)
	counselRouter.Handle("/conclusions/{id:[^/]*}", conclusionHandler)
	return nil
}
