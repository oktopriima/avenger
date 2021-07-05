/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 05/07/21, 13:22
 * Copyright (c) 2021
 */

package httpresponse

import (
	"github.com/labstack/echo/v4"
	"github.com/oktopriima/avenger/errors"
	"net/http"
	"os"
)

var httpStatusCode = map[int]int{
	errors.Undefined:       http.StatusBadRequest,
	errors.ValidationError: http.StatusBadRequest,
	errors.ErrorQueryDB:    http.StatusUnprocessableEntity,
	errors.PageNotFound:    http.StatusNotFound,
	errors.Forbidden:       http.StatusForbidden,
	errors.InvalidRequest:  http.StatusBadRequest,
	errors.TimeOut:         http.StatusRequestTimeout,
}

type errorStruct struct {
	Code             int    `json:"code"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message,omitempty"`
}

type successStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONErr(ctx echo.Context, err error) error {
	env := os.Getenv("OS_ENV")
	var s errorStruct

	switch e := err.(type) {
	case *errors.Error:
		if env != "production" {
			s.DeveloperMessage = e.GetDeveloperMessage()
		}

		s.Code = e.GetCode()
		s.Message = e.GetMessage()
		return ctx.JSON(httpStatusCode[e.Code], s)
	default:
		if env != "production" {
			s.DeveloperMessage = err.Error()
		}

		s.Code = errors.Undefined
		s.Message = errors.ErrorMap[errors.Undefined]

		return ctx.JSON(httpStatusCode[errors.Undefined], s)
	}
}

func JSONSuccess(c echo.Context, data interface{}) error {
	res := successStruct{
		Code:    errors.Success,
		Message: errors.ErrorMap[errors.Success],
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

func PageNotFound(ctx echo.Context) error {
	return ctx.String(http.StatusNotFound, errors.ErrorMap[errors.PageNotFound])
}
