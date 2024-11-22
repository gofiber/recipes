package server

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"

	"app/datasources"

	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {
	app := NewServer(context.Background(), &datasources.DataSources{})

	resp, err := app.Test(httptest.NewRequest("GET", "/api/status", nil))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "ok", string(body))
}
