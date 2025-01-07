package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to the Web Page Analyzer.",
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

    c.JSON(http.StatusOK, gin.H{
        "message": "Received URL for analysis",
        "url":     req.URL,
    })
}