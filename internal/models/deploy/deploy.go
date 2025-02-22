package deploy

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/jira"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/slack"
)

type Thread struct {
	starter slack.User
	issues  []jira.Issue
	Title   string
	Status  Status
}

func (t *Thread) Issues() []jira.Issue {
	return t.issues
}

func (t *Thread) SetIssues(issues []jira.Issue) {
	t.issues = issues
}

func NewThread(starter slack.User) *Thread {
	return &Thread{
		starter: starter,
		issues:  make([]jira.Issue, 0),
	}
}
