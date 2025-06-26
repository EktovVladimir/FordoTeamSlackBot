package db

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
)

type CodeReview struct {
	Id                types.UniqId `json:"id" bson:"_id"`
	ThreadTs          string       `json:"thread_ts" bson:"thread_ts"`
	PullRequestNumber string       `json:"pull_request_number" bson:"pull_request_number"`
	Status            string       `json:"status" bson:"status"`
	AuditableFields   `bson:",inline"`
}

func (u *CodeReview) GetId() types.UniqId {
	return u.Id
}

func (u *CodeReview) SetId(id types.UniqId) {
	u.Id = id
}
