package repositories

import (
	"context"

	"example.com/expenses-tracker/pkg/models"
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
	UpdateUser(context.Context, *models.User) error
	GetUserByEmailAddress(context.Context, string) (*models.User, error)
	DeleteUser(context.Context, *models.User) error
	GetUserByAuthToken(context.Context, string) (*models.User, error)
}

type UserAuthRepository interface {
	CreateAuthToken(context.Context, *models.UserToken) error
	DeleteAuthToken(context.Context, *models.UserToken) error
	DeleteAllForUser(context.Context, *models.User) error
	GetByAuthToken(context.Context, string) (*models.UserToken, error)
}

type ResetTokenRepository interface {
	CreateResetTokenForUser(context.Context, *models.ResetToken) error
	GetResetToken(context.Context, string) (*models.ResetToken, error)
}
