package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"example.com/expenses-tracker/api/internal/repositories"
	"example.com/expenses-tracker/pkg/encryption"
	"example.com/expenses-tracker/pkg/models"
	"example.com/expenses-tracker/pkg/requests"
	"example.com/expenses-tracker/pkg/responses"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	errInvalidToken            = "invalid token"
	errFailedToCheckToken      = "failed to check token: %w"
	errFailedToCreateToken     = "failed to create token: %w"
	errFailedToGetTokenByValue = "failed to get token by value: %w"
	errFailedToGetUserByEmail  = "failed to get user by email address: %w"
	errFailedToGetUserByToken  = "failed to get user by token: %w"
	errInvalidCredentials      = "user credentials incorrect"
)

type AuthHandler struct {
	userTokenRepository repositories.UserAuthRepository
	userRepository      repositories.UserRepository
	encryptionHandler   *encryption.EncryptionHandler
}

// NewAuthHandler creates a new auth handler for checking an authenticated user
func NewAuthHandler(userTokenRepository repositories.UserAuthRepository, userRepository repositories.UserRepository, encryptionHandler *encryption.EncryptionHandler) *AuthHandler {
	return &AuthHandler{
		userTokenRepository: userTokenRepository,
		userRepository:      userRepository,
		encryptionHandler:   encryptionHandler,
	}
}

func (h *AuthHandler) GetUserForAuthToken(ctx context.Context, token string) (*models.User, error) {
	result, err := h.userRepository.GetUserByAuthToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetUserByToken, err)
	}

	return result, nil
}

func (h *AuthHandler) HandleLoginRequest(ctx context.Context, request *requests.LoginRequest) (*responses.AuthenticatedUserResponse, error) {
	user, err := h.userRepository.GetUserByEmailAddress(ctx, request.EmailAddress)
	if err != nil {
		return nil, err
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
	encryptedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expiryTime := time.Now().Add(time.Minute * 20)
	id, _ := uuid.NewV7()

	authToken := &models.UserToken{
		ID:         id.String(),
		CreatedAt:  &now,
		Value:      string(encryptedToken),
		User:       user,
		ExpiryTime: &expiryTime,
	}

	err = h.userTokenRepository.CreateAuthToken(ctx, authToken)
	if err != nil {
		return nil, fmt.Errorf(errFailedToCreateToken, err)
	}

	return authToken, err
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
