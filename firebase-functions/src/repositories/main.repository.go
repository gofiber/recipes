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

func (self *MainRepository) InsertMessage(body *models.MessageInputBody) error {
	id := uuid.New().String()
	_, err := self.DB.Collection("messages").Doc(id).Set(context.Background(), body)
	return err
}
