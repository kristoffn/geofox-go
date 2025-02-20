package geofox

import (
	"fmt"
)

type GeofoxReturnCode string

const (
	RCOK                GeofoxReturnCode = "OK"
	RCCNTooMany         GeofoxReturnCode = "ERROR_CN_TOO_MANY"
	RCCommError         GeofoxReturnCode = "ERROR_COMM"
	RCRouteError        GeofoxReturnCode = "ERROR_ROUTE"
	RCErrorText         GeofoxReturnCode = "ERROR_TEXT"
	RCStartDestTooClose GeofoxReturnCode = "START_DEST_TOO_CLOSE"
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
	ReturnCode GeofoxReturnCode
}

func (e *ErrorUnauthorized) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorForbidden struct {
	StatusCode int
	ReturnCode GeofoxReturnCode
}

func (e *ErrorForbidden) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorNotFound struct {
	StatusCode int
	ReturnCode GeofoxReturnCode
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorTooManyRequests struct {
	StatusCode int
	ReturnCode GeofoxReturnCode
}

func (e *ErrorTooManyRequests) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorInternalServerError struct {
	StatusCode int
	ReturnCode GeofoxReturnCode
}

func (e *ErrorInternalServerError) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}

type ErrorGeneric struct {
	StatusCode int
	ReturnCode GeofoxReturnCode
}

func (e *ErrorGeneric) Error() string {
	return fmt.Sprintf("status code: %d, error type: %s",
		e.StatusCode, e.ReturnCode)
}
