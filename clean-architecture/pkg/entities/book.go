package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Book Constructs your Book model under entities.
type Book struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Author    string             `json:"author" bson:"author,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Reqeusts for Books
type DeleteRequest struct {
	ID string `json:"id"`
}
