package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

//Construct your model under entities.
type Book struct {
	ID     primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Title  string             `json:"title" binding:"required,min=2" bson:"title"`
	Author string             `json:"author" bson:"author,omitempty"`
}

// Only one struct per file should exists unless another struct is closely related with the one defined in this file.
type DeleteRequest struct {
	ID string `json:"id" binding:"required,gte=1"`
}
