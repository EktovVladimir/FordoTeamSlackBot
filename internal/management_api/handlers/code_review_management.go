package handlers

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/gorilla/mux"
	"net/http"
)

type CodeReviewsManagement struct {
	repo repository.CodeReviewRepository
}

func NewCodeReviewsManagement(repo repository.CodeReviewRepository) *CodeReviewsManagement {
	return &CodeReviewsManagement{repo}
}

func (h *CodeReviewsManagement) SetupRoutes(router *mux.Router) {
	sub := router.PathPrefix("/reviews").Subrouter()
	sub.HandleFunc("", h.getCodeReviews).Methods("GET")
	sub.HandleFunc("/{id}", h.getCodeReview).Methods("GET")
	sub.HandleFunc("", h.createCodeReview).Methods("POST")
	sub.HandleFunc("/{id}", h.updateCodeReview).Methods("PUT")
	sub.HandleFunc("/{id}", h.deleteCodeReview).Methods("DELETE")
}

func (h *CodeReviewsManagement) getCodeReviews(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respData []api.CodeReviewResponse
	if err := api_helper.MapAndResponse(w, &data, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CodeReviewsManagement) getCodeReview(w http.ResponseWriter, r *http.Request) {
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

	var respData api.CodeReviewResponse
	if err := api_helper.MapAndResponse(w, &data, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CodeReviewsManagement) createCodeReview(w http.ResponseWriter, r *http.Request) {
	var req api.CreateCodeReviewRequest
	var dbModel db.CodeReview
	if err := api_helper.ValidateAndMapRequest(r.Body, &req, &dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if err := h.repo.Create(r.Context(), &dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.CodeReviewResponse
	if err := api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CodeReviewsManagement) updateCodeReview(w http.ResponseWriter, r *http.Request) {
	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var req api.UpdateCodeReviewRequest
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

	var respModel api.CodeReviewResponse
	if err = api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *CodeReviewsManagement) deleteCodeReview(w http.ResponseWriter, r *http.Request) {
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
