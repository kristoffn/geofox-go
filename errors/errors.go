package errors

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

var _ error = (*GeofoxError)(nil)
var _ error = (*RequestError)(nil)
var _ error = (*AuthenticationError)(nil)
var _ error = (*RateLimitError)(nil)
var _ error = (*ServerError)(nil)
var _ error = (*ServiceError)(nil)

type GeofoxError struct {
	StatusCode   int
	ReturnCode   GeofoxErrorType
	ErrorText    string
	ErrorDevInfo string
}

func (e *GeofoxError) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s, message: %s, dev information: %s",
		e.StatusCode, e.ReturnCode, e.ErrorText, e.ErrorDevInfo)
}

type RequestError struct {
	GeofoxError
}

func (e *RequestError) Error() string {
	return e.Error()
}

type AuthenticationError struct {
	GeofoxError
}

func (e *AuthenticationError) Error() string {
	return e.Error()
}

type AuthorizationError struct {
	GeofoxError
}

func (e *AuthorizationError) Error() string {
	return e.Error()
}

type RateLimitError struct {
	GeofoxError
}

func (e *RateLimitError) Error() string {
	return e.Error()
}

type ServerError struct {
	GeofoxError
}

func (e *ServerError) Error() string {
	return e.Error()
}

type ServiceError struct {
	GeofoxError
}

func (e *ServiceError) Error() string {
	return e.Error()
}
