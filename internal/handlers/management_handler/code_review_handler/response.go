package code_review_handler

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type CodeReviewResponse struct {
	Id                types.UniqId `json:"id"`
	ThreadTs          string       `json:"thread_ts"`
	PullRequestNumber string       `json:"pull_request_number"`
	Status            string       `json:"status"`
}
