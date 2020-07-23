package database

import (
	"fiber-docker-nginx/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin : try to login in the APP*/
func TryLogin(email string, password string) (models.User, bool) {
	user, isRegister, _ := UserExist(email)
	if !isRegister {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
