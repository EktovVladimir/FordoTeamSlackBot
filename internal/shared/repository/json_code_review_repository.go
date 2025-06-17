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

type codeReviewStore interface {
	GetCodeReviews() *db_adapter.Entity[*db.CodeReview]
	SaveChanges() error
}

type jsonCodeReviewRepository struct {
	db codeReviewStore
}

func NewJsonCodeReviewRepository(db codeReviewStore) *jsonCodeReviewRepository {
	return &jsonCodeReviewRepository{db}
}

func (r *jsonCodeReviewRepository) GetAll(ctx context.Context) ([]*db.CodeReview, error) {
	data := r.db.GetCodeReviews().GetAll()

	slices.SortFunc(data, func(a, b *db.CodeReview) int {
		return cmp.Compare(a.Id, b.Id)
	})

	return data, nil
}

func (r *jsonCodeReviewRepository) GetById(ctx context.Context, id types.UniqId) (*db.CodeReview, error) {
	entity := r.db.GetCodeReviews()

	data, found := entity.Get(id)
	if !found {
		return nil, errors.New("code review record not found")
	}

	return data, nil
}

func (r *jsonCodeReviewRepository) Create(ctx context.Context, item *db.CodeReview) error {
	entity := r.db.GetCodeReviews()
	entity.Insert(item)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonCodeReviewRepository) Update(ctx context.Context, item *db.CodeReview) error {
	entity := r.db.GetCodeReviews()
	entity.Update(item)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}

func (r *jsonCodeReviewRepository) Delete(ctx context.Context, id types.UniqId) error {
	entity := r.db.GetCodeReviews()
	entity.Delete(id)

	if err := r.db.SaveChanges(); err != nil {
		return err
	}

	return nil
}
