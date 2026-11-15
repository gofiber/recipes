package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ErrRefreshTokenUnusable is returned by RotateRefreshToken when the refresh token is unknown,
// already revoked or expired, which includes a token that another request has just rotated.
var ErrRefreshTokenUnusable = errors.New("refresh token is unknown, revoked or expired")

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

// RotateRefreshToken revokes the refresh token and creates its successor, which expires when the
// old one would have, in one transaction.
//
// Revoking is a single conditional UPDATE, not a read followed by a write. When several requests
// present the same token at once, the database lets exactly one of them change the row from not
// revoked to revoked, and only that one gets RowsAffected == 1 and a successor. The others
// get ErrRefreshTokenUnusable, so a refresh token can only be used once.
func (r *RefreshTokenRepository) RotateRefreshToken(token *RefreshToken) (*RefreshToken, error) {
	var next *RefreshToken
	err := r.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&RefreshToken{}).
			Where("token = ? AND revoked = ? AND expires_at > ?", token.Token, false, time.Now()).
			Update("revoked", true)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return ErrRefreshTokenUnusable
		}

		next = &RefreshToken{
			Token:     uuid.New().String(),
			UserId:    token.UserId,
			ExpiresAt: token.ExpiresAt,
		}
		return tx.Create(next).Error
	})
	if err != nil {
		return nil, err
	}
	return next, nil
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
