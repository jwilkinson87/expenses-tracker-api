package http

import (
	"example.com/expenses-tracker/internal/repositories"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repo repositories.UserAuthRepository
}

// NewAuthHandler creates a new auth http handler
func NewAuthHandler(repo repositories.UserAuthRepository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

func (h *AuthHandler) RegisterRoutes(g *gin.Engine) {
	g.POST("/login", h.loginUser)
	g.POST("/logout", h.logoutUser)
}

func (h *AuthHandler) loginUser(c *gin.Context) {

}

func (h *AuthHandler) logoutUser(c *gin.Context) {

}
