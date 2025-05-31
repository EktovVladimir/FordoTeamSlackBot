package service

import (
	"context"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/event"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/repository"
	"time"
)

type githubEventGenerator struct {
	count      int
	intervalMs int
}

func NewGithubEventGenerator(count int, intervalMs int) *githubEventGenerator {
	return &githubEventGenerator{
		count:      count,
		intervalMs: intervalMs,
	}
}

func (s *githubEventGenerator) Start(ctx context.Context) <-chan repository.Event {
	ch := make(chan repository.Event)
	go func() {
		defer close(ch)
		defer fmt.Println("githubEventGenerator: finished")

		counter := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("githubEventGenerator: context done")
				return
			case <-time.After(time.Duration(s.intervalMs) * time.Millisecond):
				if s.count > 0 && counter >= s.count {
					return
				}

				ev := event.GenerateRandomPullRequestEvent()
				ch <- ev
				counter++
			}
		}
	}()

	return ch
}
