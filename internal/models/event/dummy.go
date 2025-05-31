package event

import (
	"fmt"
	"math/rand"
)

func GenerateRandomPullRequestEvent() GithubPullRequestEvent {
	return GithubPullRequestEvent{
		Number:     randomPRNumber(),
		Repo:       randomRepo(),
		BaseBranch: randomBranch(),
		Branch:     randomBranch(),
		Action:     randomAction(),
		Requester:  randomRequester(),
	}
}

func GenerateRandomSlackCommandEvent() SlackCommandEvent {
	return SlackCommandEvent{
		Command:  randomCommandName(),
		Args:     randomCommandArgs(),
		Channel:  randomChannel(),
		UserName: randomUserName(),
	}
}

func randomStringFromList(list []string) string {
	return list[rand.Intn(len(list))]
}

func randomPRNumber() string {
	return fmt.Sprintf("%d", rand.Intn(1000)+1)
}

func randomRepo() string {
	repos := []string{
		"FordoTeamSlackBot",
		"GolangPracticeOtus",
	}
	return randomStringFromList(repos)
}

func randomBranch() string {
	branches := []string{
		"feature/test1",
		"feature/test3",
		"bugfix/2025-05-30",
	}
	return randomStringFromList(branches)
}

func randomAction() string {
	actions := []string{
		"opened",
		"closed",
		"reopened",
		"merged",
		"synchronize",
	}
	return randomStringFromList(actions)
}

func randomRequester() string {
	requesters := []string{
		"EktovVladimir",
		"developer1",
		"coder42",
		"git-user",
	}
	return randomStringFromList(requesters)
}

func randomChannel() string {
	channels := []string{
		"general",
		"random",
		"devops",
		"backend",
		"frontend",
		"releases",
	}
	return "#" + randomStringFromList(channels)
}

func randomUserName() string {
	users := []string{
		"alice",
		"bob",
		"charlie",
		"dave",
		"eve",
		"frank",
	}
	return randomStringFromList(users)
}

func randomCommandName() string {
	users := []string{
		"cr",
		"deploy",
	}
	return randomStringFromList(users)
}

func randomCommandArgs() string {
	users := []string{
		"new",
		fmt.Sprintf("pr %s", randomPRNumber()),
	}
	return randomStringFromList(users)
}
