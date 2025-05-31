package models

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/abstract"

type Infrastructure struct {
	PrService       abstract.PullRequestService
	SlackService    abstract.MessengerService
	IssueService    abstract.IssueTrackerService
	SettingsService abstract.SettingsService
}
