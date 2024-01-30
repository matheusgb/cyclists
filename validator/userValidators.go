package validator

import (
	"github.com/go-playground/validator/v10"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
)

var mapCreateUserRequestErrors = map[string]string{
	"name":                  "Invalid name, it must have at least 3 characters and at most 255 characters",
	"email":                 "Invalid email",
	"password":              "Invalid password, it must have at least 6 characters and at most 255 characters",
	"password_confirmation": "Invalid password confirmation, it must be equal to password",
}

func CreateUser(request requests.CreateUser) map[string]string {
	err := Validator.Struct(request)
	if err != nil {
		errors := make(map[string]string)
		for _, error := range err.(validator.ValidationErrors) {
			field := error.Field()
			errors[field] = mapCreateUserRequestErrors[field]
		}
		return errors
	}
	return nil
}

func UpdateUser(request requests.UpdateUser) map[string]string {
	err := Validator.Struct(request)
	if err != nil {
		errors := make(map[string]string)
		for _, error := range err.(validator.ValidationErrors) {
			field := error.Field()
			errors[field] = mapCreateUserRequestErrors[field]
		}
		return errors
	}
	return nil
}
