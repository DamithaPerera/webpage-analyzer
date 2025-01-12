package utils

import (
    "os"
    "path/filepath"
    "github.com/sirupsen/logrus"
)

var Logger = logrus.New()

// InitLogger initializes the global logger with a JSON formatter and info level.
func InitLogger() {
    Logger.SetFormatter(&logrus.JSONFormatter{})
    Logger.SetLevel(logrus.InfoLevel)

    // Set a consistent log file path
    logFilePath := filepath.Join(getBaseDir(), "webpage-analyzer.log")

    // Create or open the log file
    logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        Logger.Fatalf("Failed to open log file: %v", err)
    }

    Logger.SetOutput(logFile)
}

// getBaseDir determines the base directory for the log file.
func getBaseDir() string {
    // Ensure logs are written to the project root directory
    baseDir, err := os.Getwd()
    if err != nil {
        Logger.Fatalf("Failed to get working directory: %v", err)
    }
    return baseDir
}
