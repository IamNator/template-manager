package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	VerifiedAt *time.Time `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (Account) TableName() string {
	return "accounts"
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now().UTC()
	}
	return nil
}
