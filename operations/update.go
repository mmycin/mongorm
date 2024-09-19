package operations

import (
	"context"
	"fmt"
	"time"

	"github.com/mmycin/mongorm/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Update modifies an existing document in the specified collection based on the filter and update parameters.
func Update(db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
	collection := db.Collection(collectionName)
	updateSet := utils.Json{"$set": update}

	// Add UpdatedAt field to the update document
	if updateDoc, ok := update.(utils.Json); ok {
		updateDoc["updated_at"] = time.Now()
	} else if updateDoc, ok := update.(bson.D); ok {
		updateDoc = append(updateDoc, bson.E{Key: "updated_at", Value: time.Now()})
		updateSet["$set"] = updateDoc
	}

	_, err := collection.UpdateOne(context.Background(), filter, updateSet)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}

	return nil
}

