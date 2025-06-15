package deploy

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/issues"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/shared"
)

type Starter struct {
	Author      shared.UserLink
	Issues      []issues.Issue
	PullRequest git.PullRequest
}
