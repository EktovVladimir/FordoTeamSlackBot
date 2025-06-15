package management_api

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/middlewares"
	"github.com/gorilla/mux"
)

func (a *ManagementApi) setupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	apiRouter := r.PathPrefix("/management").Subrouter()

	a.userHandler.SetupRoutes(apiRouter)

	return r
}
