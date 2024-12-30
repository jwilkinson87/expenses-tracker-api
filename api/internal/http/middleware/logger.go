package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Debug("request received", "method", c.Request.Method, "path", c.Request.URL.Path, "request_id", c.GetString("request_id"), "status_code", c.Writer.Status())
		c.Next()
	}
}
