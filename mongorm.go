package mongorm

import (
	"context"
	"github.com/mmycin/mongorm/client"
	"github.com/mmycin/mongorm/operations"
)

// Initialize sets up the MongoDB client and database.
func Initialize(uri, dbName string) (*mongo.Client, error) {
	config := client.MongoConfig{
		URI:    uri,
		DBName: dbName,
	}
	return client.Initialize(config)
}

// CreateOne inserts a single document.
func CreateOne(ctx context.Context, collectionName string, doc interface{}) error {
	return operations.CreateOne(ctx, client.Database, collectionName, doc)
}

// ReadOne retrieves a single document.
func ReadOne(ctx context.Context, collectionName string, filter interface{}, result interface{}) error {
	return operations.ReadOne(ctx, client.Database, collectionName, filter, result)
}

// ReadAll retrieves multiple documents.
func ReadAll(ctx context.Context, collectionName string, filter interface{}, results interface{}) error {
	return operations.ReadAll(ctx, client.Database, collectionName, filter, results)
}

// Update modifies an existing document.
func Update(ctx context.Context, collectionName string, filter interface{}, update interface{}) error {
	return operations.Update(ctx, client.Database, collectionName, filter, update)
}

// DeleteOne removes a single document.
func DeleteOne(ctx context.Context, collectionName string, filter interface{}) error {
	return operations.DeleteOne(ctx, client.Database, collectionName, filter)
}

// DeleteAll removes multiple documents.
func DeleteAll(ctx context.Context, collectionName string, filter interface{}) error {
	return operations.DeleteAll(ctx, client.Database, collectionName, filter)
}

// HandleError panics if an error is encountered.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
