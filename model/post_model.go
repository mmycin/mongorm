package model

// Post model for post collection.
type Post struct {
	BaseModel
	Title   string `bson:"title"`
	Content string `bson:"content"`
	Author  string `bson:"author"`
}
