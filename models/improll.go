package models

import (
	"time"
)

type SharedCategory struct {
	ID        *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title     *string   `json:"title"`
	Category  Category  `json:category`
	Downloads *bool     `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	Tags      [*string] `json:"tags"`
}

type Category struct {
}