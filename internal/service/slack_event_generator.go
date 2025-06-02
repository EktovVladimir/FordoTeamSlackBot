package service

import (
	"context"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/event"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/repository"
	"time"
)

type slackEventGenerator struct {
	count      int
	intervalMs int
}

func NewSlackEventGenerator(count int, intervalMs int) *slackEventGenerator {
	return &slackEventGenerator{
		count:      count,
		intervalMs: intervalMs,
	}
}

func (s *slackEventGenerator) Start(ctx context.Context) <-chan repository.Event {
	ch := make(chan repository.Event)
	go func() {
		defer close(ch)
		defer fmt.Println("slackEventGenerator: finished")

		counter := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("slackEventGenerator: context done")
				return
			case <-time.After(time.Duration(s.intervalMs) * time.Millisecond):
				if s.count > 0 && counter >= s.count {
					return
				}

				ev := event.GenerateRandomSlackCommandEvent()
				ch <- ev
				counter++
			}
		}
	}()

	return ch
}
