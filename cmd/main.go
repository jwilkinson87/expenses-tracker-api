package cmd

import (
	"database/sql"
	"fmt"
	"sync"

	"example.com/expenses-tracker/internal/config"
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
	EncryptionHandler  *handlers.EncryptionHandler
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
		setupContainer(db)
	})

	setupMiddleware(gin)
	setupHttpHandlers(gin)

	gin.Run()
	defer db.Close()
}

func setupContainer(db *sql.DB) {
	encryptionHandler, err := handlers.NewEncryptionHandler(config.NewEncryptionConfigFromEnvironmentVariables())
	if err != nil {
		panic(err)
	}

	container.UserAuthRepository = repositories.NewAuthRepository(db)
	container.UserRepository = repositories.NewUserRepository(db)
	container.ExpenseRepository = repositories.NewExpensesRepository(db)
	container.EncryptionHandler = encryptionHandler
	container.AuthHandler = handlers.NewAuthHandler(container.UserAuthRepository, container.UserRepository, encryptionHandler)
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

	authHandler := http.NewAuthHandler(*container.AuthHandler)
	authHandler.RegisterRoutes(g)
}
