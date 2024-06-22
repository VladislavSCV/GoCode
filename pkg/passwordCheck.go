package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"regexp"
)

func shaHashing(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)
	return hex.EncodeToString(sha256Hash[:])
}

func CheckAndHashPassword(password string) (string, error) {
	CheckPassword(password)
	return HashFuncPassword(password)
}


func HashFuncPassword(password string) (string, error) {
	// todo: implement password hashing
	return shaHashing(password), nil
}


func CheckPassword(password string) (bool, error) {
	// todo: implement password check strong
	// Check password length
	if len(password) < 8 {
		return false, errors.New("password must be at least 8 characters long")
	}

	// Check if password contains at least one digit
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return false, errors.New("password must contain at least one digit")
	}

	// Check if password contains at least one uppercase letter
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false, errors.New("password must contain at least one uppercase letter")
	}

	

	// Check if password contains at least one lowercase letter
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false, errors.New("password must contain at least one lowercase letter")
	}

	// Check if password contains at least one special character
	if !regexp.MustCompile(`[!@#$%^&*()\-_=+[\]{}|\\:;"'<>?,./]`).MatchString(password) {
		return false, errors.New("password must contain at least one special character")
	}
	return true, nil
}
