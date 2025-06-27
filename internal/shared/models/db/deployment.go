package db

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"github.com/uptrace/bun"
)

type Deployment struct {
	bun.BaseModel `bun:"table:deployments,alias:d" bson:"-"`

	Id                types.UniqId `json:"id" bson:"_id" bun:",pk,autoincrement"`
	ThreadTs          string       `json:"thread_ts" bson:"thread_ts" bun:",notnull"`
	PullRequestNumber string       `json:"pull_request_number" bson:"pull_request_number" bun:",notnull"`
	WorkflowRunId     string       `json:"workflow_run_id" bson:"workflow_run_id" bun:",notnull"`
	Status            string       `json:"status" bson:"status" bun:",notnull"`
	AuditableFields   `bson:",inline" bun:",embed"`
}

func (u *Deployment) GetId() types.UniqId {
	return u.Id
}

func (u *Deployment) SetId(id types.UniqId) {
	u.Id = id
}
