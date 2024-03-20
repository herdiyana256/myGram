package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "myGram/controllers"
    "myGram/models"
    "myGram/routes"
)

func main() {
    //  koneksi ke database PostgreSQL : myGram {}
    dsn := "user=postgres password=Sitcomindo@123 dbname=mygram port=5432 sslmode=disable TimeZone=Asia/Jakarta"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // Auto migrasi model-model ....
    db.AutoMigrate(&models.User{}) //   membuat tabel user

    r := gin.Default()

    // Serve static files (like index.html, CSS, JS)
    r.Static("/static", "./static")

    // Configure routes  ....
    routes.SetupRoutes(r, db) //  koneksi database ke fungsi SetupRoutes

    // Run the server
    r.Run(":8080")
}
