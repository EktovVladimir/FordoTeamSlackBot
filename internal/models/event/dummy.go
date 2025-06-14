package event

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/helpers"
)

func GenerateRandomPullRequestEvent() GithubPullRequestEvent {
	return GithubPullRequestEvent{
		Number:     helpers.RandomPRNumber(),
		Repo:       helpers.RandomRepo(),
		BaseBranch: helpers.RandomBranch(),
		Branch:     helpers.RandomBranch(),
		Action:     helpers.RandomAction(),
		Requester:  helpers.RandomRequester(),
	}
}

func GenerateRandomSlackCommandEvent() SlackCommandEvent {
	return SlackCommandEvent{
		Command:  helpers.RandomCommandName(),
		Args:     helpers.RandomCommandArgs(),
		Channel:  helpers.RandomChannel(),
		UserName: helpers.RandomUserName(),
	}
}
