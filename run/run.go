package run

import (
	"context"
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

func RunHw16(ctx context.Context) *sync.WaitGroup {
	wg := new(sync.WaitGroup)

	ch := service.StartGenerators(
		ctx,
		wg,
		service.NewSlackEventGenerator(-1, 1000),
		service.NewGithubEventGenerator(10, 500))

	service.StartEventReader(ch, wg)
	service.StartEventLogger(ctx, 200)

	return wg
}
