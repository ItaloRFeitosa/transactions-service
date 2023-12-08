package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	c.Next()

	if raw != "" {
		path = path + "?" + raw
	}

	timestamp := time.Now()
	latency := timestamp.Sub(start)
	clientIP := c.ClientIP()
	method := c.Request.Method
	statusCode := c.Writer.Status()
	bodySize := c.Writer.Size()

	slog.InfoContext(c.Request.Context(), "request finished",
		"latency", latency,
		"client_ip", clientIP,
		"method", method,
		"status_code", statusCode,
		"body_size", bodySize,
		"path", path,
	)
}
