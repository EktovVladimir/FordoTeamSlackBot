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

type deploymentStore interface {
	GetDeployments() *db_adapter.Entity[*db.Deployment]
	SaveChanges() error
}

type jsonDeploymentRepository struct {
	db deploymentStore
}

func NewJsonDeploymentRepository(db deploymentStore) *jsonDeploymentRepository {
	return &jsonDeploymentRepository{db}
}

func (r *jsonDeploymentRepository) GetAll(ctx context.Context) ([]*db.Deployment, error) {
	data := r.db.GetDeployments().GetAll()

	slices.SortFunc(data, func(a, b *db.Deployment) int {
		return cmp.Compare(a.Id, b.Id)
	})

	return data, nil
}

func (r *jsonDeploymentRepository) GetById(ctx context.Context, id types.UniqId) (*db.Deployment, error) {
	entity := r.db.GetDeployments()

	data, found := entity.Get(id)
	if !found {
		return nil, errors.New("deployment record not found")
	}

	return data, nil
}

func (r *jsonDeploymentRepository) Create(ctx context.Context, item *db.Deployment) error {
	entity := r.db.GetDeployments()
	entity.Insert(item)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonDeploymentRepository) Update(ctx context.Context, item *db.Deployment) error {
	entity := r.db.GetDeployments()
	entity.Update(item)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonDeploymentRepository) Delete(ctx context.Context, id types.UniqId) error {
	entity := r.db.GetDeployments()
	entity.Delete(id)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}
