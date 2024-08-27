package mongorm

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect establishes a connection to the MongoDB server and returns the client.
func Connect(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client, nil
}

// Create inserts a new document into the specified collection.
func (m *Model) Create(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {
	collection := db.Collection(collectionName)

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}

	// Ensure InsertedID is of type primitive.ObjectID
	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		m.ID = id
	} else {
		return fmt.Errorf("unexpected type for InsertedID: %T", res.InsertedID)
	}

	return nil
}

// Read retrieves a single document from the specified collection based on the filter.
func (m *Model) Read(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return fmt.Errorf("failed to find document: %w", err)
	}

	return nil
}

// Update modifies an existing document in the specified collection based on the filter and update parameters.
func (m *Model) Update(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
	collection := db.Collection(collectionName)

	m.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}

	return nil
}

// Delete removes a document from the specified collection based on the filter.
func (m *Model) Delete(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}

	return nil
}
