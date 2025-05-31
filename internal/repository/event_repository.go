package repository

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/abstract"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/event"
)

var (
	slackCommandEvents []event.SlackCommandEvent
	pullRequestEvents  []event.GithubPullRequestEvent
	initiators         = make(map[string]struct{})
	id                 = 1
)

func StoreEvent(ev abstract.Event) {
	fmt.Printf("%d. StoreEvent: %s\n", id, ev.Print())

	initiators[ev.GetInitiator()] = struct{}{}

	switch ev.(type) {
	case event.GithubPullRequestEvent:
		pullRequestEvents = append(pullRequestEvents, ev.(event.GithubPullRequestEvent))
		break
	case event.SlackCommandEvent:
		slackCommandEvents = append(slackCommandEvents, ev.(event.SlackCommandEvent))
		break
	}

	fmt.Printf("slackCommandEvents count: %d, pullRequestEvents count: %d\n",
		len(slackCommandEvents), len(pullRequestEvents))
	fmt.Println("All initiators: ", initiators)
	fmt.Println()

	id++
}
