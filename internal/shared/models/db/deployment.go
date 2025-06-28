package db

import (
	"github.com/uptrace/bun"
)

type Deployment struct {
	bun.BaseModel   `bun:"table:deployments,alias:d"`
	UniqFields      `bun:",embed"`
	AuditableFields `bun:",embed"`

	ThreadTs          string `bun:",notnull"`
	PullRequestNumber string `bun:",notnull"`
	WorkflowRunId     string `bun:",notnull"`
	Status            string `bun:",notnull"`
}
