package routes

import (
    "github.com/gin-gonic/gin"
    "myGram/controllers"
    "myGram/middleware"
    "gorm.io/gorm"
)

// SetupRoutes mengatur semua rute aplikasi.
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    // Definisikan rute untuk register dan login pengguna
    router.POST("/users/register", func(c *gin.Context) {
        controllers.RegisterUser(c, db)
    })
    router.POST("/users/login", func(c *gin.Context) {
        controllers.LoginUser(c, db)
    })

    // Inisialisasi grup rute pengguna dengan middleware otentikasi
    userRoutes := router.Group("/users")
    userRoutes.Use(middleware.AuthMiddleware())
    {
        userRoutes.PUT("/:userId", func(c *gin.Context) {
            // Implementasi UpdateUser
        })
        userRoutes.DELETE("/:userId", func(c *gin.Context) {
            // Implementasi DeleteUser
        })
    }

    // Inisialisasi grup rute foto dengan middleware otentikasi
    photoRoutes := router.Group("/photos")
    photoRoutes.Use(middleware.AuthMiddleware())
    {
        photoRoutes.POST("/", func(c *gin.Context) {
            controllers.CreatePhoto(c)
        })
        photoRoutes.GET("/", func(c *gin.Context) {
            controllers.GetPhotos(c)
        })
        photoRoutes.PUT("/:photoId", func(c *gin.Context) {
            // Implementasi UpdatePhoto
        })
        photoRoutes.DELETE("/:photoId", func(c *gin.Context) {
            // Implementasi DeletePhoto
        })
    }

    // Inisialisasi grup rute komentar dengan middleware otentikasi
    commentRoutes := router.Group("/comments")
    commentRoutes.Use(middleware.AuthMiddleware())
    {
        commentRoutes.POST("/", func(c *gin.Context) {
            controllers.CreateComment(c)
        })
        commentRoutes.GET("/", func(c *gin.Context) {
            controllers.GetComments(c)
        })
        commentRoutes.PUT("/:commentId", func(c *gin.Context) {
            // Implementasi UpdateComment
        })
        commentRoutes.DELETE("/:commentId", func(c *gin.Context) {
            // Implementasi DeleteComment
        })
    }

    // Inisialisasi grup rute media sosial dengan middleware otentikasi
    socialMediaRoutes := router.Group("/socialmedias")
    socialMediaRoutes.Use(middleware.AuthMiddleware())
    {
        socialMediaRoutes.POST("/", func(c *gin.Context) {
            // Implementasi CreateSocialMedia
        })
        socialMediaRoutes.GET("/", func(c *gin.Context) {
            // Implementasi GetSocialMedias
        })
        socialMediaRoutes.PUT("/:socialMediaId", func(c *gin.Context) {
            // Implementasi UpdateSocialMedia
        })
        socialMediaRoutes.DELETE("/:socialMediaId", func(c *gin.Context) {
            // Implementasi DeleteSocialMedia
        })
    }
}
