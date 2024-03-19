package main

import (
	"github.com/gin-gonic/gin"  // Import package gin untuk membuat server HTTP
	"myGram/routes"             // Import package routes untuk mengatur rute HTTP
)

func main() {
	r := gin.Default()  // Membuat instance router Gin dengan konfigurasi default

	// Konfigurasi routes menggunakan fungsi SetupRoutes dari package routes
	routes.SetupRoutes(r)

	// Menjalankan server HTTP pada port 8080
	r.Run(":8080")
}
