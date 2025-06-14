package service

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"sync"
)

func StartEventReader(ch <-chan repository.Event, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ev := range ch {
			repository.StoreEvent(ev)
		}
		fmt.Println("EventReader: finished")
	}()
}
