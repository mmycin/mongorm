package operations

import (
	"context"
	"fmt"

	"github.com/mmycin/mongorm/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteOne(db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)

	// Ensure filter is of type bson.M
	filterBson, ok := filter.(utils.Json)
	if !ok {
		return fmt.Errorf("filter must be of type bson.M")
	}

	_, err := collection.DeleteOne(context.Background(), filterBson)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}

	return nil
}

func DeleteAll(db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)

	// Ensure filter is of type bson.M
	filterBson, ok := filter.(utils.Json)
	if !ok {
		return fmt.Errorf("filter must be of type bson.M")
	}

	_, err := collection.DeleteMany(context.Background(), filterBson)
	if err != nil {
		return fmt.Errorf("failed to delete documents: %w", err)
	}

	return nil
}
