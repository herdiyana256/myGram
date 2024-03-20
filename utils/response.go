package utils

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// BadRequest mengirimkan respon dengan status HTTP 400 Bad Request.
// Digunakan ketika terjadi kesalahan dalam permintaan yang dikirimkan oleh klien.
func BadRequest(c *gin.Context, message string) {
    c.JSON(http.StatusBadRequest, gin.H{"message": message})
}

// Created mengirimkan respon dengan status HTTP 201 Created.
// Dipanggil setelah sebuah entitas berhasil dibuat dalam sistem.
func Created(c *gin.Context, data interface{}) {
    c.JSON(http.StatusCreated, data)
}

// InternalServerError mengirimkan respon dengan status HTTP 500 Internal Server Error.
// Digunakan ketika terjadi kesalahan internal yang tidak diharapkan dalam server.
func InternalServerError(c *gin.Context, message string) {
    c.JSON(http.StatusInternalServerError, gin.H{"message": message})
}

func Unauthorized(c *gin.Context, message string) {
    c.JSON(http.StatusUnauthorized, gin.H{"message": message})
}