package review

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/issues"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/shared"
)

type Starter struct {
	Requester    shared.UserLink
	Reviewers    []shared.UserLink
	PullRequests []git.PullRequest
	Issues       []issues.Issue
}
