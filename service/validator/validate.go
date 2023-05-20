package validator

import (
	"wb_task1/entity"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(data entity.Order) error {
	err := validate.Struct(data)
	if err != nil {
		return err
	}
	return nil
}
