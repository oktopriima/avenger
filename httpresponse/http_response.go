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
	errors2 "github.com/oktopriima/avenger/errors"
	"net/http"
	"os"
)

var httpStatusCode = map[int]int{
	errors2.Undefined:       http.StatusBadRequest,
	errors2.ValidationError: http.StatusBadRequest,
	errors2.ErrorQueryDB:    http.StatusUnprocessableEntity,
	errors2.PageNotFound:    http.StatusNotFound,
	errors2.Forbidden:       http.StatusForbidden,
	errors2.InvalidRequest:  http.StatusBadRequest,
	errors2.TimeOut:         http.StatusRequestTimeout,
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
	case *errors2.Error:
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

		s.Code = errors2.Undefined
		s.Message = errors2.ErrorMap[errors2.Undefined]

		return ctx.JSON(httpStatusCode[errors2.Undefined], s)
	}
}

func JSONSuccess(c echo.Context, data interface{}) error {
	res := successStruct{
		Code:    errors2.Success,
		Message: errors2.ErrorMap[errors2.Success],
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

func PageNotFound(ctx echo.Context) error {
	return ctx.String(http.StatusNotFound, errors2.ErrorMap[errors2.PageNotFound])
}
