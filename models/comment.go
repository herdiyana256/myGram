package models

import (
    "errors"
    "time"
    "gorm.io/gorm"
)

// Comment adalah model untuk menyimpan data komentar.
type Comment struct {
    gorm.Model
    UserID    uint      `json:"userId" gorm:"not null"` // UserID adalah ID pengguna yang membuat komentar.
    PhotoID   uint      `json:"photoId" gorm:"not null"` // PhotoID adalah ID foto yang dikomentari.
    Message   string    `json:"message" gorm:"not null"` // Message adalah pesan yang dituliskan dalam komentar.
    CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // CreatedAt adalah waktu pembuatan komentar.
    UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // UpdatedAt adalah waktu terakhir kali komentar diperbarui.
}

// Membuat komentar baru
func CreateComment(db *gorm.DB, userID, photoID uint, message string) (*Comment, error) {
    comment := &Comment{
        UserID:  userID,
        PhotoID: photoID,
        Message: message,
    }
    result := db.Create(&comment)
    if result.Error != nil {
        return nil, result.Error
    }
    return comment, nil
}

// Mendapatkan semua komentar berdasarkan ID foto
func GetCommentsByPhotoID(db *gorm.DB, photoID string) ([]*Comment, error) {
    var comments []*Comment
    result := db.Where("photo_id = ?", photoID).Find(&comments)
    if result.Error != nil {
        return nil, result.Error
    }
    return comments, nil
}

// Validasi pesan komentar
func (c *Comment) Validate() error {
    if c.Message == "" {
        return errors.New("message is required")
    }
    return nil
}
