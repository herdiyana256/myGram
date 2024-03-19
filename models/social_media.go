// package models: Package yang berisi definisi struktur data yang merepresentasikan objek dalam aplikasi.

import (
	"errors" // Mengimport package errors untuk menangani error dalam aplikasi.
	"time"   // Mengimport package time untuk penggunaan tipe data waktu.
)

// SocialMedia adalah struktur data yang merepresentasikan media sosial dalam aplikasi.
type SocialMedia struct {
	ID             uint      `gorm:"primaryKey"` // ID adalah identifikasi unik dari media sosial.
	Name           string    `gorm:"not null"`   // Name adalah nama dari media sosial.
	SocialMediaURL string    `gorm:"not null"`   // SocialMediaURL adalah URL atau tautan dari media sosial.
	UserID         uint      `gorm:"not null"`   // UserID adalah ID pengguna yang memiliki media sosial ini.
	CreatedAt      time.Time `gorm:"autoCreateTime"` // CreatedAt adalah waktu pembuatan media sosial.
	UpdatedAt      time.Time `gorm:"autoUpdateTime"` // UpdatedAt adalah waktu terakhir kali media sosial diperbarui.
}

// Validate adalah metode untuk melakukan validasi terhadap media sosial.
// Metode ini memeriksa apakah nama dan URL media sosial tidak kosong.
func (s *SocialMedia) Validate() error {
	if s.Name == "" {
		return errors.New("name is required") // Jika nama media sosial kosong, maka akan dikembalikan error.
	}
	if s.SocialMediaURL == "" {
		return errors.New("social media URL is required") // Jika URL media sosial kosong, maka akan dikembalikan error.
	}
	return nil // Jika validasi berhasil, tidak ada error yang dikembalikan.
}

// Implement database operations if needed
// Implementasikan operasi-operasi database jika diperlukan.
// Operasi-operasi database seperti membuat, membaca, memperbarui, dan menghapus media sosial dapat diimplementasikan di sini.
