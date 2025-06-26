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
	deploymentCollectionName = "deployments"
	deploymentIdSequenceName = "deployment_id"
)

type MongoDeploymentRepository struct {
	db *mongo.Database
}

func NewMongoDeploymentRepository(db *mongo.Database) *MongoDeploymentRepository {
	return &MongoDeploymentRepository{db}
}

func (r MongoDeploymentRepository) GetAll(ctx context.Context) ([]*db.Deployment, error) {
	cursor, err := r.db.Collection(deploymentCollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var res []*db.Deployment
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r MongoDeploymentRepository) GetById(ctx context.Context, id types.UniqId) (*db.Deployment, error) {
	var deployment db.Deployment
	err := r.db.Collection(deploymentCollectionName).
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&deployment)
	if err != nil {
		return nil, err
	}
	return &deployment, nil
}

func (r MongoDeploymentRepository) GetUpdatedBetween(ctx context.Context, start, end time.Time) ([]*db.Deployment, error) {
	filter := bson.M{
		"updated_at": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	opts := options.Find().SetSort(bson.D{{"updated_at", 1}})

	cursor, err := r.db.Collection(deploymentCollectionName).Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Join(err, ErrorNotFound)
	}
	defer cursor.Close(ctx)

	var items []*db.Deployment
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (r MongoDeploymentRepository) Create(ctx context.Context, deployment *db.Deployment) error {
	deployment.Id = getMongoNextSequence(r.db, deploymentIdSequenceName)

	deployment.CreatedAt = time.Now()
	deployment.UpdatedAt = time.Now()

	_, err := r.db.Collection(deploymentCollectionName).InsertOne(ctx, deployment)
	return err
}

func (r MongoDeploymentRepository) Update(ctx context.Context, deployment *db.Deployment) error {
	deployment.UpdatedAt = time.Now()

	_, err := r.db.Collection(deploymentCollectionName).
		UpdateOne(
			ctx,
			bson.M{"_id": deployment.Id},
			bson.M{"$set": deployment},
		)
	return err
}

func (r MongoDeploymentRepository) Delete(ctx context.Context, id types.UniqId) error {
	_, err := r.db.Collection(deploymentCollectionName).
		DeleteOne(ctx, bson.M{"_id": id})
	return err
}
