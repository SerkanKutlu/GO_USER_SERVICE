package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"userService/config"
)

type MongoService struct {
	Users *mongo.Collection
}

func GetMongoService(mongoConfig *config.MongoConfig) *MongoService {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConfig.ConnectionString))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	database := client.Database(mongoConfig.Database)
	userCollection := database.Collection(mongoConfig.Collection["user"])
	service := MongoService{
		userCollection,
	}
	return &service
}
