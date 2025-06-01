package service

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/event"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/shared"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/repository"
	"time"
)

type loggedSlice[T repository.Event] struct {
	slice     *shared.ConcurrentSlice[T]
	prevCount int
	name      string
}

func StartEventLogger(intervalMs int) {
	go func() {
		exit := false

		loggedGhEvents := &loggedSlice[*event.GithubPullRequestEvent]{
			slice: repository.GetPullRequestEvents(),
			name:  "pr events",
		}

		loggedSlackEvents := &loggedSlice[*event.SlackCommandEvent]{
			slice: repository.GetSlackCommandEvents(),
			name:  "slack events",
		}

		for !exit {
			currentLog := make([]string, 0)

			if logs := processLog[*event.GithubPullRequestEvent](loggedGhEvents); len(logs) > 0 {
				currentLog = append(currentLog, logs...)
				fmt.Printf("Found %d new pr events\n", len(logs))
			}

			if logs := processLog[*event.SlackCommandEvent](loggedSlackEvents); len(logs) > 0 {
				currentLog = append(currentLog, logs...)
				fmt.Printf("Found %d new slack events\n", len(logs))
			}

			if len(currentLog) > 0 {
				for _, l := range currentLog {
					fmt.Println(l)
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

func processLog[T repository.Event](ls *loggedSlice[T]) []string {
	res := make([]string, 0)

	newLen := ls.slice.Len()
	if newLen > ls.prevCount {

		subSlice := ls.slice.Slice(ls.prevCount, newLen)

		for _, i := range subSlice {
			res = append(res, i.Print())
		}

		ls.prevCount = newLen
	}

	return res
}
