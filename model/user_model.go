package model

type User struct {
	BaseModel `bson:",inline"` // Embed BaseModel
	Name      string
	Email     string
}