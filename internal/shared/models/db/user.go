package db

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type User struct {
	Id         types.UniqId `json:"id"`
	SlackName  string       `json:"slack_name"`
	GithubName string       `json:"github_name"`
	Email      string       `json:"email"`
}

func (u *User) GetId() types.UniqId {
	return u.Id
}

func (u *User) SetId(id types.UniqId) {
	u.Id = id
}
