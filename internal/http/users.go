package http

import (
	"encoding/json"
	"log"
	"net/http"

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
	token := c.MustGet("user_token").(string) // At this point, it's already been validated
	user, err := u.repo.GetUserByAuthToken(c, token)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve authenticated user"})
		return
	}

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve authenticated user"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}
