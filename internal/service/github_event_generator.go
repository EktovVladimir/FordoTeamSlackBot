package service

import (
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

func (s *githubEventGenerator) Start() <-chan repository.Event {
	ch := make(chan repository.Event)
	go func() {
		for i := 0; i < s.count; i++ {
			ev := event.GenerateRandomPullRequestEvent()
			ch <- ev
			time.Sleep(time.Duration(s.intervalMs) * time.Millisecond)
		}

		fmt.Println("githubSlackGenerator finished")

		close(ch)
	}()

	return ch
}
