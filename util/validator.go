package util

import "github.com/go-playground/validator/v10"

func GetErrorValidateMessageStruct(validationErrors validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)

	for _, validationError := range validationErrors {
		errors[validationError.Field()] = validationError.Error()
	}

	return errors
}
