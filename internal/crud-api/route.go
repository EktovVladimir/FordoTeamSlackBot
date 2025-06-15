package crud_api

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/api/middleware"
	"github.com/gorilla/mux"
)

func (a *CrudApi) setupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", a.getUsers).Methods("GET")
	usersRouter.HandleFunc("", a.createUser).Methods("POST")

	return r
}
