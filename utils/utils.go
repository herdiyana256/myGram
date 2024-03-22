package utils

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// GetDB   untuk mendapatkan koneksi database.
func GetDB() (*gorm.DB, error) {
    // DSN (Data Source Name) untuk koneksi ke database PostgreSQL
    dsn := "user=postgres password=Sitcomindo@123 dbname=mygram port=5432 sslmode=disable TimeZone=Asia/Jakarta"

    // Membuat koneksi ke database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Mengembalikan koneksi database yang berhasil dibuka
    return db, nil
}
