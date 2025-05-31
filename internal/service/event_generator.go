package service

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/repository"
	"sync"
)

type eventGenerator interface {
	Start() <-chan repository.Event
}

func StartGenerators(wg *sync.WaitGroup, generators ...eventGenerator) <-chan repository.Event {
	out := make(chan repository.Event)
	closedChannels := make(chan struct{}, len(generators))

	for _, gen := range generators {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch := gen.Start()
			for val := range ch {
				out <- val
			}
			closedChannels <- struct{}{}
		}()
	}

	go func() {
		for i := 0; i < len(generators); i++ {
			<-closedChannels
		}
		close(out)
	}()

	return out
}
