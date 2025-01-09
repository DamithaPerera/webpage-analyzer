package handlers

import (
	"net/http"
	"webpage-analyzer/internal/analyzer"
	"webpage-analyzer/internal/services"
	"webpage-analyzer/internal/utils"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Web Page Analyzer. Use POST /analyze to analyze a webpage.",
	})
}

func AnalyzePage(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Please provide a valid URL."})
		return
	}

	if !utils.IsValidURL(req.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format."})
		return
	}

	client := &services.DefaultHTTPClient{}
	result, err := analyzer.Analyze(req.URL, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}