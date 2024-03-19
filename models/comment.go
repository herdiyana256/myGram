// package models: Package yang berisi definisi struktur data yang merepresentasikan objek dalam aplikasi.

import (
	"errors" // Mengimport package errors untuk menangani error dalam aplikasi.
	"time"   // Mengimport package time untuk penggunaan tipe data waktu.
)


type Comment struct {
	ID        uint      `gorm:"primaryKey"` // ID adalah identifikasi unik dari komentar.
	UserID    uint      `gorm:"not null"`   // UserID adalah ID pengguna yang membuat komentar.
	PhotoID   uint      `gorm:"not null"`   // PhotoID adalah ID foto yang dikomentari.
	Message   string    `gorm:"not null"`   // Message adalah pesan yang dituliskan dalam komentar.
	CreatedAt time.Time `gorm:"autoCreateTime"` // CreatedAt adalah waktu pembuatan komentar.
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // UpdatedAt adalah waktu terakhir kali komentar diperbarui.
}


// Metode ini memeriksa apakah pesan komentar tidak kosong.
func (c *Comment) Validate() error {
	if c.Message == "" {
		return errors.New("message is required") // Jika pesan komentar kosong, maka akan dikembalikan error.
	}
	return nil // Jika validasi berhasil, tidak ada error yang dikembalikan.
}

// Implement database operations if needed
// Implementasikan operasi-operasi database jika diperlukan.
// Operasi-operasi database seperti membuat, membaca, memperbarui, dan menghapus komentar dapat diimplementasikan di sini.
