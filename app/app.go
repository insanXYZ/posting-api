package app

import (
	"fmt"
	"os"
	"posting-api/config"
	"posting-api/controller"
	"posting-api/repository"
	"posting-api/route"
	"posting-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	echo      *echo.Echo
	db        *gorm.DB
	validator *validator.Validate
}

func Init() *App {
	return &App{
		echo:      config.NewEcho(),
		db:        config.NewGorm(),
		validator: config.NewValidator(),
	}
}

func (a *App) Run() error {
	// repository
	userRepository := repository.NewUserRepository(a.db)

	// service
	authService := service.NewAuthService(a.validator, userRepository)

	// controller
	authController := controller.NewAuthController(authService)

	route.SetRoute(&route.RouteConfig{
		Echo:           a.echo,
		AuthController: authController,
	})

	return a.echo.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
