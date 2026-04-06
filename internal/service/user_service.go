package service

import (
	"context"
	"errors"

	"GinGolang/internal/dto"
	"GinGolang/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// 1. KONSUMEN MENDIFINISIKAN INTERFACE
// Service butuh Repository, jadi Service mendefinisikan apa saja yang dia butuhkan.
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}

// 2. STRUCT CONCRETE
type userServiceImpl struct {
	repo UserRepository
}

// 3. CONSTRUCTOR
func NewUserService(repo UserRepository) *userServiceImpl {
	return &userServiceImpl{repo: repo}
}

// 4. BUSINESS LOGIC: REGISTER
func (s *userServiceImpl) Register(ctx context.Context, req dto.UserRegister) (dto.UserResponse, error) {
	// Aturan Bisnis 1: Cek apakah email sudah terdaftar
	existingUser, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.UserResponse{}, err // Error dari database (misal DB mati)
	}
	if existingUser != nil {
		return dto.UserResponse{}, errors.New("email sudah digunakan") // Error validasi bisnis
	}

	// Aturan Bisnis 2: Hash Password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, errors.New("gagal memproses password")
	}

	// Mapping dari DTO Request ke Model Database
	userModel := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Simpan ke database via Repository
	err = s.repo.Create(ctx, &userModel)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// Mapping dari Model Database ke DTO Response (Sembunyikan password!)
	response := dto.UserResponse{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		CreatedAt: userModel.CreatedAt,
	}

	return response, nil
}
