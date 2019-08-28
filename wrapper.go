package errors

import (
	"strings"
)

// ErrorWrapper defines the interface for an error wrapper that extends an error with additional information
type ErrorWrapper interface {
	Error() string
	GetOriginalError() error
}

// wrappedError holds an error wrapped with a context message
type wrappedError struct {
	originalError error
	path          string
	messages      []string
}

// Error returns the string representation of the error
func (err *wrappedError) Error() string {
	builder := strings.Builder{}

	builder.WriteString(err.path)

	if len(err.messages) > 0 {
		builder.WriteString(": ")

		for _, message := range err.messages {
			builder.WriteString(message)
			builder.WriteString("; ")
		}
	}

	builder.WriteString(" ➡︎ ")
	builder.WriteString(err.originalError.Error())

	return builder.String()
}

// GetOriginalError returns the original error
func (err *wrappedError) GetOriginalError() error {
	if err.originalError != nil {
		originalError, ok := (err.originalError).(ErrorWrapper)
		if ok {
			return originalError.GetOriginalError()
		}
	}

	return err.originalError
}
