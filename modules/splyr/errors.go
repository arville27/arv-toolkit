package splyr

import (
	"errors"
	"fmt"
)

var (
	InvalidTrackId              = errors.New("Invalid Track ID")
	LyricsNotFound              = errors.New("Lyrics not found")
	FailedFetchLyrics           = errors.New("Failed to fetch lyrics")
	FailedRequestingAccessToken = errors.New("Failed while requesting access token")
	UnknownError                = errors.New("Unknown error")
)

type SplyrError struct {
	errorType error
	reason    string
	cause     error
}

func (e SplyrError) ErrorType() error {
	return e.errorType
}

func (e SplyrError) Reason() string {
	return e.reason
}

func (e SplyrError) Cause() error {
	return e.cause
}

func NewSplyrError(reason string, errorType error, cause error) *SplyrError {
	return &SplyrError{
		errorType: errorType,
		reason:    reason,
		cause:     cause,
	}
}

func (e SplyrError) Error() string {
	if e.Cause() != nil {
		return fmt.Sprintf("SplyrError: %s: %s cause: %v", e.ErrorType(), e.Reason(), e.Cause())
	}
	return fmt.Sprintf("SplyrError: %s: %s", e.ErrorType(), e.Reason())
}
