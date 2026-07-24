package doiitshared

import (
	"log"

	shared_error "github.com/doiit/shared/error"
	"github.com/doiit/shared/events"
	"github.com/doiit/shared/helper"
	"github.com/doiit/shared/logger"
	"github.com/doiit/shared/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
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

// GetUserContext retrieves the user context from the Gin context.
func UserContext(c *gin.Context) (helper.UserContext, error) {
	return helper.GetUserContext(c)
}

// Get environment variable
func GetEnv(key, defaultValue string) string {
	return helper.GetEnv(key, defaultValue)
}

func GetPagination(c *gin.Context) helper.Pagination {
	return helper.GetPagination(c)
}

// Send Response
func SendResponse(c *gin.Context, statusCode int, errorCode string, message string, success bool, data any) {
	helper.SendResponse(c, statusCode, errorCode, message, success, data)
}

// Send Success Response
func SendSuccessResponse(c *gin.Context, statusCode int, message string, data any) {
	SendResponse(c, statusCode, "", message, true, data)
}

// Send Error Response
func SendErrorResponse(c *gin.Context, statusCode int, errorCode, message string) {
	SendResponse(c, statusCode, errorCode, message, false, nil)
}

// Generate UUID
func UUID() uuid.UUID {
	return helper.GenerateUUID()
}

// Generate Error
func Error(message, context string) error {
	return shared_error.Error(message, context)
}

// Ensure that a user have the neccessary permissions
func RoleGuard(allowedRoles ...string) gin.HandlerFunc {
	return middleware.RoleGuard(allowedRoles...)
}

// Get IP Adrress from the request gin context
func GetIPAddress(c *gin.Context) string {
	return helper.GetIP(c)
}

// Instantiate an event producer.
// The event producer/publisher sends an event to the broker and the consumer grabs the specific topic and process it
// this allows for communication to be very easy and handles complex scenarios easily.
func NewEventProducer(brokers []string, topic string) (*events.EventProducer, error) {
	return events.NewEventProducer(brokers, topic)
}
