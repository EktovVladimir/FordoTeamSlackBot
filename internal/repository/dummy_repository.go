package repository

import (
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/settings"
)

type dummyRepository struct {
}

func NewDummyRepository() *dummyRepository {
	return &dummyRepository{}
}

func (r *dummyRepository) GetItem(source string, key string) (db.Setting, error) {
	//TODO implement me
	if key == settings.IssuePrefixKey {
		return db.Setting{Key: key, Source: "*", Value: "OTAB"}, nil
	}

	return db.Setting{}, errors.New("not found")
}

func (r *dummyRepository) SetItem(item db.Setting) error {
	//TODO implement me
	panic("implement me")
}
