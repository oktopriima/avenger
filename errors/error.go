/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 05/07/21, 13:05
 * Copyright (c) 2021
 */

package errors

import (
	"fmt"
	"github.com/ztrue/tracerr"
)

const (
	Success = iota
	Undefined
	ValidationError
	InvalidRequest
	ErrorQueryDB
	Forbidden
	PageNotFound
	TimeOut
)

var ErrorMap = map[int]string{
	Success:         "success",
	Undefined:       "undefined error",
	ValidationError: "validation errors",
	InvalidRequest:  "invalid request",
	ErrorQueryDB:    "error while query into database",
	Forbidden:       "forbidden action",
	PageNotFound:    "404 page not found",
	TimeOut:         "request timeout",
}

func New(code int, err error) *Error {
	var msg string
	if err != nil {
		tracerr.PrintSourceColor(tracerr.Wrap(err), 1)
		msg = err.Error()
	}

	return &Error{
		Code:             code,
		Message:          ErrorMap[code],
		DeveloperMessage: msg,
	}
}

type Error struct {
	Code             int
	Message          string
	DeveloperMessage string
}

func (e *Error) Error() string {
	return fmt.Sprintf(" errors found with code %d. message %s", e.Code, e.Message)
}

func (e *Error) GetDeveloperMessage() string {
	return e.DeveloperMessage
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) GetCode() int {
	return e.Code
}
