package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"example.com/expenses-tracker/internal/models"
)

const (
	errFailedToCreateUser     = "failed to create user: %w"
	errFailedToGetUserByEmail = "failed to get user by email: %w"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository sets up a new user repository
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	sql := "SELECT * FROM users WHERE email_address = $1"
	var user models.User
	err := u.db.QueryRowContext(ctx, sql, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetUserByEmail, err)
	}

	return &user, nil
}
