package hw

import (
	"context"
	service2 "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/service"
	"sync"
)

func RunHw16(ctx context.Context) *sync.WaitGroup {
	wg := new(sync.WaitGroup)

	ch := service2.StartGenerators(
		ctx,
		wg,
		service2.NewSlackEventGenerator(-1, 1000),
		service2.NewGithubEventGenerator(10, 500))

	service2.StartEventReader(ch, wg)
	//service.StartEventLogger(ctx, 200)

	return wg
}
