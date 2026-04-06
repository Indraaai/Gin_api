package handler

import (
	"GinGolang/internal/dto"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. KONSUMEN MENDIFINISIKAN INTERFACE
// Handler butuh Service, jadi Handler mendefinisikan fungsinya di sini.
// Perhatikan: Handler tidak butuh tahu soal Repository atau GORM.
type UserService interface {
	Register(ctx context.Context, req dto.UserRegister) (dto.UserResponse, error)
}

// 2. STRUCT CONCRETE
type UserHandler struct {
	svc UserService
}

// 3. CONSTRUCTOR
func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// 4. METHOD HANDLER: Menerima request dari Gin Router
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.UserRegister

	// Langkah 1: Parsing JSON dari Body HTTP ke struct DTO
	// Fungsi ShouldBindJSON ini SANGAT SAKTI. Dia otomatis mengecek aturan
	// binding:"required,email" yang ada di DTO kita.
	if err := c.ShouldBindJSON(&req); err != nil {
		// Jika validasi gagal (misal email salah format, atau password kurang dari 6 huruf)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Format data tidak valid",
			"details": err.Error(),
		})
		return
	}

	// Langkah 2: Panggil logika bisnis di layer Service
	// c.Request.Context() meneruskan context bawaan dari HTTP request Gin
	res, err := h.svc.Register(c.Request.Context(), req)
	if err != nil {
		// Cek apakah errornya karena email duplikat (biasanya ini Bad Request 400)
		if err.Error() == "email sudah digunakan" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		// Jika error karena sistem (misal database mati)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendaftarkan user"})
		return
	}

	// Langkah 3: Kembalikan Response Sukses ke Client
	// 201 Created adalah status code HTTP standar untuk pembuatan data baru
	c.JSON(http.StatusCreated, gin.H{
		"message": "User berhasil didaftarkan",
		"data":    res,
	})
}
