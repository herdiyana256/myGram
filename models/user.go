package models

import (
    // "gorm.io/gorm"
    "time"
)

//  struktur data untuk user  aplikasi.
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"not null;unique"`
    Email     string    `gorm:"not null;unique"`
    Password  string    `gorm:"not null"`
    Age       int       `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

//  metode untuk memvalidasi data pengguna.
func (u *User) Validate() error {
    // Implementasi validasi sesuai kebutuhan aplikasi.
    return nil
}

// HashPassword -->  metode untuk meng-hash password pengguna.
func (u *User) HashPassword() error {
    // Implementasi peng-hash-an password 
    return nil
}
