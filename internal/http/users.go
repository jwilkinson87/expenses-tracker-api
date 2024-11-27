package http

import (
	"example.com/expenses-tracker/internal/repositories"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	repo repositories.UserRepository
}

func NewUsersHandler(repo repositories.UserRepository) *UsersHandler {
	return &UsersHandler{repo: repo}
}

func (u *UsersHandler) RegisterRoutes(g *gin.Engine) {
	g.GET("/whoami", u.getAuthenticatedUser)
}

func (u *UsersHandler) getAuthenticatedUser(c *gin.Context) {

}
