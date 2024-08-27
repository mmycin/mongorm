package mongorm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
