package operations

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteOne(db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}

	return nil
}

func DeleteAll(db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete documents: %w", err)
	}

	return nil
}
