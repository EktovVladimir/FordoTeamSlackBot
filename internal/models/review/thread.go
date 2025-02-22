package review

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/common"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/jira"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/slack"
)

type Thread struct {
	request Request
	Status  Status
	Thread  slack.TheadInfo
	Issues  []jira.Issue
	Created common.DateTime
}

func (t Thread) Request() Request {
	return t.request
}

func NewThread(request Request) *Thread {
	return &Thread{
		request: request,
		Status:  Status{},
		Issues:  make([]jira.Issue, 0),
		Created: common.Now(),
	}
}
