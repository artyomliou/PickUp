package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type SelectAnswer struct {
	Qid     uint   `form:"qid" binding:"required"`
	Options []uint `form:"options" binding:"required"`
}

type SelectAnswers []SelectAnswer

func (data *SelectAnswers) Value() (driver.Value, error) {
	b, err := json.Marshal(*data)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
func (data *SelectAnswers) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	return nil
}
