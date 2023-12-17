package entity

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	bycrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	ID             string     `json:"id"`
	Email          string     `json:"email"`
	HashSalt       string     `json:"hash_salt"`
	HashedPassword string     `json:"hashed_password"`
	VerifiedAt     *time.Time `json:"verified_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
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

func (a *Account) SetPassword(password string) (err error) {
	a.HashSalt = time.Now().Format(time.RFC3339Nano)
	saltedPassword := a.HashSalt + password
	hashedPasswordByte, err := bycrypt.GenerateFromPassword([]byte(saltedPassword), bycrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.HashedPassword = string(hashedPasswordByte)
	return nil
}

func (c Account) ComparePassword(password string) bool {
	saltedPassword := c.HashSalt + password
	return bycrypt.CompareHashAndPassword([]byte(c.HashedPassword), []byte(saltedPassword)) == nil
}

func GenerateRandomPassword() string {
	// generate random password of length 8
	long := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = long[rand.Int63()%int64(len(long))]
	}
	return string(b)
}
