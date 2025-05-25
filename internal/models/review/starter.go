package review

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/github"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/jira"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/shared"
)

type Starter struct {
	Requester    shared.UserLink
	Reviewers    []shared.UserLink
	PullRequests []github.PullRequest
	JiraIssues   []jira.Issue
}
