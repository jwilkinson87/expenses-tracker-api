package handlers

import "example.com/expenses-tracker/internal/repositories"

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

func (a *AuthHandler) validateToken(token string) (error, bool) {
	return nil, false
}
