package http

import (
	"log/slog"
	"net/http"

	"example.com/expenses-tracker/api/internal/repositories"
	"example.com/expenses-tracker/api/internal/validation"
	"example.com/expenses-tracker/pkg/models"
	"example.com/expenses-tracker/pkg/requests"
	"example.com/expenses-tracker/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UsersHandler struct {
	repo repositories.UserRepository
}

func NewUsersHandler(repo repositories.UserRepository) *UsersHandler {
	return &UsersHandler{repo: repo}
}

func (u *UsersHandler) RegisterRoutes(g *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	g.POST("", u.createUser)
	g.GET("/whoami", append(middlewares, u.getAuthenticatedUser)...)
}

func (u *UsersHandler) createUser(c *gin.Context) {
	var request requests.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := validation.FormatValidationMessages(request, err.(validator.ValidationErrors))
		c.JSON(http.StatusBadRequest, responses.NewErrorJsonHttpResponse(http.StatusBadRequest, validationErrors))
		return
	}

	user := &models.User{}
	if err := user.FromUserRequest(&request); err != nil {
		slog.Debug("failed to create user", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	}

	if err := u.repo.CreateUser(c, user); err != nil {
		slog.Debug("failed to create user", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.Header("Location", "/users/"+user.ID)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (u *UsersHandler) getAuthenticatedUser(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	c.JSON(http.StatusOK, user)
}
