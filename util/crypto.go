package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

const secretKey = "passphrasewhichneedstobe32bytes!"

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

func EncryptText(text []byte) []byte {
	//as secret key is 32 byte, AES-256 used
	c, err := aes.NewCipher([]byte(secretKey))

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	return gcm.Seal(nonce, nonce, text, nil)
}

func DecryptText(encryptedText []byte) []byte {
	ciphertext := encryptedText

	c, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return plaintext
}

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
