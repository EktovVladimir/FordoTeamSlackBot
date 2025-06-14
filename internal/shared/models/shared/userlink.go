package shared

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/messenger"
)

type UserLink struct {
	SlackUser  messenger.User
	GithubUser git.User
}
