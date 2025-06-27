package db

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u" bson:"-"`

	Id              types.UniqId `json:"id" bson:"_id" bun:",pk,autoincrement"`
	SlackName       string       `json:"slack_name" bson:"slack_name" bun:",notnull"`
	SlackId         string       `json:"slack_id" bson:"slack_id" bun:",notnull"`
	GithubName      string       `json:"github_name" bson:"github_name" bun:",notnull"`
	Email           string       `json:"email" bson:"email" bun:",notnull"`
	AuditableFields `bson:",inline" bun:",embed"`
}

func (u *User) GetId() types.UniqId {
	return u.Id
}

func (u *User) SetId(id types.UniqId) {
	u.Id = id
}

func (u *User) IsFullFilled() bool {
	return u.SlackName != "" && u.GithubName != "" && u.Email != ""
}
