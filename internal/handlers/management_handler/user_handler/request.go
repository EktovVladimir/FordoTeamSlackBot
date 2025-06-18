package user_handler

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
