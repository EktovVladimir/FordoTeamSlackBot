package repository

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/event"
)

var (
	slackCommandEvents []*event.SlackCommandEvent
	pullRequestEvents  []*event.GithubPullRequestEvent
	initiators         = make(map[string]struct{})
	id                 = 1
)

type Event interface {
	GetInitiator() string
	Print() string
}

func StoreEvent(ev Event) {
	switch concreteEv := ev.(type) {
	case event.GithubPullRequestEvent:
		pullRequestEvents = append(pullRequestEvents, &concreteEv)
	case event.SlackCommandEvent:
		slackCommandEvents = append(slackCommandEvents, &concreteEv)
	default:
		return
	}

	fmt.Printf("%d. StoreEvent: %s\n", id, ev.Print())

	initiators[ev.GetInitiator()] = struct{}{}

	fmt.Printf("slackCommandEvents count: %d, pullRequestEvents count: %d\n",
		len(slackCommandEvents), len(pullRequestEvents))
	fmt.Println("All initiators: ", initiators)
	fmt.Println()

	id++
}
