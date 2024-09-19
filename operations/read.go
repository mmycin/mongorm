package operations

import (
	"context"
	"fmt"

	"github.com/mmycin/mongorm/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadOne(db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return fmt.Errorf("failed to find document: %w", err)
	}

	// Ensure the ID is correctly set in the result
	if idSetter, ok := result.(interface{ SetID(id interface{}) }); ok {
		idSetter.SetID(filter.(utils.Json)["_id"])
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

	// Ensure the IDs are correctly set in the results
	if idSetter, ok := results.(interface{ SetIDs(ids []interface{}) }); ok {
		var ids []interface{}
		for cursor.Next(ctx) {
			var doc bson.M
			if err := cursor.Decode(&doc); err != nil {
				return fmt.Errorf("failed to decode document: %w", err)
			}
			ids = append(ids, doc["_id"])
		}
		idSetter.SetIDs(ids)
	}

	fmt.Printf("\n\n\n%v", results)
	
	return nil
}
