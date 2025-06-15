package api

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type CreateUserRequest struct {
	SlackName  string `json:"slack_name" validate:"required,min=2"`
	GithubName string `json:"github_name" validate:"omitempty,min=2"`
	Email      string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	SlackName  *string `json:"slack_name" validate:"omitempty,min=2"`
	GithubName *string `json:"github_name" validate:"omitempty,min=2"`
	Email      *string `json:"email" validate:"omitempty,email"`
}

type UserResponse struct {
	Id         types.UniqId `json:"id"`
	SlackName  string       `json:"slack_name"`
	GithubName string       `json:"github_name"`
	Email      string       `json:"email"`
}
