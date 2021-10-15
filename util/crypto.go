package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func NewSHA256(text []byte) []byte {
	hash := sha256.Sum256(text)
	return hash[:]
}

func CompareSHA256(text, hash []byte) error {
	result := NewSHA256(text)
	if hex.EncodeToString(result) != hex.EncodeToString(hash) {
		return fmt.Errorf("sha256 did not match")
	}
	return nil
}
