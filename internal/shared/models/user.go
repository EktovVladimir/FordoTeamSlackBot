package models

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"

type User struct {
	Id          db.UniqId
	Email       string
	SlackId     string
	SlackName   string
	GithubLogin string
}
