package main

import (
	"log"
	"os"

	"sewa-lapangan-voli/config"
	"sewa-lapangan-voli/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env (diabaikan jika tidak ada, misal di production pakai env langsung)
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("[ENV] file .env tidak ditemukan, pakai environment variable sistem")
	}

	// Inisialisasi database + jalankan migrasi
	config.InitDB()

	// Set Gin mode dari env (debug / release)
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Setup semua route
	router.SetupRouter(r)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("[APP] server berjalan di port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("[APP] gagal menjalankan server: %v", err)
	}
}
