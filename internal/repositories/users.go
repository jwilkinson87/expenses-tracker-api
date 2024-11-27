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

func (u *userRepository) DeleteUser(ctx context.Context, user *models.User) error {
	return nil
}

func (u *userRepository) GetUserByEmailAddress(ctx context.Context, email string) (*models.User, error) {
	sql := `
		SELECT
			u.id, u.first_name, u.last_name, u.email, u.password, u.created_at
		FROM
			users u
		WHERE u.email = $1 LIMIT 1
	`
	var user models.User
	err := u.db.QueryRowContext(ctx, sql, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetUserByEmail, err)
	}

	return &user, nil
}
