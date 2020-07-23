package database

import (
	"context"
	"time"

	"fiber-docker-nginx/database"e"

	"fiber-docker-nginx/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*UserRegister : Register a user in the database*/
func UserRegister(user models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := database.MongoCN.Database("yendo")
	col := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
