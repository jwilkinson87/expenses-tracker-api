package http

import (
	"database/sql"

	"example.com/expenses-tracker/internal/repository"
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (r *Router) Setup(g *gin.Engine, db *sql.DB) error {
	repo := repository.NewExpensesRepository(db)
	expenseHandler := NewExpensesHandler(repo)

	expenseHandler.registerRoutes(g)
}
