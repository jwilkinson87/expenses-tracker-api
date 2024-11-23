package repository

import (
	"context"

	"example.com/expenses-tracker/internal/models"
	"example.com/expenses-tracker/internal/requests"
)

type ExpenseRepository interface {
	CreateExpense(context.Context, *models.Expense) error
	GetExpense(context.Context, string) (*models.Expense, error)
	UpdateExpense(context.Context, *models.Expense) error
	DeleteExpense(context.Context, *models.Expense) error
	GetAllForUser(context.Context, *models.User) (models.Expenses, error)
}

type UserRepository interface {
	CreateUser(context.Context, *models.User) error
	LoginUser(context.Context, *requests.LoginRequest) string
}
