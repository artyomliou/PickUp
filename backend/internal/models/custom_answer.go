package models

type CustomAnswer struct {
	Qid  uint   `binding:"required"`
	Text string `binding:"required"`
}

type CustomAnswers []CustomAnswer
