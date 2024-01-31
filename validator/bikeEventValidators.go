package validator

import (
	"github.com/go-playground/validator/v10"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/bikeEvent"
)

var mapBikeEventRequestErrors = map[string]string{
	"name":                    "Invalid name, it must have at least 3 characters and at most 255 characters",
	"start_place":             "Invalid start place, it must have at least 2 characters and at most 255 characters",
	"start_date":              "Invalid start date, it must be a valid date",
	"start_date_registration": "Invalid start date registration, it must be a valid date",
	"end_date_registration":   "Invalid end date registration, it must be a valid date",
	"organizer":               "Invalid organizer, it must be a number greater than 0",
	"participants_limit":      "Invalid participants limit, it must be a number between 2 and 1000",
	"aditional_information":   "Invalid aditional information, it must have at least 3 characters and at most 255 characters",
}

func BikeEvent[requestBikeEvent requests.CreateBikeEvent | requests.UpdateBikeEvent](request requestBikeEvent) map[string]string {
	err := Validator.Struct(request)
	if err != nil {
		errors := make(map[string]string)
		for _, error := range err.(validator.ValidationErrors) {
			field := error.Field()
			errors[field] = mapBikeEventRequestErrors[field]
		}
		return errors
	}
	return nil
}
