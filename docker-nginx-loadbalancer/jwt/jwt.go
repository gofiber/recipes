package jwt

import (
	"time"

	"fiber-docker-nginx/models"

	jwt "github.com/dgrijalva/jwt-go"
)

/*GenerateJWT : Generate jwt with user info*/
func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("sshh!")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastname":  user.Lastname,
		"birthdate": user.Birthdate,
		"location":  user.Location,
		"avatar":    user.Avatar,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 1500).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
