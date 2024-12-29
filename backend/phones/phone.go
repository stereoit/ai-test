package phones

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phone struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Name      string             `json:"name"`
	Price     float64            `json:"price"`
	Brand     string             `json:"brand"`
}
