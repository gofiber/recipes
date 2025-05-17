package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/datasources"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetStatus(t *testing.T) {
	app := NewServer(&datasources.DataSources{})

	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/status", nil))
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, "ok", string(body))
}
