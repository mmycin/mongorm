package client

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect establishes a connection to the MongoDB server.
func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to MongoDB: %w", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("unable to ping MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB successfully")
	return client, nil
}
