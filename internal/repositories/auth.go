package repositories

import (
	"context"
	"database/sql"

	"example.com/expenses-tracker/internal/models"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *authRepository {
	return &authRepository{db}
}

func (a *authRepository) CreateAuthToken(ctx context.Context, token *models.UserToken) error {
	return nil
}

func (a *authRepository) DeleteAllForUser(ctx context.Context, user *models.User) error {
	return nil
}

func (a *authRepository) DeleteAuthToken(ctx context.Context, token *models.UserToken) error {
	return nil
}

func (a *authRepository) GetByAuthToken(ctx context.Context, tokenValue string) (*models.UserToken, error) {
	return nil, nil
}
