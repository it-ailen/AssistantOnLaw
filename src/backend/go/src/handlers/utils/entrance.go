package utils

import "github.com/gorilla/mux"

func Register(router *mux.Router, ugcDir, ugcPrefix string) error {
	adminRouter := router.PathPrefix("/utils").Subrouter()

	imagesHandler := NewImagesHandler(ugcDir, ugcPrefix)
	adminRouter.Handle("/images/{tag:[^/]+}", imagesHandler)

	return nil
}
