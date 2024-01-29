package validator

import (
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

type ErrorResponse struct {
	Tag string
}

func Init() {
	Validator = validator.New(validator.WithRequiredStructEnabled())
}

func Struct(s interface{}) []ErrorResponse {
	err := Validator.Struct(s)
	if err != nil {
		var errors []ErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ErrorResponse{
				Tag: err.Tag(),
			})
		}
		return errors
	}
	return nil
}
