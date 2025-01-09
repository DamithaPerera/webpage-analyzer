package analyzer

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHTTPClient struct {
	Response *http.Response
	Error    error
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.Response, m.Error
}

func TestAnalyze_Success(t *testing.T) {
	mockHTML := "<html><head><title>Example Title</title></head><body></body></html>"
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(mockHTML)),
		},
		Error: nil,
	}

	result, err := Analyze("http://example.com", mockClient)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Example Title", result.Title)
	assert.Equal(t, "HTML5", result.HTMLVersion) // Assuming HTML5 detection
}

func TestAnalyze_FetchError(t *testing.T) {
	mockClient := &MockHTTPClient{
		Response: nil,
		Error:    errors.New("failed to fetch"),
	}

	result, err := Analyze("http://example.com", mockClient)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "unable to fetch the URL", err.Error())
}

func TestAnalyze_Non200Response(t *testing.T) {
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(bytes.NewBufferString("")),
		},
		Error: nil,
	}

	result, err := Analyze("http://example.com", mockClient)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "non-success HTTP status received: Internal Server Error", err.Error())
}

func TestAnalyze_InvalidHTML(t *testing.T) {
	// Simulate malformed HTML that might still partially parse
	mockClient := &MockHTTPClient{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("<html><invalid></html>")), // Malformed but parseable
		},
		Error: nil,
	}

	// Call the Analyze function
	result, err := Analyze("http://example.com", mockClient)

	// Validate the result
	if err == nil {
		// If no error, ensure the result is partially populated or invalid
		assert.Equal(t, "", result.Title, "Expected title to be empty for malformed HTML")
		assert.NotNil(t, result)
	} else {
		// If an error is returned, ensure it matches expectations
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "error parsing HTML document", err.Error())
	}
}

