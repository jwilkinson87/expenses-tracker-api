package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

type EncryptionHandler struct {
	SecretKey []byte
}

func NewEncryptionHandlerFromEnvVars() *EncryptionHandler {
	return &EncryptionHandler{
		SecretKey: []byte(os.Getenv("ENCRYPTION_KEY")),
	}
}

func (h *EncryptionHandler) Encrypt(token string) (string, error) {
	block, err := aes.NewCipher(h.SecretKey)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(token), nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (h *EncryptionHandler) Decrypt(encryptedToken string) (string, error) {
	block, err := aes.NewCipher(h.SecretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedToken)
	if err != nil {
		return "", err
	}

	nonce, ciphertext := ciphertext[:12], ciphertext[12:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
