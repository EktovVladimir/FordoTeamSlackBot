package repository

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"github.com/uptrace/bun"
	"time"
)

type PostgresCodeReviewRepository struct {
	db *bun.DB
}

func NewPostgresCodeReviewRepository(db *bun.DB) *PostgresCodeReviewRepository {
	return &PostgresCodeReviewRepository{db: db}
}

func (r *PostgresCodeReviewRepository) GetAll(ctx context.Context) ([]*db.CodeReview, error) {
	var reviews []*db.CodeReview
	err := r.db.NewSelect().
		Model(&reviews).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *PostgresCodeReviewRepository) GetById(ctx context.Context, id types.UniqId) (*db.CodeReview, error) {
	var review db.CodeReview
	err := r.db.NewSelect().
		Model(&review).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *PostgresCodeReviewRepository) GetUpdatedBetween(ctx context.Context, start, end time.Time) ([]*db.CodeReview, error) {
	var reviews []*db.CodeReview
	err := r.db.NewSelect().
		Model(&reviews).
		Where("updated_at BETWEEN ? AND ?", start, end).
		Order("updated_at ASC").
		Scan(ctx)

	if err != nil {
		return nil, errors.Join(err, ErrorNotFound)
	}

	return reviews, nil
}

func (r *PostgresCodeReviewRepository) Create(ctx context.Context, review *db.CodeReview) error {
	_, err := r.db.NewInsert().
		Model(review).
		Returning("*").
		Exec(ctx)

	return err
}

func (r *PostgresCodeReviewRepository) Update(ctx context.Context, review *db.CodeReview) error {
	review.UpdatedAt = time.Now()

	_, err := r.db.NewUpdate().
		Model(review).
		WherePK().
		Exec(ctx)

	return err
}

func (r *PostgresCodeReviewRepository) Delete(ctx context.Context, id types.UniqId) error {
	_, err := r.db.NewDelete().
		Model((*db.CodeReview)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}
