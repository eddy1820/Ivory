package error_code

import "net/http"

var (
	Success                   = NewErrorData(http.StatusOK, 0, "Success")
	InternalServerError       = NewErrorData(http.StatusInternalServerError, 10000000, "Internal server error")
	InvalidParams             = NewErrorData(http.StatusBadRequest, 10000001, "Invalid params")
	NotFound                  = NewErrorData(http.StatusNotFound, 10000002, "Not found")
	UnauthorizedAuthNotExist  = NewErrorData(http.StatusUnauthorized, 10000003, "Unauthorized token error")
	UnauthorizedTokenError    = NewErrorData(http.StatusUnauthorized, 10000004, "Unauthorized token error")
	UnauthorizedTokenTimeout  = NewErrorData(http.StatusUnauthorized, 10000005, "Unauthorized token error")
	UnauthorizedTokenGenerate = NewErrorData(http.StatusUnauthorized, 10000006, "Unauthorized token error")
	TooManyRequests           = NewErrorData(http.StatusTooManyRequests, 10000007, "Too many requests")
)
