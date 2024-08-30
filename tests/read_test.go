package test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/mmycin/mongorm"
	"github.com/mmycin/mongorm/model"
	"github.com/mmycin/mongorm/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestReadOne(t *testing.T) {
	err := mongorm.Initialize("url_string", "test2db")
	utils.HandleError(err)

	// Specify the ObjectID of the user to find
	userID, err := primitive.ObjectIDFromHex("66d1eb345dc13732bd3e6fed")
	if err != nil {
		log.Fatalf("Invalid ObjectID format: %v", err)
	}

	// Define the filter to match the specific user by ID
	filter := bson.M{"_id": userID}

	// Read the user with the specified ID
	var user model.User
	err = mongorm.ReadOne(context.Background(), "users", filter, &user)
	if err != nil {
		utils.HandleError(err)
	}

	// Print the user details
	if user.ID.IsZero() {
		fmt.Println("No user found with the specified ID.")
	} else {
		fmt.Printf("ID: %s, Name: %s, Email: %s\n", user.ID.Hex(), user.Name, user.Email)
	}
}