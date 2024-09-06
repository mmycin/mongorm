package mongorm

import (
	"github.com/mmycin/mongorm/client"
	"github.com/mmycin/mongorm/operations"
	"go.mongodb.org/mongo-driver/mongo"
)

// Initialize sets up the MongoDB client and database
func Initialize(uri, dbName string) (*mongo.Client, error) {
	config := client.MongoConfig{
		URI:    uri,
		DBName: dbName,
	}
	return client.Initialize(config)
}


// CreateOne inserts a single document.
func CreateOne(collectionName string, doc interface{}) error {
	return operations.CreateOne(client.Database, collectionName, doc)
}

// ReadOne retrieves a single document.
func ReadOne(collectionName string, filter interface{}, result interface{}) error {
	return operations.ReadOne(client.Database, collectionName, filter, result)
}

// ReadAll retrieves multiple documents.
func ReadAll(collectionName string, results interface{}) error {
	return operations.ReadAll(client.Database, collectionName, results)
}

// Update modifies an existing document.
func Update(collectionName string, filter interface{}, update interface{}) error {
	return operations.Update(client.Database, collectionName, filter, update)
}

// DeleteOne removes a single document.
func DeleteOne( collectionName string, filter interface{}) error {
	return operations.DeleteOne(client.Database, collectionName, filter)
}

// DeleteAll removes multiple documents.
func DeleteAll(collectionName string, filter interface{}) error {
	return operations.DeleteAll(client.Database, collectionName, filter)
}

// HandleError panics if an error is encountered.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
