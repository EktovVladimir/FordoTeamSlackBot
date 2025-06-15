package review

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/issues"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/shared"
)

type Starter struct {
	Requester    shared.UserLink
	Reviewers    []shared.UserLink
	PullRequests []git.PullRequest
	Issues       []issues.Issue
}
