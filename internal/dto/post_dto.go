package dto

import (
	"time"

	"github.com/google/uuid"
)

// PostCreateReq adalah format JSON dari client saat membuat artikel/post
type PostCreateReq struct {
	Title   string `json:"title" binding:"required,min=5,max=255"`
	Content string `json:"content" binding:"required"`
	// Catatan: Kita TIDAK meminta user_id dari JSON request body.
	// Kenapa? Karena user_id yang valid nanti akan kita ambil dari Token JWT
	// saat user sudah login, supaya user tidak bisa memalsukan post atas nama orang lain.
}

// PostResponse adalah format JSON saat menampilkan artikel
type PostResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
