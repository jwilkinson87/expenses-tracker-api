package repositories

import (
	"context"
	"database/sql"

	"example.com/expenses-tracker/pkg/models"
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
	sqlQuery := `DELETE FROM users_auth_tokens WHERE user_id = $1`
	_, err := a.db.ExecContext(ctx, sqlQuery, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *authRepository) DeleteAuthToken(ctx context.Context, token *models.UserToken) error {
	sqlQuery := `DELETE FROM users_auth_tokens WHERE id = $1`
	_, err := a.db.ExecContext(ctx, sqlQuery, token.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *authRepository) GetByAuthToken(ctx context.Context, tokenValue string) (*models.UserToken, error) {
	return nil, nil
}
