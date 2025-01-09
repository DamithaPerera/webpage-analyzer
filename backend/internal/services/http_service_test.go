package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultHTTPClient_Get_Success(t *testing.T) {
	client := &DefaultHTTPClient{}

	// Mock a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	resp, err := client.Get(server.URL)
	assert.NoError(t, err, "Expected no error from Get request")
	assert.NotNil(t, resp, "Expected non-nil response")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
}

func TestDefaultHTTPClient_Get_Error(t *testing.T) {
	client := &DefaultHTTPClient{}

	// Attempt to connect to an invalid server URL
	resp, err := client.Get("http://nonexistent.server.invalid")

	// Validate the response
	assert.Error(t, err, "Expected an error for invalid server URL")
	assert.Nil(t, resp, "Expected nil response for invalid server URL")
}

