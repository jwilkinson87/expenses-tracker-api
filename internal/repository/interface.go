package repository

import (
	"context"

	models "example.com/expenses-tracker/internal/pkg"
)

type ExpenseRepository interface {
	CreateExpense(context.Context, *models.Expense) error
	GetExpense(context.Context, string) (*models.Expense, error)
	UpdateExpense(context.Context, *models.Expense) error
	DeleteExpense(context.Context, *models.Expense) error
	GetAllForUser(context.Context, *models.User) (models.Expenses, error)
}
