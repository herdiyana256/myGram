// package controllers: Package yang berisi fungsi-fungsi untuk meng-handle request HTTP terkait dengan manajemen pengguna (user) dalam aplikasi.

import (
	"github.com/gin-gonic/gin" // Mengimport library gin yang digunakan untuk membuat server HTTP dengan Go.
	"myGram/models"             // Mengimport package models yang berisi definisi model data aplikasi.
	"myGram/utils"              // Mengimport package utils yang berisi utilitas yang digunakan dalam aplikasi.
)

// RegisterUser adalah fungsi untuk mendaftarkan pengguna baru.
func RegisterUser(c *gin.Context) {
	var user models.User // Mendeklarasikan variabel user yang akan menampung data pengguna yang diterima dari client.

	// Memeriksa dan melakukan binding data JSON yang diterima dari client ke dalam variabel user.
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequest(c, "Invalid request body") // Jika terjadi error dalam binding data, maka respon error dikirimkan kembali ke client.
		return
	}

	// Memeriksa validitas data pengguna yang diterima.
	if err := user.Validate(); err != nil {
		// Perbaiki pemanggilan utils.ValidationError
		utils.BadRequest(c, err.Error()) // Jika terdapat kesalahan dalam validasi, maka respon error dikirimkan kembali ke client.
		return
	}

	// Melakukan hashing terhadap password pengguna.
	if err := user.HashPassword(); err != nil {
		utils.InternalServerError(c, "Error hashing password") // Jika terjadi kesalahan dalam proses hashing password, maka respon error dikirimkan kembali ke client.
		return
	}

	// Menyimpan data pengguna baru ke dalam database.
	if err := models.CreateUser(&user); err != nil {
		utils.InternalServerError(c, "Error creating user") // Jika terjadi kesalahan dalam penyimpanan data pengguna ke dalam database, maka respon error dikirimkan kembali ke client.
		return
	}

	utils.Created(c, user) // Mengirimkan respon berhasil ke client setelah pengguna berhasil didaftarkan.
}

// Implement other controller functions (LoginUser, UpdateUser, DeleteUser, etc.)
// Di sini dapat diimplementasikan fungsi-fungsi lain untuk manajemen pengguna seperti login, update, dan delete user.
