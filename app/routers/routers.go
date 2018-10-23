package routers

import (
	"github.com/gorilla/mux"
	"../handlers"
)

func NewRouter() *mux.Router {

	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.KeepContext = true
	mainRouter.Methods("GET").Path("/api/info").HandlerFunc(handlers.GetAPIInfo)
	mainRouter.Methods("GET").Path("/api/users").HandlerFunc(handlers.GetAllUsersHandler)
	mainRouter.Methods("POST").Path("/api/users").HandlerFunc(handlers.CreateUserHandler)
	mainRouter.Methods("GET").Path("/api/users/{id}").HandlerFunc(handlers.GetUserByIdHandler)
	mainRouter.Methods("PUT").Path("/api/users/{id}").HandlerFunc(handlers.UpdateUserHandler)
	mainRouter.Methods("DELETE").Path("/api/users/{id}").HandlerFunc(handlers.DeleteUserHandler)

	return mainRouter
}