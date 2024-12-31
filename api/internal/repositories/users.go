package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"example.com/expenses-tracker/pkg/models"
)

const (
	errFailedToCreateUser         = "failed to create user: %w"
	errFailedToGetUserByEmail     = "failed to get user by email: %w"
	errFailedToGetUserByAuthToken = "failed to get user by auth token: %w"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository sets up a new user repository
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	sql := `
		INSERT INTO users (id, first_name, last_name, email, password, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	result, err := u.db.ExecContext(ctx, sql, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return fmt.Errorf(errFailedToCreateUser, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf(errFailedToCreateUser, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf(errFailedToCreateUser, errors.New("no rows affected"))
	}

	return nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	return nil
}

func (u *userRepository) DeleteUser(ctx context.Context, user *models.User) error {
	return nil
}

func (u *userRepository) GetUserByEmailAddress(ctx context.Context, email string) (*models.User, error) {
	sqlStmt := `
		SELECT
			u.id, u.first_name, u.last_name, u.email, u.password, u.created_at
		FROM
			users u
		WHERE u.email = $1 LIMIT 1
	`
	var user models.User
	err := u.db.QueryRowContext(ctx, sqlStmt, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf(errFailedToGetUserByEmail, err)
	}

	return &user, nil
}

func (u *userRepository) GetUserBySessionID(ctx context.Context, token string) (*models.User, error) {
	sql := `
		SELECT
			u.id, u.first_name, u.last_name, u.email, u.password, u.created_at
		FROM
			users u
		JOIN
			users_sessions s ON s.user_id = u.id
		WHERE 
			s.id = $1
		LIMIT 1
	`

	var user models.User
	err := u.db.QueryRowContext(ctx, sql, token).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetUserByAuthToken, err)
	}

	return &user, nil
}
