package main

import (
    "github.com/gin-gonic/gin"
    "webpage-analyzer/internal/handlers"
)

func main() {
    r := gin.Default()

    r.GET("/", handlers.HomePage)
    r.POST("/analyze", handlers.AnalyzePage)

    if err := r.Run(":8080"); err != nil {
        panic(err)
    }
}