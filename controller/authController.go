package controller

import (
	"net/http"
	"posting-api/dto"
	"posting-api/service"
	"posting-api/util"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	userService *service.AuthService
}

func NewAuthController(userService *service.AuthService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (a *AuthController) Register(ctx echo.Context) error {
	req := new(dto.RegisterRequest)

	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	err = a.userService.HandleRegister(ctx.Request().Context(), req)
	if err != nil {
		return util.HttpResponseError(ctx, "failed register", err)
	}

	return util.HttpResponseSuccess(ctx, "success register", nil)
}

func (a *AuthController) Login(ctx echo.Context) error {
	req := new(dto.LoginRequest)

	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	accToken, err := a.userService.HandleLogin(ctx.Request().Context(), req)
	if err != nil {
		return util.HttpResponseError(ctx, "failed login", err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = "Bearer " + accToken
	cookie.Path = "/"

	ctx.SetCookie(cookie)

	return util.HttpResponseSuccess(ctx, "success login", nil)
}
