package db

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type Deployment struct {
	Id                types.UniqId `json:"id"`
	ThreadTs          string       `json:"thread_ts"`
	PullRequestNumber string       `json:"pull_request_number"`
	WorkflowRunId     string       `json:"workflow_run_id"`
	Status            string       `json:"status"`
}

func (u *Deployment) GetId() types.UniqId {
	return u.Id
}

func (u *Deployment) SetId(id types.UniqId) {
	u.Id = id
}
