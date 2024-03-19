// package routes: Package yang berisi definisi routing untuk API endpoints.

import (
	"github.com/gin-gonic/gin" // Mengimport package gin untuk membuat router HTTP.
	"myGram/controllers" // Mengimport package controllers yang berisi fungsi-fungsi untuk menangani permintaan HTTP.
	"myGram/middleware" // Mengimport package middleware untuk menambahkan middleware ke rute.
)

// SetupRoutes mengonfigurasi semua rute API.
func SetupRoutes(router *gin.Engine) {
	// Rute untuk registrasi pengguna dan login.
	router.POST("/users/register", controllers.RegisterUser)
	router.POST("/users/login", controllers.LoginUser)

	// Rute untuk operasi pengguna (update dan hapus) dengan autentikasi middleware.
	userRoutes := router.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.PUT("/:userId", controllers.UpdateUser) // Rute untuk memperbarui pengguna.
		userRoutes.DELETE("/", controllers.DeleteUser)    // Rute untuk menghapus pengguna.
	}

	// Rute untuk operasi foto (buat, baca, update, hapus) dengan autentikasi middleware dan otorisasi foto middleware.
	photoRoutes := router.Group("/photos")
	photoRoutes.Use(middleware.AuthMiddleware())
	{
		photoRoutes.POST("/", controllers.CreatePhoto) // Rute untuk membuat foto baru.
		photoRoutes.GET("/", controllers.GetPhotos)    // Rute untuk mendapatkan daftar foto.
		photoRoutes.PUT("/:photoId", middleware.AuthorizePhoto(), controllers.UpdatePhoto) // Rute untuk memperbarui foto dengan otorisasi.
		photoRoutes.DELETE("/:photoId", middleware.AuthorizePhoto(), controllers.DeletePhoto) // Rute untuk menghapus foto dengan otorisasi.
	}

	// Rute untuk operasi komentar (buat, baca, update, hapus) dengan autentikasi middleware dan otorisasi komentar middleware.
	commentRoutes := router.Group("/comments")
	commentRoutes.Use(middleware.AuthMiddleware())
	{
		commentRoutes.POST("/", controllers.CreateComment) // Rute untuk membuat komentar baru.
		commentRoutes.GET("/", controllers.GetComments)    // Rute untuk mendapatkan daftar komentar.
		commentRoutes.PUT("/:commentId", middleware.AuthorizeComment(), controllers.UpdateComment) // Rute untuk memperbarui komentar dengan otorisasi.
		commentRoutes.DELETE("/:commentId", middleware.AuthorizeComment(), controllers.DeleteComment) // Rute untuk menghapus komentar dengan otorisasi.
	}

	// Rute untuk operasi media sosial (buat, baca, update, hapus) dengan autentikasi middleware dan otorisasi media sosial middleware.
	socialMediaRoutes := router.Group("/socialmedias")
	socialMediaRoutes.Use(middleware.AuthMiddleware())
	{
		socialMediaRoutes.POST("/", controllers.CreateSocialMedia) // Rute untuk membuat media sosial baru.
		socialMediaRoutes.GET("/", controllers.GetSocialMedias)    // Rute untuk mendapatkan daftar media sosial.
		socialMediaRoutes.PUT("/:socialMediaId", middleware.AuthorizeSocialMedia(), controllers.UpdateSocialMedia) // Rute untuk memperbarui media sosial dengan otorisasi.
		socialMediaRoutes.DELETE("/:socialMediaId", middleware.AuthorizeSocialMedia(), controllers.DeleteSocialMedia) // Rute untuk menghapus media sosial dengan otorisasi.
	}
}
