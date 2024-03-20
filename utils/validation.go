package utils

import (
	"regexp"              
	"golang.org/x/crypto/bcrypt" // hashing password
)

// IsValidEmail digunakan untuk memeriksa apakah alamat email yang diberikan valid.
func IsValidEmail(email string) bool {
	// Pola ekspresi reguler untuk memvalidasi alamat email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// HashPassword --> untuk meng-hash password yang diberikan menggunakan bcrypt.
func HashPassword(password string) (string, error) {
	// Generate hashed password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
