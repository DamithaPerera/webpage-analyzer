package handlers

import (
    "net/http"
    "webpage-analyzer/internal/analyzer"
    "webpage-analyzer/internal/services"
    "webpage-analyzer/internal/utils"

    "github.com/gin-gonic/gin"
)

// HomePage handles the root endpoint and provides a welcome message.
func HealthCheck(c *gin.Context) {
    utils.Logger.Info("HealthCheck endpoint accessed")
    c.JSON(http.StatusOK, gin.H{
        "status":  "OK",
        "message": "Web Page Analyzer service is running.",
    })
}

// AnalyzePage handles the /analyze endpoint to analyze a given webpage.
func AnalyzePage(c *gin.Context) {
    utils.Logger.Info("AnalyzePage endpoint accessed")

    var req struct {
        URL string `json:"url" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        utils.Logger.Warn("Invalid input: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Please provide a valid URL."})
        return
    }

    if !utils.IsValidURL(req.URL) {
        utils.Logger.Warn("Invalid URL format: ", req.URL)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format."})
        return
    }

    utils.Logger.Info("Starting analysis for URL: ", req.URL)
    client := &services.DefaultHTTPClient{}
    result, err := analyzer.Analyze(req.URL, client)
    if err != nil {
        if err.Error() == "503 Service Unavailable: The server is currently unable to handle the request" {
            c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
            return
        }

        if err.Error() == "504 Gateway Timeout: The server, while acting as a gateway, did not receive a timely response" {
            c.JSON(http.StatusGatewayTimeout, gin.H{"error": err.Error()})
            return
        }

        utils.Logger.Error("Error analyzing URL: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    utils.Logger.Info("Successfully analyzed URL: ", req.URL)
    c.JSON(http.StatusOK, result)
}
