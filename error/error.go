package error

import (
	"errors"
	"fmt"
)

const (
	InvalidTokenMessage    = "Invalid token. Kindly login again"
	ExpireTokenMessage     = "Token expired. Login again"
	InvalidClaimsMessage   = "Invalid token claims"
	AccessDeniedMessage    = "Access denied"
	SessionExpiredMessage  = "Session has expired"
	SessionRevokedMessage  = "Session was revoked"
	ValidationErrorMessage = "Validation error"
	UnauthorizedMessage    = "Unauthorized access"
)

const (
	AccessDenied          = "invalid_token"
	BindingError          = "binding_error"
	BadRequestError       = "bad_request"
	ExpireToken           = "token_expired"
	RefreshToken          = "refresh_token_error"
	HashingError          = "hashing_error"
	InvalidToken          = "invalid_token"
	InvalidClaims         = "invalid_token"
	InvalidCredentials    = "invalid_credentials"
	SessionExpired        = "session_error"
	SessionRevoked        = "session_revoked"
	ValidationError       = "validation_error"
	DuplicateError        = "duplicate_error"
	ResourceCreationError = "resource_creation_error"
	ResourceUpdateError   = "resource_update_error"
	ResourceDeletionError = "resource_deletion_error"
	NotFound              = "not_found"
	Unauthorized          = "unauthorized"
)

// Error returns a new error with the given message.
func Error(message, context string) error {
	if context == "" {
		return errors.New(message)
	}

	err := fmt.Sprintf("[%s]: %s", context, message)
	return errors.New(err)
}
