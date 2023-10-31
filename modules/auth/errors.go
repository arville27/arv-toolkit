package auth

import (
	"errors"
	"fmt"
)

var (
	InvalidCredential = errors.New("Bad credentials")
)

type AuthError struct {
	errorType error
	reason    string
	cause     error
}

func (e AuthError) ErrorType() error {
	return e.errorType
}

func (e AuthError) Reason() string {
	return e.reason
}

func (e AuthError) Cause() error {
	return e.cause
}

func NewAuthError(reason string, errorType error, cause error) *AuthError {
	return &AuthError{
		errorType: errorType,
		reason:    reason,
		cause:     cause,
	}
}

func (e AuthError) Error() string {
	if e.Cause() != nil {
		return fmt.Sprintf("AuthError: %s: %s cause: %v", e.ErrorType(), e.Reason(), e.Cause())
	}
	return fmt.Sprintf("AuthError: %s: %s", e.ErrorType(), e.Reason())
}
