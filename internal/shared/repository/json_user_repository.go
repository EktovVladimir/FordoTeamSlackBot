package repository

import (
	"cmp"
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"slices"
)

type userStore interface {
	GetUsers() *db_adapter.Entity[*db.User]
	SaveChanges() error
}

type jsonUserRepository struct {
	db userStore
}

func NewJsonUserRepository(db userStore) *jsonUserRepository {
	return &jsonUserRepository{db}
}

func (r *jsonUserRepository) GetAll(ctx context.Context) ([]*db.User, error) {
	data := r.db.GetUsers().GetAll()

	slices.SortFunc(data, func(a, b *db.User) int {
		return cmp.Compare(a.Id, b.Id)
	})

	return data, nil
}

func (r *jsonUserRepository) GetById(ctx context.Context, id types.UniqId) (*db.User, error) {
	entity := r.db.GetUsers()

	data, found := entity.Get(id)
	if !found {
		return nil, ErrorNotFound
	}

	return data, nil
}

func (r *jsonUserRepository) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	entity := r.db.GetUsers()
	data := entity.GetAll()

	for _, user := range data {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, ErrorNotFound
}

func (r *jsonUserRepository) Create(ctx context.Context, item *db.User) error {
	entity := r.db.GetUsers()
	entity.Insert(item)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonUserRepository) Update(ctx context.Context, item *db.User) error {
	entity := r.db.GetUsers()
	entity.Update(item)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonUserRepository) Delete(ctx context.Context, id types.UniqId) error {
	entity := r.db.GetUsers()
	entity.Delete(id)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}
