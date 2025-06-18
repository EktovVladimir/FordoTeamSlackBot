package user_handler

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type UserResponse struct {
	Id         types.UniqId `json:"id"`
	SlackName  string       `json:"slack_name"`
	GithubName string       `json:"github_name"`
	Email      string       `json:"email"`
}
