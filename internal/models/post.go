package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID     uuid.UUID `gorm:"primaryKey;type:uuid,default:gen_random_uuid()"`
	UserId uuid.UUID `gorm:"type:uuid;not null"`

	Title     string
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relasi balik (Belongs-To): Menandakan Post ini milik satu User
	User User `gorm:"foreignKey:UserID"` // GORM otomatis nge-map ini ke kolom 'user_id'
}
