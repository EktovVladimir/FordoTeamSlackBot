package code_review_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
)

// CodeReviewResponse model
//
//	@Description	Информация о процессе код-ревью
type CodeReviewResponse struct {
	//	@Description	Уникальный ID
	//	@Example:		1
	Id db.UniqId `json:"id"`

	//	@Description	Идентификатор треда в Slack
	//	@Example:		"1234567890.123456"
	ThreadTs string `json:"thread_ts"`

	//	@Description	Номер Pull Request
	//	@Example:		"42"
	PullRequestNumber string `json:"pull_request_number"`

	//	@Description	Статус код-ревью
	//	@Example:		"approved"
	//	@Enum			pending,approved,rejected
	Status string `json:"status"`
}
