package repository

import (
	event2 "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/event"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/shared"
	"sync"
)

type Event interface {
	GetInitiator() string
	Print() string
}

var (
	pullRequestEvents  = shared.NewConcurrentSlice[*event2.GithubPullRequestEvent]()
	slackCommandEvents = shared.NewConcurrentSlice[*event2.SlackCommandEvent]()
	initiators         = sync.Map{}
)

func StoreEvent(ev Event) {
	switch concreteEv := ev.(type) {
	case event2.GithubPullRequestEvent:
		pullRequestEvents.Append(&concreteEv)
	case event2.SlackCommandEvent:
		slackCommandEvents.Append(&concreteEv)
	default:
		return
	}

	initiators.Store(ev.GetInitiator(), struct{}{})
}

func GetPullRequestEvents() *shared.ConcurrentSlice[*event2.GithubPullRequestEvent] {
	return pullRequestEvents
}

func GetSlackCommandEvents() *shared.ConcurrentSlice[*event2.SlackCommandEvent] {
	return slackCommandEvents
}
