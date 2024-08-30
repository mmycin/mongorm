package client

import "go.mongodb.org/mongo-driver/mongo"

type MongoConfig struct {
	URI    string
	DBName string
}

var Client *mongo.Client
var Database *mongo.Database

// Initialize sets up the MongoDB client and database.
func Initialize(config MongoConfig) error {
	var err error
	Client, err = Connect(config.URI)
	if err != nil {
		return err
	}

	Database = Client.Database(config.DBName)
	return nil
}
