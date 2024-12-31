package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"

	"example.com/expenses-tracker/pkg/models"
)

type TokenHandler struct {
	secretKey []byte
}

func NewTokenHandler(secretKey []byte) *TokenHandler {
	return &TokenHandler{
		secretKey: secretKey,
	}
}

func (h *TokenHandler) GenerateForSession(session *models.UserSession, expiryTime time.Time) string {
	data := session.ID

	handler := hmac.New(sha256.New, h.secretKey)
	handler.Write([]byte(data))
	signature := handler.Sum(nil)

	return base64.StdEncoding.EncodeToString([]byte(data + "." + string(signature)))
}

func (h *TokenHandler) ValidateToken(token string) (bool, *string) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false, nil
	}

	// Split the token into raw data (session ID) and signature
	parts := string(data)
	separatorIndex := strings.LastIndex(parts, ".")
	if separatorIndex == -1 {
		return false, nil
	}

	rawData := parts[:separatorIndex]
	signature := parts[separatorIndex+1:]

	// Recompute the HMAC signature
	handler := hmac.New(sha256.New, h.secretKey)
	handler.Write([]byte(rawData))
	expectedSignature := handler.Sum(nil)

	if !hmac.Equal(expectedSignature, []byte(signature)) {
		return false, nil
	}

	return true, &rawData
}
