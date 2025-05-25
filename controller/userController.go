package controller

import (
	"posting-api/dto"
	"posting-api/dto/converter"
	"posting-api/dto/message"
	"posting-api/service"
	"posting-api/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) GetUser(ctx echo.Context) error {
	claims := ctx.Get("user").(jwt.MapClaims)

	user, err := u.userService.HandleGetUser(ctx.Request().Context(), claims)
	if err != nil {
		return util.HttpResponseError(ctx, message.FAILED_GET_USER, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_GET_USER, converter.UserToResponseDto(user))
}

func (u *UserController) UpdateUser(ctx echo.Context) error {
	claims := ctx.Get("user").(jwt.MapClaims)
	req := new(dto.UpdateUserRequest)

	err := ctx.Bind(req)
	if err != nil {
		return err
	}
	err = u.userService.HandleUpdateUser(ctx.Request().Context(), claims, req)
	if err != nil {
		return util.HttpResponseError(ctx, message.FAILED_UPDATE_USER, err)
	}
	return util.HttpResponseSuccess(ctx, message.SUCCESS_UPDATE_USER, nil)
}

func (u *UserController) DeleteUser(ctx echo.Context) error {
	claims := util.GetClaims(ctx)

	err := u.userService.HandleDeleteUser(ctx.Request().Context(), claims)
	if err != nil {
		return util.HttpResponseError(ctx, message.FAILED_DELETE_USER, err)
	}
	return util.HttpResponseSuccess(ctx, message.SUCCESS_DELETE_USER, nil)
}
