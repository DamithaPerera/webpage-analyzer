package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
    "testing"
	"webpage-analyzer/internal/models"
	"webpage-analyzer/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	HealthCheck(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Web Page Analyzer service is running.")
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

// MockAnalyzer simulates the behavior of the Analyze function
type MockAnalyzer struct {
	AnalyzeFunc func(url string, client services.HTTPClient) (*models.AnalysisResult, error)
}

func (m *MockAnalyzer) Analyze(url string, client services.HTTPClient) (*models.AnalysisResult, error) {
	return m.AnalyzeFunc(url, client)
}

// Updated AnalyzePage function for testing
func AnalyzePageWithMock(c *gin.Context, mockAnalyzer *MockAnalyzer) {
	var req struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Please provide a valid URL."})
		return
	}

	client := &services.DefaultHTTPClient{}
	result, err := mockAnalyzer.Analyze(req.URL, client)
	if err != nil {
		if err.Error() == "504 Gateway Timeout: The server, while acting as a gateway, did not receive a timely response" {
			c.JSON(http.StatusGatewayTimeout, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}


func TestAnalyzePage_504Error(t *testing.T) {
	mockAnalyzer := &MockAnalyzer{
		AnalyzeFunc: func(url string, client services.HTTPClient) (*models.AnalysisResult, error) {
			return nil, errors.New("504 Gateway Timeout: The server, while acting as a gateway, did not receive a timely response")
		},
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	reqBody := `{"url":"http://example.com"}` // Simulate a request triggering a 504 error
	c.Request = httptest.NewRequest(http.MethodPost, "/analyze", bytes.NewBufferString(reqBody))
	c.Request.Header.Set("Content-Type", "application/json")

	AnalyzePageWithMock(c, mockAnalyzer)

	assert.Equal(t, http.StatusGatewayTimeout, w.Code)
	assert.Contains(t, w.Body.String(), "504 Gateway Timeout")
}