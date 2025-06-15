package repository

import (
	"cmp"
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"slices"
)

type settingsStore interface {
	GetSettings() *db_adapter.Entity[*db.Setting]
	SaveChanges() error
}

type jsonSettingRepository struct {
	db settingsStore
}

func NewJsonSettingRepository(db settingsStore) *jsonSettingRepository {
	return &jsonSettingRepository{db}
}

func (r *jsonSettingRepository) GetAll(ctx context.Context) ([]*db.Setting, error) {
	data := r.db.GetSettings().GetAll()

	slices.SortFunc(data, func(a, b *db.Setting) int {
		return cmp.Compare(a.Id, b.Id)
	})

	return data, nil
}

func (r *jsonSettingRepository) GetById(ctx context.Context, id types.UniqId) (*db.Setting, error) {
	entity := r.db.GetSettings()

	data, found := entity.Get(id)
	if !found {
		return nil, errors.New("setting not found")
	}

	return data, nil
}

func (r *jsonSettingRepository) GetByKey(ctx context.Context, key string) (*db.Setting, error) {
	data := r.db.GetSettings().GetAll()

	for _, setting := range data {
		if setting.Key == key {
			return setting, nil
		}
	}

	return nil, errors.New("setting not found")
}

func (r *jsonSettingRepository) Create(ctx context.Context, item *db.Setting) error {
	entity := r.db.GetSettings()
	entity.Insert(item)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonSettingRepository) Update(ctx context.Context, item *db.Setting) error {
	entity := r.db.GetSettings()
	entity.Update(item)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonSettingRepository) Delete(ctx context.Context, id types.UniqId) error {
	entity := r.db.GetSettings()
	entity.Delete(id)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonSettingRepository) DeleteByKey(ctx context.Context, key string) error {
	item, err := r.GetByKey(ctx, key)
	if err != nil {
		return err
	}

	entity := r.db.GetSettings()
	entity.Delete(item.Id)

	if err = r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}
