package repository

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollectionName = "users"
	userIdSequenceName = "user_id"
)

type MongoUserRepository struct {
	db *mongo.Database
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{db}
}

func (r MongoUserRepository) GetAll(ctx context.Context) ([]*db.User, error) {
	cursor, err := r.db.Collection(userCollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var res []*db.User
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r MongoUserRepository) GetById(ctx context.Context, id types.UniqId) (*db.User, error) {
	var user db.User
	err := r.db.Collection(userCollectionName).
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r MongoUserRepository) GetByEmail(ctx context.Context, email string) (*db.User, error) {
	var user db.User
	err := r.db.Collection(userCollectionName).
		FindOne(ctx, bson.M{"email": email}).
		Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r MongoUserRepository) Create(ctx context.Context, user *db.User) error {
	user.Id = getMongoNextSequence(r.db, userIdSequenceName)
	_, err := r.db.Collection(userCollectionName).InsertOne(ctx, user)
	return err
}

func (r MongoUserRepository) Update(ctx context.Context, user *db.User) error {
	_, err := r.db.Collection(userCollectionName).
		UpdateOne(
			ctx,
			bson.M{"_id": user.Id},
			bson.M{"$set": user},
		)
	return err
}

func (r MongoUserRepository) Delete(ctx context.Context, id types.UniqId) error {
	_, err := r.db.Collection(userCollectionName).
		DeleteOne(ctx, bson.M{"_id": id})
	return err
}
