package models

import (
	"pick-up/backend/internal/db"
	"time"

	"gorm.io/gorm"
)

type SelectQuestion struct {
	ID            uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId       uint            `json:"-"`
	Store         Store           `json:"-"`
	Products      []*Product      `json:"-" gorm:"many2many:select_question_product"`
	SelectOptions []*SelectOption `json:"options"`
	Name          string          `json:"name" gorm:"not null"`
	IsRequired    bool            `json:"isRequired" gorm:"not null;default=false"`
	AtMost        uint            `json:"atMost" gorm:"not null;default=0"`
	CreatedAt     time.Time       `json:"-"`
	UpdatedAt     time.Time       `json:"-"`
	DeletedAt     gorm.DeletedAt  `json:"-"`
}

func SeedSelectQuestion(storeId uint) (*SelectQuestion, []uint, error) {
	question := &SelectQuestion{
		StoreId:    storeId,
		Name:       "冰量",
		IsRequired: true,
		AtMost:     1,
	}
	return NewSelectQuestion(question, []SelectOption{
		{
			Name: "正常冰 Regular Ice",
		},
		{
			Name: "少冰 Less Ice",
		},
		{
			Name: "微冰 Easy Ice",
		},
		{
			Name: "去冰 Ice-Free",
		},
	})
}

func NewSelectQuestion(question *SelectQuestion, options []SelectOption) (*SelectQuestion, []uint, error) {
	if tx := db.Conn().Save(&question); tx.Error != nil {
		return nil, nil, tx.Error
	}

	for _, opt := range options {
		opt.SelectQuestionId = question.ID
		if tx := db.Conn().Save(&opt); tx.Error != nil {
			return nil, nil, tx.Error
		}
	}

	optionIds := make([]uint, len(options))
	for _, opt := range options {
		optionIds = append(optionIds, opt.ID)
	}

	return question, optionIds, nil
}

func AssociateProductWithSelectQuestion(product *Product, question *SelectQuestion) error {
	return db.Conn().Model(&product).Association("SelectQuestions").Append(question)
}
