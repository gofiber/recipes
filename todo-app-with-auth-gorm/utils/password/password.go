package password

import "golang.org/x/crypto/bcrypt"

// Generate return a hashed password
func Generate(raw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

// Verify compares a hashed password with plaintext password
func Verify(hash string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}
