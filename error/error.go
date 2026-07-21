package shared_error

import (
	"errors"
	"fmt"
)

const (
	InvalidTokenMessage    string = "Invalid token. Kindly login again"
	ExpireTokenMessage     string = "Token expired. Login again"
	InvalidClaimsMessage   string = "Invalid token claims"
	AccessDeniedMessage    string = "Access denied"
	SessionExpiredMessage  string = "Session has expired"
	SessionRevokedMessage  string = "Session was revoked"
	ValidationErrorMessage string = "Validation error"
	UnauthorizedMessage    string = "Unauthorized access"
)

const (
	AccessDenied          string = "invalid_token"
	BindingError          string = "binding_error"
	BadRequestError       string = "bad_request"
	ExpireToken           string = "token_expired"
	RefreshToken          string = "refresh_token_error"
	HashingError          string = "hashing_error"
	InvalidToken          string = "invalid_token"
	InvalidClaims         string = "invalid_token"
	InvalidCredentials    string = "invalid_credentials"
	SessionExpired        string = "session_error"
	SessionRevoked        string = "session_revoked"
	ValidationError       string = "validation_error"
	DuplicateError        string = "duplicate_error"
	ResourceCreationError string = "resource_creation_error"
	ResourceUpdateError   string = "resource_update_error"
	ResourceDeletionError string = "resource_deletion_error"
	NotFound              string = "not_found"
	Unauthorized          string = "unauthorized"
)

// Error returns a new error with the given message.
func Error(message, context string) error {
	if context == "" {
		return errors.New(message)
	}

	err := fmt.Sprintf("[%s]: %s", context, message)
	return errors.New(err)
}
