package code_review_handler

// CreateCodeReviewRequest model
//
//	@Description	Запрос на создание код-ревью
type CreateCodeReviewRequest struct {
	//	@Description	Идентификатор треда в Slack
	//	@Example:		"1234567890.123456"
	//	@Required		true
	ThreadTs string `json:"thread_ts" validate:"required"`

	//	@Description	Номер Pull Request
	//	@Example:		"42"
	//	@Required		true
	PullRequestNumber string `json:"pull_request_number" validate:"required,numeric"`

	//	@Description	Статус код-ревью
	//	@Example:		"pending"
	//	@Enum			pending,approved,rejected
	//	@Required		true
	Status string `json:"status" validate:"required,oneof=pending approved rejected"`
}

// UpdateCodeReviewRequest model
//
//	@Description	Запрос на обновление код-ревью
type UpdateCodeReviewRequest struct {
	//	@Description	Идентификатор треда в Slack
	//	@Example:		"1234567890.123456"
	ThreadTs *string `json:"thread_ts" validate:"omitempty"`

	//	@Description	Номер Pull Request
	//	@Example:		"42"
	PullRequestNumber *string `json:"pull_request_number" validate:"omitempty,numeric"`

	//	@Description	Статус код-ревью
	//	@Example:		"approved"
	//	@Enum			pending,approved,rejected
	Status *string `json:"status" validate:"omitempty,oneof=pending approved rejected"`
}
