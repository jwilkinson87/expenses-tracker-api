package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := uuid.NewV7()
		if err != nil {
			log.Default().Println("failed to generate uuid")
			c.Next()
			return
		}

		c.Set("request_id", uuid.String())

		c.Next()
	}
}
