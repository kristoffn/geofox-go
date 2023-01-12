package geofox

import (
	"fmt"
)

type GeofoxErrorType string

const (
	ErrorCNTooMany     GeofoxErrorType = "ERROR_CN_TOO_MANY"
	ErrorCommunication GeofoxErrorType = "ERROR_COMM"
	ErrorRoute         GeofoxErrorType = "ERROR_ROUTE"
	ErrorText          GeofoxErrorType = "ERROR_TEXT"
)

var _ error = (*Error)(nil)
var _ error = (*RequestError)(nil)
var _ error = (*AuthenticationError)(nil)
var _ error = (*RateLimitError)(nil)
var _ error = (*ServerError)(nil)
var _ error = (*ServiceError)(nil)

type Error struct {
	StatusCode   int
	ReturnCode   GeofoxErrorType
	ErrorText    string
	ErrorDevInfo string
}

func (e *Error) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s, message: %s, dev information: %s",
		e.StatusCode, e.ReturnCode, e.ErrorText, e.ErrorDevInfo)
}

type RequestError struct {
	geofoxError *Error
}

func (e *RequestError) Error() string {
	return e.geofoxError.Error()
}

type AuthenticationError struct {
	geofoxError *Error
}

func (e *AuthenticationError) Error() string {
	return e.geofoxError.Error()
}

type AuthorizationError struct {
	geofoxError *Error
}

func (e *AuthorizationError) Error() string {
	return e.geofoxError.Error()
}

type RateLimitError struct {
	geofoxError *Error
}

func (e *RateLimitError) Error() string {
	return e.geofoxError.Error()
}

type ServerError struct {
	geofoxError *Error
}

func (e *ServerError) Error() string {
	return e.geofoxError.Error()
}

type ServiceError struct {
	geofoxError *Error
}

func (e *ServiceError) Error() string {
	return e.geofoxError.Error()
}
