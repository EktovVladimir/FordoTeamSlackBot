package repository

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/event"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/shared"
	"sync"
)

type Event interface {
	GetInitiator() string
	Print() string
}

var (
	pullRequestEvents  = shared.NewConcurrentSlice[*event.GithubPullRequestEvent]()
	slackCommandEvents = shared.NewConcurrentSlice[*event.SlackCommandEvent]()
	initiators         = sync.Map{}
)

func StoreEvent(ev Event) {
	switch concreteEv := ev.(type) {
	case event.GithubPullRequestEvent:
		pullRequestEvents.Append(&concreteEv)
	case event.SlackCommandEvent:
		slackCommandEvents.Append(&concreteEv)
	default:
		return
	}

	initiators.Store(ev.GetInitiator(), struct{}{})
}

func GetPullRequestEvents() *shared.ConcurrentSlice[*event.GithubPullRequestEvent] {
	return pullRequestEvents
}

func GetSlackCommandEvents() *shared.ConcurrentSlice[*event.SlackCommandEvent] {
	return slackCommandEvents
}
