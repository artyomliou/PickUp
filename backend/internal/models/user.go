package models

import (
	"database/sql"
	"fmt"
	"pick-up/backend/internal/db"
	"time"

	"github.com/google/uuid"
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

func (u *User) IsPasswordMatched(inputPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(inputPassword)) == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SeedUser() (*User, error) {
	return NewUser(fmt.Sprintf("%s@not-exist.com", uuid.New()), "ultra_secret")
}

func NewUser(email string, password string) (*User, error) {
	passwordHash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := User{
		Email:    email,
		Password: passwordHash,
	}
	tx := db.Conn().Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	u := &User{}
	if tx := db.Conn().Where("email = ?", email).First(u); tx.Error != nil {
		return nil, tx.Error
	}
	return u, nil
}
