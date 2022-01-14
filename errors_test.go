// Package errors has utilities to create and extend errors, easing the error handling process
package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func wrap1() error {
	return Wrap(errors.New("some error"))
}

func wrap2() error {
	return Wrap(wrap1())
}

func TestWrap(t *testing.T) {
	assert.Equal(t, "go-errors.TestWrap ➡︎ some error", Wrap(errors.New("some error")).Error())
	assert.Equal(t, "go-errors.wrap1 ➡︎ some error", wrap1().Error())
	assert.Equal(t, "go-errors.wrap2 ➡︎ go-errors.wrap1 ➡︎ some error", wrap2().Error())
}

func TestIs(t *testing.T) {
	t.Parallel()

	errOne := errors.New("err 1")
	errTwo := errors.New("err 2")

	tests := []struct {
		lhs      error
		rhs      error
		expected bool
	}{
		{
			lhs:      nil,
			rhs:      NewValidationError("a", "b"),
			expected: false,
		},
		{
			lhs:      NewValidationError("a", "b"),
			rhs:      nil,
			expected: false,
		},
		{
			lhs:      nil,
			rhs:      Wrap(NewValidationError("a", "b")),
			expected: false,
		},
		{
			lhs:      Wrap(NewValidationError("a", "b")),
			rhs:      nil,
			expected: false,
		},
		{
			lhs:      NewValidationError("a", "b"),
			rhs:      NewValidationError("b", "c"),
			expected: false,
		},
		{
			lhs:      NewValidationError("a", "b"),
			rhs:      NewValidationError("a", "b"),
			expected: false,
		},
		{
			lhs:      errOne,
			rhs:      errTwo,
			expected: false,
		},
		{
			lhs:      errOne,
			rhs:      errOne,
			expected: true,
		},
		{
			lhs:      Wrap(errOne),
			rhs:      errOne,
			expected: true,
		},
		{
			lhs:      Wrap(errOne),
			rhs:      Wrap(errOne),
			expected: true,
		},
		{
			lhs:      Wrap(Wrap(Wrap(errTwo))),
			rhs:      Wrap(errTwo),
			expected: true,
		},
		{
			lhs:      Wrap(Wrap(Wrap(errOne))),
			rhs:      Wrap(errTwo),
			expected: false,
		},
	}

	for _, tt := range tests {
		actual := Is(tt.lhs, tt.rhs)

		assert.Equal(t, tt.expected, actual)
	}
}

func TestEquals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lhs      error
		rhs      error
		expected bool
	}{
		{
			lhs:      nil,
			rhs:      nil,
			expected: true,
		},
		{
			lhs:      nil,
			rhs:      nil,
			expected: true,
		},
		{
			lhs:      NewValidationError("card_id", "Cartão deve ser informado!"),
			rhs:      nil,
			expected: false,
		},
		{
			lhs:      nil,
			rhs:      NewValidationError("card_id", "Cartão deve ser informado!"),
			expected: false,
		},
		{
			lhs:      NewValidationError("card_id", "Cartão deve ser informado!"),
			rhs:      NewValidationError("card_id", "Cartão deve ser informado!"),
			expected: true,
		},
		{
			lhs:      NewValidationError("card_id", "Cartão deve ser informado!"),
			rhs:      NewValidationError("card_id", "Cartão deve ser informado!"),
			expected: true,
		},
	}

	for _, tt := range tests {
		actual := Equals(tt.lhs, tt.rhs)

		assert.Equal(t, tt.expected, actual)
	}
}

func TestWrapf(t *testing.T) {
	t.Parallel()

	t.Run("when error is nil", func(t *testing.T) {
		t.Run("returns nil", func(t *testing.T) {
			actual := Wrapf(nil, "%s", "test")
			assert.Equal(t, nil, actual)
		})
	})

	t.Run("when error is not nil", func(t *testing.T) {
		t.Run("adds formatted string to wrapped error messages", func(t *testing.T) {
			actual := Wrapf(New("oops"), "%s %d", "test", 1)
			assert.Equal(t, "go-errors.TestWrapf.func2.1: test 1;  ➡︎ oops", actual.Error())
		})
	})
}
