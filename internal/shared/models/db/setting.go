package db

import "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"

type Setting struct {
	Id    types.UniqId `json:"id"`
	Key   string       `json:"key"`
	Value string       `json:"value"`
}

func (u *Setting) GetId() types.UniqId {
	return u.Id
}

func (u *Setting) SetId(id types.UniqId) {
	u.Id = id
}
