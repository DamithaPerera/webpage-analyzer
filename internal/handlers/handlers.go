package handlers

import (
    "net/http"
    "webpage-analyzer/internal/analyzer"
    "webpage-analyzer/internal/services"
    "webpage-analyzer/internal/utils"

    "github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
    utils.Logger.Info("HomePage endpoint accessed")
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to the Web Page Analyzer. Use POST /analyze to analyze a webpage.",
    })
}

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
        utils.Logger.Error("Error analyzing URL: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    utils.Logger.Info("Successfully analyzed URL: ", req.URL)
    c.JSON(http.StatusOK, result)
}