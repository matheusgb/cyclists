package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

type ErrorResponse struct {
	Tag string
}

func Init() {
	Validator = validator.New(validator.WithRequiredStructEnabled())
	makeValidatorCatchJsonTagFromRequestStruct()
}

func makeValidatorCatchJsonTagFromRequestStruct() {
	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}
