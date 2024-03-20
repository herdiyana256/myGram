package controllers

import (
    "github.com/gin-gonic/gin"
    "myGram/models" 
    "myGram/utils"
    "gorm.io/gorm"
)

// RegisterUser untuk  fungsi  mendaftarkan pengguna baru.
func RegisterUser(c *gin.Context, db *gorm.DB) {
    var user models.User // Mendeklarasikan variabel user yang akan menampung data pengguna yang diterima dari client.

    // Memeriksa dan melakukan binding data JSON yang diterima dari client ke dalam variabel user.
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.BadRequest(c, "Invalid request body") // Jika terjadi error dalam binding data, maka respon error dikirimkan kembali ke client.
        return
    }

    // Memeriksa validitas data pengguna yang diterima.
    if err := user.Validate(); err != nil {
        utils.BadRequest(c, err.Error()) // Jika terdapat kesalahan dalam validasi, maka respon error dikirimkan kembali ke client.
        return
    }

    // Melakukan hashing terhadap password pengguna.
    if err := user.HashPassword(); err != nil {
        utils.InternalServerError(c, "Error hashing password") // Jika terjadi kesalahan dalam proses hashing password, maka respon error dikirimkan kembali ke client.
        return
    }

    // Menyimpan data pengguna baru ke dalam database.
    if err := db.Create(&user).Error; err != nil {
        utils.InternalServerError(c, "Error creating user") // Jika terjadi kesalahan dalam penyimpanan data pengguna ke dalam database, maka respon error dikirimkan kembali ke client.
        return
    }

    utils.Created(c, user) // send  respon berhasil ke client setelah pengguna berhasil didaftarkan.
}

// LoginUser untuk  fungsi  melakukan autentikasi pengguna.
func LoginUser(c *gin.Context, db *gorm.DB) {
    // Implementasi autentikasi pengguna (login) di sini
}

// UpdateUser adalah fungsi untuk memperbarui informasi pengguna.
func UpdateUser(c *gin.Context, db *gorm.DB) {
    // Implementasi pembaruan informasi pengguna di sini
}

// DeleteUser adalah fungsi untuk menghapus pengguna.
func DeleteUser(c *gin.Context, db *gorm.DB) {
    // Implementasi penghapusan pengguna di sini
}
