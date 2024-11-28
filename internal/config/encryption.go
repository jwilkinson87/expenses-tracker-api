package config

import "os"

type EncryptionConfig struct {
	Key string
}

// NewEncryptionConfig creates a new encryption config type
func NewEncryptionConfig(key string, IV string) *EncryptionConfig {
	return &EncryptionConfig{Key: key}
}

func NewEncryptionConfigFromEnvironmentVariables() *EncryptionConfig {
	return &EncryptionConfig{
		Key: os.Getenv("ENCRYPTION_KEY"),
	}
}
