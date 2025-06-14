package crud_api

import (
	"encoding/json"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dummyRes []db.User

	dummyRes = append(dummyRes, db.User{
		Id:         1,
		SlackName:  "user",
		GithubName: "user",
		Email:      "user@example.com",
	})

	dummyRes = append(dummyRes, db.User{
		Id:         2,
		SlackName:  "user2",
		GithubName: "user2",
		Email:      "user2@example.com",
	})

	json.NewEncoder(w).Encode(dummyRes)
}
