package models

import (
    "errors"
    "time"
)

// Struktur data Comment
type Comment struct {
    ID        uint      `gorm:"primaryKey"` // ID adalah identifikasi unik dari komentar.
    UserID    uint      `gorm:"not null"`   // UserID adalah ID pengguna yang membuat komentar.
    PhotoID   uint      `gorm:"not null"`   // PhotoID adalah ID foto yang dikomentari.
    Message   string    `gorm:"not null"`   // Message adalah pesan yang dituliskan dalam komentar.
    CreatedAt time.Time `gorm:"autoCreateTime"` // CreatedAt adalah waktu pembuatan komentar.
    UpdatedAt time.Time `gorm:"autoUpdateTime"` // UpdatedAt adalah waktu terakhir kali komentar diperbarui.
}

// Struktur data Photo
type Photo struct {
    ID        uint      `gorm:"primaryKey"` // ID adalah identifikasi unik dari foto.
    Title     string    `gorm:"not null"`   // Title adalah judul atau nama dari foto.
    Caption   string    `gorm:"not null"`   // Caption adalah deskripsi atau keterangan dari foto.
    PhotoURL  string    `gorm:"not null"`   // PhotoURL adalah URL atau lokasi file foto.
    UserID    uint      `gorm:"not null"`   // UserID adalah ID pengguna yang memiliki foto ini.
    CreatedAt time.Time `gorm:"autoCreateTime"` // CreatedAt adalah waktu pembuatan foto.
    UpdatedAt time.Time `gorm:"autoUpdateTime"` // UpdatedAt adalah waktu terakhir kali foto diperbarui.
}

// Struktur data SocialMedia
type SocialMedia struct {
    ID             uint      `gorm:"primaryKey"` // ID adalah identifikasi unik dari media sosial.
    Name           string    `gorm:"not null"`   // Name adalah nama dari media sosial.
    SocialMediaURL string    `gorm:"not null"`   // SocialMediaURL adalah URL atau tautan dari media sosial.
    UserID         uint      `gorm:"not null"`   // UserID adalah ID pengguna yang memiliki media sosial ini.
    CreatedAt      time.Time `gorm:"autoCreateTime"` // CreatedAt adalah waktu pembuatan media sosial.
    UpdatedAt      time.Time `gorm:"autoUpdateTime"` // UpdatedAt adalah waktu terakhir kali media sosial diperbarui.
}

// Validasi pesan komentar
func (c *Comment) Validate() error {
    if c.Message == "" {
        return errors.New("message is required") // Jika pesan komentar kosong, maka akan dikembalikan error.
    }
    return nil // Jika validasi berhasil, tidak ada error yang dikembalikan.
}
