package operations

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Update modifies an existing document in the specified collection based on the filter and update parameters.
func Update(db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
	collection := db.Collection(collectionName)
	updateSet := bson.M{"$set": update}
	// Check if update is a map (bson.M) and add UpdatedAt field
	if updateDoc, ok := update.(bson.M); ok {
		updateDoc["$set"].(bson.M)["updated_at"] = time.Now()
	} else if updateDoc, ok := update.(bson.D); ok {
		for i, elem := range updateDoc {
			if elem.Key == "$set" {
				updateDoc[i].Value.(bson.M)["updated_at"] = time.Now()
				break
			}
		}
	}

	_, err := collection.UpdateOne(context.Background(), filter, updateSet)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}

	return nil
}
