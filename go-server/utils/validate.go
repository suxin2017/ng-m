package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func CheckStruct(s interface{}) error {
	if err := Validate.Struct(s); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			return validationError
		}
	}
	return nil
}
