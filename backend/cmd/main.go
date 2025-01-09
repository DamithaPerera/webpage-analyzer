package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "webpage-analyzer/internal/handlers"
    "webpage-analyzer/internal/utils"
)

func main() {
    utils.InitLogger()
    utils.Logger.Info("Starting Web Page Analyzer server")

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type"},
        AllowCredentials: true,
    }))

    r.GET("/", handlers.HomePage)
    r.POST("/analyze", handlers.AnalyzePage)

    if err := r.Run(":8080"); err != nil {
        utils.Logger.Fatal("Failed to start server: ", err)
    }
}