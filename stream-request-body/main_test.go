package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestStreamRequestBody(t *testing.T) {
	tests := []struct {
		description  string
		body         []byte
		expectedCode int
	}{
		{
			description:  "small body",
			body:         []byte("hello world"),
			expectedCode: http.StatusOK,
		},
		{
			description:  "empty body",
			body:         []byte{},
			expectedCode: http.StatusOK,
		},
		{
			description:  "binary body",
			body:         bytes.Repeat([]byte{0x01, 0x02, 0x03}, 100),
			expectedCode: http.StatusOK,
		},
	}

	app := setup()

	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewReader(test.body))
		req.Header.Set("Content-Type", "application/octet-stream")

		res, err := app.Test(req, fiber.TestConfig{Timeout: 0, FailOnTimeout: false})

		assert.Nilf(t, err, test.description)
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
	}
}
