package admin

import (
	"errors"
	"fmt"

	"goUniAdmin/internal/db"

	"goUniAdmin/internal/config"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AdminService manages admin CRUD operations with GORM
type AdminService struct {
	db  *db.DB
	cfg *config.Config // Pointer to config
}

// NewAdminService initializes the service with a GORM database connection and config
func NewAdminService(db *db.DB, cfg *config.Config) *AdminService {
	return &AdminService{
		db:  db,
		cfg: cfg,
	}
}

func (s *AdminService) Create(admin Admin) (Admin, error) {
	if err := ValidateAdmin(admin); err != nil {
		return Admin{}, err
	}

	var existingAdmin Admin
	if err := s.db.Where("email_id = ? AND	is_deleted=?", admin.EmailID, false).First(&existingAdmin).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return Admin{}, err // Database error
	}

	if existingAdmin.ID != uuid.Nil {
		return Admin{}, errors.New("Email already exists")
	}

	if err := s.db.Create(&admin).Error; err != nil {
		return Admin{}, err
	}
	return admin, nil
}

// Read retrieves an admin by ID
func (s *AdminService) Read(id uuid.UUID) (Admin, error) {
	var admin Admin
	result := s.db.Where("id = ? AND is_deleted = ?", id, false).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Admin{}, errors.New("Admin not found")
		}
		return Admin{}, result.Error
	}
	return admin, nil
}

// Update modifies an existing admin
func (s *AdminService) Update(id uuid.UUID, updated Admin) (Admin, error) {
	if err := ValidateAdmin(updated); err != nil {
		return Admin{}, err
	}

	var existing Admin
	result := s.db.Where("id = ? AND is_deleted = ?", id, false).First(&existing)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Admin{}, errors.New("admin not found")
		}
		return Admin{}, result.Error
	}

	updated.ID = id // Ensure ID isn't changed
	if err := s.db.Model(&existing).Updates(updated).Error; err != nil {
		return Admin{}, err
	}
	return updated, nil
}

// Delete performs a soft delete by setting IsDeleted to true
func (s *AdminService) Delete(id uuid.UUID) error {
	result := s.db.Model(&Admin{}).Where("id = ? AND is_deleted = ?", id, false).Update("is_deleted", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("admin not found")
	}
	return nil
}

// List returns all non-deleted admins
func (s *AdminService) List(limit, offset int) ([]Admin, int64, error) {
	var admins []Admin
	// Query with filter, limit, offset, and ordering by created_at in descending order

	result := s.db.Where("is_deleted = ?", false).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&admins)

	// Get the total count of records matching the filter (without limit/offset)
	var totalCount int64
	err := s.db.Model(&Admin{}).Where("is_deleted = ?", false).Count(&totalCount).Error

	// If there was an error fetching the total count, return the error
	if err != nil {
		return nil, 0, err
	}

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return admins, totalCount, nil
}

// Read amin by email
func (s *AdminService) ReadByEmail(email string) (Admin, error) {
	var admin Admin
	result := s.db.Where("email_id =? AND is_deleted =?", email, false).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Admin{}, errors.New("admin not found")
		}
		return Admin{}, result.Error
	}
	return admin, nil
}

// GenerateToken generates a JWT token for the admin
func (s *AdminService) GenerateToken(id uuid.UUID) (string, error) {
	// Check if the service or config is nil
	if s == nil || s.cfg.JWTSecret == "" {
		return "", fmt.Errorf("service or JWT secret is not initialized")
	}

	// Log the secret length and ID for debugging
	fmt.Printf("Generating token for ID: %s, JWTSecret length: %d\n", id.String(), len(s.cfg.JWTSecret))

	// Set the expiration time (e.g., 1 hour from now)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the claims with standard fields
	claims := jwt.MapClaims{
		"id":  id.String(),           // Convert UUID to string for safety
		"exp": expirationTime.Unix(), // Set expiration time as a Unix timestamp
		"iat": time.Now().Unix(),     // Set issued at time as a Unix timestamp
	}

	// Generate the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	signedToken, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return signedToken, nil
}
