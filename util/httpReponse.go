package util

import (
	"net/http"
	"posting-api/dto"

	"github.com/labstack/echo/v4"
)

func HttpResponseSuccess(ctx echo.Context, msg string, data any, statusCode ...int) error {
	if len(statusCode) == 0 {
		statusCode = append(statusCode, http.StatusOK)
	}

	return ctx.JSON(statusCode[0], dto.Response{
		Data:    data,
		Message: msg,
	})
}

func HttpResponseError(ctx echo.Context, msg string, err error, statusCode ...int) error {
	if len(statusCode) == 0 {
		statusCode = append(statusCode, http.StatusBadRequest)
	}

	return ctx.JSON(statusCode[0], GetErrorResponse(msg, err))
}
