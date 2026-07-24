package helper

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
)

// UserContext represents the context of a user in the application.
type UserContext struct {
	SessionID uuid.UUID `json:"session_id"`
	UserID    uuid.UUID `json:"user_id"`
	DeviceID  uuid.UUID `json:"device_id"`
	IPAddress string    `json:"ip_address"`
	Roles     []string  `json:"roles"`
}

// GetUserContext retrieves the user context from the Gin context.
func GetUserContext(c *gin.Context) (UserContext, error) {
	ip := GetIP(c)
	var deviceId uuid.UUID
	var userId uuid.UUID
	var sessionId uuid.UUID

	var roles, exists = c.Get("roles")
	if !exists {
		return UserContext{}, errors.New("roles not found in context")
	}

	userID, exists := c.Get("userID")
	if !exists {
		return UserContext{}, errors.New("userID not found in context")
	}
	deviceID, exists := c.Get("deviceID")
	if !exists {
		return UserContext{}, errors.New("deviceID not found in context")
	}
	sessionID, exists := c.Get("sessionID")
	if !exists {
		return UserContext{}, errors.New("sessionID not found in context")
	}

	// casting userID type
	if id, ok := userID.(uuid.UUID); !ok {

		return UserContext{}, errors.New("Could not convert userID to UUID type")

	} else {
		userId = id
	}

	// casting deviceID type
	if id, ok := deviceID.(uuid.UUID); !ok {

		return UserContext{}, errors.New("Could not convert deviceID to UUID type")

	} else {
		deviceId = id
	}

	// casting sessionID type
	if id, ok := sessionID.(uuid.UUID); !ok {

		return UserContext{}, errors.New("Could not convert sessionID to UUID type")

	} else {
		sessionId = id
	}

	return UserContext{
		SessionID: sessionId,
		DeviceID:  deviceId,
		UserID:    userId,
		Roles:     roles.([]string),
		IPAddress: ip,
	}, nil

}

// GetContextWithTimeout creates a new context with a specified timeout duration in seconds.
// This function is useful for operations that need to be completed within a certain time frame, allowing for cancellation if the operation takes too long.
func GetContextWithTimeout(duration int64) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
}

// Get IP address from the gin context
func GetIP(c *gin.Context) string {

	removePort := func(ip string) string {
		ip = strings.TrimSpace(ip)
		if strings.Contains(ip, ":") {
			parts := strings.Split(ip, ":")
			return parts[0]
		}

		return ip
	}

	// Check X-Forwarded-For header first (for proxies/load balancers)
	forwarded := c.GetHeader("X-Forwarded-For")
	if forwarded != "" {
		// X-Forwarded-For can contain multiple IPs, take the first one
		ips := strings.Split(forwarded, ",")
		if len(ips) > 0 {
			return removePort(ips[0])
		}
	}

	// Check X-Real-IP header
	realIP := c.GetHeader("X-Real-IP")
	if realIP != "" {
		return removePort(realIP)
	}

	// Fallback to Gin's ClientIP() method
	ip := c.ClientIP()
	return removePort(ip)
}
