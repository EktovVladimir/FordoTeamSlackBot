package repository

import (
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/db"
)

type dataStore interface {
	GetUsers() []*db.User
	GetSettings() []*db.Setting
	GetDeployments() []*db.Deployment
	GetCodeReviews() []*db.CodeReview
	Insert(item db.Entity) error
	/*SaveUsers() error
	SaveSettings() error
	SaveDeployments() error
	SaveCodeReviews() error*/
}

type dataRepository struct {
	db dataStore
}

func NewDataRepository(db dataStore) *dataRepository {
	return &dataRepository{db: db}
}

func (r *dataRepository) InsertUser(value *db.User) error {
	if err := r.db.Insert(value); err != nil {
		return err
	}

	return nil
	//return r.db.SaveUsers()
}

func (r *dataRepository) GetSetting(source string, key string) (*db.Setting, error) {
	settings := r.db.GetSettings()

	for _, item := range settings {
		if item.Key == key && item.Source == source {
			return item, nil
		}
	}

	return nil, errors.New("setting not found")
}

func (r *dataRepository) InsertSetting(value *db.Setting) error {
	if err := r.db.Insert(value); err != nil {
		return err
	}

	return nil
	//return r.db.SaveSettings()
}

func (r *dataRepository) InsertDeployment(value *db.Deployment) error {
	if err := r.db.Insert(value); err != nil {
		return err
	}

	return nil
	//return r.db.SaveUsers()
}

func (r *dataRepository) InsertCodeReview(value *db.CodeReview) error {
	if err := r.db.Insert(value); err != nil {
		return err
	}

	return nil
	//return r.db.SaveCodeReviews()
}
