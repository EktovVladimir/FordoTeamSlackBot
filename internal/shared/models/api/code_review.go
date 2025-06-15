package api

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type CreateCodeReviewRequest struct {
	ThreadTs          string `json:"thread_ts" validate:"required"`
	PullRequestNumber string `json:"pull_request_number" validate:"required,numeric"`
	Status            string `json:"status" validate:"required,oneof=pending approved rejected"`
}

type UpdateCodeReviewRequest struct {
	ThreadTs          *string `json:"thread_ts" validate:"omitempty"`
	PullRequestNumber *string `json:"pull_request_number" validate:"omitempty,numeric"`
	Status            *string `json:"status" validate:"omitempty,oneof=pending approved rejected"`
}

type CodeReviewResponse struct {
	Id                types.UniqId `json:"id"`
	ThreadTs          string       `json:"thread_ts"`
	PullRequestNumber string       `json:"pull_request_number"`
	Status            string       `json:"status"`
}
