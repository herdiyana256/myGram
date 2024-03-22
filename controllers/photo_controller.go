package controllers

import (
    "github.com/gin-gonic/gin"
    "myGram/models"
    "myGram/utils"
)

// CreatePhoto adalah fungsi untuk membuat foto baru.
func CreatePhoto(c *gin.Context) {
    var photo models.Photo

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
    db, err := utils.GetDB()
    if err != nil {
        utils.InternalServerError(c, "Failed to connect to database")
        return
    }
    _, err = models.CreatePhoto(db, photo.Title, photo.Caption, photo.PhotoURL, photo.UserID)
    if err != nil {
        utils.InternalServerError(c, "Failed to create photo")
        return
    }

    utils.OK(c) // Mengirimkan status OK tanpa data
}

// GetPhotos adalah fungsi untuk mendapatkan daftar foto.
func GetPhotos(c *gin.Context) {
    db, err := utils.GetDB()
    if err != nil {
        utils.InternalServerError(c, "Failed to connect to database")
        return
    }
    _, err = models.GetPhotos(db)
    if err != nil {
        utils.InternalServerError(c, "Failed to retrieve photos")
        return
    }

    utils.OK(c) // Perbaiki pemanggilan utils.OK agar tidak memiliki argumen
}
