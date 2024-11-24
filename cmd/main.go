package cmd

import (
	"fmt"

	"example.com/expenses-tracker/internal/database"
	"example.com/expenses-tracker/internal/http"
	"github.com/gin-gonic/gin"
)

const (
	errFailedToConnectToDatabase = "failed to connect to database: %w"
)

// Setup prepares this application
func Setup() {
	gin := gin.New()
	db, err := database.NewDatabase()
	if err != nil {
		panic(fmt.Errorf(errFailedToConnectToDatabase, err))
	}

	router := &http.Router{}
	router.Setup(gin, db)

	gin.Run()
	defer db.Close()
}
