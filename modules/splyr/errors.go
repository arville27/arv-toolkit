package splyr

import "fmt"

type SplyrError struct {
	Reason string
	Cause  error
}

func (e SplyrError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("SplyrError: %s: %v", e.Reason, e.Cause)
	}
	return fmt.Sprintf("SplyrError: %s", e.Reason)
}
