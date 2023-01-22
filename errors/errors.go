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

var _ GeofoxError = (*ErrorUnauthorized)(nil)
var _ GeofoxError = (*ErrorForbidden)(nil)
var _ GeofoxError = (*ErrorNotFound)(nil)
var _ GeofoxError = (*ErrorTooManyRequests)(nil)
var _ GeofoxError = (*ErrorInternalServerError)(nil)
var _ GeofoxError = (*ErrorGeneric)(nil)

type GeofoxError interface {
	Error() string
}

type ErrorUnauthorized struct {
	StatusCode int
	ReturnCode GeofoxErrorType
}

func (e *ErrorUnauthorized) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorForbidden struct {
	StatusCode int
	ReturnCode GeofoxErrorType
}

func (e *ErrorForbidden) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorNotFound struct {
	StatusCode int
	ReturnCode GeofoxErrorType
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorTooManyRequests struct {
	StatusCode int
	ReturnCode GeofoxErrorType
}

func (e *ErrorTooManyRequests) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorInternalServerError struct {
	StatusCode int
	ReturnCode GeofoxErrorType
}

func (e *ErrorInternalServerError) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorGeneric struct {
	StatusCode int
	ReturnCode GeofoxErrorType
}

func (e *ErrorGeneric) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}
