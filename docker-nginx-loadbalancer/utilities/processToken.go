package utilities

import (
	"errors"
	"strings"

	database "fiber-docker-nginx/database/user""

	jwt "github.com/dgrijalva/jwt-go"
	"fiber-docker-nginx/models"
)

/*Email : value of the user email*/
var Email string

/*UserID : value of the user id*/
var UserID string

/*ProcessToken : Process jwt and extract values*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("thisisnotyendo")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("jwt format not valid")
	}

	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, exist, _ := database.UserExist(claims.Email)
		if exist {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, exist, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid jwt")
	}
	return claims, false, string(""), err
}
