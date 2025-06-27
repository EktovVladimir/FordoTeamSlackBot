package repository

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"github.com/uptrace/bun"
	"time"
)

type PostgresDeploymentRepository struct {
	db *bun.DB
}

func NewPostgresDeploymentRepository(db *bun.DB) *PostgresDeploymentRepository {
	return &PostgresDeploymentRepository{db: db}
}

func (r *PostgresDeploymentRepository) GetAll(ctx context.Context) ([]*db.Deployment, error) {
	var deployments []*db.Deployment
	err := r.db.NewSelect().
		Model(&deployments).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return deployments, nil
}

func (r *PostgresDeploymentRepository) GetById(ctx context.Context, id types.UniqId) (*db.Deployment, error) {
	var deployment db.Deployment
	err := r.db.NewSelect().
		Model(&deployment).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &deployment, nil
}

func (r *PostgresDeploymentRepository) GetUpdatedBetween(ctx context.Context, start, end time.Time) ([]*db.Deployment, error) {
	var deployments []*db.Deployment
	err := r.db.NewSelect().
		Model(&deployments).
		Where("updated_at BETWEEN ? AND ?", start, end).
		Order("updated_at ASC").
		Scan(ctx)

	if err != nil {
		return nil, errors.Join(err, ErrorNotFound)
	}

	return deployments, nil
}

func (r *PostgresDeploymentRepository) Create(ctx context.Context, deployment *db.Deployment) error {
	_, err := r.db.NewInsert().
		Model(deployment).
		Returning("*").
		Exec(ctx)

	return err
}

func (r *PostgresDeploymentRepository) Update(ctx context.Context, deployment *db.Deployment) error {
	deployment.UpdatedAt = time.Now()

	_, err := r.db.NewUpdate().
		Model(deployment).
		WherePK().
		Exec(ctx)

	return err
}

func (r *PostgresDeploymentRepository) Delete(ctx context.Context, id types.UniqId) error {
	_, err := r.db.NewDelete().
		Model((*db.Deployment)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}
