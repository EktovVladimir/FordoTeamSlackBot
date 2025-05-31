package run

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/request_source"
)

func RunTest1(app *AppContext) {
	dummyGhRepo, _ := app.Infrastructure.PrService.Get("FordoTeamSlackBot", "1")
	source := request_source.Workflow{
		Repo:   "FordoTeamSlackBot",
		Branch: "develop",
		Team:   "backoffice",
	}

	app.DeployService.CreateThread(dummyGhRepo, source)
}
