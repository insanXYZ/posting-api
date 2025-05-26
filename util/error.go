package util

import (
	"posting-api/dto"

	"github.com/go-playground/validator/v10"
)

func GetErrorResponse(msg string, err error) dto.Response {
	res := dto.Response{
		Message: msg,
	}

	if err != nil {
		if ValidationErrors, ok := err.(validator.ValidationErrors); ok {
			res.Error = GetErrorValidateMessageStruct(ValidationErrors)
		} else {
			res.Message = err.Error()
		}
	}

	return res
}
