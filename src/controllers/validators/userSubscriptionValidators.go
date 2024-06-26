package validators

import (
	"github.com/go-playground/validator/v10"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/userSubscription"
)

var mapUserSubscriptionRequestErrors = map[string]string{
	"bike_event_id": "invalid bike event id, it must be a number greater than 0",
	"user_id":       "invalid user id, it must be a number greater than 0",
}

func UserSubscription(request requests.UserSubscription) map[string]string {
	err := Validator.Struct(request)
	if err != nil {
		errors := make(map[string]string)
		for _, error := range err.(validator.ValidationErrors) {
			field := error.Field()
			errors[field] = mapUserSubscriptionRequestErrors[field]
		}
		return errors
	}
	return nil
}
