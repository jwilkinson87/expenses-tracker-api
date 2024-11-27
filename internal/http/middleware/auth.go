package middleware

import (
	"net/http"

	"example.com/expenses-tracker/internal/handlers"
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
			return
		}

		a.handler.ValidateToken(c, token)

		c.Next()
	}
}
