package crud_api

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/middlewares"
	"github.com/gorilla/mux"
)

func (a *CrudApi) setupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	apiRouter := r.PathPrefix("/api").Subrouter()

	usersRouter := apiRouter.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", a.getUsers).Methods("GET")
	usersRouter.HandleFunc("/{id}", a.getUser).Methods("GET")
	usersRouter.HandleFunc("", a.createUser).Methods("POST")
	usersRouter.HandleFunc("/{id}", a.updateUser).Methods("PUT")
	usersRouter.HandleFunc("/{id}", a.deleteUser).Methods("DELETE")

	return r
}
