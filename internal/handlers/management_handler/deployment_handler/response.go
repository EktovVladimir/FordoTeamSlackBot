package deployment_handler

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type DeploymentResponse struct {
	Id                types.UniqId `json:"id"`
	ThreadTs          string       `json:"thread_ts"`
	PullRequestNumber string       `json:"pull_request_number"`
	WorkflowRunId     string       `json:"workflow_run_id"`
	Status            string       `json:"status"`
}
