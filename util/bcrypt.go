package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// generate hash out of secret
func HashSecret(secret string) (string, error) {
	hs, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash secret, %w", err)
	}
	return string(hs), nil
}

// verify hash secret
func VerifyHashSecret(secret, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
}
