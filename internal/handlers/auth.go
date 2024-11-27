package handlers

import (
	"context"
	"errors"
	"fmt"

	"example.com/expenses-tracker/internal/repositories"
)

const (
	errInvalidToken       = "invalid token"
	errfailedToCheckToken = "failed to check token: %w"
)

type AuthHandler struct {
	userTokenRepository repositories.UserAuthRepository
	userRepository      repositories.UserRepository
}

// NewAuthHandler creates a new auth handler for checking an authenticated user
func NewAuthHandler(userTokenRepository repositories.UserAuthRepository, userRepository repositories.UserRepository) *AuthHandler {
	return &AuthHandler{
		userTokenRepository: userTokenRepository,
		userRepository:      userRepository,
	}
}

// ValidateToken will check that a specified token value is valid.
func (a *AuthHandler) ValidateToken(ctx context.Context, token string) (error, bool) {
	if token == "" || len(token) == 0 {
		return errors.New(errInvalidToken), false
	}

	userToken, err := a.userTokenRepository.GetByAuthToken(ctx, token)
	if err != nil {
		return fmt.Errorf(errfailedToCheckToken, err), false
	}

	if userToken == nil {
		return nil, false
	}

	return nil, userToken.IsTokenValid()
}
