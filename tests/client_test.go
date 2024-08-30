package test

import (
	"testing"

	"github.com/mmycin/mongorm/client"
	"github.com/mmycin/mongorm/utils"
)

func TestInitialize(t *testing.T) {
	err := client.Initialize(client.MongoConfig{
		URI:    "url_string",
		DBName: "testdb",
	})
	utils.HandleError(err)
}
