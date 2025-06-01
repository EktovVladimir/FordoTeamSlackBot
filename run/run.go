package run

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service"
	"sync"
)

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
