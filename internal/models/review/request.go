package review

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/github"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/slack"
)

type Request struct {
	requester    slack.User
	reviewers    []slack.User
	pullRequests []github.PullRequest
	randomCount  int
}

func (r *Request) Requester() slack.User {
	return r.requester
}

func (r *Request) Reviewers() []slack.User {
	return r.reviewers
}

func (r *Request) PullRequests() []github.PullRequest {
	return r.pullRequests
}

func (r *Request) SetRandomReviewersCount(randomCount int) {
	r.randomCount = randomCount
}

func (r *Request) SetReviewers(reviewers []slack.User) {
	r.reviewers = reviewers
}

func NewRequest(requester slack.User, prs []github.PullRequest) *Request {
	return &Request{
		requester:    requester,
		pullRequests: prs,
		reviewers:    make([]slack.User, 0),
	}
}
