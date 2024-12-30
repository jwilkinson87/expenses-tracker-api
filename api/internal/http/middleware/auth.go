package middleware

import (
	"net/http"

	"example.com/expenses-tracker/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	handler *handlers.AuthHandler
}

func NewAuthMiddleware(handler *handlers.AuthHandler) *authMiddleware {
	return &authMiddleware{handler}
}

func (a *authMiddleware) HandleAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("authorization")
		if token == "" || len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		isValid, err := a.handler.ValidateToken(c, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to validate token"})
			c.Abort()
			return
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_token", token)
		c.Next()
	}
}
