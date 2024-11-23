package http

import (
	"net/http"

	models "example.com/expenses-tracker/internal/pkg"
	"example.com/expenses-tracker/internal/repository"
	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	repo repository.ExpenseRepository
}

// NewHandler sets up a new HTTP handler
func NewHandler(repository repository.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: repository}
}

func (e *ExpenseHandler) GetExpenses(c *gin.Context) {
	ctx := c.Request.Context()
	// TODO: This needs to be an authenticated user
	user := &models.User{
		ID: "123",
	}

	expenses, err := e.repo.GetAllForUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve expenses"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func (e *ExpenseHandler) CreateExpense(c *gin.Context) {
	var json models.Expense
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := e.repo.CreateExpense(c.Request.Context(), &json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, nil)
}
