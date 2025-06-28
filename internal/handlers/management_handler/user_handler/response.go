package user_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
)

// UserResponse model
// @Description Пользователь в системе
type UserResponse struct {
	// @Description Уникальный Id
	Id db.UniqId `json:"id"`
	// @Description Имя пользователя Slack
	SlackName string `json:"slack_name"`
	// @Description Имя пользователя в Github
	GithubName string `json:"github_name"`
	// @Description Рабочий email
	Email string `json:"email"`
}
