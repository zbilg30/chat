package pkg

import (
	"auth/internal/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	client      *mongo.Client
	mongoConfig config.Mongo
}

type IMongoClient interface {
	GetCollection(string) *mongo.Collection
}

func NewMongoClient(mongoConfig config.Mongo) (IMongoClient, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoConfig.Uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB")

	return &mongoClient{
		client:      client,
		mongoConfig: mongoConfig,
	}, nil
}

func (client *mongoClient) GetCollection(collectionName string) *mongo.Collection {
	collection := client.client.Database(client.mongoConfig.Database).Collection(collectionName)
	return collection
}