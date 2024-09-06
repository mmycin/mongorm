package test

import (
	"testing"

	"github.com/mmycin/mongorm"
	"github.com/mmycin/mongorm/model"
	"github.com/mmycin/mongorm/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func TestDeleteOne(t *testing.T) {
	// Initialize MongoDB connection
	_, err := mongorm.Initialize("mongodb+srv://Mycin:myc23084.fun@cluster0.yzel00n.mongodb.net/", "test2db")
	utils.HandleError(err)

	// Create a new user for testing
	user := model.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	err = mongorm.CreateOne("users", &user)
	utils.HandleError(err)

	// Define the filter to match the user to delete
	userID := user.ID
	filter := bson.M{"_id": userID}

	// Delete the user
	err = mongorm.DeleteOne("users", filter)
	utils.HandleError(err)

	// Verify the deletion
	var deletedUser model.User
	err = mongorm.ReadOne("users", filter, &deletedUser)
	if err == nil {
		t.Errorf("Expected error when reading deleted user, but got none")
	} else {
		t.Logf("Successfully deleted user, error: %v", err)
	}
}
