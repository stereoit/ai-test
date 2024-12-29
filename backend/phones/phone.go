package phones

import (
	"time"
)

type Phone struct {
	ID        string    `json:"id`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Brand     string    `json:"brand"`
}
