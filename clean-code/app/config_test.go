package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("DATABASE_URL", "postgres://user:pass@localhost:5432/dbname")
	defer os.Unsetenv("PORT")
	defer os.Unsetenv("DATABASE_URL")

	conf := NewConfiguration()

	assert.Equal(t, "8080", conf.Port)
	assert.Equal(t, "postgres://user:pass@localhost:5432/dbname", conf.DatabaseURL)
}

func TestNewConfiguration_Defaults(t *testing.T) {
	os.Unsetenv("PORT")
	os.Unsetenv("DATABASE_URL")

	conf := NewConfiguration()

	assert.Equal(t, "3000", conf.Port)
	assert.Equal(t, "", conf.DatabaseURL)
}

func TestGetEnvOrDefault(t *testing.T) {
	os.Setenv("TEST_ENV", "value")
	defer os.Unsetenv("TEST_ENV")

	value := getEnvOrDefault("TEST_ENV", "default")
	assert.Equal(t, "value", value)

	value = getEnvOrDefault("NON_EXISTENT_ENV", "default")
	assert.Equal(t, "default", value)
}
