# errors packages

This package holds helpers for better error handling, adding wrapped and contextualized errors, allowing easier error debugging and rastreability without the need to input method chains as part of error messages.

## Error types

### `ApplicationError`

A generic application error, as a failure in parametrization or other unexpected error.  
*This error usually translates to a HTTP **503 Service Unavailable** error.*

### `ConflictError`

Indicates an action is conflicting with another action, as for example duplicated requests.  
*This error usually translates to a HTTP **429 Conflict** error.*

### `ForbiddenError`

Indicates an action is not allowed, even if authenticated.  
*This error usually translates to a HTTP **403 Forbidden** error.*

### `NotAuthorizedError`

Indicates an action needs authorization or authentication to proceed.  
*This error usually translated to a HTTP **401 Unauthorized** error.*

### `ValidationError`

Indicates that a parameter provided is not in the correct format or not present if required.  
This error allows to set a property that is related to the error and also add sub validation errors to build a validation error chain.  
*This error usually translates to a HTTP **422 Unprocessable Entity** error.*

> **Note:** All error constructors return a wrapped version of the error, removing the need to always pair an error constructor with a call to `errors.Wrap`.

## Error Wrapping

This provides error bubbling tracking and other utility methods to work with wrapped errors.

### `ErrorWrapper` interface

An `ErrorWrapper` interface is also provided for identifying wrapped errors easyly.

This interface contains two methods only:

```go
// ErrorWrapper defines the interface for an error wrapper that extends an error with additional information
type ErrorWrapper interface {
    Error() string
    GetOriginalError() error
}
```

#### `Error() string`

This method makes this interface compatible with `error` interface, so types that implement it are also implementing `error`.

#### `GetOriginalError() error`

This method returns the original error that was wrapped, even if in a chain of wrapped errors, the original error that got wrapped first will be returned.

**Example:**

```go
// some error returned from another call as err
if wrappedError, ok := err.(errors.ErrorWrapper); ok {
    fmt.Println(wrappedError.GetOriginalError()) // original error message
}
```

### `Wrap(err error, messages ...string) error`

This package provides utility methods to work with wrapped errors to allow better error output when bubbling errors through many layers. Allowing to easily spot where the error actually ocurred.

**Example:**

```go
func SomeErrMethod() error {
    return errors.NewApplicationError("original error message")
}

func SomeMethod() error {
    // some error returned from another call as err
    return errors.Wrap(err)
}

func Main() {
    err := SomeMethod()

    fmt.Println(err) // main.SomeMethod ➡︎ main.SomeErrMethod ➡︎ original error message
}
```

## `GetOriginalError(err error) error`

There is a utility method to retrieve the original error from a chain of wrapped errors:

**Example:**

```go
// some error returned from another call as err, as from a call to a DB query
originalErr := errors.GetOriginalError(err)

if originalErr == sql.ErrNoRows {
    // record not found
}
```
