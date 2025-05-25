package review

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/github"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/slack"
)

type Request struct {
	requester    slack.User
	reviewers    []slack.User
	pullRequests []github.PullRequest
	randomCount  int
}
