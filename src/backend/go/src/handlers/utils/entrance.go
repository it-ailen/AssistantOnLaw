package utils

import "github.com/gorilla/mux"

func Register(router *mux.Router, ugcDir, ugcPrefix string) error {
	etcRouter := router.PathPrefix("/utils").Subrouter()

	imagesHandler := NewImagesHandler(ugcDir, ugcPrefix)
	etcRouter.Handle("/images/{tag:[^/]+}", imagesHandler)

	filesHandler := NewFilesHandler(ugcDir, ugcPrefix)
	etcRouter.Handle("/files/{tag:[^/]+}", filesHandler)

	return nil
}
