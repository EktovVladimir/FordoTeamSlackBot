package models

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
)

type CodeReviewFact struct {
	Id        db.UniqId
	Requester *User
	Reviewers []*User
	Issues    []*JiraIssue
	Prs       []*PullRequestRef
	Message   *CreatedMessage
}
