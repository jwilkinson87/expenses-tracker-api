package http

import (
	"log/slog"
	"net/http"

	"example.com/expenses-tracker/api/internal/repositories"
	"example.com/expenses-tracker/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExpenseHandler struct {
	repo repositories.ExpenseRepository
}

// NewExpensesHandler sets up a new HTTP handler
func NewExpensesHandler(repository repositories.ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: repository}
}

func (e *ExpenseHandler) RegisterRoutes(g *gin.RouterGroup) {
	g.GET("", e.getAll)
	g.POST("", e.create)
}

func (e *ExpenseHandler) getAll(c *gin.Context) {
	uuid, _ := uuid.NewV7()
	ctx := c.Request.Context()
	// TODO: This needs to be an authenticated user
	user := &models.User{
		ID: uuid.String(),
	}

	expenses, err := e.repo.GetAllForUser(ctx, user)
	if err != nil {
		slog.Debug("failed to retrieve expenses", "error", err.Error())
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
		slog.Debug("failed to bind json", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := e.repo.CreateExpense(c.Request.Context(), &json)
	if err != nil {
		slog.Debug("failed to create expense", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, nil)
}
