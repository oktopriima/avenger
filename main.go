package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oktopriima/avenger/constant"
	"github.com/oktopriima/avenger/errors"
	"github.com/oktopriima/avenger/httpresponse"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	echo.NotFoundHandler = func(c echo.Context) error {
		return httpresponse.PageNotFound(c)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "TIME = ${time_rfc3339} | METHOD = ${method} | URL = ${host}${uri} | STATUS = ${status}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
			http.MethodOptions,
			http.MethodPatch,
		},
	}))

	e.GET("/error", a)
	e.GET("/success", b)

	e.Logger.Fatal(e.Start(":5000"))
}

func a(ctx echo.Context) error {
	// connect to your controller or usecase
	// resp, err := callUsecaseFunction()
	// return httpresponse.JSONErr(ctx, err)

	resp, err := usecaseExampleError(ctx.Request().Context())
	if err != nil {
		return httpresponse.JSONErr(ctx, err)
	}

	return httpresponse.JSONSuccess(ctx, resp)
}

func usecaseExampleError(ctx context.Context) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	err := errors.New(constant.TimeOut, nil)
	return nil, err
}

func b(ctx echo.Context) error {
	// connect to your controller or usecase
	// resp, err := callUsecaseFunction()
	// return httpresponse.JSONErr(ctx, err)

	resp, err := usecaseExampleSuccess(ctx.Request().Context())
	if err != nil {
		return httpresponse.JSONErr(ctx, err)
	}

	return httpresponse.JSONSuccess(ctx, resp)
}

func usecaseExampleSuccess(ctx context.Context) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	return struct {
		User string `json:"user"`
	}{User: "pong"}, nil
}

func testError() error {
	r, err := http.Get("http://localhost:1323/success")
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %v", err)
	}

	fmt.Println(string(body))

	return nil
}
