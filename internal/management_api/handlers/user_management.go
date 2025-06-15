package handlers

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type UserManagement struct {
	repo repository.UserRepository
}

func NewUserManagement(repo repository.UserRepository) *UserManagement {
	return &UserManagement{repo}
}

func (h *UserManagement) SetupRoutes(router *mux.Router) {
	sub := router.PathPrefix("/users").Subrouter()
	sub.HandleFunc("", h.getUsers).Methods("GET")
	sub.HandleFunc("/{id}", h.getUser).Methods("GET")
	sub.HandleFunc("", h.createUser).Methods("POST")
	sub.HandleFunc("/{id}", h.updateUser).Methods("PUT")
	sub.HandleFunc("/{id}", h.deleteUser).Methods("DELETE")
}

func (h *UserManagement) getUsers(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.GetAll(r.Context())
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

func (h *UserManagement) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	data, err := h.repo.GetById(r.Context(), id)
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

func (h *UserManagement) createUser(w http.ResponseWriter, r *http.Request) {
	var req api.CreateUserRequest
	var dbModel db.User
	if err := api_helper.ValidateAndMapRequest(r.Body, &req, &dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := h.repo.Create(r.Context(), &dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.UserResponse
	if err := api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserManagement) updateUser(w http.ResponseWriter, r *http.Request) {
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

	dbModel, err := h.repo.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := api_helper.MapIgnoreEmpty(&req, dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := h.repo.Update(r.Context(), dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.UserResponse
	if err := api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserManagement) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	_, err = h.repo.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
