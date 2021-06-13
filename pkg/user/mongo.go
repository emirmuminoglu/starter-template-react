package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	client       *mongo.Client
	databaseName string
}

func NewMongoRepository(client *mongo.Client) Repository {
	return &mongoRepository{client: client}
}

func (*mongoRepository) Migrate() error {
	return nil
}

func (repo *mongoRepository) db() *mongo.Database {
	return repo.client.Database(repo.databaseName)
}

func (repo *mongoRepository) FindByID(ctx context.Context, id int, user *Model) error {
	return repo.db().Collection("user").FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(user)
}

func (repo *mongoRepository) FindByUsername(ctx context.Context, username string, user *Model) error {
	return repo.db().Collection("user").FindOne(ctx, bson.M{
		"username": username,
	}).Decode(user)
}
