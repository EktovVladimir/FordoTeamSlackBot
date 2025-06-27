package db

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"github.com/uptrace/bun"
)

type CodeReview struct {
	bun.BaseModel `bun:"table:code_reviews,alias:cr" bson:"-"`

	Id                types.UniqId `json:"id" bson:"_id" bun:",pk,autoincrement"`
	ThreadTs          string       `json:"thread_ts" bson:"thread_ts" bun:",notnull"`
	PullRequestNumber string       `json:"pull_request_number" bson:"pull_request_number" bun:",notnull"`
	Status            string       `json:"status" bson:"status" bun:",notnull"`
	AuditableFields   `bson:",inline" bun:",embed"`
}

func (u *CodeReview) GetId() types.UniqId {
	return u.Id
}

func (u *CodeReview) SetId(id types.UniqId) {
	u.Id = id
}
