package user_handler

import "time"

// CreateUserRequest model
// @Description Запрос на создание пользователя
type CreateUserRequest struct {
	// @Description Имя пользователя Slack
	SlackName string `json:"slack_name" validate:"required,min=2"`
	// @Description Имя пользователя в Github
	GithubName string `json:"github_name" validate:"omitempty,min=2"`
	// @Description Рабочий email
	Email string `json:"email" validate:"required,email"`
}

// UpdateUserRequest model
// @Description Запрос на изменение пользователя
type UpdateUserRequest struct {
	// @Description Имя пользователя Slack
	SlackName *string `json:"slack_name" validate:"omitempty,min=2"`
	// @Description Имя пользователя в Github
	GithubName *string `json:"github_name" validate:"omitempty,min=2"`
	// @Description Рабочий email
	Email *string `json:"email" validate:"omitempty,email"`
}

type SearchUsersRequest struct {
	UpdatedAtFrom *time.Time `json:"updated_at_from" validate:"omitempty,ltfield=UpdatedAtTo"`
	UpdatedAtTo   *time.Time `json:"updated_at_to" validate:"omitempty,gtfield=UpdatedAtFrom"`
}
