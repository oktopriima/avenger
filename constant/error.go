/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 05/07/21, 19:41
 * Copyright (c) 2021
 */

package constant

const (
	Success = iota
	Undefined
	ValidationError
	InvalidRequest
	ErrorQueryDB
	Forbidden
	PageNotFound
	TimeOut
	AuthorizationMissing
	FailedExtractedToken
	TokenExpired
)

var ErrorMap = map[int]string{
	Success:              "success",
	Undefined:            "undefined error",
	ValidationError:      "validation errors",
	InvalidRequest:       "invalid request",
	ErrorQueryDB:         "error while query into database",
	Forbidden:            "forbidden action",
	PageNotFound:         "404 page not found",
	TimeOut:              "request timeout",
	AuthorizationMissing: "authorization token is missing",
	FailedExtractedToken: "failed extracted token",
	TokenExpired:         "token already expired",
}
