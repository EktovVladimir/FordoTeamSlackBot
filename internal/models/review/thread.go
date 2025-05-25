package review

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/slack"

type Thread struct {
	Starter Starter
	Thread  slack.Thread
	Status  string
}
