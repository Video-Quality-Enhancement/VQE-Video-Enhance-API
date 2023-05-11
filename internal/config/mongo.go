package config

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slog"
)

type MongoClient interface {
	ConnectToDB() *mongo.Database
	Disconnect()
}

type mongoClient struct {
	client *mongo.Client
}

func NewMongoClient() MongoClient { // *  v.v.v.imp MongoClient and not *mongoClient, because we want to return an interface and not a struct

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		slog.Error("Failed to connect to MongoDB", "err", err)
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		slog.Error("Failed to ping MongoDB", "err", err)
		panic(err)
	}

	return &mongoClient{
		client: client,
	}
}

func (m *mongoClient) ConnectToDB() *mongo.Database {

	return m.client.Database(os.Getenv("MONGO_DB"))

}

func (m *mongoClient) Disconnect() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := m.client.Disconnect(ctx)
	if err != nil {
		slog.Error("Failed to disconnect from MongoDB", "err", err)
		panic(err)
	}

}
