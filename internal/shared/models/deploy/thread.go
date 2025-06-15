package deploy

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/messenger"
)

type Thread struct {
	Starter Starter
	Thread  messenger.Thread
	Status  string
}
