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
    c.JSON(http.StatusOK, gin.H{
        "message": "Analyzing page testign",
    })
}