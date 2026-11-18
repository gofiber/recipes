package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"example.com/GofiberFirebaseBoilerplate/src/models"
	"github.com/google/uuid"
)

type MainRepository struct {
	DB *firestore.Client
}

// InsertMessage stores the message and returns the id of the new document.
func (r *MainRepository) InsertMessage(body *models.MessageInputBody) (string, error) {
	id := uuid.New().String()
	if _, err := r.DB.Collection("messages").Doc(id).Set(context.Background(), body); err != nil {
		return "", err
	}
	return id, nil
}
