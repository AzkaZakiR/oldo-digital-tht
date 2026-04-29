package pkg

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())

		switch e.Tag() {
		case "required":
			errors[field] = field + " is required"
		case "email":
			errors[field] = "invalid email format"
		case "min":
			errors[field] = field + " is too short"
		case "numeric":
			errors[field] = field + " must be numeric"
		default:
			errors[field] = "invalid value"
		}
	}

	return errors
}