package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	errors2 "github.com/oktopriima/avenger/errors"
	httpresponse2 "github.com/oktopriima/avenger/httpresponse"
	"io/ioutil"
	"net/http"
)

func main() {
	e := echo.New()

	echo.NotFoundHandler = func(c echo.Context) error {
		return httpresponse2.PageNotFound(c)
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

	e.GET("/", a)
	e.GET("/ping", b)

	e.Logger.Fatal(e.Start(":5000"))
}

func a(ctx echo.Context) error {
	return httpresponse2.JSONErr(ctx, errors2.New(errors2.ErrorQueryDB, fmt.Errorf("test developer message")))
}

func b(ctx echo.Context) error {
	return httpresponse2.JSONSuccess(ctx, "pong")
}

func testError() error {
	r, err := http.Get("http://localhost:1323/not-found")
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
