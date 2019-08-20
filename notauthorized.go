package errors

// NewNotAuthorizedError returns a NotAuthorizedError instance
func NewNotAuthorizedError(message string) error {
	return wrap(&NotAuthorizedError{
		Message: message,
	}, 4)
}

// NotAuthorizedError represents an access restriction error structure
type NotAuthorizedError struct {
	Message string
}

// Error returns the string representation of the error
func (e *NotAuthorizedError) Error() string {
	return e.Message
}
