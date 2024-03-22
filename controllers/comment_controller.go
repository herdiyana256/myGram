package controllers

import (
    "github.com/gin-gonic/gin"
    "myGram/models"
    "myGram/utils"
)

// CreateComment adalah fungsi untuk membuat komentar baru.
func CreateComment(c *gin.Context) {
    var comment models.Comment

    // Binding data JSON yang diterima dari client ke dalam variabel comment.
    if err := c.ShouldBindJSON(&comment); err != nil {
        utils.BadRequest(c, "Invalid request body")
        return
    }

    // Simpan komentar ke database
    db, err := utils.GetDB()
    if err != nil {
        utils.InternalServerError(c, "Failed to connect to database")
        return
    }
    createdComment, err := models.CreateComment(db, comment.UserID, comment.PhotoID, comment.Message)
    if err != nil {
        utils.InternalServerError(c, "Failed to create comment")
        return
    }

    utils.Created(c, createdComment)
}

// GetComments adalah fungsi untuk mendapatkan daftar komentar.
func GetComments(c *gin.Context) {
    // Mendapatkan daftar komentar dari database
    db, err := utils.GetDB()
    if err != nil {
        utils.InternalServerError(c, "Failed to connect to database")
        return
    }
    
    // Mendapatkan photoID dari query parameter
    photoID := c.Query("photoID")
    
    _, err = models.GetCommentsByPhotoID(db, photoID)
    if err != nil {
        utils.InternalServerError(c, "Failed to retrieve comments")
        return
    }

    utils.OK(c) // Perbaiki pemanggilan utils.OK agar tidak memiliki argumen
}
