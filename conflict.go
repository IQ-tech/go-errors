package errors

// NewConflictError returns a ConflictError instance
func NewConflictError(message string) error {
	return wrap(&ConflictError{
		Message: message,
	}, 4)
}

// ConflictError represents an conflict error structure
type ConflictError struct {
	Message string
}

// Error returns the string representation of the error
func (e *ConflictError) Error() string {
	return e.Message
}
