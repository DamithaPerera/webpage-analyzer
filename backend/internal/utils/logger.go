package utils

import (
    "github.com/sirupsen/logrus"
)

var Logger = logrus.New()

//initializes the global logger with a JSON formatter and info level
func InitLogger() {
    Logger.SetFormatter(&logrus.JSONFormatter{})
    Logger.SetLevel(logrus.InfoLevel)
}
