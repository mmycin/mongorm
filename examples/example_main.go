package main

import (
	"fmt"

	"github.com/mmycin/mongorm"
	"github.com/mmycin/mongorm/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Notice struct {
	model.BaseModel `bson:",inline"`
	Date            string `json:"date"`
	Title           string `json:"title"`
	Content         string `json:"content"`
}

func main() {
	_, err := mongorm.Initialize("uri", "Notice")
	if err != nil {
		panic(err)
	}
	// var notices []Notice
	// err = mongorm.ReadAll("notice", &notices)
	// fmt.Printf("%v", notices)

	var notice1 = Notice{
		Date:    "06/09/2024",
		Title:   "Test Doc",
		Content: "A test content",
	}
	err = mongorm.CreateOne("notice", notice1)
	var notice Notice
	err = mongorm.ReadOne("notice", bson.M{"_id": notice1.ID}, &notice)
	fmt.Printf("notice: %v\n", notice)
}
