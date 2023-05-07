package models

type SelectAnswer struct {
	Qid     uint   `binding:"required"`
	Options []uint `binding:"required"`
}

type SelectAnswers []SelectAnswer
