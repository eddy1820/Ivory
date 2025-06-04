package middleware

import (
	"time"

	"gate/internal/logger"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path // fallback
		}
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		userAgent := c.Request.UserAgent()
		requestID := c.GetString("request_id")

		logEvent := logger.Logger.Info()
		if len(c.Errors) > 0 {
			logEvent = logger.Logger.Error().Err(c.Errors.Last())
		}

		logEvent.
			Str("method", method).
			Str("path", path).
			Str("request_id", requestID).
			Int("status", status).
			Dur("latency", latency).
			Str("ip", clientIP).
			Str("user-agent", userAgent)

		if raw != "" {
			logEvent.Str("query", raw)
		}

		logEvent.Msg("HTTP request log")
	}
}
