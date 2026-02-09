package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RefreshToken represents a refresh token in the system
type RefreshToken struct {
	gorm.Model
	UserId    uint      `gorm:"not null" json:"user_id"`
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Revoked   bool      `gorm:"not null" json:"revoked"`
}

// RefreshTokenRepository handles database operations for refresh tokens
type RefreshTokenRepository struct {
	db *gorm.DB
}

// NewRefreshTokenRepository creates a new refresh token repository
func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

// CreateRefreshToken creates a new refresh token for a user
func (r *RefreshTokenRepository) CreateRefreshToken(userId uint, ttl time.Duration) (*RefreshToken, error) {
	token := &RefreshToken{
		Token:     uuid.New().String(),
		UserId:    userId,
		ExpiresAt: time.Now().Add(ttl),
		Revoked:   false,
	}
	if err := r.db.Create(token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

// GetRefreshToken retrieves a refresh token by its token string
func (r *RefreshTokenRepository) GetRefreshToken(tokenString string) (*RefreshToken, error) {
	var token RefreshToken
	if err := r.db.Where("token = ?", tokenString).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

// RevokeRefreshToken marks a refresh token as revoked
func (r *RefreshTokenRepository) RevokeRefreshToken(tokenString string) error {
	if err := r.db.Model(&RefreshToken{}).Where("token = ?", tokenString).Update("revoked", true).Error; err != nil {
		return err
	}
	return nil
}

// RevokeAllUserTokens marks all refresh tokens for a user as revoked.
func (r *RefreshTokenRepository) RevokeAllUserTokens(userId uint) error {
	if err := r.db.Model(&RefreshToken{}).Where("user_id = ?", userId).Update("revoked", true).Error; err != nil {
		return err
	}
	return nil
}
