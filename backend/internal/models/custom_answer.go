package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type CustomAnswer struct {
	Qid  uint   `form:"qid" binding:"required"`
	Text string `form:"text" binding:"required"`
}

type CustomAnswers []CustomAnswer

func (data *CustomAnswers) Value() (driver.Value, error) {
	b, err := json.Marshal(*data)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
func (data *CustomAnswers) Scan(value interface{}) error {
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
