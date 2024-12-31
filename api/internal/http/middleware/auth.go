package middleware

import (
	"errors"
	"net/http"
	"strings"

	"example.com/expenses-tracker/api/internal/handlers"
	"github.com/davecgh/go-spew/spew"
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
		authHeader := c.Request.Header.Get("authorization")
		token, found := strings.CutPrefix(authHeader, "Bearer ")
		if !found {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		if token == "" || len(token) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		sessionId, err := a.handler.GetSessionIdFromToken(c, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": getErrorMessage(err)})
			c.Abort()
			return
		}

		if sessionId == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		session, err := a.handler.GetBySessionID(c, *sessionId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to verify token"})
			c.Abort()
			return
		}

		isValid, err := a.handler.ValidateDigitalFootprint(c, session, c.MustGet("digital_fingerprint").(string))
		if err != nil {
			spew.Dump(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to verify token"})
			c.Abort()
			return
		}

		if !isValid {
			// a.handler.DeleteSession(c, session)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		user, err := a.handler.GetUserBySessionID(c, *sessionId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to verify token"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func getErrorMessage(err error) string {
	if errors.Is(err, handlers.ErrExpiredToken) {
		return "Expired session"
	}

	if errors.Is(err, handlers.ErrInvalidToken) {
		return "Invalid session"
	}

	return "Unable to validate token"
}
