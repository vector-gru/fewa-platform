package utils

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt with the default cost
// Returns the hashed password as a string and an error if any occurs
func HashPassword(password string) (string, error) {
    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    return string(hashedBytes), nil
}

// VerifyPassword checks if the provided password matches the hashed password
// Returns an error if the passwords do not match
func VerifyPassword(hashedPassword, password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return fmt.Errorf("password verification failed: %w", err)
    }
    return nil
}


