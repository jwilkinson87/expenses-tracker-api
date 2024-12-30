package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"example.com/expenses-tracker/pkg/models"
)

const (
	errFailedToCreateResetToken     = "failed to create reset token: %w"
	errFailedToGetResetTokenByValue = "failed to get reset token by value: %w"
)

type resetTokenRepository struct {
	db *sql.DB
}

func NewResetTokenRepository(db *sql.DB) *resetTokenRepository {
	return &resetTokenRepository{db}
}

func (r *resetTokenRepository) CreateResetTokenForUser(ctx context.Context, model *models.ResetToken) error {
	sql := `
	INSERT INTO reset_tokens (id, reset_token, expiry_token, created_at, user_id) VALUES
	($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, sql, model.ID, model.ResetToken, model.ExpiryTime, model.CreatedAt, model.User.ID)
	if err != nil {
		return fmt.Errorf(errFailedToCreateResetToken, err)
	}

	return nil
}

func (r *resetTokenRepository) GetResetToken(ctx context.Context, resetToken string) (*models.ResetToken, error) {
	sql := `
	SELECT r.id, r.reset_token, r.expiry_token, r.created_at,
	u.id AS user_id, u.first_name, u.last_name, u.email, u.password, u.created_at AS user_created_at
	FROM reset_tokens r
	JOIN users u ON u.id = r.user_id
	WHERE r.reset_token = $1
	`

	var result models.ResetToken
	var user models.User

	err := r.db.QueryRow(sql, resetToken).Scan(
		&result.ID, &result.ResetToken, &result.ExpiryTime, &result.CreatedAt,
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf(errFailedToGetResetTokenByValue, err)
	}

	result.User = &user

	return &result, nil
}
