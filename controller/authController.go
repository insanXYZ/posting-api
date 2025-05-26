package controller

import (
	"net/http"
	"posting-api/dto"
	"posting-api/dto/message"
	"posting-api/service"
	"posting-api/util"
	"time"

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
		return util.HttpResponseError(ctx, message.FAILED_REGISTER, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_REGISTER, nil)
}

func (a *AuthController) Login(ctx echo.Context) error {
	req := new(dto.LoginRequest)

	err := ctx.Bind(req)
	if err != nil {
		return err
	}

	accToken, refToken, err := a.userService.HandleLogin(ctx.Request().Context(), req)
	if err != nil {
		return util.HttpResponseError(ctx, message.FAILED_LOGIN, err)
	}

	cookie := &http.Cookie{
		Name:     "refresh-token",
		Path:     "/api/refresh",
		Secure:   true,
		Value:    "Bearer " + refToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
	}

	ctx.SetCookie(cookie)

	return util.HttpResponseSuccess(ctx, message.SUCCESS_LOGIN, echo.Map{
		"access_token": "Bearer " + accToken,
	})
}

func (a *AuthController) Refresh(ctx echo.Context) error {
	claims := util.GetClaims(ctx)
	newToken, err := a.userService.HandleRefresh(ctx.Request().Context(), claims)
	if err != nil {
		return util.HttpResponseError(ctx, message.FAILED_REFRESH_TOKEN, err)
	}

	return util.HttpResponseSuccess(ctx, message.SUCCESS_REFRESH_TOKEN, echo.Map{
		"access_token": "Bearer " + newToken,
	})
}
