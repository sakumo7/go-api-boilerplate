package routers

import (
	"../handlers"
	"github.com/gorilla/mux"
)

func SetUserRouters(router *mux.Router) *mux.Router {

	router.Methods("GET").Path("/api/info").HandlerFunc(handlers.GetAPIInfo)
	router.Methods("GET").Path("/api/users").HandlerFunc(handlers.GetAllUsersHandler)
	router.Methods("POST").Path("/api/users").HandlerFunc(handlers.CreateUserHandler)
	router.Methods("GET").Path("/api/users/{id}").HandlerFunc(handlers.GetUserByIdHandler)
	router.Methods("PUT").Path("/api/users/{id}").HandlerFunc(handlers.UpdateUserHandler)
	router.Methods("DELETE").Path("/api/users/{id}").HandlerFunc(handlers.DeleteUserHandler)

	return router
}
