package operations

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadOne(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return fmt.Errorf("failed to find document: %w", err)
	}

	return nil
}

func ReadAll(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, results interface{}) error {
	collection := db.Collection(collectionName)

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, results); err != nil {
		return fmt.Errorf("failed to decode documents: %w", err)
	}

	return nil
}
