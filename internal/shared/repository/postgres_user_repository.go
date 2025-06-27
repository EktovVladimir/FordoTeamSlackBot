package repository

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"github.com/uptrace/bun"
	"time"
)

type PostgresUserRepository struct {
	db *bun.DB
}

func NewPostgresUserRepository(db *bun.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]*db.User, error) {
	var users []*db.User
	err := r.db.NewSelect().
		Model(&users).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *PostgresUserRepository) GetById(ctx context.Context, id types.UniqId) (*db.User, error) {
	var user db.User
	err := r.db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	var user db.User
	err := r.db.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetUpdatedBetween(ctx context.Context, start, end time.Time) ([]*db.User, error) {
	var users []*db.User
	err := r.db.NewSelect().
		Model(&users).
		Where("updated_at BETWEEN ? AND ?", start, end).
		Order("updated_at ASC").
		Scan(ctx)

	if err != nil {
		return nil, errors.Join(err, ErrorNotFound)
	}

	return users, nil
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *db.User) error {
	_, err := r.db.NewInsert().
		Model(user).
		Returning("*").
		Exec(ctx)

	return err
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *db.User) error {
	user.UpdatedAt = time.Now()

	_, err := r.db.NewUpdate().
		Model(user).
		WherePK().
		Exec(ctx)

	return err
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id types.UniqId) error {
	_, err := r.db.NewDelete().
		Model((*db.User)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}
