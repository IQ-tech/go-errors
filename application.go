package errors

// NewApplicationError returns a ApplicationError instance
func NewApplicationError(message string) error {
	return wrap(&ApplicationError{
		Message: message,
	}, 4)
}

// ApplicationError represents a common applicatino error structure
type ApplicationError struct {
	Message string
}

// Error returns the string representation of the error
func (e *ApplicationError) Error() string {
	return e.Message
}
