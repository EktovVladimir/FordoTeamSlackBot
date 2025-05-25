package review

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/messenger"
)

type Request struct {
	requester    messenger.User
	reviewers    []messenger.User
	pullRequests []git.PullRequest
	randomCount  int
}
