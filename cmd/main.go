package cmd

import (
	"database/sql"
	"fmt"
	"sync"

	"example.com/expenses-tracker/internal/database"
	"example.com/expenses-tracker/internal/handlers"
	"example.com/expenses-tracker/internal/http"
	"example.com/expenses-tracker/internal/http/middleware"
	"example.com/expenses-tracker/internal/repositories"
	"github.com/gin-gonic/gin"
)

type Container struct {
	UserRepository     repositories.UserRepository
	UserAuthRepository repositories.UserAuthRepository
	ExpenseRepository  repositories.ExpenseRepository
	AuthHandler        *handlers.AuthHandler
}

var (
	container *Container
	once      sync.Once
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

	container = &Container{}
	once.Do(func() {
		setupRepositories(db)
	})

	setupMiddleware(gin)
	setupHttpHandlers(gin)

	gin.Run()
	defer db.Close()
}

func setupRepositories(db *sql.DB) {
	container.UserAuthRepository = repositories.NewAuthRepository(db)
	container.UserRepository = repositories.NewUserRepository(db)
	container.ExpenseRepository = repositories.NewExpensesRepository(db)
	container.AuthHandler = handlers.NewAuthHandler(container.UserAuthRepository, container.UserRepository)
}

func setupMiddleware(g *gin.Engine) {
	authMiddleware := middleware.NewAuthMiddleware(container.AuthHandler)

	g.Use(middleware.RequestIdMiddleware())
	g.Use(authMiddleware.HandleAuthToken())
}

func setupHttpHandlers(g *gin.Engine) {
	expenseHandler := http.NewExpensesHandler(container.ExpenseRepository)
	expenseHandler.RegisterRoutes(g)

	userHandler := http.NewUsersHandler(container.UserRepository)
	userHandler.RegisterRoutes(g)
}
