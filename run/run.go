package run

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service"
	"sync"
)

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
