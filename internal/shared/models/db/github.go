package db

import "github.com/uptrace/bun"

type PullRequest struct {
	bun.BaseModel `bun:"table:pull_requests,alias:pr"`
	UniqFields    `bun:",embed"`

	Owner  string `bun:",notnull"`
	Repo   string `bun:",notnull"`
	Number int    `bun:",notnull"`
}
