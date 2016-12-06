package accounts

import "github.com/gorilla/mux"

func Register(router *mux.Router) error {
	accountsRouter := router.PathPrefix("/accounts").Subrouter()

	accountsRouter.Handle("/login", NewLoginHandler())
	accountsRouter.Handle("/logout", NewLogoutHandler())
	accountsRouter.Handle("/auth", NewAuthHandler())

	return nil
}
