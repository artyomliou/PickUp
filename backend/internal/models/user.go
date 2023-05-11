package models

import (
	"database/sql"
	"log"
	"pick-up/backend/internal/db"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Email           string         `json:"email" gorm:"size:100;not null;index"`
	EmailVerifiedAt sql.NullTime   `json:"-"`
	Password        string         `json:"-" gorm:"size:256"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-"`
	Carts           []*Cart        `json:"carts,omitempty"`
	Orders          []*Order       `json:"orders,omitempty"`
}

func (u *User) Create(email string, password string) {
	passwordHash, _ := HashPassword(password)
	u.Email = email
	u.Password = passwordHash

	if result := db.Conn().Create(u); result.Error != nil {
		log.Panic(result.Error)
	}
}

func (u *User) GetByEmail(email string) {
	db.Conn().Where("email = ?", email).First(u)
}

func (u *User) IsPasswordMatched(password string) bool {
	return CheckPasswordHash(u.Password, password)
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
