package crud_api

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/middlewares"
	"github.com/gorilla/mux"
)

func (a *CrudApi) setupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	apiRouter := r.PathPrefix("/api").Subrouter()

	a.userHandler.SetupRoutes(apiRouter)

	return r
}
