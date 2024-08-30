package test

import (
	"context"
	"testing"

	"github.com/mmycin/mongorm"
	"github.com/mmycin/mongorm/model"
	"github.com/mmycin/mongorm/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateOne(t *testing.T) {
	// Initialize MongoDB connection
	err := mongorm.Initialize("url_string", "test2db")
	utils.HandleError(err)

	// Create a new user
	user := model.User{
		Name:  "Musfirat jahan nafisa",
		Email: "jane@example.com",
	}

	// Insert the user into the database
	err = mongorm.CreateOne(context.Background(), "users", &user)
	utils.HandleError(err)

	// Verify the creation by reading the user back from the database
	var createdUser model.User
	err = mongorm.ReadOne(context.Background(), "users", bson.M{"_id": user.ID}, &createdUser)
	if err != nil {
		t.Errorf("Failed to read created user: %v", err)
		return
	}

	// Check if the read user matches the created user
	if createdUser.Name != user.Name || createdUser.Email != user.Email {
		t.Errorf("User data mismatch. Expected: %v, Got: %v", user, createdUser)
	}

	t.Logf("User created successfully with ID: %s", createdUser.ID.Hex())
}
