package logger

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// New creates a new logger instance with the specified label.
// The logger will prefix all log messages with the provided label and
// include standard log flags (date and time).
func New(label string) *log.Logger {
	t := time.Now().Format("2006-01-02 15:04:05")
	label = strings.ToUpper(label)

	prefix := fmt.Sprintf("[%s] %s", label, t)
	return log.New(log.Writer(), prefix, log.LstdFlags)
}
