package repository

import (
	"GinGolang/internal/models"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Struct concrete yang menyimpan koneksi database
type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository adalah constructor yang akan dipanggil di main.go / app wiring
func NewUserRepository(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

// Create: Menyimpan data User baru ke tabel users
func (r *userRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	// Menggunakan WithContext sangat disarankan agar proses query
	// bisa dibatalkan jika request HTTP dari user tiba-tiba terputus / timeout
	err := r.db.WithContext(ctx).Create(user).Error
	return err
}

// FindByEmail: Mencari user berdasarkan alamat email
func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		// Jika error-nya karena data memang tidak ada di database
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Kembalikan nil agar gampang di-if oleh Service nanti
		}
		// Jika error karena hal lain (misal database mati)
		return nil, err
	}

	return &user, nil
}

// Jangan lupa tambahkan import ini di atas

// ... (kode Create dan FindByEmail sebelumnya) ...

// FindByID: Mencari user berdasarkan UUID (Primary Key)
// Ini akan sangat sering dipakai setelah user mendapatkan Token JWT
func (r *userRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	// Mencari berdasarkan Primary Key
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil jika UUID tidak ditemukan
		}
		return nil, err
	}

	return &user, nil
}
