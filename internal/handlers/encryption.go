package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"example.com/expenses-tracker/internal/config"
)

const (
	errFailedToCreateCipherBlock = "failed to create cipher block: %w"
)

type EncryptionHandler struct {
	config *config.EncryptionConfig
	block  cipher.Block
}

// NewEncryptionHandler creates a new encryption handler instance.
func NewEncryptionHandler(config *config.EncryptionConfig) (*EncryptionHandler, error) {
	handler := &EncryptionHandler{config: config}
	block, err := aes.NewCipher([]byte(config.Key))
	if err != nil {
		return nil, fmt.Errorf(errFailedToCreateCipherBlock, err)
	}

	handler.block = block
	return handler, nil
}

func (h *EncryptionHandler) EncryptValue(value []byte) ([]byte, error) {
	ciphertext := make([]byte, aes.BlockSize+len(value))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(h.block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], value)

	return ciphertext, nil
}

func (h *EncryptionHandler) DecryptValue(value []byte) ([]byte, error) {
	if len(value) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := value[:aes.BlockSize]
	value = value[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(h.block, iv)
	cfb.XORKeyStream(value, value)

	return value, nil
}
