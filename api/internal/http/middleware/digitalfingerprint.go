package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GenerateDigitalFingerprint() gin.HandlerFunc {
	return func(c *gin.Context) {
		digitalFingerprint := fmt.Sprintf("%s:%s", c.GetHeader("User-Agent"), c.ClientIP())
		c.Set("digital_fingerprint", digitalFingerprint)

		c.Next()
	}
}
