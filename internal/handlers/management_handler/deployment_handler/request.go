package deployment_handler

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
