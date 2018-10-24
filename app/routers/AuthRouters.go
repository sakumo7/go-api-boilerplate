package routers

import (
	"../core/authentication"
	"../handlers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", handlers.Login).Methods("POST")

	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(handlers.RefreshToken),
		)).Methods("GET")

	router.Handle("/logout", negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(handlers.Logout),
	)).Methods("GET")

	return router
}
