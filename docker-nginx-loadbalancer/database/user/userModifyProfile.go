package database

import (
	"context"
	"errors"
	"time"

	"fiber-docker-nginx/database"e"

	"fiber-docker-nginx/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModifyProfile : Update the values in a profile*/
func ModifyProfile(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := database.MongoCN.Database("yendo")
	col := db.Collection("users")

	register := make(map[string]interface{})

	if len(user.Name) > 0 {
		register["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		register["lastname"] = user.Lastname
	}

	register["birthdate"] = user.Birthdate

	if len(user.Password) > 0 {
		register["password"] = user.Password
	}
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}
	if len(user.Location) > 0 {
		register["location"] = user.Location
	}

	updateString := bson.M{
		"$set": register,
	}

	objtID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objtID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, errors.New("Not updated")
	}
	return true, nil
}
