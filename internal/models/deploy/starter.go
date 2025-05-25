package deploy

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/issues"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/shared"
)

type Starter struct {
	Author      shared.UserLink
	Issues      []issues.Issue
	PullRequest git.PullRequest
}
