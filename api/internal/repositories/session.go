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
	_, err := r.db.ExecContext(ctx, "INSERT INTO user_sessions (id, user_id, digital_fingerprint, created_at, expiry_time) VALUES ($1, $2, $3, $4)", session.ID, session.User.ID, session.DigitalFingerPrint, session.CreatedAt, session.ExpiryTime)
	if err != nil {
		return err
	}

	return nil
}

func (r *userSessionRepository) DeleteSession(ctx context.Context, session *models.UserSession) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM user_sessions WHERE id = $1", session.ID)
	return err
}

func (r *userSessionRepository) DeleteAllForUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM user_sessions WHERE user_id = $1", user.ID)
	return err
}

func (r *userSessionRepository) GetBySessionID(ctx context.Context, sessionID string) (*models.UserSession, error) {
	var session models.UserSession
	var user models.User

	err := r.db.QueryRowContext(
		ctx,
		`SELECT 
        s.id,
        s.digital_fingerprint, 
        s.created_at, 
        s.expiry_time, 
        u.id AS user_id, 
        u.first_name, 
        u.last_name, 
        u.email, 
        u.password, 
        u.created_at
     FROM 
        user_sessions s
     JOIN 
        users u
     ON 
        s.user_id = u.id
     WHERE 
        s.id = $1`,
		sessionID,
	).Scan(
		&session.ID,
		&session.DigitalFingerPrint,
		&session.CreatedAt,
		&session.ExpiryTime,
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	session.User = &user

	return &session, nil
}
