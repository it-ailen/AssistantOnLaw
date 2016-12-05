package resource

import "github.com/gorilla/mux"

func Register(router *mux.Router) error {
	//resourceRouter := router.PathPrefix("/resources").Subrouter()

	fileHandler := NewFileHandler()
	router.Handle("/resources", fileHandler)
	router.Handle("/resources/{id:[^/]*}", fileHandler)

	router.Handle("/resources/tree/{id:[^/]*}", NewFileTreeHandler())

	return nil
}
