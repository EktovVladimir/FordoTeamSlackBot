package user

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/gorilla/mux"
	"net/http"
)

type store interface {
	GetUsers() *db_adapter.Entity[*db.User]
	SaveChanges() error
}

type Handler struct {
	db store
}

func New(db store) *Handler {
	return &Handler{db}
}

func (c *Handler) SetupRoutes(router *mux.Router) {
	usersRouter := router.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", c.getUsers).Methods("GET")
	usersRouter.HandleFunc("/{id}", c.getUser).Methods("GET")
	usersRouter.HandleFunc("", c.createUser).Methods("POST")
	usersRouter.HandleFunc("/{id}", c.updateUser).Methods("PUT")
	usersRouter.HandleFunc("/{id}", c.deleteUser).Methods("DELETE")
}

func (c *Handler) getUsers(w http.ResponseWriter, _ *http.Request) {
	data := c.db.GetUsers().GetAll()

	var respData []api.UserResponse
	if err := api_helper.MapAndResponse(w, &data, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	entity := c.db.GetUsers()

	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	data, found := entity.Get(id)
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var respData api.UserResponse
	if err := api_helper.MapAndResponse(w, &data, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var req api.CreateUserRequest
	var dbModel db.User
	if err := api_helper.ValidateAndMapRequest(r.Body, &req, &dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	entity := c.db.GetUsers()
	entity.Insert(&dbModel)

	if err := c.db.SaveChanges(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.UserResponse
	if err := api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var req api.UpdateUserRequest
	if err := api_helper.ValidateRequest(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	entity := c.db.GetUsers()
	dbModel, found := entity.Get(id)
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := api_helper.MapIgnoreEmpty(&req, dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	entity.Update(dbModel)

	if err := c.db.SaveChanges(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.UserResponse
	if err := api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	entity := c.db.GetUsers()
	_, found := entity.Get(id)
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	entity.Delete(id)

	if err := c.db.SaveChanges(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
