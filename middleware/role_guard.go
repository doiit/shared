package middleware

import (
	"net/http"
	"slices"

	"github.com/doiit/shared/helper"
	"github.com/gin-gonic/gin"
)

// Ensure that a user have the neccessary permissions
func RoleGuard(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user from context (set by AuthMiddleware)
		userInterface, exists := c.Get("roles")
		if !exists {
			helper.SendErrorResponse(c, http.StatusUnauthorized, "authentication_failure", "User roles not found in context")
			return
		}

		roles := userInterface.([]string)

		for _, role := range roles {
			if slices.Contains(allowedRoles, role) {
				c.Next()
				return
			}
		}

		helper.SendErrorResponse(c, http.StatusForbidden, "authorization_failure", "Insufficient permissions to access this resource")
		c.Abort()
	}
}
