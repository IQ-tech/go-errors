// Package errors has utilities to create and extend errors, easing the error handling process
package errors

import (
	stderrors "errors"
	"runtime"
	"strings"
)

func New(message string) error {
	return stderrors.New(message)
}

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
func getCallerFunction(skip int) string {
	// get caller function path
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	path := strings.Split(frame.Function, "/")

	return path[len(path)-1]
}

// Returns true when two errors at the same.
//
// err := errors.New("oops")
// errors.Is(err, err) => true
//
// err2 = errors.New("oops")
// errors.Is(err, err2) => false
func Is(a, b error) bool {
	return GetOriginalError(a) == GetOriginalError((b))
}

// Returns true when two errors have the same error message.
//
// err := errors.New("oops")
// errors.Equals(err, err) => true
//
// err2 := errors.New("oops...")
// errors.Equals(err, err2) => false
func Equals(a, b error) bool {
	if a == nil || b == nil {
		return Is(a, b)
	}

	return GetOriginalError(a).Error() == GetOriginalError(b).Error()
}
