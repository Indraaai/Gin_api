package routes

import (
	"github.com/gin-gonic/gin"

	// Ingat: Ganti "GinGolang" dengan nama module yang ada di baris pertama file go.mod kamu
	"GinGolang/internal/handler"
)

// SetupUserRoutes bertugas mendaftarkan semua endpoint yang berhubungan dengan User
func SetupUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	// Kita buat prefix /api/users agar URL-nya rapi
	// Nanti URL-nya jadi: POST http://localhost:8080/api/users/register
	userGroup := router.Group("/api/users")
	{
		userGroup.POST("/register", userHandler.Register)

		// Nanti rute Login juga akan masuk ke sini:
		// userGroup.POST("/login", userHandler.Login)
	}
}
