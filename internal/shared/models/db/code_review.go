package db

import (
	"github.com/uptrace/bun"
)

type CodeReview struct {
	bun.BaseModel   `bun:"table:code_reviews,alias:cr"`
	UniqFields      `bun:",embed"`
	AuditableFields `bun:",embed"`

	ThreadTs          string `bun:",notnull"`
	PullRequestNumber string `bun:",notnull"`
	Status            string `bun:",notnull"`
}
