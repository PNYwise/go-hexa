package utils

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
	Error       bool
}

var validate = validator.New()

func Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	if errs := validate.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
