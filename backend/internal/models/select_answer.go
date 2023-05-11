package models

type SelectAnswer struct {
	SelectQuestionId uint   `json:"selectQuestionId" binding:"required"`
	Options          []uint `json:"options" binding:"required"`
}

type SelectAnswers []SelectAnswer
