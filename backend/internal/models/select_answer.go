package models

type SelectAnswer struct {
	Qid     uint   `form:"qid" binding:"required"`
	Options []uint `form:"options" binding:"required"`
}
