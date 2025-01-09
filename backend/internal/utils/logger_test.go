package utils

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogger_Debug(t *testing.T) {
	var buf bytes.Buffer
	Logger.SetOutput(&buf)

	Logger.SetLevel(logrus.DebugLevel)
	Logger.Debug("Test debug message")

	// Verify the output contains the expected message
	assert.Contains(t, buf.String(), "Test debug message")
}

func TestLogger_Panic(t *testing.T) {
	var buf bytes.Buffer
	Logger.SetOutput(&buf)

	// Set log level to panic
	Logger.SetLevel(logrus.PanicLevel)

	// Recover from panic
	defer func() {
		r := recover()
		assert.NotNil(t, r, "Expected panic on Panic log")
		assert.Contains(t, buf.String(), "Test panic message")
	}()

	Logger.Panic("Test panic message")
}
