package user

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *Handler {
	return &Handler{repo}
}

func (c *Handler) SetupRoutes(router *mux.Router) {
	sub := router.PathPrefix("/users").Subrouter()
	sub.HandleFunc("", c.getUsers).Methods("GET")
	sub.HandleFunc("/{id}", c.getUser).Methods("GET")
	sub.HandleFunc("", c.createUser).Methods("POST")
	sub.HandleFunc("/{id}", c.updateUser).Methods("PUT")
	sub.HandleFunc("/{id}", c.deleteUser).Methods("DELETE")
}

func (c *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	data, err := c.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respData []api.UserResponse
	if err := api_helper.MapAndResponse(w, &data, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	data, err := c.repo.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

	if err := c.repo.Create(r.Context(), &dbModel); err != nil {
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

	dbModel, err := c.repo.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := api_helper.MapIgnoreEmpty(&req, dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := c.repo.Update(r.Context(), dbModel); err != nil {
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

	_, err = c.repo.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := c.repo.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
