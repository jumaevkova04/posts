package crypto

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

// GenerateTokenBase64 returns a unique token based on the provided key string
func GenerateTokenBase64(key string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hash), nil
}
