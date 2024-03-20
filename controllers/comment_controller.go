//implementasi dari  fungsi yang digunakan untuk meng-handle request HTTP terkait dengan fitur-fitur dalam aplikasi.
package controllers

import (
    "github.com/gin-gonic/gin"
    "myGram/models"
    "myGram/utils"
)

// CreateComment --> fungsi untuk membuat komentar baru.
func CreateComment(c *gin.Context) {
	var comment models.Comment // Mendeklarasikan variabel comment yang akan menampung data komentar yang diterima dari client.

	// Memeriksa dan melakukan binding data JSON yang diterima dari client ke dalam variabel comment.
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.BadRequest(c, "Invalid request body") // Jika terjadi error dalam binding data, maka respon error dikirimkan kembali ke client.
		return
	}

	// Implement validation and authorization logic ---> 
	// Di sini biasanya dilakukan validasi terhadap data komentar dan juga otorisasi untuk menentukan apakah pengguna memiliki izin untuk membuat komentar.
    // ...............
	// Save comment to database (models.CreateComment)
	// Di sini biasanya dilakukan penyimpanan data komentar ke dalam database dengan menggunakan fungsi CreateComment yang ada di dalam package models.

	utils.Created(c, comment) // Mengirimkan respon berhasil ke client setelah komentar berhasil dibuat.
}

// GetComments adalah fungsi untuk mendapatkan daftar komentar.
func GetComments(c *gin.Context) {
	// Implement logic to retrieve comments from database
	//  pengambilan data komentar dari database menggunakan fungsi GetComments yang ada di dalam package models.

	// comments := models.GetComments() // Mendapatkan daftar komentar dari database.

	// Return comments as response
	// daftar komentar yang telah didapatkan akan dikirimkan kembali
}