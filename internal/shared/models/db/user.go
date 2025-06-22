package db

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type User struct {
	Id         types.UniqId `json:"id"`
	SlackName  string       `json:"slack_name"`
	SlackId    string       `json:"slack_id"`
	GithubName string       `json:"github_name"`
	Email      string       `json:"email"`
}

func (u *User) GetId() types.UniqId {
	return u.Id
}

func (u *User) SetId(id types.UniqId) {
	u.Id = id
}

func (u *User) IsFullFilled() bool {
	return u.SlackName != "" && u.GithubName != "" && u.Email != ""
}
