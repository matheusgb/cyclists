package validators

import (
	"github.com/go-playground/validator/v10"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
)

var mapLoginUserRequestErrors = map[string]string{
	"email":    "invalid email",
	"password": "invalid password",
}

func LoginUser(request requests.LoginUser) map[string]string {
	err := Validator.Struct(request)
	if err != nil {
		errors := make(map[string]string)
		for _, error := range err.(validator.ValidationErrors) {
			field := error.Field()
			errors[field] = mapLoginUserRequestErrors[field]
		}
		return errors
	}
	return nil
}
