package controllers

import (
	"github.com/gin-gonic/gin"
	"myGram/models"
	"myGram/utils"
)

// CreatePhoto adalah fungsi untuk membuat foto baru.
func CreatePhoto(c *gin.Context) {
	var photo models.Photo

	// Binding data JSON yang diterima dari client ke dalam variabel photo.
	if err := c.ShouldBindJSON(&photo); err != nil {
		utils.BadRequest(c, "Invalid request body")
		return
	}

	// Validasi data foto
	if err := photo.Validate(); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// Simpan foto ke database
	if err := models.CreatePhoto(&photo); err != nil {
		utils.InternalServerError(c, "Failed to create photo")
		return
	}

	utils.Created(c, photo)
}

// GetPhotos adalah fungsi untuk mendapatkan daftar foto.
func GetPhotos(c *gin.Context) {
	// Mendapatkan daftar foto dari database
	photos, err := models.GetPhotos()
	if err != nil {
		utils.InternalServerError(c, "Failed to retrieve photos")
		return
	}

	utils.OK(c, photos)
}

// UpdatePhoto adalah fungsi untuk mengupdate informasi foto.
func UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")
	var updatedPhoto models.Photo

	// Binding data JSON yang diterima dari client ke dalam variabel updatedPhoto.
	if err := c.ShouldBindJSON(&updatedPhoto); err != nil {
		utils.BadRequest(c, "Invalid request body")
		return
	}

	// Validasi data foto yang diperbarui
	if err := updatedPhoto.Validate(); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	// Update informasi foto di database
	if err := models.UpdatePhoto(photoID, &updatedPhoto); err != nil {
		utils.InternalServerError(c, "Failed to update photo")
		return
	}

	utils.OK(c, gin.H{"message": "Photo updated successfully"})
}

// DeletePhoto adalah fungsi untuk menghapus foto.
func DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	// Hapus foto dari database
	if err := models.DeletePhoto(photoID); err != nil {
		utils.InternalServerError(c, "Failed to delete photo")
		return
	}

	utils.OK(c, gin.H{"message": "Photo deleted successfully"})
}
