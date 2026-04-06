package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid,default:gen_random_uuid()"` // GORM otomatis nge-map ini ke kolom 'id'
	Name      string    // GORM otomatis nge-map ini ke kolom 'name'
	Email     string    // GORM otomatis nge-map ini ke kolom 'email'
	Password  string    // Ingat, kita butuh password untuk fitur login nanti!
	CreatedAt time.Time `gorm:"autoCreateTime"` // Otomatis diisi GORM saat Insert
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Otomatis di-update GORM saat Update

	Posts []Post `gorm:"foreignKey:UserID"` // Relasi one-to-many dengan Post
}

// Pastikan import "github.com/google/uuid" dan "gorm.io/gorm" sudah ada di atas

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New() // Bikin UUID acak yang valid
	}
	return
}
