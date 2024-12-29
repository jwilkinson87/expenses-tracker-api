package http

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/expenses-tracker/internal/repositories"
	"example.com/expenses-tracker/pkg/models"
	"example.com/expenses-tracker/pkg/requests"
	"example.com/expenses-tracker/pkg/responses"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	repo repositories.UserRepository
}

func NewUsersHandler(repo repositories.UserRepository) *UsersHandler {
	return &UsersHandler{repo: repo}
}

func (u *UsersHandler) RegisterRoutes(g *gin.Engine) {
	g.POST("/users", u.createUser)
	g.GET("/whoami", u.getAuthenticatedUser)
}

func (u *UsersHandler) createUser(c *gin.Context) {
	var request requests.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorJsonHttpResponse(http.StatusBadRequest, request, err))
		return
	}

	user := &models.User{}
	if err := user.FromUserRequest(&request); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	}

	if err := u.repo.CreateUser(c, user); err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.Header("Location", "/users/"+user.ID)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (u *UsersHandler) getAuthenticatedUser(c *gin.Context) {
	token := c.MustGet("user_token").(string) // At this point, it's already been validated
	user, err := u.repo.GetUserByAuthToken(c, token)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve authenticated user"})
		return
	}

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve authenticated user"})
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}
