package validators

import (
	"github.com/go-playground/validator/v10"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
)

var mapUserRequestErrors = map[string]string{
	"name":                  "invalid name, it must have at least 3 characters and at most 255 characters",
	"email":                 "invalid email",
	"password":              "invalid password, it must have at least 6 characters",
	"password_confirmation": "invalid password confirmation, it must be equal to password",
}

func User[requestUser requests.CreateUser | requests.UpdateUser | requests.SendPasswordResetEmail | requests.ResetPassword](request requestUser) map[string]string {
	err := Validator.Struct(request)
	if err != nil {
		errors := make(map[string]string)
		for _, error := range err.(validator.ValidationErrors) {
			field := error.Field()
			errors[field] = mapUserRequestErrors[field]
		}
		return errors
	}
	return nil
}
