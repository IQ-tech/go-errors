package errors

import (
	"strings"
)

// NewValidationError returns a ValidationError instance with the provided parameters
func NewValidationError(field string, message string) error {
	return wrap(&ValidationError{
		Field:   field,
		Message: message,
	}, 4)
}

// ValidationError represents an input validation error
type ValidationError struct {
	Field   string            `json:"field_name,omitempty"`
	Message string            `json:"message,omitempty"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

func (e *ValidationError) Error() string {
	var builder = strings.Builder{}

	if e.Field != "" {
		builder.WriteString(e.Field)
		builder.WriteString(": ")
	}

	builder.WriteString(e.Message)

	for _, err := range e.Errors {
		builder.WriteString("\n - ")
		builder.WriteString(err.Error())
		builder.WriteRune(';')
	}

	return builder.String()
}

// AddError adds a new validation error to the chain
func (e *ValidationError) AddError(field string, message string) {
	e.Errors = append(e.Errors, ValidationError{
		Field:   field,
		Message: message,
	})
}
