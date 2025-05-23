package util

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetErrorValidateMessageStruct(validationErrors validator.ValidationErrors) map[string]string {
	errorsMap := make(map[string]string)

	for _, validationError := range validationErrors {
		field := strings.ToLower(validationError.Field())
		tag := validationError.Tag()

		switch tag {
		case "required":
			errorsMap[field] = field + " is required"
		case "email":
			errorsMap[field] = "invalid email format"
		case "min":
			errorsMap[field] = field + " must be at least " + validationError.Param() + " characters"
		default:
			errorsMap[field] = "invalid " + field
		}
	}

	return errorsMap
}
