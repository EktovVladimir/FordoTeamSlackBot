package repository

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	settingCollectionName = "settings"
	settingIdSequenceName = "setting_id"
)

type MongoSettingRepository struct {
	db *mongo.Database
}

func NewMongoSettingRepository(db *mongo.Database) *MongoSettingRepository {
	return &MongoSettingRepository{db}
}

func (r MongoSettingRepository) GetAll(ctx context.Context) ([]*db.Setting, error) {
	cursor, err := r.db.Collection(settingCollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var res []*db.Setting
	err = cursor.All(ctx, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r MongoSettingRepository) GetById(ctx context.Context, id types.UniqId) (*db.Setting, error) {
	var setting db.Setting
	err := r.db.Collection(settingCollectionName).
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&setting)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (r MongoSettingRepository) GetByKey(ctx context.Context, key string) (*db.Setting, error) {
	var setting db.Setting
	err := r.db.Collection(settingCollectionName).
		FindOne(ctx, bson.M{"key": key}).
		Decode(&setting)
	if err != nil {
		return nil, err
	}
	return &setting, nil
}

func (r MongoSettingRepository) Create(ctx context.Context, setting *db.Setting) error {
	setting.Id = getMongoNextSequence(r.db, settingIdSequenceName)
	_, err := r.db.Collection(settingCollectionName).InsertOne(ctx, setting)
	return err
}

func (r MongoSettingRepository) Update(ctx context.Context, setting *db.Setting) error {
	_, err := r.db.Collection(settingCollectionName).
		UpdateOne(
			ctx,
			bson.M{"_id": setting.Id},
			bson.M{"$set": setting},
		)
	return err
}

func (r MongoSettingRepository) Delete(ctx context.Context, id types.UniqId) error {
	_, err := r.db.Collection(settingCollectionName).
		DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r MongoSettingRepository) DeleteByKey(ctx context.Context, key string) error {
	_, err := r.db.Collection(settingCollectionName).
		DeleteOne(ctx, bson.M{"key": key})
	return err
}
