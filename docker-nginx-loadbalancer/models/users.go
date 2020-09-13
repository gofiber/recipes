package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User : The User model*/
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Lastname  string             `bson:"lastname" json:"lastname,omitempty"`
	Birthdate time.Time          `bson:"birthdate" json:"birthdate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Name,validation.Required),
		validation.Field(&u.Lastname,validation.Required),
		validation.Field(&u.Birthdate,validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password,validation.Required,validation.Length(6,32)),
		validation.Field(&u.Avatar,validation.Required,is.URL),
		validation.Field(&u.Location,validation.Required),
	)
}