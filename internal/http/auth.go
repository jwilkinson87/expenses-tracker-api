package http

import (
	"net/http"

	"example.com/expenses-tracker/internal/handlers"
	"example.com/expenses-tracker/internal/requests"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	internalHandler handlers.AuthHandler
}

// NewAuthHandler creates a new auth http handler
func NewAuthHandler(internalHandler handlers.AuthHandler) *AuthHandler {
	return &AuthHandler{internalHandler: internalHandler}
}

func (h *AuthHandler) RegisterRoutes(g *gin.Engine) {
	g.POST("/login", h.loginUser)
	g.POST("/logout", h.logoutUser)
}

func (h *AuthHandler) loginUser(c *gin.Context) {
	var loginRequest requests.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.internalHandler.HandleLoginRequest(c, &loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
}

func (h *AuthHandler) logoutUser(c *gin.Context) {

}
