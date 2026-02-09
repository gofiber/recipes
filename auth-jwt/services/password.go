package services

import "golang.org/x/crypto/bcrypt"

// HashPassword creates a bcrypt hash from a plain-text password
func HashPassword(password string) (string, error) {
	// The cost determines how computationally expensive the hash is
	// Higher is more secure but slower (default is 10)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// VerifyPassword checks if the provided password matches the stored hash
func VerifyPassword(hashedPassword, providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
}


