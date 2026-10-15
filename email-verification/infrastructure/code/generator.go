package code

import (
	"crypto/rand"
	"crypto/sha3"
	"encoding/hex"
	"fmt"
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
	sum := sha3.Sum256([]byte(code))
	return hex.EncodeToString(sum[:])
}
