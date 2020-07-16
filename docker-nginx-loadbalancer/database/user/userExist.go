package database

import (
	"context"
	"time"

	"fiber-docker-nginx/database"e"

	"fiber-docker-nginx/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*UserExist : Check if the user is alredy register*/
func UserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	defer cancel()

	db := database.MongoCN.Database("yendo")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
