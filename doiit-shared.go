package doiitshared

import (
	"log"

	"github.com/doiit/shared/logger"
)

// This package is a shared library for the DOIIT project. It contains common utilities and helpers that can be used across different services in the DOIIT ecosystem.
func GetAppName() string {
	return "doiit"
}

// Logger creates a new logger instance with the specified label.
// The logger will prefix all log messages with the provided label and include standard log flags (date and time).
func Logger(label string) *log.Logger {
	return logger.New(label)
}
