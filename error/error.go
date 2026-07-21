package shared_error

import (
	"errors"
	"fmt"
)

type errorType string
type messageType string

const (
	InvalidTokenMessage    messageType = "Invalid token. Kindly login again"
	ExpireTokenMessage     messageType = "Token expired. Login again"
	InvalidClaimsMessage   messageType = "Invalid token claims"
	AccessDeniedMessage    messageType = "Access denied"
	SessionExpiredMessage  messageType = "Session has expired"
	SessionRevokedMessage  messageType = "Session was revoked"
	ValidationErrorMessage messageType = "Validation error"
	UnauthorizedMessage    messageType = "Unauthorized access"
)

const (
	AccessDenied          errorType = "invalid_token"
	BindingError          errorType = "binding_error"
	BadRequestError       errorType = "bad_request"
	ExpireToken           errorType = "token_expired"
	RefreshToken          errorType = "refresh_token_error"
	HashingError          errorType = "hashing_error"
	InvalidToken          errorType = "invalid_token"
	InvalidClaims         errorType = "invalid_token"
	InvalidCredentials    errorType = "invalid_credentials"
	SessionExpired        errorType = "session_error"
	SessionRevoked        errorType = "session_revoked"
	ValidationError       errorType = "validation_error"
	DuplicateError        errorType = "duplicate_error"
	ResourceCreationError errorType = "resource_creation_error"
	ResourceUpdateError   errorType = "resource_update_error"
	ResourceDeletionError errorType = "resource_deletion_error"
	NotFound              errorType = "not_found"
	Unauthorized          errorType = "unauthorized"
)

// Error returns a new error with the given message.
func Error(message, context string) error {
	if context == "" {
		return errors.New(message)
	}

	err := fmt.Sprintf("[%s]: %s", context, message)
	return errors.New(err)
}
