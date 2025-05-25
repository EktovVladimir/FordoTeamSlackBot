package deploy

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/github"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/jira"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/shared"
)

type Starter struct {
	Author      shared.UserLink
	JiraIssues  []jira.Issue
	PullRequest github.PullRequest
}
