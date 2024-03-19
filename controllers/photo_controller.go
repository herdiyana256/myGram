// package controllers: Ini adalah package yang berisi implementasi dari berbagai fungsi yang digunakan untuk meng-handle request HTTP terkait dengan fitur-fitur dalam aplikasi.

import (
	"github.com/gin-gonic/gin" // Mengimport library gin yang digunakan untuk membuat server HTTP dengan Go.
	"myGram/models"             // Mengimport package models yang berisi definisi model data aplikasi.
	"myGram/utils"              // Mengimport package utils yang berisi utilitas yang digunakan dalam aplikasi.
)

// CreatePhoto adalah fungsi untuk membuat foto baru.
func CreatePhoto(c *gin.Context) {
	var photo models.Photo // Mendeklarasikan variabel photo yang akan menampung data foto yang diterima dari client.

	// Memeriksa dan melakukan binding data JSON yang diterima dari client ke dalam variabel photo.
	if err := c.ShouldBindJSON(&photo); err != nil {
		utils.BadRequest(c, "Invalid request body") // Jika terjadi error dalam binding data, maka respon error dikirimkan kembali ke client.
		return
	}

	// Implement validation and authorization logic here
	// Di sini biasanya dilakukan validasi terhadap data foto dan juga otorisasi untuk menentukan apakah pengguna memiliki izin untuk membuat foto.

	// Save photo to database (models.CreatePhoto)
	// Di sini biasanya dilakukan penyimpanan data foto ke dalam database dengan menggunakan fungsi CreatePhoto yang ada di dalam package models.

	utils.Created(c, photo) // Mengirimkan respon berhasil ke client setelah foto berhasil dibuat.
}

// GetPhotos adalah fungsi untuk mendapatkan daftar foto.
func GetPhotos(c *gin.Context) {
	// Implement logic to retrieve photos from database
	// Di sini biasanya dilakukan pengambilan data foto dari database menggunakan fungsi GetPhotos yang ada di dalam package models.

	// photos := models.GetPhotos() // Mendapatkan daftar foto dari database.

	// Return photos as response
	// Di sini biasanya daftar foto yang telah didapatkan akan dikirimkan kembali sebagai respon ke client.
}

// Implement UpdatePhoto, DeletePhoto, and AuthorizePhoto middleware similarly
// Fungsi ini biasanya digunakan untuk melakukan update, delete, dan otorisasi terhadap permintaan yang berkaitan dengan foto, seperti misalnya memastikan pengguna memiliki izin untuk mengedit atau menghapus foto.
