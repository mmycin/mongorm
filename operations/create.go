package operations

import (
	"context"
	"fmt"

	"github.com/mmycin/mongorm/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateOne inserts a single document into the specified collection.
func CreateOne(db *mongo.Database, collectionName string, doc interface{}) error {
	collection := db.Collection(collectionName)

	// Check if doc is a pointer to BaseModel or a value of BaseModel
	if baseModel, ok := doc.(*model.BaseModel); ok {
		baseModel.PreSave() // Update timestamps
	} else if baseModel, ok := doc.(model.BaseModel); ok {
		// Update timestamps on a copy of the BaseModel
		baseModel.PreSave()
		doc = &baseModel // Ensure doc is a pointer to the updated BaseModel
	}

	// Insert the document
	result, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}

	// Handle the inserted ID
	if baseModel, ok := doc.(*model.BaseModel); ok {
		if id, ok := result.InsertedID.(primitive.ObjectID); ok {
			baseModel.ID = id
		} 
	} 	

	return nil
}
