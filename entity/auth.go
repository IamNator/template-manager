package entity

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	bycrypt "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	ID             string     `json:"id" gorm:"primaryKey;column:id"`
	Email          string     `json:"email" gorm:"unique;column:email"`
	HashSalt       string     `json:"hash_salt" gorm:"column:hash_salt"`
	HashedPassword string     `json:"hashed_password" gorm:"column:hashed_password"`
	VerifiedAt     *time.Time `json:"verified_at" gorm:"column:verified_at"`
	CreatedAt      time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      *time.Time `json:"updated_at" gorm:"column:updated_at"`
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

type Session struct {
	ID        string    `json:"id" gorm:"primaryKey;column:id"`
	AccountID string    `json:"account_id" gorm:"column:account_id;not null"`
	Token     string    `json:"token" gorm:"column:token;not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"column:expires_at;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamptz"`
}

func (Session) TableName() string {
	return "sessions"
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now().UTC()
	}
	return nil
}

func (s *Session) GenerateToken(signingKey string) error {
	if s.ExpiresAt.IsZero() {
		s.ExpiresAt = time.Now().Add(time.Hour * 1)
	}
	token, err := generateJWT(s.AccountID, signingKey, s.ExpiresAt)
	if err != nil {
		return err
	}
	s.Token = token
	return nil
}

func generateJWT(accountID, signingKey string, exp time.Time) (string, error) {
	claims := jwt.MapClaims{
		"account_id": accountID,
		"exp":        exp.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}
