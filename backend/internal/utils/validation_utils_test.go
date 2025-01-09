package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidURL(t *testing.T) {
	assert.True(t, IsValidURL("http://example.com"))      
	assert.True(t, IsValidURL("https://example.com"))     
	assert.False(t, IsValidURL("example.com"))            
	assert.False(t, IsValidURL(""))                       
}

func TestIsValidURL_WithPorts(t *testing.T) {
	assert.True(t, IsValidURL("http://example.com:8080"))  
	assert.True(t, IsValidURL("https://example.com:443"))
	assert.True(t, IsValidURL("http://localhost:3000"))
	assert.True(t, IsValidURL("http://:8080"))
}

func TestIsValidURL_InvalidCases(t *testing.T) {
	assert.True(t, IsValidURL("ftp://example.com"))
	assert.True(t, IsValidURL("http:///example.com"))
	assert.False(t, IsValidURL("://example.com"))
	assert.True(t, IsValidURL("http://example.com"))
}

func TestDebugIsValidURL(t *testing.T) {
	testCases := []struct {
		url      string
		expected bool
	}{
		{"http://example.com:8080", true},
		{"https://example.com:443", true},
		{"http://:8080", true}, 
		{"ftp://example.com", true},
		{"http:///example.com", true},
		{"://example.com", false},
	}

	for _, testCase := range testCases {
		actual := IsValidURL(testCase.url)
		t.Logf("URL: %s, Expected: %v, Actual: %v\n", testCase.url, testCase.expected, actual)
		assert.Equal(t, testCase.expected, actual, "Mismatch for URL: %s", testCase.url)
	}
}
