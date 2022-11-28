package model

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func (c CustomValidator) Validate(i interface{}) error {
	if err := c.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}
