package repositories

import (
	"context"
	"database/sql"

	"example.com/expenses-tracker/pkg/models"
)

type userSessionRepository struct {
	db *sql.DB
}

func NewUserSessionRepository(db *sql.DB) *userSessionRepository {
	return &userSessionRepository{db}
}

func (r *userSessionRepository) CreateSession(ctx context.Context, session *models.UserSession) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users_sessions (id, user_id, digital_fingerprint, created_at, expires_at) VALUES ($1, $2, $3, $4, $5)", session.ID, session.User.ID, session.DigitalFingerPrint, session.CreatedAt, session.ExpiryTime)
	if err != nil {
		return err
	}

	return nil
}

func (r *userSessionRepository) DeleteSession(ctx context.Context, session *models.UserSession) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users_sessions WHERE id = $1", session.ID)
	return err
}

func (r *userSessionRepository) DeleteAllForUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users_sessions WHERE user_id = $1", user.ID)
	return err
}

func (r *userSessionRepository) GetBySessionID(ctx context.Context, sessionID string) (*models.UserSession, error) {
	var session models.UserSession

	err := r.db.QueryRowContext(
		ctx,
		`SELECT s.id, s.digital_fingerprint, s.created_at, s.expires_at FROM users_sessions s WHERE s.id = $1`,
		sessionID,
	).Scan(
		&session.ID,
		&session.DigitalFingerPrint,
		&session.CreatedAt,
		&session.ExpiryTime,
	)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
