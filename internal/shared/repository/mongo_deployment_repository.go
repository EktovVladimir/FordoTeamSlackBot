package repository

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r MongoDeploymentRepository) Create(ctx context.Context, deployment *db.Deployment) error {
	deployment.Id = getMongoNextSequence(r.db, deploymentIdSequenceName)
	_, err := r.db.Collection(deploymentCollectionName).InsertOne(ctx, deployment)
	return err
}

func (r MongoDeploymentRepository) Update(ctx context.Context, deployment *db.Deployment) error {
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
