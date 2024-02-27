package datasource

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI string
}

func NewMongoConfig() *MongoConfig {
	return &MongoConfig{
		URI: "mongodb://localhost:27017",
	}
}

func NewMongoClient(cfg *MongoConfig) (*mongo.Client, error) {
	credential := options.Credential{
		Username: "root",
		Password: "root",
	}
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(cfg.URI).SetMaxConnecting(100).SetMaxPoolSize(300).SetAuth(credential))
	if err != nil {
		return nil, err
	}
	return client, nil
}
