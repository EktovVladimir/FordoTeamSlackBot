package repository

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/uptrace/bun"
	"time"
)

type PostgresSettingRepository struct {
	db *bun.DB
}

func NewPostgresSettingRepository(db *bun.DB) *PostgresSettingRepository {
	return &PostgresSettingRepository{db: db}
}

func (r *PostgresSettingRepository) GetAll(ctx context.Context) ([]*db.Setting, error) {
	var settings []*db.Setting
	err := r.db.NewSelect().
		Model(&settings).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (r *PostgresSettingRepository) GetById(ctx context.Context, id db.UniqId) (*db.Setting, error) {
	var setting db.Setting
	err := r.db.NewSelect().
		Model(&setting).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &setting, nil
}

func (r *PostgresSettingRepository) GetByKey(ctx context.Context, key string) (*db.Setting, error) {
	var setting db.Setting
	err := r.db.NewSelect().
		Model(&setting).
		Where("key = ?", key).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &setting, nil
}

func (r *PostgresSettingRepository) GetUpdatedBetween(ctx context.Context, start, end time.Time) ([]*db.Setting, error) {
	var settings []*db.Setting
	err := r.db.NewSelect().
		Model(&settings).
		Where("updated_at BETWEEN ? AND ?", start, end).
		Order("updated_at ASC").
		Scan(ctx)

	if err != nil {
		return nil, errors.Join(err, ErrorNotFound)
	}

	return settings, nil
}

func (r *PostgresSettingRepository) Create(ctx context.Context, setting *db.Setting) error {
	_, err := r.db.NewInsert().
		Model(setting).
		Returning("*").
		Exec(ctx)

	return err
}

func (r *PostgresSettingRepository) Update(ctx context.Context, setting *db.Setting) error {
	setting.UpdatedAt = time.Now()

	_, err := r.db.NewUpdate().
		Model(setting).
		WherePK().
		Exec(ctx)

	return err
}

func (r *PostgresSettingRepository) Delete(ctx context.Context, id db.UniqId) error {
	_, err := r.db.NewDelete().
		Model((*db.Setting)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}

func (r *PostgresSettingRepository) DeleteByKey(ctx context.Context, key string) error {
	_, err := r.db.NewDelete().
		Model((*db.Setting)(nil)).
		Where("key = ?", key).
		Exec(ctx)

	return err
}
