package models

import (
    "errors"
    "time"
    "gorm.io/gorm"
)

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

// Membuat foto baru
func CreatePhoto(db *gorm.DB, title, caption, photoURL string, userID uint) (*Photo, error) {
    photo := &Photo{
        Title:    title,
        Caption:  caption,
        PhotoURL: photoURL,
        UserID:   userID,
    }
    result := db.Create(&photo)
    if result.Error != nil {
        return nil, result.Error
    }
    return photo, nil
}

// Mendapatkan semua foto
func GetPhotos(db *gorm.DB) ([]*Photo, error) {
    var photos []*Photo
    result := db.Find(&photos)
    if result.Error != nil {
        return nil, result.Error
    }
    return photos, nil
}

// Memperbarui foto
func UpdatePhoto(db *gorm.DB, id uint, title, caption, photoURL string) (*Photo, error) {
    var photo Photo
    result := db.First(&photo, id)
    if result.Error != nil {
        return nil, result.Error
    }

    photo.Title = title
    photo.Caption = caption
    photo.PhotoURL = photoURL

    result = db.Save(&photo)
    if result.Error != nil {
        return nil, result.Error
    }

    return &photo, nil
}

// Menghapus foto
func DeletePhoto(db *gorm.DB, id uint) error {
    result := db.Delete(&Photo{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// Validasi data foto
func (p *Photo) Validate() error {
    if p.Title == "" {
        return errors.New("title is required")
    }
    if p.Caption == "" {
        return errors.New("caption is required")
    }
    if p.PhotoURL == "" {
        return errors.New("photo URL is required")
    }
    return nil
}
