package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := uuid.NewV7()
		if err != nil {
			slog.Error("failed to generate request id", "error", err.Error())
			c.Next()
			return
		}

		c.Set("request_id", uuid.String())

		c.Next()
	}
}
