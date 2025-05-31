package repository

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/event"
	"sync"
	"time"
)

type Event interface {
	GetInitiator() string
	Print() string
}

type concurrentSlice struct {
	mu    sync.RWMutex
	slice []Event
}

var (
	slackCommandEvents = &concurrentSlice{mu: sync.RWMutex{}, slice: make([]Event, 0)}
	pullRequestEvents  = &concurrentSlice{mu: sync.RWMutex{}, slice: make([]Event, 0)}
	initiators         = sync.Map{}
)

func StoreEvent(ev Event) {
	switch concreteEv := ev.(type) {
	case event.GithubPullRequestEvent:
		appendEvent(pullRequestEvents, concreteEv)
	case event.SlackCommandEvent:
		appendEvent(slackCommandEvents, concreteEv)
	default:
		return
	}

	initiators.Store(ev.GetInitiator(), struct{}{})
}

func appendEvent(cs *concurrentSlice, ev Event) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.slice = append(cs.slice, ev)
}

func StartEventReader(ch <-chan Event, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ev := range ch {
			StoreEvent(ev)
		}
		fmt.Println("EventReader finished")
	}()
}

func StartEventLogger(intervalMs int) {
	go func() {
		exit := false

		loggedSlices := []*struct {
			concurrentSlice *concurrentSlice
			prevCount       int
			name            string
		}{
			{slackCommandEvents, 0, "slack events"},
			{pullRequestEvents, 0, "pr events"},
		}

		for !exit {

			currentLog := make([]Event, 0)

			for _, s := range loggedSlices {
				s.concurrentSlice.mu.RLock()

				newLen := len(s.concurrentSlice.slice)
				if newLen > s.prevCount {
					fmt.Printf("Found new events in %s (%d > %d):\n", s.name, newLen, s.prevCount)

					subSlice := s.concurrentSlice.slice[s.prevCount:newLen]
					currentLog = append(currentLog, subSlice...)

					s.prevCount = newLen
				}

				s.concurrentSlice.mu.RUnlock()
			}

			if len(currentLog) > 0 {
				for _, l := range currentLog {
					fmt.Println(l.Print())
				}
			} else {
				fmt.Println("Not found new events")
			}

			fmt.Println()

			time.Sleep(time.Duration(intervalMs) * time.Millisecond)
		}

		//Скорее всего никогда не вызовется, так как не дожидаемся завершения этой горутины
		fmt.Println("EventLogger finished")
	}()
}
