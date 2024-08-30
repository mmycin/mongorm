package test

import (
	"context"
	"testing"
	"time"

	"github.com/mmycin/mongorm"
	"github.com/mmycin/mongorm/model"
	"github.com/mmycin/mongorm/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateOne(t *testing.T) {
	// Initialize MongoDB connection
	err := mongorm.Initialize("url_string", "test2db")
	utils.HandleError(err)

	// Create a new user for testing
	user := model.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	err = mongorm.CreateOne(context.Background(), "users", &user)
	utils.HandleError(err)

	// Define the filter to match the user to update
	userID := user.ID
	filter := bson.M{"_id": userID}

	// Define the update parameters
	update := bson.M{"$set": bson.M{"email": "john.updated@example.com"}}

	// Update the user
	err = mongorm.Update(context.Background(), "users", filter, update)
	utils.HandleError(err)

	// Verify the update
	var updatedUser model.User
	err = mongorm.ReadOne(context.Background(), "users", filter, &updatedUser)
	utils.HandleError(err)

	if updatedUser.Email != "john.updated@example.com" {
		t.Errorf("Expected email to be 'john.updated@example.com', but got '%s'", updatedUser.Email)
	}

	// Verify the UpdatedAt field is updated
	if updatedUser.UpdatedAt.Before(time.Now().Add(-time.Minute)) {
		t.Errorf("Expected UpdatedAt to be recent, but got '%v'", updatedUser.UpdatedAt)
	}
}
