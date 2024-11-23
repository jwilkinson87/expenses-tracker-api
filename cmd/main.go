package cmd

import (
	"fmt"

	"example.com/expenses-tracker/internal/database"
	httpHandler "example.com/expenses-tracker/internal/http"
	"example.com/expenses-tracker/internal/repository"
	"github.com/gin-gonic/gin"
)

const (
	errFailedToConnectToDatabase = "failed to connect to database: %w"
)

// Setup prepares this application
func Setup() {
	router := gin.New()
	db, err := database.NewDatabase()
	if err != nil {
		panic(fmt.Errorf(errFailedToConnectToDatabase, err))
	}

	repo := repository.NewExpensesRepository(db)
	handler := httpHandler.NewHandler(repo)

	router.GET("/expenses", handler.GetExpenses)

	router.Run()
	defer db.Close()
}
