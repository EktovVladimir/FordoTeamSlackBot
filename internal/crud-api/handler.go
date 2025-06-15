package crud_api

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"net/http"
)

func (a *CrudApi) getUsers(w http.ResponseWriter, _ *http.Request) {
	data := a.db.GetUsers().GetAll()

	var respData []api.UserResponse
	if err := api_helper.MapAndResponse(w, &data, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *CrudApi) getUser(w http.ResponseWriter, r *http.Request) {
	entity := a.db.GetUsers()

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

func (a *CrudApi) createUser(w http.ResponseWriter, r *http.Request) {
	var req api.CreateUserRequest
	var dbModel db.User
	if err := api_helper.ValidateAndMapRequest(r.Body, &req, &dbModel); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	entity := a.db.GetUsers()
	entity.Insert(&dbModel)

	if err := a.db.SaveChanges(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.UserResponse
	if err := api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *CrudApi) updateUser(w http.ResponseWriter, r *http.Request) {
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

	entity := a.db.GetUsers()
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

	if err := a.db.SaveChanges(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var respModel api.UserResponse
	if err := api_helper.MapAndResponse(w, &dbModel, &respModel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *CrudApi) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := api_helper.GetUniqIdFromRoute(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	entity := a.db.GetUsers()
	_, found := entity.Get(id)
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	entity.Delete(id)

	if err := a.db.SaveChanges(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
