// package models: Package yang berisi definisi struktur data yang merepresentasikan objek dalam aplikasi.

import (
	"errors" // Mengimport package errors untuk menangani error dalam aplikasi.
	"regexp" // Mengimport package regexp untuk penggunaan regular expression.
	"golang.org/x/crypto/bcrypt" // Mengimport package bcrypt untuk mengenkripsi password.
	"gorm.io/gorm" // Mengimport package gorm untuk berinteraksi dengan database.
)

// IsValidEmail memeriksa apakah alamat email yang diberikan valid.
func IsValidEmail(email string) bool {
	// Pola regular expression untuk memvalidasi alamat email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// User adalah struktur data yang merepresentasikan pengguna dalam aplikasi.
type User struct {
	gorm.Model // Struktur data User memiliki Model bawaan dari gorm untuk menyimpan ID, CreatedAt, UpdatedAt, dan DeletedAt.
	Username string // Username adalah nama pengguna.
	Email    string // Email adalah alamat email pengguna.
	Password string // Password adalah kata sandi pengguna.
	Age      int    // Age adalah usia pengguna.
}

// Validate memvalidasi bidang-bidang pengguna.
// Metode ini memeriksa apakah Username, Email, Password, dan Age tidak kosong dan apakah alamat email valid.
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required") // Jika Username kosong, maka akan dikembalikan error.
	}
	if u.Email == "" {
		return errors.New("email is required") // Jika Email kosong, maka akan dikembalikan error.
	}
	if !IsValidEmail(u.Email) {
		return errors.New("invalid email address") // Jika Email tidak valid, maka akan dikembalikan error.
	}
	if u.Password == "" {
		return errors.New("password is required") // Jika Password kosong, maka akan dikembalikan error.
	}
	if u.Age < 8 {
		return errors.New("age must be at least 8") // Jika Age kurang dari 8, maka akan dikembalikan error.
	}
	return nil // Jika validasi berhasil, tidak ada error yang dikembalikan.
}

// HashPassword mengenkripsi password pengguna menggunakan bcrypt.
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err // Jika terjadi error saat mengenkripsi password, error akan dikembalikan.
	}
	u.Password = string(hashedPassword) // Password yang dienkripsi disimpan di Password pengguna.
	return nil // Jika enkripsi berhasil, tidak ada error yang dikembalikan.
}

// CreateUser membuat pengguna baru dalam database.
func CreateUser(user *User) error {
	// Implementasi Anda untuk membuat pengguna baru dalam database.
	return nil // Kembalikan nil jika operasi berhasil.
}
