package operations

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadOne(db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return fmt.Errorf("failed to find document: %w", err)
	}

	return nil
}

func ReadAll(db *mongo.Database, collectionName string, results interface{}) error {
	collection := db.Collection(collectionName)
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to find documents: %w", err)
	}
	
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, results); err != nil {
		return fmt.Errorf("failed to decode documents: %w", err)
	}
	fmt.Printf("\n\n\n%v", results)
	
	
	return nil
}
