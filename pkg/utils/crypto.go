package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

func GenerateSalt(length int) (string, error) {
	if length <= 0 {
		length = 16
	}
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func HashPassword(password, salt string) string {
	hash := md5.Sum([]byte(password + salt))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(password, salt, hashedPassword string) bool {
	return HashPassword(password, salt) == hashedPassword
}

