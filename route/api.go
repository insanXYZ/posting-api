package route

import (
	"posting-api/controller"

	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	Echo           *echo.Echo
	AuthController *controller.AuthController
}

func SetRoute(cfg *RouteConfig) {
	echo := cfg.Echo

	echo.POST("/register", cfg.AuthController.Register)
	echo.POST("/login", cfg.AuthController.Login)
}
