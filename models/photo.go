package models

// import (
// 	"errors" // Mengimport package errors untuk menangani error dalam aplikasi.
// 	"time"   // Mengimport package time untuk penggunaan tipe data waktu.
// )

// // Photo adalah struktur data yang merepresentasikan foto dalam aplikasi.
// type Photo struct {
// 	ID        uint      `gorm:"primaryKey"` // ID adalah identifikasi unik dari foto.
// 	Title     string    `gorm:"not null"`   // Title adalah judul dari foto.
// 	Caption   string    `gorm:"not null"`   // Caption adalah keterangan atau deskripsi dari foto.
// 	PhotoURL  string    `gorm:"not null"`   // PhotoURL adalah URL tempat foto disimpan.
// 	UserID    uint      `gorm:"not null"`   // UserID adalah ID pengguna yang mengunggah foto.
// 	CreatedAt time.Time `gorm:"autoCreateTime"` // CreatedAt adalah waktu pembuatan foto.
// 	UpdatedAt time.Time `gorm:"autoUpdateTime"` // UpdatedAt adalah waktu terakhir kali foto diperbarui.
// }

// // Validate adalah metode untuk melakukan validasi terhadap foto.
// // Metode ini memeriksa apakah judul dan URL foto tidak kosong.
// func (p *Photo) Validate() error {
// 	if p.Title == "" {
// 		return errors.New("title is required") // Jika judul foto kosong, maka akan dikembalikan error.
// 	}
// 	if p.PhotoURL == "" {
// 		return errors.New("photo URL is required") // Jika URL foto kosong, maka akan dikembalikan error.
// 	}
// 	return nil // Jika validasi berhasil, tidak ada error yang dikembalikan.
// }

// Implement database operations if needed
// Implementasikan operasi-operasi database jika diperlukan.
// Operasi-operasi database seperti membuat, membaca, memperbarui, dan menghapus foto dapat diimplementasikan di sini.
