package errors

// NewNotFoundError returns a NewNotFoundError instance
func NewNotFoundError(message string) error {
	return wrap(&NotFoundError{
		Message: message,
	}, 4)
}

// NotFoundError represents an not found error structure
type NotFoundError struct {
	Message string
}

// Error returns the string representation of the error
func (e *NotFoundError) Error() string {
	return e.Message
}
