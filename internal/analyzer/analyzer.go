package handlers

import (
    "net/http"
    "webpage-analyzer/internal/utils"
    "github.com/gin-gonic/gin"
)

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

    c.JSON(http.StatusOK, gin.H{"message": "URL is valid", "url": req.URL})
}