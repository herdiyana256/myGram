// package controllers: Package yang berisi fungsi-fungsi untuk meng-handle request HTTP terkait dengan fitur social media dalam aplikasi.

import (
	"github.com/gin-gonic/gin" // Mengimport library gin yang digunakan untuk membuat server HTTP dengan Go.
	"myGram/models"             // Mengimport package models yang berisi definisi model data aplikasi.
	"myGram/utils"              // Mengimport package utils yang berisi utilitas yang digunakan dalam aplikasi.
)

// CreateSocialMedia adalah fungsi untuk membuat entri social media baru.
func CreateSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia // Mendeklarasikan variabel socialMedia yang akan menampung data social media yang diterima dari client.

	// Memeriksa dan melakukan binding data JSON yang diterima dari client ke dalam variabel socialMedia.
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		utils.BadRequest(c, "Invalid request body") // Jika terjadi error dalam binding data, maka respon error dikirimkan kembali ke client.
		return
	}

	// Implement validation and authorization logic here
	// Di sini biasanya dilakukan validasi terhadap data social media dan juga otorisasi untuk menentukan apakah pengguna memiliki izin untuk membuat entri social media.

	// Save social media to database (models.CreateSocialMedia)
	// Di sini biasanya dilakukan penyimpanan data social media ke dalam database dengan menggunakan fungsi CreateSocialMedia yang ada di dalam package models.

	utils.Created(c, socialMedia) // Mengirimkan respon berhasil ke client setelah entri social media berhasil dibuat.
}

// GetSocialMedias adalah fungsi untuk mendapatkan daftar social media.
func GetSocialMedias(c *gin.Context) {
	// Implement logic to retrieve social medias from database
	// Di sini biasanya dilakukan pengambilan data social media dari database menggunakan fungsi GetSocialMedias yang ada di dalam package models.

	// socialMedias := models.GetSocialMedias() // Mendapatkan daftar social media dari database.

	// Re
