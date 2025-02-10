package code

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

type DefaultCodeGenerator struct{}

func NewCodeGenerator() *DefaultCodeGenerator {
	return &DefaultCodeGenerator{}
}

func (g *DefaultCodeGenerator) Generate() (string, error) {
	b := make([]byte, 3)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%06x", b), nil
}

func (g *DefaultCodeGenerator) Hash(code string) string {
	hash := sha3.New256()
	hash.Write([]byte(code))
	return hex.EncodeToString(hash.Sum(nil))
}
