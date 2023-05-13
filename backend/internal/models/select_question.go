package models

import (
	"pick-up/backend/internal/db"
	"time"

	"gorm.io/gorm"
)

type SelectQuestion struct {
	ID         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreId    uint           `json:"-"`
	Store      Store          `json:"-"`
	Products   []*Product     `json:"-" gorm:"many2many:select_question_product"`
	Name       string         `json:"name" gorm:"not null"`
	IsRequired bool           `json:"isRequired" gorm:"not null;default=false"`
	AtMost     uint           `json:"atMost" gorm:"not null;default=0"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

func SeedSelectQuestion(storeId uint) (*SelectQuestion, []uint, error) {
	question := &SelectQuestion{
		StoreId:    storeId,
		Name:       "冰量",
		IsRequired: true,
		AtMost:     1,
	}
	return NewSelectQuestion(question, []string{
		"正常冰",
		"少冰",
		"微冰",
		"去冰",
		"完全去冰",
	})
}

func NewSelectQuestion(question *SelectQuestion, optNames []string) (*SelectQuestion, []uint, error) {
	if tx := db.Conn().Save(&question); tx.Error != nil {
		return nil, nil, tx.Error
	}

	options := make([]uint, len(optNames))
	for _, optName := range optNames {
		optModel := SelectOption{
			SelectQuestionId: question.ID,
			Name:             optName,
			Price:            0,
		}
		if tx := db.Conn().Save(&optModel); tx.Error != nil {
			return nil, nil, tx.Error
		}
		options = append(options, optModel.ID)
	}

	return question, options, nil
}

func AssociateProductWithSelectQuestion(product *Product, question *SelectQuestion) error {
	return db.Conn().Model(&product).Association("SelectQuestions").Append(question)
}
