package models

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type User struct {
	Id          types.UniqId
	Email       string
	SlackId     string
	SlackName   string
	GithubLogin string
}
