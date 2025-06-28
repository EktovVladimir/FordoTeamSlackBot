package repository

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
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
		Relation("SlackPost").
		Relation("PullRequests").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *PostgresCodeReviewRepository) GetById(ctx context.Context, id db.UniqId) (*db.CodeReview, error) {
	var review db.CodeReview
	err := r.db.NewSelect().
		Model(&review).
		Relation("SlackPost").
		Relation("PullRequests").
		Where("cr.id = ?", id).
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
		Relation("SlackPost").
		Relation("PullRequests").
		Where("cr.updated_at BETWEEN ? AND ?", start, end).
		Order("cr.updated_at ASC").
		Scan(ctx)

	if err != nil {
		return nil, errors.Join(err, ErrorNotFound)
	}

	return reviews, nil
}

func (r *PostgresCodeReviewRepository) Create(ctx context.Context, review *db.CodeReview) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(tx bun.Tx) {
		_ = tx.Rollback()
	}(tx)

	_, err = tx.NewInsert().
		Model(review.SlackPost).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}

	review.SlackPostId = review.SlackPost.Id

	_, err = tx.NewInsert().
		Model(&review.PullRequests).
		Exec(ctx)
	if err != nil {
		return err
	}

	_, err = tx.NewInsert().
		Model(review).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}

	relations := make([]*db.CodeReviewToPR, 0)
	for _, record := range review.PullRequests {
		relations = append(relations, &db.CodeReviewToPR{
			CodeReviewId:  review.Id,
			PullRequestId: record.Id,
		})
	}

	_, err = tx.NewInsert().
		Model(&relations).
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresCodeReviewRepository) Update(ctx context.Context, review *db.CodeReview) error {
	review.UpdatedAt = time.Now()

	_, err := r.db.NewUpdate().
		Model(review).
		WherePK().
		Exec(ctx)

	return err
}

func (r *PostgresCodeReviewRepository) Delete(ctx context.Context, id db.UniqId) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(tx bun.Tx) {
		_ = tx.Rollback()
	}(tx)

	_, err = tx.NewDelete().
		Model((*db.CodeReviewToPR)(nil)).
		Where("code_review_id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	_, err = tx.NewDelete().
		Model((*db.CodeReview)(nil)).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
