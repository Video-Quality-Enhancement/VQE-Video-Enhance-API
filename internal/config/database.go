package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient interface {
	ConnectToDB() *mongo.Database
	Disconnect() error
}

type mongoClient struct {
	client *mongo.Client
}

func NewMongoClient() MongoClient { // *  v.v.v.imp MongoClient and not *mongoClient, because we want to return an interface and not a struct

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &mongoClient{
		client: client,
	}
}

func (m *mongoClient) ConnectToDB() *mongo.Database {

	return m.client.Database(os.Getenv("MONGO_DB"))

}

func (m *mongoClient) Disconnect() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := m.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
