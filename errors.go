package geofox

import "fmt"

const (
	errUnmarshalError = "error unmarshalling the JSON response"
)

type Error struct {
	StatusCode int
	Message    string
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

var _ error = (*RequestError)(nil)
var _ error = (*RatelimitError)(nil)
var _ error = (*ServiceError)(nil)
var _ error = (*NotFoundError)(nil)
var _ error = (*AuthorizationError)(nil)

type RequestError struct {
	err *Error
}

func (e RequestError) Error() string {
	return e.err.Error()
}

type RatelimitError struct {
	err *Error
}

func (e RatelimitError) Error() string {
	return e.err.Error()
}

type ServiceError struct {
	err *Error
}

func (e ServiceError) Error() string {
	return e.err.Error()
}

type NotFoundError struct {
	err *Error
}

func (e NotFoundError) Error() string {
	return e.err.Error()
}

type AuthorizationError struct {
	err *Error
}

func (e AuthorizationError) Error() string {
	return e.err.Error()
}
