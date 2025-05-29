package response

import (
	"gate/pkg/error_code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// APIResponse defines the standard response structure.
type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// Success sends a successful response with data.
func Success[T any](ctx *gin.Context, data T) {
	ctx.JSON(http.StatusOK, APIResponse[T]{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// SuccessNoData sends a success response without data.
func SuccessNoData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, APIResponse[any]{
		Code:    0,
		Message: "success",
	})
}

// Error sends an error response with code and message.
func Error(ctx *gin.Context, data *error_code.ErrorData) {
	ctx.JSON(data.HttpStatus, APIResponse[any]{
		Code:    data.Code,
		Message: data.Message,
	})
}
