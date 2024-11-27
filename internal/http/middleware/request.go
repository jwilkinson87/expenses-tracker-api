package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := uuid.GenerateUUID()
		if err != nil {
			log.Default().Println("failed to generate uuid")
			c.Next()
			return
		}

		c.Set("request_id", uuid)

		c.Next()
	}
}
