package request_source

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/messenger"

type Slack struct {
	User    messenger.User
	Channel string
}

func (w Slack) GetKey() string {
	return w.Channel
}
