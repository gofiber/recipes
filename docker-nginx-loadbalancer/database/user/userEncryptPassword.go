package database

import "golang.org/x/crypto/bcrypt"

/*EncryptPassword : Encrypt the password*/
func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
