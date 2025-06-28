package db

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel   `bun:"table:users,alias:u"`
	UniqFields      `bun:",embed"`
	AuditableFields `bun:",embed"`

	SlackName  string `bun:",notnull"`
	SlackId    string `bun:",notnull"`
	GithubName string `bun:",notnull"`
	Email      string `bun:",notnull"`
}

func (u *User) IsFullFilled() bool {
	return u.SlackName != "" && u.GithubName != "" && u.Email != ""
}
