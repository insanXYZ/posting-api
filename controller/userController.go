package controller

import (
	"posting-api/dto"
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
		return util.HttpResponseError(ctx, "failed get user", err)
	}

	return util.HttpResponseSuccess(ctx, "success get user", *user)
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
		return util.HttpResponseError(ctx, "failed update user", err)
	}
	return util.HttpResponseSuccess(ctx, "success update user", nil)
}
