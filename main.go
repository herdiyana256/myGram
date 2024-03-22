package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"myGram/models"
	"myGram/routes"
	"net/http"
)

func main() {
	DbHost := "localhost"
	DbPort := 5433
	DbUser := "postgres"
	DbPassword := "Sitcomindo@123"
	DbName := "mygram"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)

	// Koneksi ke database PostgreSQL: myGram {}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrasi model-model ....
	db.AutoMigrate(&models.User{}) // membuat tabel user

	r := gin.Default()

	// Serve static files (like CSS, JavaScript, and images) from the "static" directory
	r.StaticFS("/static", http.Dir("static"))

	// Rute untuk menampilkan halaman index.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Configure routes  ....
	routes.SetupRoutes(r, db) // koneksi database ke fungsi SetupRoutes

	// Run the server
	r.Run(":8080")
}
