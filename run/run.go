package run

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/request_source"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service"
	"sync"
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

func RunHw16() *sync.WaitGroup {
	wg := new(sync.WaitGroup)

	ch := service.StartGenerators(
		wg,
		service.NewSlackEventGenerator(25, 50),
		service.NewGithubEventGenerator(10, 50))

	service.StartEventReader(ch, wg)
	service.StartEventLogger(200)

	return wg
}
