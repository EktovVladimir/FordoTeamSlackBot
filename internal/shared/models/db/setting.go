package db

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"github.com/uptrace/bun"
)

type Setting struct {
	bun.BaseModel `bun:"table:settings,alias:s" bson:"-"`

	Id              types.UniqId `json:"id" bson:"_id" bun:",pk,autoincrement"`
	Key             string       `json:"key" bson:"key" bun:",notnull"`
	Value           string       `json:"value" bson:"value" bun:",notnull"`
	AuditableFields `bson:",inline" bun:",embed"`
}

func (u *Setting) GetId() types.UniqId {
	return u.Id
}

func (u *Setting) SetId(id types.UniqId) {
	u.Id = id
}
