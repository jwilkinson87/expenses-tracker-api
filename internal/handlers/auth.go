package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"example.com/expenses-tracker/internal/models"
	"example.com/expenses-tracker/internal/repositories"
	"example.com/expenses-tracker/internal/requests"
	"example.com/expenses-tracker/internal/responses"
	"golang.org/x/crypto/bcrypt"
)

const (
	errInvalidToken            = "invalid token"
	errFailedToCheckToken      = "failed to check token: %w"
	errFailedToCreateToken     = "failed to create token: %w"
	errFailedToGetTokenByValue = "failed to get token by value: %w"
	errFailedToGetUserByEmail  = "failed to get user by email address: %w"
	errInvalidCredentials      = "user credentials incorrect"
)

type AuthHandler struct {
	userTokenRepository repositories.UserAuthRepository
	userRepository      repositories.UserRepository
	encryptionHandler   *EncryptionHandler
}

// NewAuthHandler creates a new auth handler for checking an authenticated user
func NewAuthHandler(userTokenRepository repositories.UserAuthRepository, userRepository repositories.UserRepository, encryptionHandler *EncryptionHandler) *AuthHandler {
	return &AuthHandler{
		userTokenRepository: userTokenRepository,
		userRepository:      userRepository,
		encryptionHandler:   encryptionHandler,
	}
}

func (h *AuthHandler) HandleLoginRequest(ctx context.Context, request *requests.LoginRequest) (*responses.AuthenticatedUserResponse, error) {
	user, err := h.userRepository.GetUserByEmailAddress(ctx, request.EmailAddress)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetUserByEmail, err)
	}

	if user == nil {
		return nil, errors.New(errInvalidCredentials)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.New(errInvalidCredentials)
	}

	token, err := h.generateSecureTokenForUser(32)
	if err != nil {
		return nil, err
	}

	result, err := h.persistTokenForUser(ctx, token, user)
	if err != nil {
		return nil, fmt.Errorf(errFailedToCreateToken, err)
	}
	response := &responses.AuthenticatedUserResponse{
		Token:      token,
		ExpiryTime: *result.ExpiryTime,
	}

	return response, nil
}

func (h *AuthHandler) HandleLogout(ctx context.Context, token string) (bool, error) {
	result, err := h.userTokenRepository.GetByAuthToken(ctx, token)
	if err != nil {
		return false, fmt.Errorf(errFailedToGetTokenByValue, err)
	}

	if result == nil {
		return false, errors.New(errInvalidToken)
	}

	if !result.IsTokenValid() {
		return false, errors.New(errInvalidToken)
	}

	if err = h.userTokenRepository.DeleteAuthToken(ctx, result); err != nil {
		return false, err
	}

	return true, nil
}

func (h *AuthHandler) persistTokenForUser(ctx context.Context, token string, user *models.User) (*models.UserToken, error) {
	encryptedToken, err := h.encryptionHandler.EncryptValue([]byte(token))
	if err != nil {
		return nil, err
	}

	expiryTime := time.Now().Add(time.Minute * 20)

	authToken := &models.UserToken{
		Value:      string(encryptedToken),
		User:       user,
		ExpiryTime: &expiryTime,
	}

	h.userTokenRepository.CreateAuthToken(ctx, authToken)

	return authToken, nil
}

func (h *AuthHandler) generateSecureTokenForUser(tokenSize int64) (string, error) {
	bytes := make([]byte, tokenSize)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes), nil
}

// ValidateToken will check that a specified token value is valid.
func (a *AuthHandler) ValidateToken(ctx context.Context, token string) (bool, error) {
	if token == "" || len(token) == 0 {
		return false, errors.New(errInvalidToken)
	}

	userToken, err := a.userTokenRepository.GetByAuthToken(ctx, token)
	if err != nil {
		return false, fmt.Errorf(errFailedToCheckToken, err)
	}

	if userToken == nil {
		return false, nil
	}

	return userToken.IsTokenValid(), nil
}
