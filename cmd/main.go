package main

import (
    "github.com/gin-gonic/gin"
    "github.com/penglongli/gin-metrics/ginmetrics"
    "webpage-analyzer/internal/handlers"
    "webpage-analyzer/internal/utils"
)

func main() {
    utils.InitLogger()
    utils.Logger.Info("Starting Web Page Analyzer server")

    r := gin.Default()

    m := ginmetrics.GetMonitor()
    m.SetMetricPath("/metrics")
    m.Use(r)

    r.GET("/", handlers.HomePage)
    r.POST("/analyze", handlers.AnalyzePage)

    if err := r.Run(":8080"); err != nil {
        utils.Logger.Fatal("Failed to start server: ", err)
    }
}