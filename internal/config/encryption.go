package config

import (
	"encoding/base64"
	"os"
)

type EncryptionConfig struct {
	Key string
}

// NewEncryptionConfig creates a new encryption config type
func NewEncryptionConfig(key string, IV string) *EncryptionConfig {
	return &EncryptionConfig{Key: key}
}

func NewEncryptionConfigFromEnvironmentVariables() *EncryptionConfig {
	decoded, err := base64.StdEncoding.DecodeString(os.Getenv("ENCRYPTION_KEY"))
	if err != nil {
		panic("Unable to decode encryption key")
	}

	return &EncryptionConfig{
		Key: string(decoded),
	}
}
