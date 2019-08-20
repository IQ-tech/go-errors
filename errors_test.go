// Package errors has utilities to create and extend errors, easing the error handling process
package errors

import (
	"errors"
	"testing"
)

func wrap1() error {
	return Wrap(errors.New("some error"))
}

func wrap2() error {
	return Wrap(wrap1())
}

func TestWrap(t *testing.T) {
	err := wrap2()

	errMsg := err.Error()
	expectedErrMsg := "go-errors.wrap2 ➡︎ go-errors.wrap1 ➡︎ some error"

	if errMsg != expectedErrMsg {
		t.Errorf("Wrap() returned wrong error message: %s", errMsg)
	}
}
