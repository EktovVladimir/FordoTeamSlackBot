package repository

import (
	"errors"
	db2 "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db"
)

type dataStore interface {
	GetUsers() []*db2.User
	GetSettings() []*db2.Setting
	GetDeployments() []*db2.Deployment
	GetCodeReviews() []*db2.CodeReview
	Insert(item db2.UniqEntity) error
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

func (r *dataRepository) InsertUser(value *db2.User) error {
	if err := r.db.Insert(value); err != nil {
		return err
	}

	return nil
	//return r.db.SaveUsers()
}

func (r *dataRepository) GetSetting(source string, key string) (*db2.Setting, error) {
	settings := r.db.GetSettings()

	for _, item := range settings {
		if item.Key == key && item.Source == source {
			return item, nil
		}
	}

	return nil, errors.New("setting not found")
}

func (r *dataRepository) InsertSetting(value *db2.Setting) error {
	if err := r.db.Insert(value); err != nil {
		return err
	}

	return nil
	//return r.db.SaveSettings()
}

func (r *dataRepository) InsertDeployment(value *db2.Deployment) error {
	if err := r.db.Insert(value); err != nil {
		return err
	}

	return nil
	//return r.db.SaveUsers()
}

func (r *dataRepository) InsertCodeReview(value *db2.CodeReview) error {
	if err := r.db.Insert(value); err != nil {
		return err
	}

	return nil
	//return r.db.SaveCodeReviews()
}
