package repository

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	codeReviewCollectionName = "code_reviews"
	codeReviewIdSequenceName = "code_review_id"
)

type MongoCodeReviewRepository struct {
	db *mongo.Database
}

func NewMongoCodeReviewRepository(db *mongo.Database) *MongoCodeReviewRepository {
	return &MongoCodeReviewRepository{db}
}

func (r MongoCodeReviewRepository) GetAll(ctx context.Context) ([]*db.CodeReview, error) {
	cursor, err := r.db.Collection(codeReviewCollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var res []*db.CodeReview
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r MongoCodeReviewRepository) GetById(ctx context.Context, id types.UniqId) (*db.CodeReview, error) {
	var review db.CodeReview
	err := r.db.Collection(codeReviewCollectionName).
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&review)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r MongoCodeReviewRepository) GetUpdatedBetween(ctx context.Context, start, end time.Time) ([]*db.CodeReview, error) {
	filter := bson.M{
		"updated_at": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	opts := options.Find().SetSort(bson.D{{"updated_at", 1}})

	cursor, err := r.db.Collection(codeReviewCollectionName).Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Join(err, ErrorNotFound)
	}
	defer cursor.Close(ctx)

	var items []*db.CodeReview
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (r MongoCodeReviewRepository) Create(ctx context.Context, review *db.CodeReview) error {
	review.Id = getMongoNextSequence(r.db, codeReviewIdSequenceName)

	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	_, err := r.db.Collection(codeReviewCollectionName).InsertOne(ctx, review)
	return err
}

func (r MongoCodeReviewRepository) Update(ctx context.Context, review *db.CodeReview) error {
	review.UpdatedAt = time.Now()

	_, err := r.db.Collection(codeReviewCollectionName).
		UpdateOne(
			ctx,
			bson.M{"_id": review.Id},
			bson.M{"$set": review},
		)
	return err
}

func (r MongoCodeReviewRepository) Delete(ctx context.Context, id types.UniqId) error {
	_, err := r.db.Collection(codeReviewCollectionName).
		DeleteOne(ctx, bson.M{"_id": id})
	return err
}
