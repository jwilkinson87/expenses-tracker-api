package http

import (
	"log"
	"net/http"

	"example.com/expenses-tracker/internal/models"
	"example.com/expenses-tracker/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
)

type ExpenseHandler struct {
	repo repositories.ExpenseRepository
}

// NewExpensesHandler sets up a new HTTP handler
func NewExpensesHandler(repository repositories.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: repository}
}

func (e *ExpenseHandler) RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/", e.getAll)
	g.POST("/", e.create)
}

func (e *ExpenseHandler) getAll(c *gin.Context) {
	uuid, _ := uuid.GenerateUUID()
	ctx := c.Request.Context()
	// TODO: This needs to be an authenticated user
	user := &models.User{
		ID: uuid,
	}

	expenses, err := e.repo.GetAllForUser(ctx, user)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve expenses"})
		return
	}

	if expenses == nil {
		c.JSON(http.StatusOK, make([]models.Expense, 0))
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (e *ExpenseHandler) create(c *gin.Context) {
	var json models.Expense
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := e.repo.CreateExpense(c.Request.Context(), &json)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, nil)
}
