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
	StatusCode() int
	ReturnCode() GeofoxReturnCode
	Message() string
}

type baseError struct {
	statusCode int
	returnCode GeofoxReturnCode
	message    string
}

func (e *baseError) StatusCode() int {
	return e.statusCode
}

func (e *baseError) ReturnCode() GeofoxReturnCode {
	return e.returnCode
}

func (e *baseError) Message() string {
	return e.message
}

type ErrorUnauthorized struct {
	baseError
}

func (e *ErrorUnauthorized) Error() string {
	return fmt.Sprintf("unauthorized: status code %d, return code %s, message: %s",
		e.statusCode, e.returnCode, e.message)
}

type ErrorForbidden struct {
	baseError
}

func (e *ErrorForbidden) Error() string {
	return fmt.Sprintf("forbidden: status code %d, return code %s, message: %s",
		e.statusCode, e.returnCode, e.message)
}

type ErrorNotFound struct {
	baseError
}

func (e *ErrorNotFound) Error() string {
	return fmt.Sprintf("not found: status code %d, return code %s, message: %s",
		e.statusCode, e.returnCode, e.message)
}

type ErrorTooManyRequests struct {
	baseError
}

func (e *ErrorTooManyRequests) Error() string {
	return fmt.Sprintf("too many requests: status code %d, return code %s, message: %s",
		e.statusCode, e.returnCode, e.message)
}

type ErrorInternalServerError struct {
	baseError
}

func (e *ErrorInternalServerError) Error() string {
	return fmt.Sprintf("internal server error: status code %d, return code %s, message: %s",
		e.statusCode, e.returnCode, e.message)
}

type ErrorGeneric struct {
	baseError
}

func (e *ErrorGeneric) Error() string {
	return fmt.Sprintf("generic error: status code %d, return code %s, message: %s",
		e.statusCode, e.returnCode, e.message)
}
