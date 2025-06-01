package service

import (
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

func (s *slackEventGenerator) Start() <-chan repository.Event {
	ch := make(chan repository.Event)
	go func() {
		for i := 0; i < s.count; i++ {
			ev := event.GenerateRandomSlackCommandEvent()
			ch <- ev
			time.Sleep(time.Duration(s.intervalMs) * time.Millisecond)
		}

		fmt.Println("slackEventGenerator finished")

		close(ch)
	}()

	return ch
}
