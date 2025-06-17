package handlers

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type SettingsManagement struct {
	repo repository.SettingsRepository
}

func NewSettingsManagement(repo repository.SettingsRepository) *SettingsManagement {
	return &SettingsManagement{repo}
}

func (h *SettingsManagement) SetupRoutes(router *mux.Router) {
	sub := router.PathPrefix("/settings").Subrouter()
	sub.HandleFunc("", h.getSettings).Methods("GET")
	sub.HandleFunc("/{id}", h.getSetting).Methods("GET")
	sub.HandleFunc("", h.createSetting).Methods("POST")
	sub.HandleFunc("/{id}", h.updateSetting).Methods("PUT")
	sub.HandleFunc("/{id}", h.deleteSetting).Methods("DELETE")
}

func (h *SettingsManagement) getSettings(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respData []api.SettingResponse
	if err := api_helper.MapAndResponse(w, &data, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SettingsManagement) getSetting(w http.ResponseWriter, r *http.Request) {
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

	var respData api.SettingResponse
	if err := api_helper.MapAndResponse(w, &data, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SettingsManagement) createSetting(w http.ResponseWriter, r *http.Request) {
	var req api.CreateSettingRequest
	var dbModel db.Setting
	if err := api_helper.ValidateAndMapRequest(r.Body, &req, &dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := h.repo.Create(r.Context(), &dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.SettingResponse
	if err := api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SettingsManagement) updateSetting(w http.ResponseWriter, r *http.Request) {
	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var req api.UpdateSettingRequest
	if err = api_helper.ValidateRequest(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	dbModel, err := h.repo.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err = api_helper.MapIgnoreEmpty(&req, dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err = h.repo.Update(r.Context(), dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.SettingResponse
	if err = api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *SettingsManagement) deleteSetting(w http.ResponseWriter, r *http.Request) {
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

	if err = h.repo.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
