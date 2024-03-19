package utils

import (
	"net/http"         // Import package http untuk status HTTP
	"github.com/gin-gonic/gin" // Import package gin untuk framework Gin
)

// Created digunakan untuk memberikan respons dengan status HTTP 201 Created.
// Biasanya dipanggil setelah sebuah entitas berhasil dibuat dalam sistem.
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

// InternalServerError digunakan untuk memberikan respons dengan status HTTP 500 Internal Server Error.
// Digunakan ketika terjadi kesalahan internal yang tidak diharapkan dalam server.
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": message})
}

// BadRequest digunakan untuk memberikan respons dengan status HTTP 400 Bad Request.
// Biasanya digunakan ketika terjadi kesalahan dalam permintaan yang dikirimkan oleh klien.
func BadRequest(c *gin.Cont
