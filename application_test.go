package errors

import (
	"testing"
)

func TestNewApplicationError(t *testing.T) {
	err := NewApplicationError("test app error")

	_, ok := GetOriginalError(err).(*ApplicationError)
	if !ok {
		t.Errorf("NewApplicationError() returned an incorrect error type %t", err)
	}

	errMsg := err.Error()
	expectedErrMsg := "go-errors.TestNewApplicationError ➡︎ test app error"

	if errMsg != expectedErrMsg {
		t.Errorf("NewApplicationError() returned wrong error message: %s", errMsg)
	}
}
