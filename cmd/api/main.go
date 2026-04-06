package main

import (
	"GinGolang/internal/config"
	"GinGolang/internal/handler"
	"GinGolang/internal/repository"
	"GinGolang/internal/routes"
	"GinGolang/internal/service"
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

	userRepo := repository.NewUserRepository(db)

	// Inject Repository ke dalam Service
	userService := service.NewUserService(userRepo)

	// Inject Service ke dalam Handler
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "Server Gin berjalan dan Database siap!",
		})
	})
	routes.SetupUserRoutes(router, userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port jika tidak di-set di .env
	}

	log.Printf(" Server mulai berjalan di http://localhost:%s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}

}
