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
func (a *AuthHandler) ValidateToken(ctx context.Context, token string) (bool, error) {
	if token == "" || len(token) == 0 {
		return false, errors.New(errInvalidToken)
	}

	userToken, err := a.userTokenRepository.GetByAuthToken(ctx, token)
	if err != nil {
		return false, fmt.Errorf(errfailedToCheckToken, err)
	}

	if userToken == nil {
		return false, nil
	}

	return userToken.IsTokenValid(), nil
}
