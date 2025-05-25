package shared

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/github"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/slack"
)

type UserLink struct {
	SlackUser  slack.User
	GithubUser github.User
}
