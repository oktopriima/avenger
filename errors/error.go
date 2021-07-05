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
	"github.com/oktopriima/avenger/constant"
	"github.com/ztrue/tracerr"
)

func New(code int, err error) *Error {
	var msg string
	if err != nil {
		// trace or put your custom log here
		tracerr.PrintSourceColor(tracerr.Wrap(err), 1)
		msg = err.Error()
	}

	return &Error{
		Code:             code,
		Message:          constant.ErrorMap[code],
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
