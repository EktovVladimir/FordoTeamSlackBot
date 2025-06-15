package crud_api

import (
	"encoding/json"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (a *CrudApi) getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := a.db.GetUsers().GetAll()

	_ = json.NewEncoder(w).Encode(data)
}

func (a *CrudApi) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user *db.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logrus.Errorf("Error decoding user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	entity := a.db.GetUsers()

	entity.Insert(user)

	_ = json.NewEncoder(w).Encode(user)
}
