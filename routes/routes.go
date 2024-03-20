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
    router.POST("/users/register", gin.WrapF(func(c *gin.Context) {
        controllers.RegisterUser(c, db)
    }))
    router.POST("/users/login", gin.WrapF(func(c *gin.Context) {
        controllers.LoginUser(c, db)
    }))

    // Inisialisasi grup rute pengguna dengan middleware otentikasi
    userRoutes := router.Group("/users")
    userRoutes.Use(middleware.AuthMiddleware())
    {
        userRoutes.PUT("/:userId", gin.WrapF(func(c *gin.Context) {
            controllers.UpdateUser(c, db)
        }))
        userRoutes.DELETE("/:userId", gin.WrapF(func(c *gin.Context) {
            controllers.DeleteUser(c, db)
        }))
    }

    // Inisialisasi grup rute foto dengan middleware otentikasi
    photoRoutes := router.Group("/photos")
    photoRoutes.Use(middleware.AuthMiddleware())
    {
        photoRoutes.POST("/", controllers.CreatePhoto(db))
        photoRoutes.GET("/", controllers.GetPhotos(db))
        photoRoutes.PUT("/:photoId", middleware.AuthorizePhoto(db, "update"), gin.WrapF(func(c *gin.Context) {
            controllers.UpdatePhoto(c, db)
        }))
        photoRoutes.DELETE("/:photoId", middleware.AuthorizePhoto(db, "delete"), gin.WrapF(func(c *gin.Context) {
            controllers.DeletePhoto(c, db)
        }))
    }

    // Inisialisasi grup rute komentar dengan middleware otentikasi
    commentRoutes := router.Group("/comments")
    commentRoutes.Use(middleware.AuthMiddleware())
    {
        commentRoutes.POST("/", middleware.AuthorizeComment(db, "create"), controllers.CreateComment(db))
        commentRoutes.GET("/", controllers.GetComments(db))
        commentRoutes.PUT("/:commentId", middleware.AuthorizeComment(db, "update"), gin.WrapF(func(c *gin.Context) {
            controllers.UpdateComment(c, db)
        }))
        commentRoutes.DELETE("/:commentId", middleware.AuthorizeComment(db, "delete"), gin.WrapF(func(c *gin.Context) {
            controllers.DeleteComment(c, db)
        }))
    }

    // Inisialisasi grup rute media sosial dengan middleware otentikasi
    socialMediaRoutes := router.Group("/socialmedias")
    socialMediaRoutes.Use(middleware.AuthMiddleware())
    {
        socialMediaRoutes.POST("/", middleware.AuthorizeSocialMedia(db, "create"), controllers.CreateSocialMedia(db))
        socialMediaRoutes.GET("/", controllers.GetSocialMedias(db))
        socialMediaRoutes.PUT("/:socialMediaId", middleware.AuthorizeSocialMedia(db, "update"), gin.WrapF(func(c *gin.Context) {
            controllers.UpdateSocialMedia(c, db)
        }))
        socialMediaRoutes.DELETE("/:socialMediaId", middleware.AuthorizeSocialMedia(db, "delete"), gin.WrapF(func(c *gin.Context) {
            controllers.DeleteSocialMedia(c, db)
        }))
    }
}
