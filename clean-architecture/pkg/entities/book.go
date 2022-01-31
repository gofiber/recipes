package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Construct your model under entities.
type Book struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// Only one struct per file should exists unless another struct is closely related with the one defined in this file.
type DeleteRequest struct {
	ID string `json:"id"`
}
