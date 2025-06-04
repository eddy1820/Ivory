package middleware

import (
	"gate/internal/error_code"
	"gate/internal/logger"
	"gate/pkg/response"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log panic details and stack trace
				logger.Logger.Error().
					Any("panic", r).
					Bytes("stack", debug.Stack()).
					Str("path", c.FullPath()).
					Msg("Unhandled panic recovered")

				// Return standard error response (prevent server crash)
				response.Error(c, error_code.InternalServerError)
				// Stop request processing chain
				c.Abort()
			}
		}()
		c.Next()
	}
}
