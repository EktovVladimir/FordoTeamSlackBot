package models

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type CodeReviewFact struct {
	Id        types.UniqId
	Requester *User
	Reviewers []*User
	Issues    []*JiraIssue
	Prs       []*PullRequestRef
	Message   *CreatedMessage
}
