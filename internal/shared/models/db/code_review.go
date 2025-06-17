package db

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type CodeReview struct {
	Id                types.UniqId `json:"id"`
	ThreadTs          string       `json:"thread_ts"`
	PullRequestNumber string       `json:"pull_request_number"`
	Status            string       `json:"status"`
}

func (u *CodeReview) GetId() types.UniqId {
	return u.Id
}

func (u *CodeReview) SetId(id types.UniqId) {
	u.Id = id
}
