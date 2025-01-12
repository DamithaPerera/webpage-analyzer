package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	HomePage(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Welcome to the Web Page Analyzer")
}

func TestAnalyzePage_ValidInput(t *testing.T) {
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    reqBody := `{"url":"http://example.com"}`
    c.Request = httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBufferString(reqBody))
    c.Request.Header.Set("Content-Type", "application/json")

    AnalyzePage(c)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "html_version")
    assert.Contains(t, w.Body.String(), "title")
    assert.Contains(t, w.Body.String(), "headings")
}


func TestAnalyzePage_InvalidInput(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	reqBody := `{"url":""}` // Invalid input
	c.Request = httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBufferString(reqBody))
	c.Request.Header.Set("Content-Type", "application/json")

	AnalyzePage(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid input")
}

func TestAnalyzePage_ErrorFromAnalyzer(t *testing.T) {
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    reqBody := `{"url":"http://invalid-url.com"}` // Simulate an invalid URL
    c.Request = httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBufferString(reqBody))
    c.Request.Header.Set("Content-Type", "application/json")

    AnalyzePage(c)

    // Assert the HTTP status
    assert.Equal(t, http.StatusInternalServerError, w.Code)

    // Assert the response matches the expected JSON structure
    expectedResponse := `{"error":"unable to fetch the URL"}`
    assert.JSONEq(t, expectedResponse, w.Body.String(), "Expected error response does not match")
}
