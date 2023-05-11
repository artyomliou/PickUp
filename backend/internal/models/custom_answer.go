package models

type CustomAnswer struct {
	CustomQuestionId uint   `json:"customQuestionId" binding:"required"`
	Text             string `json:"text" binding:"required"`
}

type CustomAnswers []CustomAnswer
