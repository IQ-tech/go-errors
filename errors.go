// Package errors has utilities to create and extend errors, easing the error handling process
package errors

import (
	"runtime"
	"strings"
)

// Wrap wraps an error with a context message and adds execution path
func Wrap(err error, messages ...string) error {
	return wrap(err, 4, messages...)
}

// Wrap wraps an error with a context message and adds execution path
func wrap(err error, skipStack int, messages ...string) error {
	if err != nil {
		return &wrappedError{
			originalError: err,
			path:          getCallerFunction(skipStack),
			messages:      messages,
		}
	}

	return nil
}

// GetOriginalError returns the original error if the provided error is a WrappedError.
// Returns the provided error otherwise
func GetOriginalError(err error) error {
	wrappedErr, ok := err.(ErrorWrapper)
	if ok {
		return wrappedErr.GetOriginalError()
	}

	return err
}

// getCallerFunction returns the name of a method in the method chain
// indicated by the stackOrder index from last to first
func getCallerFunction(stackOrder int) string {
	// get caller function path
	pc := make([]uintptr, 1)
	runtime.Callers(stackOrder, pc)

	funcRef := runtime.FuncForPC(pc[0])

	funcPath := strings.Split(funcRef.Name(), "/")

	return funcPath[len(funcPath)-1]
}
