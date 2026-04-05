package main

import (
	"GinGolang/internal/config"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("gagal me load env")
	}
	db := config.ConnectDB()
	_ = db

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "Server Gin berjalan dan Database siap!",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port jika tidak di-set di .env
	}

	log.Printf(" Server mulai berjalan di http://localhost:%s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}

}
