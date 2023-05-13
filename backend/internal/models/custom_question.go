package models

import (
	"pick-up/backend/internal/db"
	"time"

	"gorm.io/gorm"
)

type CustomQuestion struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId     uint           `json:"-"`
	Store       Store          `json:"-"`
	Products    []*Product     `json:"-" gorm:"many2many:custom_question_product"`
	Name        string         `json:"name" gorm:"not null"`
	Placeholder string         `json:"placeholder" gorm:"not null"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

func SeedCustomQuestion(storeId uint) (*CustomQuestion, error) {
	q := &CustomQuestion{
		StoreId:     storeId,
		Name:        "有額外需求嗎？",
		Placeholder: "請在這邊寫您的需求",
	}
	return NewCustomQuestion(q)
}

func NewCustomQuestion(q *CustomQuestion) (*CustomQuestion, error) {
	if tx := db.Conn().Save(q); tx.Error != nil {
		return nil, tx.Error
	}
	return q, nil
}

func AssociateProductWithCustomQuestion(product *Product, question *CustomQuestion) error {
	return db.Conn().Model(&product).Association("CustomQuestions").Append(question)
}
