package api

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type CreateDeploymentRequest struct {
	ThreadTs          string `json:"thread_ts" validate:"required"`
	PullRequestNumber string `json:"pull_request_number" validate:"required,numeric"`
	WorkflowRunId     string `json:"workflow_run_id" validate:"omitempty"`
	Status            string `json:"status" validate:"required,oneof=pending success failed running"`
}

type UpdateDeploymentRequest struct {
	ThreadTs          *string `json:"thread_ts" validate:"omitempty,slack_ts"`
	PullRequestNumber *string `json:"pull_request_number" validate:"omitempty,numeric"`
	WorkflowRunId     *string `json:"workflow_run_id" validate:"omitempty"`
	Status            *string `json:"status" validate:"omitempty,oneof=pending success failed running"`
}

type DeploymentResponse struct {
	Id                types.UniqId `json:"id"`
	ThreadTs          string       `json:"thread_ts"`
	PullRequestNumber string       `json:"pull_request_number"`
	WorkflowRunId     string       `json:"workflow_run_id"`
	Status            string       `json:"status"`
}
