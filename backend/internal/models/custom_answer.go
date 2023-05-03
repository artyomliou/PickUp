package models

type CustomAnswer struct {
	Qid  uint   `form:"qid" binding:"required"`
	Text string `form:"text" binding:"required"`
}
