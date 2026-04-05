package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserRegister struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// UserLoginReq adalah format JSON yang diharapkan saat user melakukan login
type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Opsional: Jika nanti setelah login kamu ingin mengembalikan Token JWT
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"` // Menggunakan ulang UserResponse yang sudah kita buat!
}
