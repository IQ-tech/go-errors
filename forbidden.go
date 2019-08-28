package errors

// NewForbiddenError returns a ForbiddenError instance
func NewForbiddenError(message string) error {
	return wrap(&ForbiddenError{
		Message: message,
	}, 4)
}

// ForbiddenError represents an forbidden error structure
type ForbiddenError struct {
	Message string
}

// Error returns the string representation of the error
func (e *ForbiddenError) Error() string {
	return e.Message
}
